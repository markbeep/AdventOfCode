require('fs').readFile("day14.txt", "utf-8", function(err, data) {
    let arr = data.trim().split("\n");
    let temp = arr[0].trim();
    let ins = arr.slice(2).map(e => e.split(" -> "));
    main(temp, ins);
});

const DAY = 40;

function main(temp, ins) {
    const count = (e) => temp.split(e).length-1;  // count occurences of string in string
    let letters = new Array(...new Set([].concat(...ins).join("").split("")));  // all used letters in a list 
    let com = {};
    letters.forEach(a => letters.forEach(b => com[a+b]=count(a+b)));  // init char combination count
    let letc = letters.reduce((t, e) => ({...t, [e]: count(e)}), {}); // init letter count dict
    for (let i = 0; i < DAY; i++) com = sim(com, ins, letc);  // repeat simulation DAY times
    let values = Object.values(letc);
    console.log(Math.max(...values) - Math.min(...values));
}

function sim(com, ins, letc) {
    let copy = {};
    let keys = Object.keys(com);
    keys.forEach(st => {
        let res = ins.filter(e => e[0] === st);  // is valid transition?s
        if (res.length > 0) {
            let [from, to] = res[0];
            copy[from[0]+to] = (copy[from[0]+to] || 0) + com[from];
            copy[to+from[1]] = (copy[to+from[1]] || 0) + com[from];
            letc[to] += com[from];
        }
    })
    return copy;
}