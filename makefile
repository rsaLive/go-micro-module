.PHONY: all build run gotool help

BINARY="mealService"

all: gotool build

build:
	go build -ldflags "-s -w" -o ./build/${BINARY} ./cmd/app.go

run:
	@go run ./cmd/app.go
gotool:
	go fmt ./
	go vet ./
pushTest:
	DockerBuildTest.bat
pushProd:
	DockerBuildProd.bat
sonarqube:
	sonar-scanner.bat -D"sonar.projectKey=meal-server" -D"sonar.sources=." -D"sonar.host.url=http://localhost:9000" -D"sonar.login=sqp_531a78dcb522b905b9dca51270e0f1980266cf77"

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"
	@echo "make protoc - 解析proto文件"
