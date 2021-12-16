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

class Literal {
    constructor(binary) {
        this.version = BigInt("0b"+binary.slice(0, 3));
        this.totalVersion = this.version;
        this.type = BigInt("0b"+binary.slice(3, 6));
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

class Operator {
    constructor(binary) {
        this.version = BigInt("0b"+binary.slice(0, 3));
        this.binary = binary;
        this.type = BigInt("0b"+binary.slice(3, 6));
        this.totalPacketSize = 7n;
        this.totalVersion = this.version;
        this.subpackets = [];
        this.typeGotten = -1;
        this.value = 0n;
        let lengthIndicator = parseInt(binary[6]);

        if (lengthIndicator == 0) {
            this.totalPacketSize += 15n;
            if (binary.length < 22) return;
            let length = BigInt("0b"+binary.slice(7, 22));
            let str = binary.slice(22);
            let lengthAdded = 0n;
            while (str.length > 11 && lengthAdded < length) {
                let lit;
                if (getType(str) === 4n) lit = new Literal(str);
                else lit = new Operator(str);
                this.typeGotten = getType(str);
                this.subpackets.push(lit);
                str = str.slice(Number(lit.totalPacketSize));
                this.totalPacketSize += lit.totalPacketSize;
                lengthAdded += lit.totalPacketSize;
                this.totalVersion += lit.totalVersion;
            }
        } else {
            this.totalPacketSize += 11n;
            let packetsCount = BigInt("0b"+binary.slice(7, 18));
            let str = binary.slice(18);
            for (let i = 0n; i < packetsCount; i++) {
                let lit;
                if (str.length < 11) break;
                if (getType(str) === 4n) lit = new Literal(str);
                else lit = new Operator(str);
                this.subpackets.push(lit);
                str = str.slice(Number(lit.totalPacketSize));
                this.totalPacketSize += lit.totalPacketSize;
                this.totalVersion += lit.totalVersion;
            }
        }
        switch (this.type) {
            case 0n: this.value = sum(...this.subpackets); break;
            case 1n: this.value = prod(...this.subpackets); break;
            case 2n: this.value = min(...this.subpackets); break;
            case 3n: this.value = max(...this.subpackets); break;
            case 5n: this.value = (this.subpackets[0].value > this.subpackets[1].value) ? 1n: 0n; break;
            case 6n: this.value = (this.subpackets[0].value < this.subpackets[1].value) ? 1n: 0n; break;
            case 7n: this.value = (this.subpackets[0].value === this.subpackets[1].value) ? 1n: 0n; break;
        }
    }
}

function max() {
    let s = arguments[0].value;
    for (let i = 1; i < arguments.length; i++) {
        s = BigInt(Math.max(Number(arguments[i].value), Number(s)));
    }
    return s;
}

function min() {
    let s = arguments[0].value;
    for (let i = 1; i < arguments.length; i++) {
        s = BigInt(Math.min(Number(arguments[i].value), Number(s)));
    }
    return s;
}

function prod() {
    let s = 1n;
    for (let i = 0; i < arguments.length; i++) {
        s *= arguments[i].value;
    }
    return s;
}

function sum() {
    let s = 0n;
    for (let i = 0; i < arguments.length; i++) {
        s += arguments[i].value;
    }
    return s;
}

function getType(binary) {
    return BigInt("0b"+binary.slice(3, 6));
}