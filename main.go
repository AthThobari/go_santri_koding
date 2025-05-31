package main
import (
"backend-api/controllers"
"backend-api/models"
"time"
"github.com/gin-contrib/cors"
"github.com/gin-gonic/gin"
)

func main(){
//inisialisasi gin
router := gin.Default()

router.Use(cors.New(cors.Config{
AllowOrigins:
[]string{"*"},

AllowMethods:
[]string{"GET","POST","PUT","DELETE","OPTIONS"},

AllowHeaders:
[]string{"Origin","Content-Type","Accept"},

ExposeHeaders:
[]string{"Content-Length"},
	AllowCredentials: true,
	MaxAge: 12 * time.Hour,
}))

//panggil koneksi database
models.ConnectDatabase()


//membuat route dengan method GET
router.GET("/", func(c  *gin.Context) {
c.JSON(200, gin.H{
"message": "Hello world",
})
})

//membuat route create post
router.POST("api/posts", controllers.StorePost)

//membuat route get all posts
router.GET("api/posts", controllers.FindPosts)

//membuat route detail post
router.GET("api/posts/:id", controllers.FindPostById)

//membuat route update post
router.PUT("api/posts/:id", controllers.UpdatePost)

//membuat route delete post
router.DELETE("api/posts/:id", controllers.DeletePost)

//mulai server dengan route 8000
router.Run(":8000")
}
