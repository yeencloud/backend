package apps

import (
	"back/src/adapters/http/gin"
	"back/src/adapters/persistence/database/postgres"
	"back/src/core/domain"
	"back/src/core/usecases"
)

func MainBackend(bundle *domain.ApplicationBundle) error {
	httpConfig := bundle.Config.GetHTTPConfig()
	databaseConfig := bundle.Config.GetDatabaseConfig()

	database := postgres.StartGormDatabase(databaseConfig, "default")
	database.Migrate()

	ucs := usecases.NewInteractor(database, database, database, bundle.Translator)

	http := gin.NewServiceHttpServer(ucs, httpConfig, bundle.Translator)

	return http.Listen()
}
