export {}

let fs = require('fs');

let d: Array<string> = fs.readFileSync("day4.txt", "utf8").trim().split("\n");
d = d.map((e) => e.trim());

console.log(d);