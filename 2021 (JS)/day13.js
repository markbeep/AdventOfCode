let fs = require('fs');

let d = fs.readFileSync("day13.txt", "utf8")
    .trim()
    .split("\n");

let nu = d.filter(e => e[0]!=="f" && e.length>0).map(e => e.split(",").map(e=>e*1));
d = d.filter(e => e[0] === "f");
let folds = new Array(d.length).fill();
for (let i = 0; i < folds.length; i++) {
    let k = d[i].split("=");
    let v = k[0][k[0].length-1];
    folds[i] = [v, k[1]*1];
}

let val = 2000;
let a = new Array(val).fill().map(() => new Array(val).fill(false));

for (let i = 0; i < nu.length; i++) {
    a[nu[i][1]][nu[i][0]] = true;
}

for (let i = 0; i < folds.length; i++) {
    let axis = folds[i][0];
    let num = folds[i][1];
    a = fold(axis, num);
    console.log(count(a));
    draw(a);
}

function count(a) {
    let c = 0;
    for (let y = 0; y < a.length; y++) {
        for (let x = 0; x < a[0].length; x++) {
            if (a[y][x]) c++;
        }
    }
    return c;
}

function draw(a) {
    for (let y = 0; y < a.length; y++) {
        let line = "";
        for (let x = 0; x < a[0].length; x++) {
            if (a[y][x]) line += "#";
            else line += "."
        }
        console.log(line);
    }
}

function fold(axis, num) {
    if (axis === "y") {
        let n = new Array(num).fill().map(() => new Array(a[0].length).fill(false));
        for (let y = 0; y < a.length; y++) {
            for (let x = 0; x < a[0].length; x++) {
                if (a[y][x]) {
                    if (y < num) n[y][x] = true;
                    else {
                        n[2*num - y][x] = true;
                    }
                }
            }
        }
        return n;
    } else {
        let n = new Array(a.length).fill().map(() => new Array(num).fill(false));
        for (let y = 0; y < a.length; y++) {
            for (let x = 0; x < a[0].length; x++) {
                if (a[y][x]) {
                    if (x < num) n[y][x] = true;
                    else {
                        n[y][2*num - x] = true;
                    }
                }
            }
        }
        return n;
    }
}