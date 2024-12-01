cargo build --release && hyperfine -m 10000 -N --warmup 10 "target/release/aoc $1"
