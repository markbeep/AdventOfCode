let fs = require('fs');

let d = fs.readFileSync("day7.txt", "utf8")
    .split(",")
    .map(e => e*1);

let m1 = Infinity;
let m2 = Infinity;
let ma = Math.max(...d);
for (let j = 0; j <= ma; j++) {
    m1 = Math.min(d.reduce((t, v) => t += Math.abs(v-j), 0), m1);
    m2 = Math.min(d.reduce((t, v) => t += calc(v, j), 0), m2);
}

function calc(a, b) {
    let c = 0;
    for (let i = 1; i <= Math.abs(a-b); i++) {
        c += i;
    }
    return c;
}
console.log("Task 1: " + m1);
console.log("Task 2: " + m2);