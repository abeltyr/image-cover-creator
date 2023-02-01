build:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GOARM=7 CC="x86_64-linux-musl-gcc" CXX="x86_64-linux-musl-g++" CGO_LDFLAGS="-static" go build -a -ldflags "-s -w" -o ./dist/main .
build1:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GOARM=7 CC=arm-linux-gnueabihf-gcc CXX=arm-linux-musl-g++ CGO_LDFLAGS="-static" go build -a -ldflags "-s -w" -o ./dist/main .
