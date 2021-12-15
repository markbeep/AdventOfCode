const { time, timeEnd } = require("console");
let PriorityQueue = require("js-priority-queue");

require('fs').readFile("day15.txt", "utf-8", function(err, data) {
    let arr = data.trim().split("\n");
    arr = arr.map(e => e.trim().split("").map(b => b*1));
    time("Time");
    main(arr);
    timeEnd("Time");
});

function main(arr) {
    arr = copy(arr);

    let nodes = createGraph(arr);
    
    counter = 0;
    dijkstra(nodes[0][0]);
    console.log("Task 1:", nodes[99][99].distance);
    console.log("Task 2:", nodes[nodes.length-1][nodes[0].length-1].distance);
}

function dijkstra(S) {
    S.distance = 0;
    let pq = new PriorityQueue({ strategy: PriorityQueue.BinaryHeapStrategy, comparator:(a,b) => a.distance-b.distance });
    pq.queue(S, 0);
    while (pq.length > 0) { 
        let U = pq.dequeue();
        U.vis = true;
        U.edges.forEach(V => {
            if (V != undefined && !V.vis) {
                let tmpDist = U.distance + V.value;
                if (tmpDist < V.distance) {
                    V.distance = tmpDist;
                    pq.queue(V, V.distance);
                }
        }
        });
    }

}

function createGraph(arr) {
    let nodes = arr.map(a => (a.map(b => ({ edges: [], vis: false, value: b, distance: Infinity}))));
    for (let y = 0; y < nodes.length; y++) {
        for (let x = 0; x < nodes[0].length; x++) {
            if (y > 0) nodes[y][x].edges.push(nodes[y-1][x]);
            if (x > 0) nodes[y][x].edges.push(nodes[y][x-1]);
            if (y < nodes.length-1) nodes[y][x].edges.push(nodes[y+1][x]);
            if (x < nodes[0].length-1) nodes[y][x].edges.push(nodes[y][x+1]);
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