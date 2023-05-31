package main

import (
	"github.com/gin-gonic/gin"
	"API/OrdresAPI/middlewares"
	"API/OrdresAPI/entity"
	"API/OrdresAPI/initializer"
	"os"
	"fmt"
	"io"
	"github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/joho/godotenv"
)

var (
	DB *sql.DB
	FD *os.File
)

func SetupLogOutput() {
	FD, _ = os.Create("OrdresAPI.log")
	gin.DefaultWriter = io.MultiWriter(FD, os.Stdout)
}

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	DB, err = initializer.ConnectDataBase()
	if err != nil {
		panic(err)
		return 
	}
	// If you want to get the logs written in a log file .
	SetupLogOutput()
	defer FD.Close()
	r := gin.New()
	r.Use(gin.Recovery(), middlewares.Logger())
	r.GET("/Admin", middlewares.BasicAuth(), GetOrdres)
	r.POST("/Admin", middlewares.BasicAuth(), ChangeOrdreState)
	r.POST("/Commande", AddOrdre)
	r.Run(":8080")
}

func ConnectDataBase() error{
	// Capture connection properties.
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "BadrCham",
    }
    // Get a database handle.
    var err error
    DB, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        return err
    }

    pingErr := DB.Ping()
    if pingErr != nil {
        return pingErr
    }
    fmt.Println("Connected!")
	return nil
}

// The Admin can see all the commandes and also delete commandes one by one . 
func GetOrdres(ctx *gin.Context){
	ordres, err := GetAll()
	if err != nil {
		ctx.JSON(400, nil)
	} else {
		ctx.JSON(200, ordres)
	}
}

func ChangeOrdreState(ctx *gin.Context){
	err := ChangeState(ctx)
	if err != nil {
		ctx.JSON(400, nil)
	} else {
		ctx.JSON(200, nil)
	}
}

// The Client can Only Add 
func AddOrdre(ctx *gin.Context){
	err := SaveOrdre(ctx)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(400, nil)
	} else {
		ctx.JSON(200, nil)
	}
}



func GetAll() ([]entity.IDClient ,error) {
	// here -> make a request to get all the commandes from the DB
	var ordres []entity.IDClient
	rows, err := DB.Query("SELECT * FROM Ordres")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
        var ordre entity.IDClient
        if err := rows.Scan(&ordre.Id, &ordre.Name, &ordre.Phone, &ordre.Adress, &ordre.Ordre, &ordre.State); err != nil {
            return nil, err
        }
        ordres = append(ordres, ordre)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return ordres, nil
}

func SaveOrdre(ctx *gin.Context) error {
	var client entity.Client
	var err error
	err = ctx.BindJSON(&client)
	if err != nil {// InvalidUnmarshalError is returned if it contains an other type inside .
		return err
	}
	// checking the Details of the ordre like the phone number should be of length 10 and start with "0"...
	
	// Sending the ordre to be saved in the DB
	_, err = DB.Exec("INSERT INTO Ordres (Name, Phone, Adress, Ordre, State) VALUES (?,?,?,?,?)", client.Name, client.Phone, client.Adress, client.Ordre, 0)
	if err != nil {
		return err
	}
	return nil
}

func ChangeState(ctx *gin.Context) error {
	var id uint32
	var err error
	err = ctx.BindJSON(&id)
	if err != nil {
		return err
	}
	var currentState uint32

	err = DB.QueryRow("SELECT State FROM Ordres WHERE id = ?", id).Scan(&currentState)
	if err != nil {
		return err
	}
	// Change the State of an ordre in the DB
	if currentState < 2 {
		_, err = DB.Exec("UPDATE Ordres SET State = State + 1 WHERE id = ?", id)
		if err != nil {
			return err
		}
	}
	return nil
}