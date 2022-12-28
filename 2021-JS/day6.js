let fs = require('fs');

let d = fs.readFileSync("day6.txt", "utf8")
    .split(",")
    .map(e => e*1);
let x = Array(9).fill(0), y=256; // y is the days
d.forEach(e => x[e]++);
while(y-->0){
    x.push(x.shift());
    x[6] += x[8];
}
console.log(x.reduce((t, a) => t+a));