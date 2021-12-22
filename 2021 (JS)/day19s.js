require("fs").readFile("day19.txt", "utf-8", function(err, data) {
    let inp = data.trim().split(/--- \w+ \d+ ---/).slice(1);
    inp = inp.map(k => k.trim().split("\n").map(e => e.match(/-?\d+/g).map(a => a*1)));
    main(inp);
});

class Scanner {
    x = null;
    y = null;
    z = null;
    constructor(coords, id) {
        this.points = coords;
        this.id = id;
    }
}

const REQ_MATCHES = 6;

function main(inp) {
    // create all scanners
    let scans = [];
    for (let i = 0; i < inp.length; i++) {
        scans.push(new Scanner(inp[i], i));
    }

    scans[0].x = 0;
    scans[0].y = 0;
    scans[0].z = 0;


    let c = 0;
    // console.log(isMap([]))
    for (let i = 0; i < scans[0].points.length; i++) {
        for (let j = 0; j < scans[1].points.length; j++) {
            if (isMap(scans[0].points[i], scans[1].points[j], { x: -20, y: -1133, z: 1061 }, 39)) c++;;
        }
    }
    console.log(c);

    let scannersLocated = 0;
    // go through all scanner pairs
    for (let i = 0; i < scans.length; i++) {
        if (scans[i].x == null) continue;
        for (let j = 0; j < scans.length; j++) {
            if (j === i) continue;
            if (scans[j].x != null) continue;
            let mp = findCorrectMapping(scans[i], scans[j]);
            console.log("Found Mapping:", mp, "| i:",i,"j:",j);
            if (mp != null) {
                let [rotation, offset, matches] = mp;
                // console.log("\tMatches:",matches);
                scans[j].x = scans[i].x + offset.x;
                scans[j].y = scans[i].y + offset.y;
                scans[j].z = scans[i].z + offset.z;
                console.log("Scanner:", j, ":", scans[j].x, scans[j].y, scans[j].z);
            }
        }
    }

}

// finds a fitting mapping
function findCorrectMapping(s1, s2) {
    let maps = findPossibleMappings(s1, s2);
    let m = REQ_MATCHES-1;
    let top = null;
    for (let i = 0; i < maps.length; i++) {
        let [rotation, offset] = maps[i];
        let matches = 0;
        let kek = [];
        for (let j = 0; j < s1.points.length; j++) {
            for (let k = 0; k < s2.points.length; k++) {
                if (isMap(s1.points[j], s2.points[k], offset, rotation)) {
                    matches++;
                    kek.push(move(rotate(rotation, {x:s2.points[k][0], y:s2.points[k][1], z:s2.points[k][2]}), offset));
                }
            }
        }
        if (matches > m) {
            m = matches;
            top = maps[i];
            if (s1.id === 0 && s2.id === 1) {
                console.table(kek);
            }
        }
    }
    return (top==null)?null:[...top, m];
}

// is the mapping and rotation already in our list of res
function isIn(res, mp, rot) {
    for (let i = 0; i < res.length; i++) {
        let [rotation, pos] = res[i];
        if (pos.x === mp.x && pos.y === mp.y && pos.z === mp.z && rotation === rot) return true;
    }
    return false;
}

function move(mp, offset) {
    let m = {x:mp.x, y:mp.y, z:mp.z}
    m.x += offset.x;
    m.y += offset.y;
    m.z += offset.z;
    return m;
}

function findPossibleMappings(s1, s2) {
    let pos = [];
    for (let j = 0; j < s1.points.length; j++) {
        for (let k = 0; k < s2.points.length; k++) {
            let c1 = s1.points[j];
            let c2 = s2.points[k];
            for (let i = 0; i < 48; i++) {
                let tmpMap = {x: c2[0], y: c2[1], z: c2[2]};
                let m = rotate(i, tmpMap);
                let offset = {x: c1[0]-m.x, y: c1[1]-m.y, z: c1[2]-m.z};
                // console.log(c1, m);
                // if (offset.x === 0 && offset.y === 0 && offset.z === 0) console.log(i, offset);
                if (isMap(c1, c2, offset, i) && !isIn(pos, offset, i)) {
                    pos.push([i, offset]);  // potential maps are added to pos
                }
            }
        }
    }
    return pos;
}

function isMap(c1, c2, offset={x:0, y:0, z:0}, rotation) {
    let mp = rotate(rotation, {x:c2[0], y:c2[1], z:c2[2]});
    mp.x += offset.x;
    mp.y += offset.y;
    mp.z += offset.z;
    return (c1[0] === mp.x
        && c1[1] === mp.y
        && c1[2] === mp.z);
}

const rotate = (i,o) => {
    let m = [
        [1,1,1],
        [1,1,-1],
        [1,-1,1],
        [1,-1,-1],
        [-1,1,1],
        [-1,1,-1],
        [-1,-1,1],
        [-1,-1,-1]
    ];
    let p = [
        [0,1,2],
        [0,2,1],
        [1,0,2],
        [1,2,0],
        [2,0,1],
        [2,1,0]
    ];
    let y = Math.floor(i / 6);  // 0-7
    let x = i - 6 * y;  // 0-5
    let qe = [o.x, o.y, o.z];
    return { x: m[y][0]*qe[p[x][0]], y: m[y][1]*qe[p[x][1]], z: m[y][2]*qe[p[x][2]]};
}

// let xexe = {x:-1,y:-2,z:-3};
// for (let i = 0; i < 48; i++) {
//     console.log(rotate(i, xexe));
// }

// 39 => [-1,-1,1] & [1,2,0]=>[y,z,x]