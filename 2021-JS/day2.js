let fs = require('fs');

let d = fs.readFileSync("day2.txt", "utf8");

let [x, y1, y2, aim] = [0, 0, 0, 0];

d.split("\n")
    .map((a) => a.split(" "))
    .map((a) => [a[0], Number(a[1])])
    .forEach((d) => {
        console.log(d);
        let n = d[1];
        switch(d[0]) {
            case "forward":
                x += n;
                y2 += n * aim;
                break;
            case "up":
                y1 -= n
                aim -= n;
                break;
            case "down":
                y1 += n;
                aim += n;
                break;
        }
    });

console.log("Task 1: " + x*y1);
console.log("Task 2: " + x*y2);
