
build-wasm:
	GOOS=js GOARCH=wasm go build -v -o public/test.wasm wasm/wasm.go

clean-wasm:
	rm -frv public/test.wasm
