package user

import(
  "github.com/gin-gonic/gin"
  "github.com/carlqt/geodude/models"
)

func Create(c *gin.Context) {
  user := &models.User{
    Username: c.PostForm("username"),
    Email: c.PostForm("email"),
    Password: c.PostForm("password"),
    PasswordConfirmation: c.PostForm("passwordConfirmation"),
    Role: c.PostForm("type"),
  }

  if err := user.Create(); err != nil {
    c.JSON(400, gin.H{
      "error": err.Error(),
    })
  } else {
    c.JSON(200, gin.H{
      "status": "ok",
    })
  }
}