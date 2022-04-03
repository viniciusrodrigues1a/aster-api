package factory

import (
	"transaction-service/cmd/transaction-service/server/database"

	eventstorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/event-store-lib"
	statestorelib "github.com/viniciusrodrigues1a/aster-api/pkg/infrastructure/state-store-lib"
)

func makeMongoEventStoreRepository() *eventstorelib.MongoEventStoreRepository {
	return eventstorelib.New(database.MongoConn.Context, database.MongoConn.Client, "transactions")
}

func makeRedisStateStoreRepository() *statestorelib.RedisStateStoreRepository {
	return statestorelib.New(database.RedisConn.Context, database.RedisConn.Client)
}
