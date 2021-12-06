let fs = require('fs');

let d = fs.readFileSync("day6.txt", "utf8")
    .trim()
    .split(",")
    .map((e) => Number(e.trim()));

let count = new Array(9).fill(0);
d.forEach((e) => count[e]++);
let days = 256;
for (let day = 0; day < days; day++) {
    let k = count.shift();
    count.push(k);
    count[6] += k;
}
console.log(count.reduce((t, a) => t+a, 0));

