package utils

func MsgConstructor(cmd string) []byte {
	msgString := `{
        "cmd": "` + cmd + `"
    }`

	return []byte(msgString)
}
