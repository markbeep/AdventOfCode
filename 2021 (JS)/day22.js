require("fs").readFile("day22.txt", "utf-8", function(err, data) {
    let lines = data.trim().split("\n");
    let nums = data.trim().split("\n");
    let t = [];
    for (let i = 0; i < lines.length; i++){
        let inp = nums[i].match(/-?\d+/g).map(e => e*1);
        t.push([inp[0], inp[1], inp[2], inp[3], inp[4], inp[5], lines[i].slice(0, 2)]);
    }
    main(t);
});

function main(t) {
    console.log("Task 1:", bruteforce(t));
    console.table(t);
    console.log(intersect(t[0], t[3]));
}

function covered(t, ind) {
    let c = t[i].map(e=>e);
    for (let i = 0; i < t.length; i++) {
        if (i === ind) continue;
        
    }
}

function isIntersection(c1, c2) {
    if (!(c1[0] <= c2[1])) return false;
    if (!(c1[1] >= c2[0])) return false;
    if (!(c1[2] <= c2[3])) return false;
    if (!(c1[3] >= c2[2])) return false;
    if (!(c1[4] <= c2[5])) return false;
    if (!(c1[5] >= c2[4])) return false;
    return true;
}

function intersect(c1, c2) {
    // calc C
    let x1 = Math.max(c1[0], c2[0]);
    let x2 = Math.min(c1[1], c2[1]);
    let y1 = Math.max(c1[2], c2[2]);
    let y2 = Math.min(c1[3], c2[3]);
    let z1 = Math.max(c1[4], c2[4]);
    let z2 = Math.min(c1[5], c2[5]);
    

    

}

function inRange() {

}

function howMany(c) {
    return (c[1]-c[0]+1) * (c[3]-c[2]+1) * (c[5]-c[4]+1);
}



function bruteforce(t) {
    let os = 50; // ofset
    let dp = new Array(102).fill().map(_=>new Array(102).fill().map(_=>new Array(102).fill(false)));
    for (let i = 0; i < t.length; i++) {
        let k = t[i];
        if (k[0] < -50 || k[1] > 50 || k[2] < -50 || k[3] > 50 || k[4] < -50 || k[5] > 50) continue;
        for (let x = k[0]+os; x <= k[1]+os; x++) {
            for (let y = k[2]+os; y <= k[3]+os; y++) {
                for (let z = k[4]+os; z <= k[5]+os; z++) {
                    if (k[6] === "on") dp[x][y][z] = true;
                    else dp[x][y][z] = false;
                }
            }
        }
    }
    let c = 0;
    for (let x = 0; x < 102; x++) {
        for (let y = 0; y < 102; y++) {
            for (let z = 0; z < 102; z++) {
                if (dp[x][y][z]) c++;
            }
        }
    }
    return c;
}