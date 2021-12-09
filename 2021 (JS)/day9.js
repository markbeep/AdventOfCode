const { time } = require('console');
let fs = require('fs');

let d = fs.readFileSync("day9.txt", "utf8")
    .split("\n")
    .map(e => e.split("").map(e => 1*e));

let create2DArray = (f) => Array(d.length).fill().map((() => Array(d[0].length).fill(f)));
let c = create2DArray(0);
let dp = create2DArray([-1, -1]);
let vis = create2DArray(0);
let m = create2DArray(0);

let max = 1;  // always store max so we draw depending on it
let t = 0;
for (let y = 0; y < d.length; y++) {
    for (let x = 0; x < d[0].length; x++) {
        if (isLowest(y, x)) {  // TASK 1
            t += d[y][x]+1;
        };
        let res = dfs(y, x);  // TASK 2
        if (!equal(res)) {
            let val = ++c[res[0]][res[1]];
            max = Math.max(max, val);
            drawTable();
            m[y][x] = res;
        }
    }
}
console.log("Task 1: " + t);

// turns 2D array into 1D, filters out zeroes, sorts and then multiples highest 3 values
let k = [].concat(...c).filter(e => e > 0).sort((a, b) => b-a);
console.log("Task 2: " + k[0]*k[1]*k[2]);


// ----------- HELPER FUNCS -----------
function sleep(ms) {
    const stop = new Date().getTime() + ms;
    while (new Date().getTime() < stop);
}
function drawTable() {
    console.log("\033[0;0H");
    let ENDC = '\033[0m';
    let msg = [];
    for (let y = 0; y < dp.length; y++) {
        let line = "";
        for (let x = 0; x < dp[0].length; x++) {
            if (!equal(dp[y][x])) line += "\033[38;5;" + color(y,x) + "m" + "█" + ENDC;
            else line += '\033[2m' + " " + ENDC;
        }
        msg.push(line);
    }
    console.log(msg.join("\n"));
    console.log("\x1B[?25l");
    sleep(0.3);
}

function color(y, x) {
    let p = Math.floor(14.0*c[dp[y][x][0]][dp[y][x][1]] / max) - d[y][x];
    return 246-p;
}

function dfs(y, x) {
    if (d[y][x] === 9 || vis[y][x] !== 0) return dp[y][x];  // ignore 9's or already visited
    if (isLowest(y, x)) return dp[y][x] = [y, x];  // lowest point? return coordinates
    vis[y][x] = 1;
    let l = [];  // array to store all potential bases
    if (x>0 && d[y][x-1] < d[y][x]) l.push(dfs(y, x-1));
    if (x<d[0].length-1 && d[y][x+1] < d[y][x]) l.push(dfs(y, x+1));
    if (y>0 && d[y-1][x] <= d[y][x]) l.push(dfs(y-1, x));
    if (y<d.length-1 && d[y+1][x] <= d[y][x]) l.push(dfs(y+1, x));
    l = l.filter(e => !equal(e));  // filter out [-1,-1]
    for (let i = 0; i < l.length-1; i++) {  // do all paths lead to rome?
        if (!equal(l[i], l[i+1])) return [-1,-1];
    }
    return dp[y][x] = l[0];
    
}
function isLowest(y, x) {
    if (x>0 && d[y][x-1] <= d[y][x]) return false;
    if (x<d[0].length-1 && d[y][x+1] <= d[y][x]) return false;
    if (y>0 && d[y-1][x] <= d[y][x]) return false;
    if (y<d.length-1 && d[y+1][x] <= d[y][x]) return false;
    return true;
}
function equal(a, b=[-1,-1]) {
    return a[0] === b[0] && a[1] == b[1];
}