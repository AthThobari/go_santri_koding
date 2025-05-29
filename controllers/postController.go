package controllers

import (
"errors"
"net/http"
"backend-api/models"

"gorm.io/gorm"
"github.com/go-playground/validator/v10"
"github.com/gin-gonic/gin"
)

// type validation post input
type ValidatePostInput struct {
    Title string `json:"title" binding:"required"`
    Content string `json:"content" binding:"required"`
}

// type error message
type ErrorMsg struct {
    Field string `json:"field"`
    Message string `json:"message"`
}

// func get error message
func GetErrorMsg(fe validator.FieldError) string {
switch fe.Tag() {
    case "required":
    return "This field is required"
}
return "Unknown error"
}

//get all posts
func FindPosts(c *gin.Context){
var posts []models.Post
models.DB.Find(&posts)

//return json
c.JSON(200, gin.H{
"success": true,
"message": "List Data Post",
"data": posts,
})
}

// store a post
func StorePost(c *gin.Context) {
// validate input
var input ValidatePostInput
if err := c.ShouldBindJSON(&input); err != nil {
var ve validator.ValidationErrors
if errors.As(err, &ve) {
out := make([]ErrorMsg, len(ve))

for i, fe := range ve {
out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
}

c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
"errors":out,
})
}

return
}

//create post
post := models.Post{
Title : input.Title,
Content : input.Content,
}

models.DB.Create(&post)

//return response JSON
c.JSON(201, gin.H{
"success": true,
"messsage": "Post Created Successfully",
})
}

// func get post by id
func FindPostById(c *gin.Context) {
var post models.Post
if err := models.DB.Where("id = ?", c.Param("id")).
First(&post).Error; err != nil {
c.JSON(http.StatusBadRequest,
gin.H{
"error":"Record not found!",
})
return
}

c.JSON(200, gin.H{
"success": true,
"message":"Detail data post by id : " + c.Param("id"),
"data":post,
})
}

// func update post
func UpdatePost(c *gin.Context) {
var post models.Post
if err := models.DB.Where("id = ?", c.Param("id")).
First(&post).Error;
err != nil {
c.JSON(http.StatusBadRequest, gin.H{
"error":"Record not found!",
})
}

//validate input
var input ValidatePostInput
if err := c.ShouldBindJSON(&input);
err != nil {
var ve validator.ValidationErrors
if errors.As(err, &ve) {
out := make([]ErrorMsg, len(ve))
for i, fe := range ve {
out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
}
c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
"errors":out,
})
}
return

}

//update post
models.DB.Model(&post).Updates(input)

c.JSON(200, gin.H{
"success":true,
"message":"Post updated successfully",
"data":post,
})
}

// Delete Post
func DeletePost(c *gin.Context) {

var post models.Post

if err := models.DB.Session(&gorm.Session{
PrepareStmt: false,
}).Where("id = ?", c.Param("id")).First(&post).
Error; err != nil {
c.JSON(http.StatusBadRequest, gin.H{
"error" : "Record not found!",
})
return
}

//delete post
if err := models.DB.Session(&gorm.Session{
PrepareStmt: false,
}).Delete(&post).Error; err !=
nil {
c.JSON(500, gin.H{
"error": "Failed to delete post",
})
return
}

c.JSON(200, gin.H{
"success" : true,
"message" : "Post Deleted Successfully",
})
}

