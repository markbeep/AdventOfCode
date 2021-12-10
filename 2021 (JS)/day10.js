let fs = require('fs');

let d = fs.readFileSync("day10.txt", "utf8")
    .trim()
    .split("\n");

let t = 0;
let l = [];
let boo = new Array(d.length).fill(true);
const m1 = {"}": 1197, ")": 3, "]": 57, ">": 25137};
const m2 = {"(": 1, "[": 2, "{": 3, "<": 4};
for (let y = 0; y < d.length; y++) {
    let s = [];
    for (let x = 0; x < d[y].length; x++) {
        if (["(", "[", "{", "<"].includes(d[y][x])) s.push(d[y][x]);
        else {
            let o = d[y][x];
            let v = s.pop();
            if (o === "}" && v === "{") continue;
            if (o === ")" && v === "(") continue;
            if (o === "]" && v === "[") continue;
            if (o === ">" && v === "<") continue;
            t += m1[o];
            boo[y] = false;
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