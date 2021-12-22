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
}

function covered(t, ind) {
    let c = t[i].map(e=>e);
    for (let i = 0; i < t.length; i++) {
        if (i === ind) continue;
        
    }
}

// subtracts c2 from c1 (both should be the same state "on/off")
function subtraction(c1, c2) {
    let x1 = t[i].map(e=>e);
    let x2 = t[i].map(e=>e);
    let y1 = t[i].map(e=>e);
    let y2 = t[i].map(e=>e);
    let z1 = t[i].map(e=>e);
    let z2 = t[i].map(e=>e);

    let diff = c2[0]-c1[0]  // if diff is pos, c1 is bigger in this direction
    if (diff > 0) {
        x1[0] = c2[0];
        x1[1] = c1[0]-1;
    }
    // x2 should be 13..13 11..13 11..13
    diff = c1[1]-c2[1]  // if diff is pos, c1 is bigger
    if (diff > 0) {
        x2[0] = c1[0]+1;
        x2[1] = c2[0];
    }
    diff = c2[2]-c1[2]  // if diff is pos, c1 is bigger in this direction
    if (diff > 0) {
        y1[2] = c2[2];
        y1[3] = c1[2]-1;
    }
    // x2 should be 13..13 11..13 11..13
    diff = c1[3]-c2[3]  // if diff is pos, c1 is bigger
    if (diff > 0) {
        y2[2] = c1[2]+1;
        y2[3] = c2[2];
    }

    

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