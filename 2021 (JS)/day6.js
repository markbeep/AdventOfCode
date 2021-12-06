let fs = require('fs');

let d = fs.readFileSync("day6.txt", "utf8").trim().split(",");
d = d.map((e) => Number(e.trim()));

let all = [];
for (let i = 0; i < d.length; i++) {
    all.push(d[i]);
}

let count = new Array(10).fill(0);
all.forEach((e) => {
    count[e]++;
})
let days = 256;
for (let day = 0; day < days-1; day++) {
    for (let i = 0; i < 9; i++) {
        count[i] = count[i+1];
    }
    count[9] = 0;
    if (count[0] > 0) {
        count[7] += count[0];
        count[9] += count[0];
        count[0] = 0;
    }
}

let x = count.reduce((t, a) => t+a, 0);
console.log(x);

