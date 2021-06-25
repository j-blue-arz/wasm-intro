#!/bin/bash
rm -rf build/
mkdir build
emcc heron.cpp --no-entry -Os -s EXPORTED_FUNCTIONS=[_heron,_heron_int] -o build/heron.wasm
~/wabt/build/wasm2wat build/heron.wasm -o build/heron.wat
rm serve/*
mv build/heron.wasm serve/heron.wasm
cp heron.html serve/index.html