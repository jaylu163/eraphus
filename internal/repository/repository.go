package repository

// InitRepository 初始化所有repository (如数据库连接) redis mongodb等
func InitRepository() {

	InitMysql()

	//InitRedis()
}
