let fs = require('fs');

let d = fs.readFileSync("day7.txt", "utf8")
    .split(",")
    .map(e => e*1);

let m1 = Infinity;
let m2 = Infinity;
let ma = Math.max(...d);
for (let j = 0; j <= ma; j++) {
    m1 = Math.min(d.reduce((t, v) => t += Math.abs(v-j), 0), m1);
    m2 = Math.min(d.reduce((t, v) => t += calc(Math.abs(v-j)), 0), m2);
}

function calc(n) {
    return (n*(n+1))/2;
}
console.log("Task 1: " + m1);
console.log("Task 2: " + m2);