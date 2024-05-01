## 开发环境启动

```sh
GOPROXY=https://goproxy.cn,direct
GO111MODULE=on
go install github.com/swaggo/swag/cmd/swag@latest
go mod tidy
go run .
```
