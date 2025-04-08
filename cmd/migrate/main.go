package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/anurag/shortenurl/internal/db/migrations"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/migrate"
)

func main() {
	// Initialize SQLite database
	sqlite, err := sql.Open(sqliteshim.ShimName, "file:shortenurl.db?cache=shared")
	if err != nil {
		log.Fatal(err)
	}
	db := bun.NewDB(sqlite, sqlitedialect.New())

	// Create migration instance
	// migrations := migrate.NewMigrations()

	// Register your migrations here
	// Example:
	// migrations.Register(func(ctx context.Context, db *bun.DB) error {
	//     _, err := db.NewCreateTable().Model((*YourModel)(nil)).Exec(ctx)
	//     return err
	// })

	// Create migrator
	migrator := migrate.NewMigrator(db, migrations.Migrations)

	// Run migrations
	err = migrator.Init(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	group, err := migrator.Migrate(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	if group.IsZero() {
		fmt.Printf("there are no new migrations to run\n")
		return
	}

	fmt.Printf("migrated to %s\n", group)
}
