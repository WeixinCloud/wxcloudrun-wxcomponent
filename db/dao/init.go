package dao

// Init 初始化
func Init() error {
	go startClearExpiredRecordTask()
	return nil
}
