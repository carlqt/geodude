package user

import(
  // "encoding/json"
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


  if err := user.Validate(); err != nil {
    c.JSON(400, gin.H{
      "error": err.Error(),
    })
  } else {
    if err := user.Create(); err != nil {
      c.JSON(400, gin.H{
        "error": err.Error(),
      })
    } else {
      // jsonEncoded, _ := json.Marshal(user)
      c.JSON(200, user)
    }
  }
}