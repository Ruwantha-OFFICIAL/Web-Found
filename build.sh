#!/bin/bash

APP_NAME="web-found"
VERSION="1.0.0"
COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
LDFLAGS="-X main.version=$VERSION -X main.commit=$COMMIT -s -w"

rm -rf dist release
mkdir -p dist release
cp webco.txt /release

TARGETS=("linux/amd64" "linux/arm64" "linux/386" "windows/amd64" "windows/arm64" "windows/386")

for TARGET in "${TARGETS[@]}"; do
    GOOS=${TARGET%/*}
    GOARCH=${TARGET#*/}
    EXT=$([ "$GOOS" == "windows" ] && echo ".exe" || echo "")
    FILE="${APP_NAME}_${GOOS}_${GOARCH}"
    
    echo "Building $GOOS/$GOARCH..."
    
    GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "$LDFLAGS" -o "dist/$FILE$EXT" ./cmd/

    if [ "$GOOS" == "windows" ]; then
        zip -j "release/$FILE.zip" "dist/$FILE$EXT"
    else
        tar -czvf "release/$FILE.tar.gz" -C dist "$FILE"
    fi
done

echo "Build complete! Files are in the 'release' directory."
