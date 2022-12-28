{pkgs ? import <nixpkgs> {}, lib ? pkgs.lib}:
with lib;
let x = pipe ./01.txt [
    readFile
    (splitString "\n\n")
    (map (s: (filter (s: s != "")) (splitString "\n" s)))
    (map (map toInt))
    (map (foldl add 0))
    (sort (a: b: b < a))
    (lists.sublist 0 3)
]; in "Part 1: ${toString (lists.head x)} | Part 2:${toString (foldl add 1 x)}"
