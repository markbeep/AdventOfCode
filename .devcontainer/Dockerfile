FROM debian:trixie-slim

# Dev tools
RUN apt-get update && apt-get install -y \
    git \
    curl \
    xz-utils \
    hyperfine \
    vim

# Zig
RUN curl -o zig.tar.xz https://ziglang.org/builds/zig-x86_64-linux-0.16.0-dev.1484+d0ba6642b.tar.xz
RUN tar -xf zig.tar.xz && mv zig-* /usr/zig && rm -rf zig.tar.xz zig-*
ENV PATH="/usr/zig:$PATH"
