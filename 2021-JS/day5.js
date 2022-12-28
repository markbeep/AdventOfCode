let fs = require('fs');

let d = fs.readFileSync("day5.txt", "utf8").trim().split("\n");
d = d.map((e) => e.trim());
d = d.map(e => e.split(" -> "));
d = d.map(e => e.map(a => a.split(",").map(x => Number(x))));

// get max values (to create array)
let mx = d.reduce((m, a) => Math.max(m, a[0][0], a[1][0]), 0);
let my = d.reduce((m, a) => Math.max(m, a[0][1], a[0][1]), 0);

// ------------------------ TASK 1 ------------------------
let bo = Array.from(Array(my+1), () => new Array(mx+1).fill(0));
d.forEach(b => {
    let coords = [b[0][0], b[0][1], b[1][0], b[1][1]];
    draw(...coords, bo, 1);
});
console.log("Task 1: " + count(bo));

// ------------------------ TASK 2 ------------------------
bo = Array.from(Array(my+1), () => new Array(mx+1).fill(0));
d.forEach(b => {
    let coords = [b[0][0], b[0][1], b[1][0], b[1][1]];
    bo = draw(...coords, bo, 2);
});
console.log("Task 2: " + count(bo));

// ------------------------ Helper Functions ------------------------
function count(bo) {
    let x = 0;
    bo.forEach(b => {
        b.forEach(e => {
            if (e > 1) x++;
        });
    });
    return x;
}

function draw(x1, y1, x2, y2, bo, star) {
    if (x1 == x2) {
        for (let i = Math.min(y1, y2); i <= Math.max(y1, y2); i++) {
            bo[i][x1]++;
        }
    }
    else if (y1 == y2) {
        for (let i = Math.min(x1, x2); i <= Math.max(x1, x2); i++) {
            bo[y1][i]++;
        }
    }
    else if (star == 2 && Math.abs(x1 - x2) == Math.abs(y1 - y2)) {  // draw diag
        let y0, x0;
        for (let i = 0; i <= Math.abs(x2-x1); i++) {
            if (y1 > y2) y0 = y1-i;
            else y0 = y1+i;
            if (x1 > x2) x0 = x1-i;
            else x0 = x1+i;
            bo[y0][x0]++;
        }
    }
    return bo;
}