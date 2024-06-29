package mongodb

import (
	"context"
	"os"

	"github.com/wandermaia/aulas-pos-golang/lab-leilao/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGODB_URL = "MONGODB_URL"
	MONGODB_DB  = "MONGODB_DB"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongoURL := os.Getenv(MONGODB_URL)
	mongoDatabase := os.Getenv(MONGODB_DB)

	client, err := mongo.Connect(
		ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		logger.Error("Error trying to connect to mongodb database", err)
		return nil, err
	}

	//Antes de retornar a conexão, podemos testar se o banco realemtne está acessível
	// ou se está apenas conectado
	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("Error trying to ping mongodb database", err)
		return nil, err
	}

	return client.Database(mongoDatabase), nil

}
