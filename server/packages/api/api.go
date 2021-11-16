package api

import (
	"goapp/packages/config"
	"goapp/packages/db"

	"github.com/apex/log"
)

func StartServer() {
	conn, err := db.ConnectDB()
	if err != nil {
		log.WithField("reason", err.Error()).Fatal("Db connection error occurred")
	}
	defer conn.Close()

	runMigration := config.Config[config.RUN_MIGRATION]
	dbName := config.Config[config.POSTGRES_DB]
	port := config.Config[config.SERVER_PORT]

	if runMigration == "true" && conn != nil {
		if err := db.Migrate(conn, dbName); err != nil {
			log.WithField("reason", err.Error()).Fatal("db migration failed")
		}
	}

	server := httpServer(conn)
	server.Listen(port)
}
