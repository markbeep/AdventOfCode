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
    d = d.map(e => e.map(v => (v+1)%10)); // increments all
    pos = d.map((e, y) => e.map((v, x) => [y,x,v]).filter(v => v[2] === 0)); // makes list of all 0's xy coords
    pos = [].concat(...pos);
    while (pos.length > 0) flashInc(...pos.pop());
}

function flashInc(y, x) {
    if (d[y][x] !== 0) return;
    f++;
    for (let yi = Math.max(y-1, 0); yi <= Math.min(y+1, 9); yi++) {
        for (let xi = Math.max(x-1, 0); xi <= Math.min(x+1, 9); xi++) {
            if (xi === x && yi === y || !d[yi][xi]) continue;
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