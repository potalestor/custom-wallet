package main

import (
	"fmt"
	"log"
	"os"

	"github.com/potalestor/custom-wallet/pkg/api"
	"github.com/potalestor/custom-wallet/pkg/app"
	"github.com/potalestor/custom-wallet/pkg/cfg"
	"github.com/potalestor/custom-wallet/pkg/db"
	"github.com/potalestor/custom-wallet/pkg/logger"
	"github.com/potalestor/custom-wallet/pkg/repo"
	"github.com/spf13/cobra"
)

var (
	config cfg.Config

	rootCmd = &cobra.Command{
		Use:   "custom-wallet",
		Short: "A payment system that uses custom wallets.",

		Run: func(cmd *cobra.Command, args []string) {
			if err := config.Validate(); err != nil {
				log.Fatalf("configuration is invalid: %+v", config)
			}

			logger.Initialize(&config)

			log.Println("initialize migration")
			{
				migrationdb := db.NewPostgresDB(&config)
				if err := migrationdb.Open(); err != nil {
					log.Fatalf("migration does not initialize: %v\n%+v", err, config)
				}
				defer migrationdb.Close()

				if err := migrationdb.Migrate(); err != nil {
					log.Fatalf("migration does not perform: %v\n%+v", err, config)
				}
			}

			log.Println("initialize repository")
			repository := repo.NewPgStorage(&config)
			if err := repository.Open(); err != nil {
				log.Fatalf("repository does not initialize: %v\n%+v", err, config)
			}
			defer repository.Close()

			log.Println("initialize wallet")
			wallet := app.NewWallet(repository)

			log.Println("initialize web-service")
			if err := api.NewGraceful(
				api.NewAPI(wallet).Build(),
			).Run(config.Web.Adddress); err != nil {
				log.Fatalf("web-server does not initialize: %v\n%+v", err, config)
			}

		},
	}
)

func main() {
	rootCmd.PersistentFlags().StringVarP(&config.Database.DB, "dbname", "d", "wallet", "DB name")
	rootCmd.PersistentFlags().StringVarP(&config.Database.User, "user", "u", "postgres", "DB username ")
	rootCmd.PersistentFlags().StringVarP(&config.Database.Password, "pass", "p", "postgres", "DB password")
	rootCmd.PersistentFlags().StringVarP(&config.Database.Host, "host", "s", "localhost", "DB address")
	rootCmd.PersistentFlags().IntVarP(&config.Database.Port, "port", "o", 5432, "DB port")
	rootCmd.PersistentFlags().StringVarP(&config.Migration.Path, "mpath", "m", "../../scripts/", "Migration path scripts")
	rootCmd.PersistentFlags().BoolVarP(&config.Migration.Enabled, "mon", "n", true, "Migration enabled")
	rootCmd.PersistentFlags().BoolVarP(&config.Logger.File, "log", "l", false, "Default log to Console")
	rootCmd.PersistentFlags().StringVarP(&config.Web.Adddress, "addr", "a", ":8080", "WEB-address")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
