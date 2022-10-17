rm -rf build/
mkdir build
GOOS=js GOARCH=wasm go build -o build/convolve.wasm main.go
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" build/wasm_exec.js
cp fish.jpg index.html build/