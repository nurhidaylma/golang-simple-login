package config

// Global variables
var (
	SecretKey string = "secretkeyjwt"
)

// Database configuration
var (
	Database          = "postgres"
	Database_name     = "postgres-login"
	Database_password = "postgres"
	Database_host     = "localhost"
	Database_user     = "postgres"
	Database_port     = "5432"
	Database_url      = "host=" + Database_host +
		" user=" + Database_user +
		" password=" + Database_password +
		" dbname=" + Database_name +
		" port=" + Database_port + " sslmode=disable TimeZone=Asia/Shanghai"
)
