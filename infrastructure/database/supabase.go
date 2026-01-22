package supabase

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func NewSupabase() (*sql.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		os.Getenv("SUPABASE_HOST"),
		os.Getenv("SUPABASE_PORT"),
		os.Getenv("SUPABASE_USER"),
		os.Getenv("SUPABASE_PASSWORD"),
		os.Getenv("SUPABASE_DB"),
	)

	//OpenはDBの種類を入れるから　supabaseとかじゃなくpostgres
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("supabase接続OK")

	return db, nil
}