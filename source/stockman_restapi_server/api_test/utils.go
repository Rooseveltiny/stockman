package apitest

import (
	"context"
	core "stockman/source/stockman_core"
	"stockman/source/stockman_dbs/client/postgresql"
	postgresutils "stockman/source/stockman_dbs/postgres_utils"
	stockmanrestapiserver "stockman/source/stockman_restapi_server"
)

const testBaseURL string = "http://localhost:8080"

// const applicationsJSON string = "application/json"
const textPlain string = "text/plain"

func runStockmanService() {
	core.InitAndRunStockmanService()
}

func prepareDB(ctx context.Context) {
	c, _ := postgresql.GetPostgresClient(ctx)
	postgresutils.PrepareTestPostgresSQL(ctx, c)
}

func dropDB(ctx context.Context) {
	c, _ := postgresql.GetPostgresClient(ctx)
	postgresutils.DropPreparedTestPostgresSQL(ctx, c)
}

func startRestAPIServer(ctx context.Context) (shoutDownServerFn func()) {
	return stockmanrestapiserver.StartRestAPIServer(ctx)
}

func runTest(ctx context.Context, test_f func(context.Context)) {
	runStockmanService()
	prepareDB(ctx)
	shoutDownRestAPIServer := startRestAPIServer(ctx)
	test_f(ctx)
	shoutDownRestAPIServer()
	dropDB(ctx)
}
