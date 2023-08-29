package configs

func Init() error {
	// 初始化数据库
	if err := DbInit(); err != nil {
		return err
	}

	if err := RedisInit(); err != nil {
		return err
	}

	// 初始化Fast DFS分布式存储
	FastDFSInit()

	return nil
}
