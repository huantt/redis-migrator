package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"redis-migrator/config"
	"redis-migrator/migrator"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Runs redis-migrator to run migration",
	Long:  `Runs redis-migrator to run migration`,
	Run: func(cmd *cobra.Command, args []string) {
		data := config.ParseConfig(configFilePath)
		if err := migrator.MigrateRedisData(context.Background(), data); err != nil {
			panic(err)
		}
	},
}

func init() {
	migrateCmd.PersistentFlags().StringVarP(&configFilePath, "config.file", "c", "config.yaml", "Location of configuration file to run migration.")
	rootCmd.AddCommand(migrateCmd)
}
