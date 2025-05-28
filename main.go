package main
import (
"backend-api/controllers"
"backend-api/models"

"github.com/gin-gonic/gin"
)

func main(){
//inisialisasi gin
router := gin.Default()

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

//mulai server dengan route 8000
router.Run(":8000")
}
