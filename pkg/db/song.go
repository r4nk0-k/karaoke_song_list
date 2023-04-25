package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/avast/retry-go"
	"github.com/r4nk0-k/karaoke_song_list/pkg/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SongDBController struct {
	DB *gorm.DB
}

func NewSongDBController() (*SongDBController, error) {
	dsn := GenerateDSNFromEnv()
	sqlDB, err := sql.Open("mysql", dsn)

	fmt.Println("opened")

	var gormDB *gorm.DB
	err = retry.Do(
		func() error {
			var e error
			gormDB, e = gorm.Open(mysql.New(mysql.Config{
				Conn: sqlDB,
			}), &gorm.Config{})
			return e
		},
		retry.Delay(5*time.Second),
		retry.Attempts(5),
	)
	if err != nil {
		return nil, err
	}

	return &SongDBController{DB: gormDB}, nil
}

func (c *SongDBController) Migrate(models ...interface{}) error {
	return c.DB.AutoMigrate(
		models...,
	)
}

func ListSong() ([]entity.Song, error) {
	return []entity.Song{{ID: "test"}}, nil
}
