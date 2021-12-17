require('fs').readFile("day17.txt", "utf-8", function(err, data) {
    let inp = data.trim().slice(15).replace("y=","").replaceAll("..", ", ").split(", ").map(e => e*1);
    main(...inp);
});

function main(x1, x2, y1, y2) {
    let h = -Infinity;
    let c = 0;
    for (let y = Math.min(y1, y2); y < Math.max(Math.abs(y1), Math.abs(y2)); y++) {
        for (let x = 1; x < 2*Math.max(Math.abs(x1), Math.abs(x2)); x++) {
            let res = sim(x, y, x1, y1, x2, y2);
            if (res.hit) {
                h = Math.max(h, res.height);
                c++;
            }
        }
    }
    console.log("Max Height:", h);
    console.log("Number of hitters:", c);
}

function sim(xVel, yVel, x1, y1, x2, y2) {
    let curX = 0;
    let curY = 0;
    let h = -Infinity;
    while (true) {
        curX += xVel;
        curY += yVel;
        h = Math.max(h, curY);
        xVel = Math.max(0, xVel-1);
        yVel--;
        if (x1 <= curX && curX <= x2 && y1 <= curY && curY <= y2) {
            return {hit: true, height: h};
        }
        if (curX > Math.max(x1, x2) || curY < Math.min(y1, y2)) break;
    }
    return {hit: false, height: 0};
}