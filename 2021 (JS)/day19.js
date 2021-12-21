require('fs').readFile("day19.txt", "utf-8", function(err, data) {
    let inp = data.trim().split(/--- \w+ \d+ ---/).slice(1);
    inp = inp.map(k => k.trim().split("\n").map(e => e.match(/-?\d+/g).map(a => a*1)));
    // console.table(inp);
    main(inp);
});

const MATCHES_REQ = 6;

class Scanner {
    x;
    y;
    z;
}

function main(scanners) {
    let scans = [];
    for (let i = 0; i < scanners.length; i++) {
        scans.push(new Scanner());
    }
    scans[0].x = 0;
    scans[0].y = 0;
    scans[0].z = 0;
    for (let i = 0; i < scanners.length; i++) {
        for (let j = i+1; j < scanners.length; j++) {
            let m = compare(scanners[i], scanners[j]);
            if (m == undefined) continue;
            scans[j].x = scans[i].x + m.x;
            scans[j].y = scans[i].y + m.y;
            scans[j].z = scans[i].z + m.z;
        }
    }
    console.log("Result");
    console.table(scans);
}

const rotations = [
    {x: 1, y: 1, z: 1},
    {x: 1, y: 1, z: -1},
    {x: 1, y: -1, z: 1},
    {x: 1, y: -1, z: -1},
    {x: -1, y: 1, z: 1},
    {x: -1, y: 1, z: -1},
    {x: -1, y: -1, z: 1},
    {x: -1, y: -1, z: -1},
]

function compare(scanner1, scanner2) {
    let maps;
    for (let i = 0; i < scanner1.length; i++) {
        for (let j = 0; j < scanner2.length; j++) {
            let p1 = scanner1[i];
            let p2 = scanner2[j];
            maps = genMapping(p1, p2);
            console.table(maps);
            for (let k = 0; k < maps.length; k++) {
                let c = 0;
                for (let a = 0; a < scanner1.length; a++) {
                    for (let b = 0; b < scanner2.length; b++) {
                        let res = test(scanner1[a], scanner2[b], maps[k]);
                        if (res) c++;
                    }
                }
                if (c >= 2) console.log(maps[k], c);
                if (c >= MATCHES_REQ) {
                    return maps[k];
                }
            }
        }
    }
    return null;
    
}

// function gen(c1, c2) {
//     let l = [];
//     let a = [];
//     for (let i = 0; i < 3; i++) {
//         for (let j = 0; j < 3; j++) {
//             for (let k = 0; k < 3; k++) {
//                 if (j == i || j == k || i == k) continue;
//                 let st = `${i}-${}-${}`;
//             }
//         }
//     }
//     l.push({x: })
// }

function genMapping(c1, c2) {
    let l = [];
    let p1 = [];
    let p2 = [];
    // all possible directions
    let a1 = [], a2 = []; // added
    for (let i = 0; i < 3; i++) {
        for (let j = 0; j < 3; j++) {
            for (let k = 0; k < 3; k++) {
                if (j == i || j == k || i == k) continue;
                let w1 = [c1[i], c1[j], c1[k]];
                if (!a1.includes(w1.join(","))) {
                    a1.push(w1.join(","));
                    p1.push(w1);
                }
                let w2 = [c2[i], c2[j], c2[k]]
                if (!a2.includes(w2.join(","))) {
                    a2.push(w2.join(","));
                    p2.push(w2);
                }
            }
        }
    }
    for (let i = 0; i < p1.length; i++) {
        for (let j = 0; j < p2.length; j++) {
            for (let k = 0; k < rotations.length; k++) {
                let r = rotations[k];
                l.push({x: p1[i][0]-p2[j][0] * r.x, y: p1[i][1]-p2[j][1] * r.y, z: p1[i][2]-p2[j][2] * r.z });
            }
        }
    }
    return l;
}

function test(c1, c2, m={x:0, y:0, z:0}) {
    // console.log(c1, c2);
    return (c1[0] === c2[0] + m.x
        && c1[1] === c2[1] + m.y
        && c1[2] === c2[2] + m.z);
}