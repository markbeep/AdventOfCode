require('fs').readFile("day14.txt", "utf-8", function(err, data) {
    let arr = data.trim().split("\n");
    let temp = arr[0].trim();
    let ins = arr.slice(2).map(e => e.split(" -> "));
    main(temp, ins);
});

function main(temp, ins) {
    let letters = new Array(...new Set([].concat(...ins).join("").split("")));  // all used letters in a list 
    let com = {};
    letters.forEach(a => letters.forEach(b => com[a+b]=0));
    for (let i = 0; i < temp.length-1; i++) {
        com[temp.slice(i, i+2)]++;
    }
    let letc = {};
    for (let i = 0; i < letters.length; i++) {
        letc[letters[i]] = 0;
    }
    for (let i = 0; i < temp.length; i++) {
        letc[temp[i]]++;
    }
    for (let i = 0; i < 40; i++) {
        com = sim(com, ins, letc);
    }
    let values = Object.values(letc);
    console.log(Math.max(...values) - Math.min(...values));
}

function sim(com, ins, letc) {
    let copy = {};
    let keys = Object.keys(com);
    for (let i = 0; i < keys.length; i++) {
        let st = keys[i];
        let res = ins.filter(e => e[0] === st);
        if (res.length > 0) {
            let [from, to] = res[0];
            copy[from[0]+to] = (copy[from[0]+to] || 0) + com[from];
            copy[to+from[1]] = (copy[to+from[1]] || 0) + com[from];
            letc[to] += com[from];
        }
    }
    return copy;
}