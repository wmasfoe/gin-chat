package respModel

const (
	SystemError int = 1001
)

// CodeMsgMapping code码和message映射
var CodeMsgMapping = map[int]string{
	SystemError: "系统错误",
}
