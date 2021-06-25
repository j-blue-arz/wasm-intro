#!/bin/bash
rm -rf build/
mkdir build
emcc hello.cpp -o build/hello.html
~/wabt/build/wasm2wat build/hello.wasm -o build/hello.wat