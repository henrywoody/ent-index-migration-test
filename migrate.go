package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/henrywoody/ent-index-migration-test/ent"
	"github.com/henrywoody/ent-index-migration-test/ent/migrate"
	_ "github.com/lib/pq"
)

func main() {
	connStr := getConnectionString()
	client, err := ent.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx := context.Background()

	migrationFilePath := "migration.sql"
	file, err := os.Create(migrationFilePath)
	if err != nil {
		log.Fatalf("failed creating auto migration file: %v\n", err)
	}
	defer file.Close()

	var builder strings.Builder

	err = client.Schema.WriteTo(
		ctx,
		&builder,
		migrate.WithGlobalUniqueID(true),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed writing schema changes: %v\n", err)
	}

	migrationText := builder.String()
	file.WriteString(migrationText)
}

func getConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_SSL_MODE"),
	)
}
