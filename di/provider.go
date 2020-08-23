// +build wireinject

package di

import (
	"github.com/google/wire"
	"os"
	"tkame123-net/worldcup-gq-server/adapter"
	"tkame123-net/worldcup-gq-server/infra/mongodb"
	"tkame123-net/worldcup-gq-server/infra/mongodb/competition"
)

var providerSet = wire.NewSet(
	provideMongodbClient,
	competition.NewRepository,
)

func provideMongodbClient() adapter.MongoClient {
	return mongodb.NewClient(os.Getenv("MONGODB_URI"), os.Getenv("MONGODB_DATABASE"))
}

// Todo:
//func ResolveServer() *grpc.Server {
//	wire.Build(providerSet)
//	return nil
//}
