GOOS=linux GOARCH=386 go build -o blindpaint-32
echo compiled for linux x86
GOOS=linux GOARCH=amd64 go build -o blindpaint-64
echo compiled for linux x64
GOOS=windows GOARCH=386 go build -o blindpaint-win32.exe
echo compiled for windows x86
GOOS=windows GOARCH=amd64 go build -o blindpaint-win64.exe
echo compiled for windows x64
