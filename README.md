## 开发环境启动

开发环境使用了 [air](https://github.com/cosmtrek/air) 支持热更新。

```sh
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
GOPROXY=https://goproxy.cn,direct
GO111MODULE=on
go install github.com/swaggo/swag/cmd/swag@latest
go mod tidy
go run .
```
