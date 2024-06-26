package main

import (
	"encoding/json"
	"golang-currency-service/models"
	"log"
	"net/http"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const apiURL = "https://api.nbrb.by/exrates/rates?periodicity=0"

func init() {
	dbUser, _ := web.AppConfig.String("db_user")
	dbPass, _ := web.AppConfig.String("db_pass")
	dbHost, _ := web.AppConfig.String("db_host")
	dbPort, _ := web.AppConfig.String("db_port")
	dbName, _ := web.AppConfig.String("db_name")
	dataSource := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True"

	orm.RegisterDataBase("default", "mysql", dataSource)
}

func main() {
	o := orm.NewOrm()

	resp, err := http.Get(apiURL)

	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	defer resp.Body.Close()

	var rates []models.Rate
	if err := json.NewDecoder(resp.Body).Decode(&rates); err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}

	for _, rate := range rates {
		log.Printf("Error inserting rate: %v", rate)
		_, err := o.Insert(&rate)
		if err != nil {
			log.Printf("Error inserting rate: %v", err)
		}
	}

	log.Println("Data updated successfully")
}
