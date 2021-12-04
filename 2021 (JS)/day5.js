export {}

let fs = require('fs');

let d = fs.readFileSync("day5.txt", "utf8").trim().split("\n");
d = d.map((e) => e.trim());

console.log(d);