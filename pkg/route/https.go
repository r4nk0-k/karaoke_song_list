package route

import (
	"fmt"
	"log"
	"net/http"

	"github.com/r4nk0-k/karaoke_song_list/pkg/db"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	controller, err := db.NewSongDBController()
	if err != nil {
		log.Fatal(err)
	}

	songs := r.Group("/v1/api/songs")
	{
		songs.GET("", listSong(*controller))
	}
}

func listSong(controller db.SongDBController) func(c *gin.Context) {
	return func(c *gin.Context) {
		// TODO: query parameterの追加
		res, err := db.ListSong()
		if err != nil {
			fmt.Println("getSong error occured: ", err)
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
