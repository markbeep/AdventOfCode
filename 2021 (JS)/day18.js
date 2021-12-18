require('fs').readFile("day18.txt", "utf-8", function(err, data) {
    let inp = data.trim().split("\n");
    main(inp);
});

function main(inp) {
    let tabs = [];
    for (let i = 0; i < inp.length; i++) {
        tabs.push(JSON.parse(inp[i]));
    }
    let tab = new Tab(tabs[0]);
    for (let i = 1; i < tabs.length; i++) {
        tab = add(tab, new Tab(tabs[i]));
        while (findExplode(tab) || findSplit(tab));
    }
    console.log("Task 1:", mag(tab));

    let l = -Infinity;
    for (let i = 0; i < inp.length; i++) {
        for (let j = 0; j < inp.length; j++) {
            if (i === j) continue;
            let t = add(new Tab(tabs[i]), new Tab(tabs[j]));
            while (findExplode(t) || findSplit(t));
            l = Math.max(mag(t), l);
        }
    }
    console.log("Task 2:", l);

}

function mag(tab) {
    if (tab instanceof Node) return tab.value;
    let m = 3 * mag(tab.left);
    m += 2 * mag(tab.right);
    return m;
}

function add(x, y) {
    let t = new Tab();
    t.left = x;
    t.right = y;
    x.parent = t;
    y.parent = t;
    x.incLevel();
    y.incLevel();
    return t;
}

function findRight(tab, startTab, strict=false) {
    if (tab == null) return null;
    // if tab.left is where we're coming from
    if (tab.right == startTab) return findRight(tab.parent, tab, strict);
    if (tab.right instanceof Node) return tab.right;  // left is a node, return that
    let res;
    if (strict) res = findRight(tab.right, tab, true);
    else res = findLeft(tab.right, tab, true);
    if (res != null) return res;
    return null;  // nothing found

}

function findLeft(tab, startTab, strict=false) {
    if (tab == null) return null;
    // if tab.left is where we're coming from
    if (tab.left == startTab) return findLeft(tab.parent, tab, strict);
    if (tab.left instanceof Node) return tab.left;  // left is a node, return that
    let res;
    if (strict) res = findLeft(tab.left, tab, true);
    else res = findRight(tab.left, tab, true);
    if (res != null) return res;
    return null;  // nothing found
}

function findSplit(tab) {
    if (tab.left instanceof Tab && findSplit(tab.left)) return true;
    else if (tab.left.value >= 10 && split(tab.left.value, "l", tab)) return true;
    if (tab.right instanceof Tab && findSplit(tab.right)) return true;
    else if (tab.right.value >= 10 && split(tab.right.value, "r", tab)) return true;
    return false;
}

function split(val, dir, tab) {
    let l = Math.floor(val/2);
    let r = Math.ceil(val/2);
    let ln = new Node(l);
    let rn = new Node(r);
    let n = new Tab();
    n.parent = tab;
    n.level = tab.level+1;
    n.left = ln;
    n.right = rn;
    if (dir === "l") tab.left = n;
    else tab.right = n;
    return true;
    
}

function findExplode(tab, dir="") {
    if (tab.left instanceof Tab && findExplode(tab.left, "l")) return true;
    if (tab.right instanceof Tab && findExplode(tab.right, "r")) return true;
    if (tab.level === 4) {  // exploding happening here
        let left = findLeft(tab.parent, tab);
        let right = findRight(tab.parent, tab);
        if (left != null) left.value += tab.left.value;
        if (right != null) right.value += tab.right.value;
        if (dir === "l") tab.parent.left = new Node(0);
        else tab.parent.right = new Node(0);
        return true;
    }
    return false;
}

class Tab {
    constructor(table=[0,0], level=0, parent=null) {
        this.parent = parent;
        this.left = table[0];
        this.right = table[1];
        this.level = level;
        if (Array.isArray(this.left)) this.left = new Tab(table[0], level+1, this);
        else this.left = new Node(table[0]);
        if (Array.isArray(this.right)) this.right = new Tab(table[1], level+1, this);
        else this.right = new Node(table[1]);
    }
    incLevel() {
        this.level++;
        if (this.left instanceof Tab) this.left.incLevel();
        if (this.right instanceof Tab) this.right.incLevel();
    }
    toArray() {
        let k = [0,0];
        if (this.left instanceof Node) k[0] = this.left.value;
        else k[0] = this.left.toArray();
        if (this.right instanceof Node) k[1] = this.right.value;
        else k[1] = this.right.toArray();
        return k;
    }
    toString() {
        return JSON.stringify(this.toArray());
    }
}

class Node {
    constructor(val) {
        this.value = val;
    }
}