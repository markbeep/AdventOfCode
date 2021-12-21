require('fs').readFile("day19.txt", "utf-8", function(err, data) {
    let inp = data.trim().split(/--- \w+ \d+ ---/).slice(1);
    inp = inp.map(k => k.trim().split("\n").map(e => e.match(/-?\d+/g).map(a => a*1)));
    // console.table(inp);
    main(inp);
});

const MATCHES_REQ = 3;

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
    let finalMap;
    outer: for (let i = 0; i < scanner1.length; i++) {
        for (let j = 0; j < scanner2.length; j++) {
            let p1 = scanner1[i];
            let p2 = scanner2[j];
            maps = genMapping(p1, p2);
            // console.table(maps);
            for (let k = 0; k < maps.length; k++) {
                for (let r = 0; r < rotations.length; r++) {
                    let c = 0;
                    for (let a = 0; a < scanner1.length; a++) {
                        for (let b = 0; b < scanner2.length; b++) {
                            let res = test(scanner1[a], scanner2[b], maps[k], rotations[r]);
                            if (res) c++;
                        }
                    }
                    if (c >= MATCHES_REQ) {
                        finalMap = maps[k];
                        break outer;
                    }
                }
            }
        }
    }
    return finalMap;
    
}

// returns if the coord is too far from the scanner, so we can ignore it
function isFar(scanner, coord, mapping) {

}

function genMapping(c1, c2) {
    let l = [];
    let p1 = [];
    let p2 = [];
    // all possible directions
    for (let i = 0; i < 3; i++) {
        for (let j = 0; j < 3; j++) {
            for (let k = 0; k < 3; k++) {
                if (j == i || j == k || i == k) continue;
                p1.push([c1[i], c1[j], c1[k]]);
                p2.push([c2[i], c2[j], c2[k]]);
            }
        }
    }
    for (let i = 0; i < p1.length; i++) {
        for (let j = 0; j < p2.length; j++) {
            // 0 0 0
            l.push({x: p1[i][0]-p2[j][0], y: p1[i][1]-p2[j][1], z: p1[i][2]-p2[j][2]});
            l.push({x: p1[i][0]-p2[j][0], y: p1[i][1]-p2[j][1], z: p1[i][2]+p2[j][2]});
            l.push({x: p1[i][0]-p2[j][0], y: p1[i][1]+p2[j][1], z: p1[i][2]-p2[j][2]});
            l.push({x: p1[i][0]-p2[j][0], y: p1[i][1]+p2[j][1], z: p1[i][2]+p2[j][2]});
            l.push({x: p1[i][0]+p2[j][0], y: p1[i][1]-p2[j][1], z: p1[i][2]-p2[j][2]});
            l.push({x: p1[i][0]+p2[j][0], y: p1[i][1]-p2[j][1], z: p1[i][2]+p2[j][2]});
            l.push({x: p1[i][0]+p2[j][0], y: p1[i][1]+p2[j][1], z: p1[i][2]-p2[j][2]});
            l.push({x: p1[i][0]+p2[j][0], y: p1[i][1]+p2[j][1], z: p1[i][2]+p2[j][2]});
            // 0 0 1
            // l.push({x: p1[i][0]-p2[j][0], y: p1[i][1]-p2[j][1], z: p2[j][2]-p1[i][2]});
            // // 0 1 0
            // l.push({x: p1[i][0]-p2[j][0], y: p2[j][1]-p1[i][1], z: p1[i][2]-p2[j][2]});
            // // 0 1 1
            // l.push({x: p1[i][0]-p2[j][0], y: p2[j][1]-p1[i][1], z: p2[j][2]-p1[i][2]});
            // // 1 0 0s
            // l.push({x: p2[j][0]-p1[i][0], y: p1[i][1]-p2[j][1], z: p1[i][2]-p2[j][2]});
            // // 1 0 1
            // l.push({x: p2[j][0]-p1[i][0], y: p1[i][1]-p2[j][1], z: p2[j][2]-p1[i][2]});
            // // 1 1 0
            // l.push({x: p2[j][0]-p1[i][0], y: p2[j][1]-p1[i][1], z: p1[i][2]-p2[j][2]});
            // // 1 1 1
            // l.push({x: p2[j][0]-p1[i][0], y: p2[j][1]-p1[i][1], z: p2[j][2]-p1[i][2]});
        }
    }
    
    return l;
}

// bruteforce all combinations
function test(c1, c2, m={x:0, y:0, z:0}, rot={x:1, y:1, z:1}) {
    // console.log(c1, c2);
    return (c1[0] === rot.x * c2[0] + m.x
        && c1[1] === rot.y * c2[1] + m.y
        && c1[2] === rot.z * c2[2] + m.z);
}