package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

type Data struct {
	date_posted time.Time `json:"date_posted"`
	date_string string    `json:"date_string"`
	current_cfs string    `json:"current_cfs"`
	time_posted string    `json:"time_posted`
	forecast    string    `json:"forecast"`
	expires     string    `json:"expires"`
	scheduled_release bool `json:"scheduled_release"`
}

func New() (*Data, error)  {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var supabaseUrl string = os.Getenv("SUPABASE_URL")

	conn, err := pgx.Connect(context.Background(), supabaseUrl)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	} else {
		log.Println("connected to DB")
	}
	defer conn.Close(context.Background())

	//check db's recent posting
	var existingData Data
	query := `SELECT date_posted, date_string, current_cfs, time_posted, forecast, expires, scheduled_release  FROM daily_data ORDER BY id DESC LIMIT 1`
	row := conn.QueryRow(context.Background(), query)
	err = row.Scan(&existingData.date_posted, &existingData.date_string, &existingData.current_cfs, &existingData.time_posted, &existingData.forecast, &existingData.expires, &existingData.scheduled_release)
	if err != nil {
		log.Fatalf("Error querying database: %v", err)
		return nil, err
	}

	return &existingData, nil
}
