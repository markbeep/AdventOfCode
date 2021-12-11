let fs = require('fs');

let d = fs.readFileSync("day11.txt", "utf8")
    .trim()
    .split("\n")
    .map(e => e.trim().split("").map(v => Number(v)));

let [f, i, flashes] = [0,0,0];
drawTable();
while (++i) {
    console.log(`Step ${i}`);
    increment();
    drawTable();
    if (i === 100) flashes = f;  // Task 1
    if (countZero() === 100) break;  // Task 2
}
console.log(`Flashes at step 100: ${flashes}`);
console.log(`Synchronized: ${i}`);

function countZero() {
    let t = d.reduce((tot, v) => tot += v.reduce((to, e) => to += Number(e === 0), 0), 0);
    return t;
}

function increment() {
    let pos = []
    d = d.map(e => e.map(v => (v+1)%10));
    for (let y = 0; y < 10; y++) {
        for (let x = 0; x < 10; x++) {
            if (d[y][x] === 0) pos.push([y, x]);
        }
    }
    while (pos.length > 0) flashInc(...pos.pop());
}

function flashInc(y, x) {
    if (d[y][x] !== 0) return;
    f++;
    for (let yi = y-1; yi <= y+1; yi++) {
        for (let xi = x-1; xi <= x+1; xi++) {
            if (xi === x && yi === y || yi<0 || yi>9 || xi<0 || xi>9 || !d[yi][xi]) continue;
            d[yi][xi] = (d[yi][xi]+1)%10;
            flashInc(yi, xi);
        }
    }
}

function drawTable() {
    d.forEach(e => {
        let line = "";
        e.forEach(v => {
            if (v !== 0) line += v;
            else line += "\033[1;31m" + v + "\033[0m";
        });
        console.log(line);
    });
}