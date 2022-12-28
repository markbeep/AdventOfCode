let fs = require('fs');

function tr(a, del="\n") {
    return a.trim().split(del);
}
function nu(a) {
    return a.map((k) => Number(k));
}

let d = fs.readFileSync("day3.txt", "utf8");
d = tr(d);
console.log(d);

let gamma = "";
let eps = "";
let ne = d;
let tmp;

for (let i = 0; i < d[0].length; i++) {
    let k = ne.reduce((tot, a) => Number(tot)+Number(a[i]), 0);
    let b;
    if (k >= ne.length / 2) {
        b = "1";
    } else {
        b = "0";
    }
    if (ne.length <= 1) break;
    tmp = [];
    ne.forEach(element => {
        if (b == element[i]) tmp.push(element);
    });
    ne = tmp;
}
gamma = ne[0];

ne = d;
tmp = [];

for (let i = 0; i < d[0].length; i++) {
    let k = ne.reduce((tot, a) => Number(tot)+Number(a[i]), 0);
    let b;
    if (k >= ne.length / 2) {
        b = 0;
    } else {
        b = 1;
    }
    if (ne.length <= 1) break;
    tmp = [];
    ne.forEach(element => {
        if (b == element[i]) tmp.push(element);
    });
    ne = tmp;
}
eps = ne[0];

console.log(parseInt(gamma, 2) * parseInt(eps, 2));