package main

// only need mysql OR sqlite
// both are included here for reference
import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

//Calculation comment for the linter to stop yelling
type Calculation struct {
	ID          uint   `json:"id"`
	Calculation string `json:"calculation"`
}

func main() {
	db, err = gorm.Open("sqlite3", "./gorm.db")

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	db.AutoMigrate(&Calculation{})
	//Router
	r := gin.Default()
	//GET LAST 10
	r.GET("/calc/", GetLast10)
	//CREATE
	r.POST("/calc/", AddCalc)
	//RUN
	r.Run(":8080")
}

//GetLast10 Comment for linter
func GetLast10(c *gin.Context) {
	var calcs []Calculation
	if err := db.Find(&calcs).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, calcs)
	}
}

//AddCalc here
func AddCalc(c *gin.Context) {
	var calc Calculation
	c.BindJSON(&calc)
	db.Create(&calc)
	c.JSON(200, calc)
}
