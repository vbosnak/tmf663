package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "Toronto#2024@#"
	dbname   = "tmf663"
)

var (
	db *gorm.DB
)

func InitDB() {
	var err error
	//  "<username>:<password>@tcp(<url>:<port>)/<database_name>?charset=utf8mb4&parseTime=True"
	dsn := "root:Toronto#2024@#@tcp(localhost:3306)/tmf663?charset=utf8mb4&parseTime=True"
	// Open a connection to the MySQL database using the provided DSN.
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = d

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("Connected to MySQL!")
}

// GetDB returns the initialized database connection.
func GetDB() *gorm.DB {
	InitDB()
	return db
}
