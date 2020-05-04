package main

import (
	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go/relay"
	"graphql-go-gin/models"
	"graphql-go-gin/provider/sqlite"
	"graphql-go-gin/schema"
	"math/rand"
)

func main() {
	// init sqlite data
	initDB()
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.POST("/query", GraphqlHandler())

	r.Run("0.0.0.0:8080")
}

// GraphqlHandler gin graphql handler
func GraphqlHandler() gin.HandlerFunc {
	h := relay.Handler{
		Schema: schema.Schema,
	}
	// gin HandlerFunc
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// initDB init sqlite data
func initDB() {

	// TEST DATA TO BE PUT INTO THE DB
	var users = []models.User{
		models.User{Name: "小明"},
		models.User{Name: "小芳"},
		models.User{Name: "小雨"},
	}

	// Since the sqlite is torn down and created on every run, I know the above users will have
	// ID's 1, 2, 3
	var books = []models.Book{
		models.Book{Name: "《时间简史》", OwnerID: 1},
		models.Book{Name: "《解忧杂货铺》", OwnerID: 1},
		models.Book{Name: "《悟空传》", OwnerID: 1},
		models.Book{Name: "《三体》", OwnerID: 1},
		models.Book{Name: "《追风筝的人》", OwnerID: 1},
		models.Book{Name: "《摆渡人》", OwnerID: 1},
		models.Book{Name: "《岁月的泡沫》", OwnerID: 2},
		models.Book{Name: "《世界因你而不同》", OwnerID: 2},
		models.Book{Name: "《撒哈拉的故事》", OwnerID: 3},
		models.Book{Name: "《且听风吟》", OwnerID: 3},
	}

	// Tags to be put in the database
	var tags = []models.Tag{
		models.Tag{Title: "文学"},
		models.Tag{Title: "传记"},
		models.Tag{Title: "历史"},
		models.Tag{Title: "科学"},
		models.Tag{Title: "小说"},
	}

	db := sqlite.DB
	// drop older test data
	db.DropTableIfExists(&models.User{}, &models.Book{}, &models.Tag{})
	db.AutoMigrate(&models.User{}, &models.Book{}, &models.Tag{})

	// add users
	for _, u := range users {
		if err := db.Create(&u).Error; err != nil {
			panic(err)
		}
	}

	var tg = []models.Tag{}
	for _, t := range tags {
		if err := db.Create(&t).Error; err != nil {
			panic(err)
		}

		tg = append(tg, t)
	}
	// add book
	for _, b := range books {
		b.Tags = tg[:rand.Intn(5)]
		if err := db.Create(&b).Error; err != nil {
			panic(err)
		}
	}
}
