require('fs').readFile("day21.txt", "utf-8", function(err, data) {
    let inp = data.match(/(?<=: )\d+/g).map(e => e*1);
    // console.log(inp);
    main(inp[0], inp[1]);
});

function main(p1, p2) {
    let d = 1;
    let a = 0, b = 0;
    // while (a < 1000 && b < 1000) {
    //     poi1 += d++ + d++ + d++;
    //     poi1 = (poi1-1)%10 + 1;
    //     a += poi1;
    //     if (a >= 1000) break;
    //     poi2 += d++ + d++ + d++;
    //     poi2 = (poi2-1)%10+1;
    //     b += poi2;
    // }
    // console.log(a, b, d);
    // console.log(Math.min(a, b) * (d-1));
    rec(0, 0, p1, p2, 0);
    console.log(dp[0][0][0]);

}

let vis = [new Array(22).fill().map(_ => new Array(22).fill(false)), new Array(22).fill().map(_ => new Array(22).fill(false))];
let dp = [new Array(22).fill().map(_ => new Array(22).fill([0,0])), new Array(22).fill().map(_ => new Array(22).fill([0,0]))];


function rec(a, b, p1, p2, t) {
    if (a >= 21 || b >= 21) {
        if (a > b) return [1, 0];
        else if (a < b) return [0, 1];
        return [0, 0];
    }
    if (vis[t][a][b]) return dp[t][a][b];
    let kek = [0,0];
    for (let i = 1; i <= 3; i++) {
        for (let j = 1; j <= 3; j++) {
            for (let k = 1; k <= 3; k++) {
                let n_p1 = (p1+i+j+k-1)%10+1;
                let n_p2 = (p2+i+j+k-1)%10+1;
                let res = 0;
                if (t === 0) res = rec(a+n_p1, b, n_p1, p2, 1);
                if (t === 1) res = rec(a, b+n_p2, p1, n_p2, 0);
                kek[0] += res[0];
                kek[1] += res[1];
            }
        }
    }
    dp[t][a][b] = kek;
    vis[t][a][b] = true;
    return kek;
}