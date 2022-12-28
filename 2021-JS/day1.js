let fs = require('fs');

let d = fs.readFileSync("day1.txt", "utf8");

// 151
// let N = (a) => Number(a);
// let k = d.split('\n')
//     .reduce((t, _, i, b) => t + N(N(b[i])+N(b[i+3])!=NaN && N(b[i])<N(b[i+3])), 0);
// console.log(k)

// 145
let k = d.split('\n')
    .map((a) => Number(a))
    .reduce((t, _, i, b) => t+Number(b[i]+b[i+3]!=NaN && b[i]<b[i+3]), 0);
console.log(k)
