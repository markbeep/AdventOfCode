let fs = require('fs');

function tr(a, del="\n") {
    return a.trim().split(del);
}
function nu(a) {
    return a.map((k) => Number(k));
}

let d = fs.readFileSync("day3.txt", "utf8");
d = tr(d);
d = nu(d);
console.log(d);