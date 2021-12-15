const { time, timeEnd } = require('console');

require('fs').readFile("day15.txt", "utf-8", function(err, data) {
    let arr = data.trim().split("\n");
    arr = arr.map(e => e.trim().split("").map(b => b*1));
    main(arr);
});

let counter = 0;

function main(arr) {
    time("Copy");
    arr = copy(arr);
    timeEnd("Copy");

    time("Create Node Graph");
    let nodes = createGraph(arr);
    timeEnd("Create Node Graph");
    
    counter = 0;
    dijkstra(nodes[0][0]);
    console.log("Task 1:", nodes[99][99].distance);
    console.log("Task 2:", nodes[nodes.length-1][nodes[0].length-1].distance);
}

function dijkstra(S) {
    S.distance = 0;
    let pq = new PriorityQueue();
    pq.enqueue(S, 0);
    while (!pq.isEmpty()) { 
        let U = pq.dequeue();
        U.vis = true;
        U.edges.forEach(V => {
            if (!V.vis) {
                let tmpDist = U.distance + V.value;
                if (tmpDist < V.distance) {
                    V.distance = tmpDist;
                    pq.enqueue(V, V.distance);
                }
        }
        });
    }

}

function createGraph(arr) {
    console.group("create Graph");
    time("node")
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
    console.groupEnd("create Graph");
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

class PriorityQueue {
    constructor() {
        this.items = []
    }
    enqueue(item, priority) {
        if (this.isEmpty()) {
            this.items.push({ item:item, priority:priority });
            return;
        }
        let index;
        for (let i = 0; i < this.items.length; i++) {
            if (this.items[i].priority <= priority) {
                index = i;
                break;
            }
            index = i+1;
        }
        this.items.splice(index, 0, {  item:item, priority:priority });
    }

    dequeue() {
        return this.items.pop().item;
    }

    isEmpty() {
        return this.items.length === 0;
    }
}