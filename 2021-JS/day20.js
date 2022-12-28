require('fs').readFile("day20.txt", "utf-8", function(err, data) {
    let inp = data.trim().split("\n");
    let line = inp[0].trim();
    let pic = inp.slice(2).map(e => e.split(""));
    main(line, pic);
});

function get(x, y, pic, border) {
    let s = "0";
    for (let yi = y-1; yi <= y+1; yi++) {
        for (let xi = x-1; xi <= x+1; xi++) {
            if (yi < 0 || yi >= pic.length || xi < 0 || xi >= pic[0].length)
                s += (border === "#") ? "1": "0";
            else
                s += (pic[yi][xi] === "#") ? "1": "0";
        }
    }
    return parseInt(s, 2);
}

function main(enc, pic) {
    let bordSymb = enc[0];
    for (let i = 0; i < 50; i++) {
        let border = (i % 2 == 0) ? ".": bordSymb;
        pic = iter(enc, pic, border);
        if (i == 1) console.log("Task 1:", count(pic));
    }
    console.log("Task 2:", count(pic));
}

function iter(enc, pic, border) {
    let n = 2;
    let c = new Array(pic.length+n).fill().map(_=> new Array(pic[0].length+n).fill(0));
    for (let y = 0; y < c.length; y++) {
        for (let x = 0; x < c[0].length; x++) {
            let index = get(x-n/2, y-n/2, pic, border);
            c[y][x] = enc[index];
        }
    }
    return c;
}

function count(pic) {
    return pic.reduce((t, line) => t+=line.reduce((ta, e) => ta+=Number(e==="#"), 0), 0);
}