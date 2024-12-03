cargo build --release && hyperfine -N --warmup 10 "target/release/aoc $1"
