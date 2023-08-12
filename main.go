package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	env "github.com/caarlos0/env/v6"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

type Config struct {
	REDIS RDSConfig `envPrefix:"REDIS_"`
	MYSQL SQLConfig `envPrefix:"MYSQL_"`
}
type RDSConfig struct {
	DB_URL    string `env:"DB_URL"`
	DB_INDEX  int    `env:"DB_INDEX"`
	KEY_VALUE string `env:"KEY_VALUE"`
}
type SQLConfig struct {
	DB_HOST       string `env:"DB_HOST"`
	DB_USER       string `env:"DB_USER"`
	DB_PWD        string `env:"DB_PWD"`
	DB_PORT       int    `env:"DB_PORT"`
	DB_NAME       string `env:"DB_NAME"`
	TIME_TERMINAL int    `env:"TIME_TERMINAL"`
}

type Report struct {
	Provider string `json:"provider"`
	Campaign string `json:"campaign_id"`

	AvgBids   float64 `json:"bids_count"`
	AvgWins   float64 `json:"wins_count"`
	AvgSpends float64 `json:"spends"`
}

func main() {

	//config Setting
	var config Config
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}
	err = env.Parse(&config)

	if err != nil {
		fmt.Println("Env load err:", err)
	}
	// //Time ticker for refresh
	// ticker := time.NewTicker(config.MYSQL.TIME_TERMINAL * int(time.Minute))
	// defer ticker.Stop()
	//MySQL connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.MYSQL.DB_USER, config.MYSQL.DB_PWD, config.MYSQL.DB_HOST, config.MYSQL.DB_PORT, config.MYSQL.DB_NAME)

	//Open MySQL
	sql_db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = sql_db.Ping()
	if err != nil {
		log.Fatal("Error connecting to MySQL:", err)
	}
	defer sql_db.Close()

	fmt.Println("Successfully connected to MySQL!")
	//Redis connection
	rds_db := redis.NewClient(&redis.Options{
		Addr:     config.REDIS.DB_URL,
		Password: "",
		DB:       0,
	})

	pong, err := rds_db.Ping().Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}
	rds_db.Do("SELECT", config.REDIS.DB_INDEX)
	fmt.Println("Connected to Redis:", pong)
	// Get data from mysql

	//Set timer
	// timer := time.Now().Add(time.Duration(-config.MYSQL.TIME_TERMINAL * int(time.Minute)))
	timer := time.Date(2021, 7, 28, 12, 17, 25, 0, time.UTC).Add(time.Duration(-config.MYSQL.TIME_TERMINAL * int(time.Minute)))
	fmt.Printf("Loading data from %s\n", timer.Format("2006-01-02 15:04:05"))

	//Get Query
	rowsForCampaign, err := sql_db.Query("SELECT provider, campaign_id, AVG(bids_count),AVG(wins_count), AVG(spends) FROM prodata_openrtb_bidder_reports WHERE date_created >= ? GROUP BY provider, campaign_id", timer)
	if err != nil {
		log.Fatal("Error connecting to MySQL:", err)
	}
	rowsForNetwork, err := sql_db.Query("SELECT provider, AVG(bids_count),AVG(wins_count), AVG(spends) FROM prodata_openrtb_bidder_reports WHERE date_created >= ? GROUP BY provider", timer)
	if err != nil {
		log.Fatal("Error connecting to MySQL:", err)
	}

	//Add data to redis
	dataBid := make(map[string]map[string]float64)
	dataWin := make(map[string]map[string]float64)
	dataSpend := make(map[string]map[string]float64)
	for rowsForCampaign.Next() {
		var report Report
		// for each row, scan the result into our tag composite object
		err = rowsForCampaign.Scan(&report.Provider, &report.Campaign, &report.AvgBids, &report.AvgWins, &report.AvgSpends)
		if err != nil {
			log.Fatal("Error connecting to MySQL:", err)
		}

		//Data processing
		if dataBid[report.Provider] == nil {
			dataBid[report.Provider] = make(map[string]float64)
		}
		dataBid[report.Provider][report.Campaign] = report.AvgBids
		if dataWin[report.Provider] == nil {
			dataWin[report.Provider] = make(map[string]float64)
		}
		dataWin[report.Provider][report.Campaign] = report.AvgWins
		if dataSpend[report.Provider] == nil {
			dataSpend[report.Provider] = make(map[string]float64)
		}
		dataSpend[report.Provider][report.Campaign] = report.AvgSpends
	}

	for field, dataMap := range dataBid {
		jsonData, err := json.Marshal(dataMap)
		if err != nil {
			fmt.Println("JSON marshal error:", err)
			return
		}
		err = rds_db.HSet(field, "bids_campagin", string(jsonData)).Err()
	}
	for field, dataMap := range dataWin {
		jsonData, err := json.Marshal(dataMap)
		if err != nil {
			fmt.Println("JSON marshal error:", err)
			return
		}
		err = rds_db.HSet(field, "wins_campagin", string(jsonData)).Err()
	}
	for field, dataMap := range dataSpend {
		jsonData, err := json.Marshal(dataMap)
		if err != nil {
			fmt.Println("JSON marshal error:", err)
			return
		}
		err = rds_db.HSet(field, "spends_campagin", string(jsonData)).Err()
	}

	for rowsForNetwork.Next() {
		var report Report
		// for each row, scan the result into our tag composite object and add redis
		err = rowsForNetwork.Scan(&report.Provider, &report.AvgBids, &report.AvgWins, &report.AvgSpends)
		err = rds_db.HSet(report.Provider, "bids_network", report.AvgBids).Err()
		err = rds_db.HSet(report.Provider, "wins_network", report.AvgWins).Err()
		err = rds_db.HSet(report.Provider, "spends_network", report.AvgSpends).Err()
	}

}
