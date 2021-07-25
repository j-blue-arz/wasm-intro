#!/bin/bash
rm -f heron.wasm
rm -f heron.wat
emcc heron.cpp --no-entry -Os -s EXPORTED_FUNCTIONS=[_heron,_heron_int] -o heron.wasm
~/wabt/build/wasm2wat heron.wasm -o heron.wat