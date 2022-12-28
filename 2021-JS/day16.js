require('fs').readFile("day16.txt", "utf-8", function(err, data) {
    let inp = data.trim();
    let binaryLength = inp.length * 4;
    inp = BigInt("0x"+inp);
    main(inp.toString(2).padStart(binaryLength, "0"));
});

function main(binary) {
    let lit = new Operator(binary);
    console.log("Task 1:", Number(lit.totalVersion));
    console.log("Task 2:", Number(lit.value));
}

class Packet {
    constructor(binary) {
        this.version = BigInt("0b"+binary.slice(0, 3));
        this.totalVersion = this.version;
        this.type = BigInt("0b"+binary.slice(3, 6));
        this.totalPacketSize = 7n;  // per default we assume an OP
        this.value = 0n;
    }
}

class Literal extends Packet {
    constructor(binary) {
        super(binary);
        let l = [];
        let cursor = 0n;
        let str = binary.slice(6);
        while (true) {
            l.push(str.slice(Number(cursor)+1, Number(cursor)+5));
            if (str[cursor] === "0") break;
            cursor += 5n;
        }
        this.totalPacketSize = 6n + 5n + cursor;
        this.value = BigInt("0b"+l.join(""));
    }
}

class Operator extends Packet {
    constructor(binary) {
        super(binary);
        this.totalVersion = this.version;
        this.subpackets = [];
        
        let str, length = Infinity, packetsCount = Infinity;
        let lengthIndicator = parseInt(binary[6]);
        if (lengthIndicator == 0) {
            this.totalPacketSize += 15n;
            length = BigInt("0b"+binary.slice(7, 22));
            str = binary.slice(22);
        } else {
            this.totalPacketSize += 11n;
            packetsCount = parseInt(binary.slice(7, 18), 2);
            str = binary.slice(18);
        }
        let res = createPackets(str, length, packetsCount);
        this.subpackets = res.subpackets;
        this.totalPacketSize += res.lengthAdded;
        this.totalVersion += res.totalVersion;

        // handle operation
        switch (this.type) {
            case 0n: this.value = this.subpackets.reduce((t, a) => t+=a.value, 0n); break;
            case 1n: this.value = this.subpackets.reduce((t, a) => t*=a.value, 1n); break;
            case 2n: this.value = BigInt(Math.min(...this.subpackets.map(e => Number(e.value)))); break;
            case 3n: this.value = BigInt(Math.max(...this.subpackets.map(e => Number(e.value)))); break;
            case 5n: this.value = (this.subpackets[0].value > this.subpackets[1].value) ? 1n: 0n; break;
            case 6n: this.value = (this.subpackets[0].value < this.subpackets[1].value) ? 1n: 0n; break;
            case 7n: this.value = (this.subpackets[0].value === this.subpackets[1].value) ? 1n: 0n; break;
        }
    }
}

/*
 * Creates a list of packets for an Operator and additionally
 * sums up all versions and returns the size of the packet.
 */
function createPackets(binary, length, iterations) {
    let str = binary;
    let subpackets = [];
    let lengthAdded = 0n;
    let totalVersion = 0n;
    let i = 0;
    while (lengthAdded < length && i < iterations) {
        let lit;
        if (getType(str) === 4n) lit = new Literal(str);
        else lit = new Operator(str);
        subpackets.push(lit);
        str = str.slice(Number(lit.totalPacketSize));
        lengthAdded += lit.totalPacketSize;
        totalVersion += lit.totalVersion;
        i++;
    }
    return { lengthAdded, subpackets, totalVersion }
}

function getType(binary) {
    return BigInt("0b"+binary.slice(3, 6));
}