package utils

func MsgConstructor(cmd string, param string) []byte {
	msgString := `{"cmd": "` + cmd + `","param": "` + param + `"}`

	return []byte(msgString)
}
