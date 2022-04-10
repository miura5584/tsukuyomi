package ent

import (
	"context"
	"entgo.io/ent/dialect"
	"log"
	"tsukuyomi/ent/migrate"
)

var Database, DatabaseErr = Open(dialect.Postgres,
	"host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")

func init() {
	for i := 0; i < 20; i++ {
		if DatabaseErr != nil {
			log.Printf("fialed connection database: %v", DatabaseErr)
			log.Printf("try connect database: %d / 20", i)
		} else {
			break
		}
	}

	defer func(Database *Client) {
		err := Database.Close()
		if err != nil {
			log.Fatalf("fialed instance database: %v", err)
		}
	}(Database)
	ctx := context.Background()
	if err := Database.Schema.Create(ctx, migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed create schema: %v", err)
	}
}
