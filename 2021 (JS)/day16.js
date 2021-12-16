require('fs').readFile("day16.txt", "utf-8", function(err, data) {
    let inp = data.trim();
    let binaryLength = inp.length * 4;
    inp = BigInt("0x"+inp);
    // console.log("VALUE:", inp);
    main(inp.toString(2).padStart(binaryLength, "0"));
});

function main(binary) {
    let res = rec(binary);
    console.log("Result:", res);
}

function printBinary(binary) {
    console.log("\033[31m"+binary.slice(0,3)+"\033[32m"+binary.slice(3,6)+"\033[0m"+binary.slice(6));
}

function rec(binary) {
    let lit = new Literal(binary);
    console.log(lit);
    let val = 0n;
    let index = 0n;
    let version = lit.version;

    if (lit.type == 4) {
        val += lit.literal;
        index = lit.literalEnding;
    }
    else {
        if (lit.lengthIndicator == 0) {
            let cursor = 0n;
            let str = lit.rest;
            while (cursor < lit.length) {
                console.log("CURSOR", cursor, "BINARY", str, "LENGTH", str.length, "LITLEN", lit.length);
                let res = rec(str);
                val += res.val;
                version += res.version;
                cursor += res.index;
                console.log(res.index);
                str = str.slice(Number(res.index));
                break;
            }
            index = 7n + 15n + lit.length;
        }
        else {
            let str = lit.rest;
            for (let i = 0n; i < lit.packets; i++) {
                let res = rec(str);
                val += res.val;
                version += res.version;
                str = str.slice(Number(res.index));
            }
        }
    }
    return { val, index, version };
}  

class Literal {
    constructor(binary) {
        printBinary(binary);
        this.version = BigInt("0b"+binary.slice(0, 3));
        this.type = BigInt("0b"+binary.slice(3, 6));
        this.lengthIndicator = -1;
        this.length = 0n;
        this.literalEnding;
        this.literal;
        this.rest = binary.slice(6);
        this.packets;
        if (this.type == 4) {
            let l = [];
            let cursor = 0;
            while (true) {
                l.push(this.rest.slice(cursor+1, cursor+5));
                if (this.rest[cursor] === "0") break;
                cursor += 5;
            }
            this.literalEnding = BigInt(cursor+5+6);
            this.literal = BigInt("0b"+l.join(""));
        } 
        else {
            this.lengthIndicator = parseInt(binary[6]);
            if (this.lengthIndicator == 0) {
                this.length = BigInt("0b"+binary.slice(7, 22));
                this.rest = binary.slice(22);
            } else {
                this.packets = BigInt("0b"+binary.slice(7, 18));
                this.rest = binary.slice(18);
            }
        }
    }
}
