package message

type SnowMsgRWer struct {
}

func (smrw *SnowMsgRWer) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (smrw *SnowMsgRWer) Write(p []byte) (n int, err error) {
	return 0, nil
}
