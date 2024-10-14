package cmd

import (
	"context"
	"log"
	"time"

	"github.com/arvinpaundra/technical-test-api-mnc/config"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/factory"
	"github.com/arvinpaundra/technical-test-api-mnc/internal/http"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/database"
	"github.com/arvinpaundra/technical-test-api-mnc/pkg/util"
	"github.com/gofiber/fiber/v2"

	"github.com/spf13/cobra"
)

var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "The rest command to handle RESTful operations",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadEnv("config", "yaml", ".")

		f := factory.NewFactory()
		app := fiber.New()

		http.NewHttpRouter(app, f)

		go func() {
			if err := app.Listen(":" + config.GetAppPort()); err != nil {
				log.Printf("failed start server: %s", err.Error())
			}
		}()

		wait := util.GracefulShutdown(context.Background(), 30*time.Second, map[string]func(ctx context.Context) error{
			"http-server": func(ctx context.Context) error {
				return app.ShutdownWithContext(ctx)
			},
			"mongo": func(ctx context.Context) error {
				conn := database.GetConnection()

				db, _ := conn.DB()

				return db.Close()
			},
		})

		<-wait
	},
}

func init() {
	rootCmd.AddCommand(restCmd)
}
