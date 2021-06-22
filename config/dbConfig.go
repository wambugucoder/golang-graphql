package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/wambugucoder/golang-graphql/graph/model"
	"log"
	"strconv"
)

// Connect ->Connect to database and migrate models.
func Connect() {
	var err error

	portString := ExtractKey("PSQL_PORT")
	port, err := strconv.ParseUint(portString, 10, 32)

	DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", ExtractKey("PSQL_HOST"), port,
		ExtractKey("PSQL_USER"), ExtractKey("PSQL_PASS"), ExtractKey("PSQL_DB")))

	if err != nil {
		panic("Failed To Connect To Db")
	}
	DB.LogMode(true)
	log.Println("Connected To Database")

	//PERFORM MIGRATIONS
	DB.DropTableIfExists(&model.Question{}, &model.Choice{}).AutoMigrate(&model.Question{}, &model.Choice{})
	log.Println("Database has been migrated")

}
