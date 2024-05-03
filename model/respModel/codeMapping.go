package respModel

const (
	SystemError int = 1001
)

// CodeMsgMapping TODO @dev 这里记录 code码和message映射
var CodeMsgMapping = map[int]string{
	SystemError: "系统错误",
}
