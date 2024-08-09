SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build stable-diffusion-server.exe main.go
