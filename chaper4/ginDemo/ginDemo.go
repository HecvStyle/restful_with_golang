package main

import (
	"database/sql"
	"log"
	"net/http"
	"restful_with_golang/chaper4/RailDemo/dbutils"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type StationResource struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	OpeningTime  string `json:"opening_time"`
	CloseingTime string `json:"closing_time"`
}

// GetStation ...
func GetStation(c *gin.Context) {
	var station StationResource
	id := c.Param("station_id")
	err := DB.QueryRow("select ID, Name, CAST(OPENING_TIME as CHAR),CAST(CLOSING_TIME as CHAR) from station where id=?", id).Scan(&station.ID, &station.Name, &station.OpeningTime, &station.CloseingTime)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"result": station,
		})
	}
}

// CreatStation ...
func CreatStation(c *gin.Context) {
	var station StationResource
	if err := c.BindJSON(&station); err == nil {
		statement, _ := DB.Prepare("insert into station(NAME,OPENING_TIME,CLOSING_TIME) values(?,?,?)")
		result, _ := statement.Exec(station.Name, station.OpeningTime, station.CloseingTime)
		if err == nil {
			newID, _ := result.LastInsertId()
			station.ID = int(newID)
			c.JSON(http.StatusOK, gin.H{
				"result": station,
			})
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func main() {
	var err error
	DB, err = sql.Open("sqlite3", "../RailDemo/railAPI/railapi.db")
	if err != nil {
		log.Println("Driver creation faild!")
	}
	dbutils.Initialize(DB)
	r := gin.Default()
	r.GET("/v1/station/:station_id", GetStation)
	r.Run(":8000")
}
