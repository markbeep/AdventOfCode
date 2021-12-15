const { time, timeEnd } = require('console');

require('fs').readFile("day15.txt", "utf-8", function(err, data) {
    let arr = data.trim().split("\n");
    arr = arr.map(e => e.trim().split("").map(b => b*1));
    time("Task 2");
    main(arr);
    timeEnd("Task 2")
});

let counter = 0;

function main(arr) {
    arr = copy(arr);
    // let dp = arr.map(e => e.map(a => -1));
    // dp[0][0] = Number(arr[0][0]);
    // for (let y = 0; y < arr.length; y++) {
    //     for (let x = 0; x < arr[0].length; x++) {
    //         if (dp[y][x] === -1) {
    //             let left = (y > 0) ? dp[y-1][x]+arr[y][x]: 100000000;
    //             let top = (x > 0) ? dp[y][x-1]+arr[y][x]: 100000000;
    //             dp[y][x] = Math.min(left, top);
    //         }
    //     }
    // }
    // console.table(dp);
    // console.log(dp[dp.length-1][dp[0].length-1] - dp[0][0]);    console.time(i);
    let nodes = createGraph(arr);
    counter = 0;
    dijkstra(nodes[0][0]);
    console.log(nodes[5][0]);
}

function getFromQ(Q) {
    let n;
    let m = 100000000000;
    Q.forEach(node => {
        if (m > node.distance && !node.vis) {
            n = node;
            m = n.distance;
        };
    });
    return n;
}

function dijkstra(S) {
    // let k = [].concat(...nodes);
    let Q = [S];
    S.distance = 0;

    let U;
    while ((U = getFromQ(Q)) != null) { 
        counter++;
        if (counter % 40 === 0) {
            Q = Q.filter(e => !e.vis);
        }
        U.vis = true;
        // console.group([U.x, U.y]);
        U.edges.forEach(V => {
            // console.log([V.x, V.y]);
            if (!V.vis) {
                let tmpDist = U.distance + V.value;
                if (tmpDist < V.distance) {
                    V.distance = tmpDist;
                    Q.push(V);
                }
                // console.table(Q);
        }
        });
        // console.groupEnd([U.x, U.y]);
    }

}

function createGraph(arr) {
    let nodes = arr.map(a => a.map(b => new Node()));
    for (let y = 0; y < nodes.length; y++) {
        for (let x = 0; x < nodes[0].length; x++) {
            if (y > 0) nodes[y][x].edges.push(nodes[y-1][x]);
            if (x > 0) nodes[y][x].edges.push(nodes[y][x-1]);
            if (y < nodes.length-1) nodes[y][x].edges.push(nodes[y+1][x]);
            if (x < nodes[0].length-1) nodes[y][x].edges.push(nodes[y][x+1]);
            nodes[y][x].value = arr[y][x];
            nodes[y][x].x = x;
            nodes[y][x].y = y;
        }
    }
    return nodes;
}

function copy(arr) {
    let n = new Array(5*arr.length).fill().map(_=> new Array(5*arr[0].length).fill(-1));
    for (let y = 0; y < n.length; y++) {
        for (let x = 0; x < n[0].length; x++) {
            let adder = Math.floor(y / arr.length) + Math.floor(x / arr[0].length);
            n[y][x] = arr[y%arr.length][x%arr[0].length] + adder;
            n[y][x] = (n[y][x]-1) % 9 + 1;
        }
    }
    return n;
}

class Node {
    edges = [];
    vis = false;
    value;
    distance = Infinity;
}