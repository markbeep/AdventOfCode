{pkgs ? import <nixpkgs> {}, lib ? pkgs.lib}:
with lib;
let
    score = s: if s == "X" then 1 else if s == "Y" then 2 else 3;
    translate = s: if s == "X" then "A" else if s == "Y" then "B" else "C";
    won = a: b: if a == b then 3 else if a == "A" && b == "C" || a == "B" && b == "A" || a == "C" && b == "B" then 6 else 0;
    two = a: b:
        if a == "A" then
            if b == "X" then 3 else
            if b == "Y" then 4 else 8 else
        if a == "B" then
            if b == "X" then 1 else
            if b == "Y" then 5 else 9 else
        if b == "X" then 2 else
        if b == "Y" then 6 else 7;
    x = pipe ./02.txt [
        readFile
        (splitString "\n")
        (filter (s: s != ""))
        (map (splitString " "))
        (map (s: {you=head s; me=last s;}))
        (map (s: {one=score s.me + (won (translate s.me) s.you); two=two s.you s.me;}))
    ];
in "Part 1 ${toString (foldl (a: b: a + b.one) 0 x)} | Part 2 ${toString (foldl (a: b: a + b.two) 0 x)}"
