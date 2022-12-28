let fs = require('fs');
let asciimo = require('asciimo').Figlet;

let text = "skius chill";
let font = "Banner";

asciimo.write(text, font, function(art) {
    word = art.trim();
    handleWord(word);
});

function handleWord(word) {
    let folds = [];
    let a = toArray(word);

    for (let i = 0; i < 5; i++) {
        a = unfold(a, folds);
        // draw(a);
    }

    let coords = []
    for (let y = 0; y < a.length; y++) {
        for (let x = 0; x < a[0].length; x++) {
            if (a[y][x]) coords.push(`${x},${y}`);
        }
    }
    msg = coords.join("\n") + "\n\n";
    folds.reverse();
    msg += folds.join("\n");

    fs.writeFile("day13.txt", msg, (err, cont) => {
        console.log("Finished");
        console.log(word);
        console.log(`Size ${coords.length}`);
    });
}


function unfold(a, folds) {
    if (Math.random() > 0.5) {  // unfold y
        let n = new Array(2*a.length+1).fill().map(() => new Array(a[0].length).fill(false));
        let num = a.length;
        folds.push(`fold along y=${num}`);
        for (let y = 0; y < a.length; y++) {
            for (let x = 0; x < a[0].length; x++) {
                let r = Math.random();
                if (r < 0.4) n[y][x] = a[y][x];
                else if (r < 0.7) n[2*num-y][x] = a[y][x];
                else {
                    n[y][x] = a[y][x];
                    n[2*num-y][x] = a[y][x];
                }
            }
        }
        return n;
    }
    else {
        let n = new Array(a.length).fill().map(() => new Array(2*a[0].length+1).fill(false));
        let num = a[0].length;
        folds.push(`fold along x=${num}`);
        for (let y = 0; y < a.length; y++) {
            for (let x = 0; x < a[0].length; x++) {
                let r = Math.random();
                if (r < 0.4) n[y][x] = a[y][x];
                else if (r < 0.7) n[y][2*num-x] = a[y][x];
                else {
                    n[y][x] = a[y][x];
                    n[y][2*num-x] = a[y][x];
                }
            }
        }
        return n;
    }
}

function toArray(word) {
    let lines = word.split("\n");
    let a = new Array(lines.length).fill().map(() => new Array(lines[0].length).fill(false));
    for (let y = 0; y < a.length; y++) {
        for (let x = 0; x < a[0].length; x++) {
            a[y][x] = lines[y][x] === "#";
        }
    }
    return a;
}

function draw(a) {
    for (let y = 0; y < a.length; y++) {
        let line = "";
        for (let x = 0; x < a[0].length; x++) {
            if (a[y][x]) line += "#";
            else line += "."
        }
        console.log(line);
    }
}
