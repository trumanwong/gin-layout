package data

import (
	"ariga.io/atlas/sql/migrate"
	atlas "ariga.io/atlas/sql/schema"
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"fmt"
	"gin-layout/internal/conf"
	"gin-layout/internal/data/ent"
	entMigrate "gin-layout/internal/data/ent/migrate"
	_ "gin-layout/internal/data/ent/runtime"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"log"
	"time"
)

type Data struct {
	client *ent.Client
}

func NewEntClient(config *conf.Configs) *ent.Client {
	drv, err := sql.Open(config.Data.Database.Driver, config.Data.Database.Source)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
		return nil
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)
	return ent.NewClient(ent.Driver(drv))
}

func (d *Data) Migrate() error {
	// Run migration.
	err := d.client.Schema.Create(
		context.Background(),
		entMigrate.WithDropIndex(true),
		entMigrate.WithDropColumn(true),
		entMigrate.WithForeignKeys(false), // Disable foreign keys.
		// Hook into Atlas Diff process.
		schema.WithDiffHook(func(next schema.Differ) schema.Differ {
			return schema.DiffFunc(func(current, desired *atlas.Schema) ([]atlas.Change, error) {
				// Before calculating changes.
				changes, err := next.Diff(current, desired)
				if err != nil {
					return nil, err
				}
				// After diff, you can filter
				// changes or return new ones.
				return changes, nil
			})
		}),
		// Hook into Atlas Apply process.
		schema.WithApplyHook(func(next schema.Applier) schema.Applier {
			return schema.ApplyFunc(func(ctx context.Context, conn dialect.ExecQuerier, plan *migrate.Plan) error {
				// Example to hook into the apply process, or implement
				// a custom applier. For example, write to a file.
				//
				//  for _, c := range plan.Changes {
				//      fmt.Printf("%s: %s", c.Comment, c.Cmd)
				//      if err := conn.Exec(ctx, c.Cmd, c.Args, nil); err != nil {
				//          return err
				//      }
				//  }
				//
				return next.Apply(ctx, conn, plan)
			})
		}),
	)
	if err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}
	return nil
}

func NewData(client *ent.Client) *Data {
	return &Data{
		client: client,
	}
}

// withTx Reusable function that runs callbacks in a transaction:
func withTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewEntClient,

	NewGreeterRepo,
)
