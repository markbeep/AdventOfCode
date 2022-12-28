let fs = require('fs');

let d = fs.readFileSync("day8.txt", "utf8").split("\n").map(e => e.split("|").map(b => b.trim()));

// TASK 1
// for every line, only look at numbers on the right side of |, then count the amount of words fulfilling the requirement
let t1 = d.reduce((t, line) => t += line[1].split(" ").reduce((t, v) => t += Number([2, 3, 4, 7].includes(v.length)), 0), 0);
console.log(t1);
// TASK 2
let t2 = d.reduce((t, line) => t += 1 * per([], ["a", "b", "c", "d", "e", "f", "g"], line), 0);
console.log(t2);

// Calculates all permutations of and if any accepts, returns the numbers for that (otherwise -1)
function per(un, rem, line) {
    if (rem.length == 0) {
        let m = {a: un[0], b: un[1], c: un[2], d: un[3], e: un[4], f: un[5], g: un[6]};
        let words = line[0].split(" ");
        for (let i = 0; i < words.length; i++) {
            if (check(words[i], m) === -1) return -1;
        }
        let word = line[1].split(" ")
            .reduce((t, v) => t += check(v, m), "");  // turns a list of segment words into a number string
        return word;
    }
    for (let i = 0; i < rem.length; i++) {
        let copy = un.map(e => e);
        copy.push(rem[i]);
        let remCopy = rem.filter((_, j) => j!=i);  // copy array without i-th element
        let ret = per(copy, remCopy, line);
        if(ret != -1) return ret;
    }
    return -1;
}

function check(a, m) {
    if (a.length == 2 && helper(a, [m.c, m.f])) return 1;
    if (a.length == 3 && helper(a, [m.a, m.c, m.f])) return 7;
    if (a.length == 4 && helper(a, [m.b, m.c, m.d, m.f])) return 4;
    if (a.length == 7 && helper(a, [m.a, m.b, m.c, m.d, m.e, m.f, m.g])) return 8;
    if (a.length == 6 && helper(a, [m.a, m.b, m.c, m.e, m.f, m.g])) return 0;
    if (a.length == 5 && helper(a, [m.a, m.c, m.d, m.e, m.g])) return 2;
    if (a.length == 5 && helper(a, [m.a, m.c, m.d, m.f, m.g])) return 3;
    if (a.length == 5 && helper(a, [m.a, m.b, m.d, m.f, m.g])) return 5;
    if (a.length == 6 && helper(a, [m.a, m.b, m.d, m.e, m.f, m.g])) return 6;
    if (a.length == 6 && helper(a, [m.a, m.b, m.c, m.d, m.f, m.g])) return 9;
    return -1;
}

function helper(a, vals) {
    for (let i = 0; i < vals.length; i++){
        if (!a.includes(vals[i])) return false;
    }
    return true;
}