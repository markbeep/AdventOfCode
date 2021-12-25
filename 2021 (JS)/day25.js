require("fs").readFile("day25.txt", "utf-8", function(err, data) {
    let lines = data.trim().replaceAll("\r", "").split("\n");
    let a = new Array(lines.length).fill().map(_=>new Array(lines[0].length));
    for (let i = 0; i < a.length; i++) {
        for (let j = 0; j < a[0].length; j++) {
            a[i][j] = lines[i][j];
        }
    }
    main(a);
});

function main(a) {
    let steps = 0, b;
    do {
        b = a;
        steps++;
        a = doStep(a);
    } while(isDif(a, b));
    console.log("Task 1:", steps);
}

function isDif(a, b) {
    for (let i = 0; i < a.length; i++) {
        for (let j = 0; j < a[0].length; j++) {
            if (a[i][j] !== b[i][j]) return true;
        }
    }
    return false;

}

function doStep(a) {
    let n = JSON.parse(JSON.stringify(a));
    // first east
    for (let i = 0; i < a.length; i++) {
        for (let j = 0; j < a[0].length; j++) {
            if (a[i][j] === "."|| n[i][j] === "v") continue;
            let c = (j+1)%a[0].length;
            if (a[i][c] === ".") {
                n[i][j] = ".";
                n[i][c] = ">";
            }
        }
    }
    let p = JSON.parse(JSON.stringify(n));
    for (let i = 0; i < a.length; i++) {
        for (let j = 0; j < a[0].length; j++) {
            if (n[i][j] === "." || n[i][j] === ">") continue;
            let c = (i+1)%a.length;
            if (n[c][j] === ".") {
                p[i][j] = ".";
                p[c][j] = "v";
            }
        }
    }
    return p;
}