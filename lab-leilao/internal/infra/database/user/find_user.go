package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/wandermaia/aulas-pos-golang/lab-leilao/configuration/logger"
	"github.com/wandermaia/aulas-pos-golang/lab-leilao/internal/entity/user_entity"
	"github.com/wandermaia/aulas-pos-golang/lab-leilao/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Essa struct é necessária para comunicação com o mongodb. Ele necessita dessas tags diferentes
type UserEntityMongo struct {
	Id   string `bson:"_id"`  // Esses campos serão representados dentro do mongo dessa forma.
	Name string `bson:"name"` // A tag para esse caso é um pouco diferente do json
}

type UserRespository struct {
	Collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) *UserRespository {
	return &UserRespository{
		Collection: database.Collection("users"),
	}
}

// Função anexada no UserRespository
// Neste caso é retornado um user_entity porque quem vai chamar essa função é a aplicação
// Implementa a interface UserRepositoryInterface do UserEntity
func (ur *UserRespository) FindUserById(ctx context.Context, userId string) (*user_entity.User, *internal_error.InternalError) {

	filter := bson.M{"_id": userId}

	var userEntityMongo UserEntityMongo
	err := ur.Collection.FindOne(ctx, filter).Decode(userEntityMongo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(fmt.Sprintf("User not found with this id = %s", userId), err)
			return nil, internal_error.NewNotFoundError(
				fmt.Sprintf("User not found with this id = %s", userId))
		}
		logger.Error("Error trying to find user by userId", err)
		return nil, internal_error.NewInternalServerError("Error trying to find user by userId")

	}

	userEntity := &user_entity.User{
		Id:   userEntityMongo.Id,
		Name: userEntityMongo.Name,
	}

	return userEntity, nil

}
