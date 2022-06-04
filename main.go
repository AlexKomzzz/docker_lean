package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var port string
var db *sqlx.DB

type Database struct {
	dbs *sqlx.DB
}

type Car struct {
	Model string `json:"model"`
	Price int    `json:"price"`
}

func main() {
	DB := Database{dbs: db}
	defer db.Close()

	if err := DB.CreateDB(); err != nil {
		log.Fatal(err.Error(), " Error create table")
	}
	if err := DB.dbs.Ping(); err != nil {
		log.Fatal(err.Error(), " Error connect")
		return
	}
	mux := gin.Default()

	mux.GET("/car", DB.GetCar)
	mux.POST("/newcar", DB.CreateCar)
	mux.GET("/car/del", DB.DropDB)

	if err := mux.Run(port); err != nil {
		log.Fatal(err.Error())
	}

}

func init() {
	var dbConnString string

	if err := godotenv.Load(); err != nil {

		new_port, exists := os.LookupEnv("PORT")
		if exists {
			port = new_port
		} else {
			log.Fatal(err.Error(), ": not found PORT")
			return
		}

		db_host, exists := os.LookupEnv("DBHOST")
		if !exists {
			log.Fatal(err.Error(), ": not found DBHOST")
			return
		}

		db_port, exists := os.LookupEnv("DBPORT")
		if !exists {
			log.Fatal(err.Error(), ": not found DBPORT")
			return
		}

		db_user, exists := os.LookupEnv("DBUSER")
		if !exists {
			log.Fatal(err.Error(), ": not found DBUSER")
			return
		}

		db_name, exists := os.LookupEnv("DBNAME")
		if !exists {
			log.Fatal(err.Error(), ": not found DBNAME")
			return
		}

		db_pass, exists := os.LookupEnv("DBPASSWORD")
		if !exists {
			log.Fatal(err.Error(), ": not found DBPASSWORD")
			return
		}

		dbConnString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", db_host, db_port, db_user, db_name, db_pass)
	} else {
		port = os.Getenv("PORT")

		dbConnString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			os.Getenv("DBHOST"),
			os.Getenv("DBPORT"),
			os.Getenv("DBUSER"),
			os.Getenv("DBNAME"),
			os.Getenv("DBPASSWORD"),
		)
	}

	db1, err := sqlx.Open("postgres", dbConnString)
	if err != nil {
		log.Fatal("error: no connect DB ", err.Error())
		return
	}

	db = db1
}

func (d *Database) selectDB() ([]Car, error) {
	var car []Car

	query := "SELECT * FROM cars"

	err := d.dbs.Select(&car, query)

	return car, err
}

func (d *Database) GetCar(c *gin.Context) {
	log.Println("Start GET")
	log.Print("Start GET")
	car, err := d.selectDB()
	if err != nil {
		log.Fatal("error: select ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, car)
}

func (d *Database) CreateCar(c *gin.Context) {
	var car Car

	if err := c.BindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	createQuery := "INSERT INTO cars (model, price) VALUES ($1, $2)"
	_, err := d.dbs.Exec(createQuery, car.Model, car.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func (d *Database) DropDB(c *gin.Context) {
	stmt := "DROP TABLE cars"
	_, err := d.dbs.Exec(stmt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Fatal(err.Error(), " Error drop table")
	}
}

func (d *Database) CreateDB() error {
	stmt := "create table cars (model varchar(255), price int)"
	_, err := d.dbs.Exec(stmt)

	return err
}
