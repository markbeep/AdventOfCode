let fs = require('fs');

let d = fs.readFileSync("day10.txt", "utf8")
    .trim()
    .split("\n");

let t = 0;
let l = [];
let valid = new Array(d.length).fill().map((_,i) => new Array(d[i].length).fill(0));  // to draw table
let boo = new Array(d.length).fill(true);
const m1 = {"}": 1197, ")": 3, "]": 57, ">": 25137};
const m2 = {"(": 1, "[": 2, "{": 3, "<": 4};
for (let y = 0; y < d.length; y++) {
    let s = [];
    for (let x = 0; x < d[y].length; x++) {
        drawTable();
        valid[y][x] = 1;
        if (["(", "[", "{", "<"].includes(d[y][x])) s.push(d[y][x]);
        else {
            let o = d[y][x];
            let v = s.pop();
            if (o === "}" && v === "{") continue;
            if (o === ")" && v === "(") continue;
            if (o === "]" && v === "[") continue;
            if (o === ">" && v === "<") continue;
            valid[y][x] = -1;
            t += m1[o];
            boo[y] = false;
            valid[y] = valid[y].map(_ => -2);
            valid[y][x] = -1;
            break;
        }
    }
    if (boo[y]) {  // task 2  | aka its an incomplete line
        let t2 = 0;
        while(s.length > 0) {
            let v = s.pop();
            t2 = t2 * 5 + m2[v];
        }
        l.push(t2);
    }
}
// Task 1
console.log(t);

// Task 2
l.sort((a,b) => a-b);
console.log(l[Math.floor(l.length/2)]);

function drawTable() {
    console.log("\033[0;0H");  // reset cursor
    msg = []
    for (let y = 0; y < valid.length; y++) {
        let line = "";
        for (let x = 0; x < valid[y].length; x++) {
            if (valid[y][x] === 1) line += "\033[42m" + d[y][x];
            else if (valid[y][x] === -1) line += "\033[101m" + d[y][x];
            else if (valid[y][x] === -2) line += "\033[40m" + d[y][x];
            else line += "\033[47m" + d[y][x];
        }
        msg.push(line);
    }
    console.log(msg.join("\n"));

}