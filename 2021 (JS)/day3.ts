let fs = require('fs');

let d: Array<string> = fs.readFileSync("day3.txt", "utf8").trim().split("\n");
d = d.map((e) => e.trim());

let gamma = "", gamma1 = "";
let eps = "", eps1 = "";
let ne: Array<string> = d; // array for gamma
let nex: Array<string> = d; // array for eps

for (let i = 0; i < d[0].length; i++) {
    // ------------------------ TASK 1 ------------------------
    // count number of 1's
    let t = d.reduce((tot, a) => Number(tot)+Number(a[i]), 0);
    gamma1 += (t >= d.length / 2) ? "1": "0";
    eps1 += (t >= d.length / 2) ? "0": "1";

    // ------------------------ TASK 2 ------------------------
    let k = ne.reduce((tot, a) => Number(tot)+Number(a[i]), 0);
    let kx = nex.reduce((tot, a) => Number(tot)+Number(a[i]), 0);
    
    let b = (k >= ne.length / 2) ? "1": "0";
    let bx = (kx >= nex.length / 2) ? "0": "1";
    
    // Remove all elements that are invalid
    ne = ne.filter((e) => b == e[i]);
    nex = nex.filter((e) => bx == e[i]);

    if (ne.length == 1) gamma = ne[0];
    if (nex.length == 1) eps = nex[0];
}

console.log(`Task 1: ${parseInt(gamma1, 2) * parseInt(eps1, 2)}`);
console.log(`Task 2: ${parseInt(gamma, 2) * parseInt(eps, 2)}`);