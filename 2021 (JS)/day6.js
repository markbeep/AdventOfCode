let fs = require('fs');

let d = fs.readFileSync("day6.txt", "utf8")
    .split(",")
    .map(e => e*1);

let count = new Array(9).fill(0);
d.forEach(e => count[e]++);
for (let day = 0; day < 256; day++) {
    count.push(count.shift());
    count[6] += count[8];
}
console.log(count.reduce((t, a) => t+a, 0));

