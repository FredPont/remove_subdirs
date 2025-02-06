echo "Start compilation for windows 64bits"

GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -o rmsd-x86_64_windows.exe .


echo "Start compilation for Mac 64bits"
GOOS=darwin GOARCH=amd64 go build -o rmsd-x86_64_Mac.bin .
GOOS=darwin GOARCH=arm64 go build -o rmsd-arm_64_Mac.bin .

echo "Start compilation for linux 64bits"
go build -o rmsd-x86_64_linux.bin .