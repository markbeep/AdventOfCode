class Node {
    edges;
    name;
    constructor(name) {
        this.edges = [];
        this.name = name;
    }
    addEdge(toNode) {
        this.edges.push(toNode);
    }
}

let fs = require('fs');

let d = fs.readFileSync("day12.txt", "utf8")
    .trim()
    .split("\n")
    .map(e => e.trim().split("-"));

let nodes = [], st = [];
for (let i = 0; i < d.length; i++) {
    let a = d[i][0], b = d[i][1];
    let an = getNode(a);
    let bn = getNode(b);
    an.addEdge(bn);
    bn.addEdge(an);
}
let start = nodes.filter(e => e.name==="start")[0];

// Task 1
let c = 0;
dfs(start, [], false, 1);
console.log("Task 1:", c);

// Task 2
c = 0;
dfs(start, [], false, 2);
console.log("Task 2:", c);

function dfs(node, vis, visitedSmall, taskValue) {
    let v = count(node.name, vis);
    if (node.name === "start" && count("start", vis) >= 1 || isLower(node.name) && (!visitedSmall && v >= taskValue  || visitedSmall && v >= 1)) return;
    vis.push(node.name);
    if (node.name === "end") {
        c++;
        return;
    }
    for (let i = 0; i < node.edges.length; i++) {
        let goto = node.edges[i];
        let copy = vis.map(e => e);
        dfs(goto, copy, visitedSmall || (isLower(node.name) && v >= 1), taskValue);
    }
}

function count(a, vis) {
    return vis.filter(e => e === a).length;
}

function isLower(a) {
    return a.toLowerCase() === a;
}

function getNode(a) {
    if (st.includes(a)) {  // find the node
        return nodes.filter(e => e.name === a)[0];
    };
    st.push(a); // create new node
    let n = new Node(a);
    nodes.push(n);
    return n;
}