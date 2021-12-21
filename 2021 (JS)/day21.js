require('fs').readFile("day21.txt", "utf-8", function(err, data) {
    let inp = data.match(/(?<=: )\d+/g).map(e => e*1);
    main(inp[0], inp[1]);
});

function main(p1, p2) {
    t1(p1, p2);
    
    rec(0, 0, p1, p2, 0, 0);
    console.log(dp[0][0][0][0]);

}

let vis = [new Array(100).fill().map(_=> new Array(22).fill().map(_ => new Array(22).fill(false))), new Array(100).fill().map(_=> new Array(22).fill().map(_ => new Array(22).fill(false)))];
let dp = [new Array(100).fill().map(_=>new Array(22).fill().map(_ => new Array(22).fill([0,0]))), new Array(100).fill().map(_=>new Array(22).fill().map(_ => new Array(22).fill([0,0])))];

function rec(a, b, p1, p2, t, player) {
    if (a >= 21 || b >= 21) {
        if (a > b) return [1, 0];
        return [0, 1];
    }
    if (vis[player][t][a][b]) return dp[player][t][a][b];
    let kek = [0,0];
    for (let i = 1; i <= 3; i++) {
        for (let j = 1; j <= 3; j++) {
            for (let k = 1; k <= 3; k++) {
                let n1 = (p1+i+j+k-1)%10+1;
                let n2 = (p2+i+j+k-1)%10+1;
                let res;
                if (player === 0) res = rec(a+n1, b, n1, p2, t+1, 1);
                if (player === 1) res = rec(a, b+n2, p1, n2, t+1, 0);
                kek[0] += res[0];
                kek[1] += res[1];
            }
        }
    }
    dp[player][t][a][b] = kek;
    vis[player][t][a][b] = true;
    return kek;
}

function t1(p1, p2) {
    let a = 0, b = 0, d = 1;
    while (a < 1000 && b < 1000) {
        p1 += d++ + d++ + d++;
        p1 = (p1-1)%10 + 1;
        a += p1;
        if (a >= 1000) break;
        p2 += d++ + d++ + d++;
        p2 = (p2-1)%10+1;
        b += p2;
    }
    console.log("Task 1:", Math.min(a, b) * (d-1));
}