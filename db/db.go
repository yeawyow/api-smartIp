package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var db5 *gorm.DB
func GetDB() *gorm.DB{
	return db
}

func GetDB5() *gorm.DB{
	return db5
}
type SqlLogger struct {
  logger.Interface	
}
func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error){
	sql , _ :=fc() 
	fmt.Printf("%v\n==================\n",sql)
}



func SetupDB(){

	dsn := os.Getenv("MYSQL_DNS")
	database,err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: &SqlLogger{},
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //nochang table name
		},
	})
	dsn5 := os.Getenv("MYSQL_DNS5")
	database5,err := gorm.Open(mysql.Open(dsn5), &gorm.Config{
		Logger: &SqlLogger{},
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //nochang table name
		},
	})
	db5=database5
	//database, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil{
          panic("fail to connet databese")
	}
  
	database.AutoMigrate()
	
	db=database

}