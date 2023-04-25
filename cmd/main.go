package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	myenv "github.com/r4nk0-k/karaoke_song_list/pkg/env"
	"github.com/r4nk0-k/karaoke_song_list/pkg/route"
	"github.com/spf13/cobra"
)

func main() {
	// mainの動作を引数で切り替えたいのでなんかコマンドを増やす
	rootCommand.AddCommand(runServer)
	rootCommand.AddCommand(initMysqlContainer)
	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}

var (
	rootCommand = &cobra.Command{
		Use:   "karaoke_song_list",
		Short: "run karaoke_song_list server process",
	}

	// useに指定した引数が来たらこのコマンドを実行する. ex. go run cmd/main.go run
	runServer = &cobra.Command{
		Use: "run",

		Run: func(cmd *cobra.Command, args []string) {
			if err := myenv.Parse(); err != nil {
				log.Fatal(err)
			}
			g := gin.New()
			//TODO: CORS設定
			route.Routes(g)
			fmt.Println(myenv.Get().Port)
			g.Run(myenv.Get().Port)
		},
	}

	initMysqlContainer = &cobra.Command{
		Use:   "init",
		Short: "init mysql container",
		Run: func(cmd *cobra.Command, args []string) {
			port := os.Getenv("MYSQL_PORT")
			newMysqlContainer("mysql_local", port)
		},
	}
)

const (
	DefaultMysqlUsername     = "root"
	DefaultMysqlPassword     = "password"
	DefaultMysqlDatabaseName = "testdb"
	DefaultMysqlHost         = "localhost"
)

// mysqlのdocker containerを起動する
func newMysqlContainer(containerName string, hostMysqlPort string) {
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 5
	if err != nil {
		log.Fatal(err)
	}

	// containerがすでにあったら削除する
	if err := pool.RemoveContainerByName(containerName); err != nil {
		log.Fatal(err)
	}

	opt := &dockertest.RunOptions{
		Name:       containerName,
		Repository: "mysql",
		Tag:        "8.0",
		Env: []string{
			fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", DefaultMysqlPassword),
			fmt.Sprintf("MYSQL_ROOT_USER=%s", DefaultMysqlUsername),
			fmt.Sprintf("MYSQL_DATABASE=%s", DefaultMysqlDatabaseName),
		},
		Cmd: []string{
			"mysqld",
			"--character-set-server=utf8mb4",
		},
	}

	if hostMysqlPort != "" {
		opt.PortBindings = map[docker.Port][]docker.PortBinding{
			"3306/tcp": {{
				HostIP:   DefaultMysqlHost,
				HostPort: fmt.Sprintf("%s/tcp", hostMysqlPort),
			}},
		}
	}

	if _, err := pool.RunWithOptions(opt); err != nil {
		log.Fatal(err)
	}
}
