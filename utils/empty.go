package utils

func IsEmpty(value interface{}) bool {
	res := value == nil || value == 0 || value == ""
	return res
}
