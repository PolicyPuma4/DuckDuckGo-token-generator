SET GOOS=windows
SET GOARCH=amd64
go build -o ddgtokengenerator_windows_amd64.exe main.go

SET GOOS=windows
SET GOARCH=arm
go build -o ddgtokengenerator_windows_arm.exe main.go

SET GOOS=darwin
SET GOARCH=amd64
go build -o ddgtokengenerator_darwin_amd64 main.go

SET GOOS=darwin
SET GOARCH=arm64
go build -o ddgtokengenerator_darwin_arm64 main.go

SET GOOS=linux
SET GOARCH=amd64
go build -o ddgtokengenerator_linux_amd64 main.go

SET GOOS=linux
SET GOARCH=arm
go build -o ddgtokengenerator_linux_arm main.go
