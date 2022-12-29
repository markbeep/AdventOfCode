{pkgs ? import <nixpkgs> {}, lib ? pkgs.lib}:
with lib;
let
    match = {fst, snd}: foldl (o: c: if elem c snd then c else o) "" fst;
    half = a: let s = stringToCharacters a; in {fst=(take (builtins.div (length s) 2) s); snd=(drop (builtins.div (length s) 2) s);};
in
pipe ./03.txt [
    readFile
    (splitString "\n")
    (filter (s: s != ""))
    (map half)
    (map match)
    (map strings.charToInt)
    (map (s: if s > 90 then s - 96 else s - 38))
    (foldl add 0)
]
