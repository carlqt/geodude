package properties

import(
  "github.com/shopspring/decimal"
  "github.com/gin-gonic/gin"
  "github.com/fatih/color"
  "github.com/carlqt/geodude/geocode"
  "net/http"
  "strconv"
  "github.com/carlqt/geodude/models"
)

func Pong(c *gin.Context) {
  c.String(http.StatusOK, "Property pong")
}

func PropertySearch(c *gin.Context) {
  point, err := geocode.Geocode(c.Query("location"))
  if err != nil {
    color.Red(err.Error())
    c.JSON(400, gin.H{
      "error": err.Error(),
    })
  } else {
    p := models.NearbyProperty(point.Geometry.Location["lat"], point.Geometry.Location["lng"])
    c.JSON(http.StatusOK, p)
  }
}

func PropertyIndex(c *gin.Context) {
  p := models.AllProperties()
  c.JSON(http.StatusOK, p)
}

func PropertyCreate(c *gin.Context) {
  var err error

  property := &models.Property{Address: c.PostForm("address"),
    Price: strToCurrency(c.PostForm("price")),
    Description: c.PostForm("description"),
    Type: c.PostForm("type"),
    Name: c.PostForm("name"),
  }

  property, err = property.GeocodeAndCreate()

  if err != nil {
    color.Red(err.Error())
    c.JSON(500, gin.H{
      "error": err.Error(),
    })
  } else {
    c.JSON(http.StatusCreated, property)
  }
}

func PropertyGeocode(c *gin.Context) {
  result, err := geocode.Geocode(c.Query("location"))

  if err != nil {
    color.Red(err.Error())
    c.JSON(500, gin.H{
      "error": err.Error(),
    })
  } else {
    c.JSON(http.StatusOK, gin.H{
      "lng": result.Geometry.Location["lng"],
      "lat": result.Geometry.Location["lat"],
    })
  }
}

func PropertyDelete(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))
  err := models.PropertyDelete(id)

  if err != nil {
    c.JSON(400, gin.H{
      "error": err.Error(),
    })
  } else {
    c.JSON(200, gin.H{
      "status": "deleted",
    })
  }
}

func strToCurrency(str string) decimal.Decimal {
  price, _ := decimal.NewFromString(str)
  return price
}