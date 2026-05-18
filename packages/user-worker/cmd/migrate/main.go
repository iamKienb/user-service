package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"
	"user-shared-module/indexing"
	"user-worker-module/internal/bootstrap/config"
	"user-worker-module/internal/bootstrap/module"

	"github.com/elastic/go-elasticsearch/v8"
	configx "github.com/iamKienb/shopify-go-platform/config"
)

type MigrationLog struct {
	Filename  string    `json:"filename"`
	AppliedAt time.Time `json:"applied_at"`
}

const MigrationIndex = ".migration_history"

func main() {
	ctx := context.Background()

	cfg, err := configx.Loader[config.UserWorkerConfig]()
	if err != nil {
		log.Fatal("failed to load config")
	}

	infra, err := module.NewInfraModule(ctx, cfg)
	if err != nil {
		log.Fatal("failed connect to infra")
	}

	client := infra.ESService.GetClient()
	ensureMigrationIndex(ctx, client)

	validAliases := []string{indexing.UserAlias, indexing.ShopAlias}

	for _, alias := range validAliases {
		pattern := fmt.Sprintf("es_migrations/%s/*.json", alias)
		files, err := filepath.Glob(pattern)
		if err != nil {
			log.Fatalf("failed to list migration files for alias %s: %v", alias, err)
		}

		sort.Strings(files)

		for _, file := range files {
			filename := filepath.Base(file)
			fullHistoryKey := fmt.Sprintf("%s_%s", alias, filename)

			applied, err := isApplied(ctx, client, fullHistoryKey)
			if err != nil {
				log.Fatalf("failed to check history file %s: %v", filename, err)
			}
			if applied {
				log.Printf("⏩ Skip: %s (applied)", fullHistoryKey)
				continue
			}

			content, err := os.ReadFile(file)
			if err != nil {
				log.Fatalf("cannot read file %s: %v", filename, err)
			}

			log.Printf("Migrating [%s]: %s", alias, filename)
			if err = infra.ESService.BootstrapIndex(ctx, alias, string(content)); err != nil {
				log.Fatalf("failed to init bootstrap for %s: %v", alias, err)
			}

			if err := saveLogHisFile(ctx, client, fullHistoryKey); err != nil {
				log.Fatalf("failed to write log for file %s: %v", filename, err)
			}

			log.Printf("Success: %s", fullHistoryKey)
		}
	}
	log.Println("All migrations done!")
}

func isApplied(ctx context.Context, client *elasticsearch.TypedClient, fileName string) (bool, error) {
	res, err := client.Get(MigrationIndex, fileName).Do(ctx)
	if err != nil {
		return false, nil
	}

	return res.Found, nil
}

func saveLogHisFile(ctx context.Context, client *elasticsearch.TypedClient, fileName string) error {
	logEntry := MigrationLog{
		Filename:  fileName,
		AppliedAt: time.Now(),
	}
	_, err := client.Index(MigrationIndex).Id(fileName).Document(logEntry).Do(ctx)

	return err
}

func ensureMigrationIndex(ctx context.Context, client *elasticsearch.TypedClient) {
	exists, _ := client.Indices.Exists(MigrationIndex).Do(ctx)
	if !exists {
		_, _ = client.Indices.Create(MigrationIndex).Do(ctx)
		log.Printf("create index [%s] managed migration history", MigrationIndex)
	}
}
