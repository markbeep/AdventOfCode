export {}

let fs = require('fs');

let d: Array<string> = fs.readFileSync("day4.txt", "utf8").trim().split("\n");
d = d.map((e) => e.trim());

// all picked elements
let ele: Array<number> = d[0].split(",").map(e => Number(e.trim()));

// creates a 2d array out of the input numbersr
let k = d.filter((a, i) => i > 1 && a.length > 0)
    .map(e => e.split(" ")).map(e => {
        return e.filter(a => a.length > 0)
            .map(e => Number(e));
    });

// create the 3d array of boards
let boards: Array<Array<Array<number>>> = [];
for (let i = 0; i < k.length; i+=5) {
    boards.push(k.slice(i, i+5));
}

// ------------------------ TASK 1 ------------------------
function task1(ele: Array<number>, boards: Array<Array<Array<number>>>) {
    let f: Array<number> = [];
    for (let e = 0; e < ele.length; e++) {
        f.push(ele[e]);
        for (let i = 0; i < boards.length; i++) {
            let b = createB(f, boards[i]);
            if (kek(b)) {
                return su(ele[e], boards[i], b);
            };
        }
    }
}
console.log("Task 1: " + task1(ele, boards));


// ------------------------ TASK 2 ------------------------
let f: Array<number> = [];
for (let e = 0; e < ele.length; e++) {
    // filter out all completed boards
    f.push(ele[e]);
    boards = boards.filter(b => !kek(createB(f, b)));
    if (boards.length === 1) {
        console.log("Task 2: " + task1(ele, boards));
        break;
    }
}

// ------------------------ Helper Functions ------------------------
function createB(elements: Array<number>, array: Array<Array<number>>) {
    let b: Array<Array<boolean>> = new Array(5).fill(new Array(5).fill(false));
    return b.map((o, y) => o.map((e, x) => elements.includes(array[y][x])));
}

function kek(a: Array<Array<boolean>>) {
    // checks if any row or column is all true
    return a.reduce((b, _, i) =>
        b || (a[i][0] && a[i][1] && a[i][2] && a[i][3] && a[i][4]
            || a[0][i] && a[1][i] && a[2][i] && a[3][i] && a[4][i]), false);
}

function su(a: number, b: Array<Array<number>>, bol: Array<Array<boolean>>) {
    // sums up all non-true values and returns the value multiplied by the last value
    let z = 0;
    for (let y = 0; y < 5; y++) {
        for (let x = 0; x < 5; x++) {
            if (!bol[y][x]) z += b[y][x];
        }
    }
    return z * a;
}