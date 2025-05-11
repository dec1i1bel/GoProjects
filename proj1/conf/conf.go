package conf

import "os"

func SetDbAccess() {
	os.Setenv("USER", "root")
	os.Setenv("PASSWD", "vkshmuk0707")
	os.Setenv("NET_TYPE", "tcp")
	os.Setenv("HOST_PORT", "127.0.0.1:3306")
	os.Setenv("DB_NAME", "proj1")
}
