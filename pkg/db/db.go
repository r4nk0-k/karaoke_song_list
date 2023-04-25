package db

import (
	"fmt"
	"time"

	driver "github.com/go-sql-driver/mysql"
	"github.com/r4nk0-k/karaoke_song_list/pkg/env"
)

func GetDBInstanceSingleton() {
	fmt.Println("ここでDBのインスタンスを返したりする")
}

func GenerateDSNFromEnv() string {
	v := env.Get()

	cfg := driver.Config{
		User:      v.Mysql.Username,
		Passwd:    v.Mysql.Password,
		DBName:    v.Mysql.DatabaseName,
		Net:       "tcp",
		Addr:      fmt.Sprintf("%s:%s", v.Mysql.Host, v.Mysql.Port),
		ParseTime: true,
		Loc:       time.Local,
		Params: map[string]string{
			"charset": "utf8mb4",
		},
		AllowNativePasswords: true,
	}

	return cfg.FormatDSN()
}
