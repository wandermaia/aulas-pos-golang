package bid

import (
	"context"
	"sync"

	"github.com/wandermaia/aulas-pos-golang/lab-leilao/configuration/logger"
	"github.com/wandermaia/aulas-pos-golang/lab-leilao/internal/entity/auction_entity"
	"github.com/wandermaia/aulas-pos-golang/lab-leilao/internal/entity/bid_entity"
	"github.com/wandermaia/aulas-pos-golang/lab-leilao/internal/infra/database/auction"
	"github.com/wandermaia/aulas-pos-golang/lab-leilao/internal/internal_error"
	"go.mongodb.org/mongo-driver/mongo"
)

type BidEntityMongo struct {
	Id        string  `bson:"_id"`
	UserId    string  `bson:"user_id"`
	AuctionId string  `bson:"auction_id"`
	Amount    float64 `bson:"amount"`
	Timestamp int64   `bson:"timestamp"`
}

// Possui o método CreateBid
type BidRepository struct {
	Collection        *mongo.Collection
	AuctionRepository *auction.AuctionRepository
}

// Parte do objeto BidRepository
func (bd BidRepository) CreateBid(
	ctx context.Context,
	bidEntities []bid_entity.Bid) *internal_error.InternalError {

	var wg sync.WaitGroup

	// Pode receber vários bids (lances) de vários leilões diferentes.
	for _, bid := range bidEntities {
		wg.Add(1)

		go func(bidValue bid_entity.Bid) {
			defer wg.Done()

			// Coletando os dados do leilão. Problema: se chegar muitos lances para o mesmo leilao, vai
			// todas as vezes no banco de dados para saber se está ativo e quando vai fechar
			//- como saber isso localmente.
			//-
			auctionEntity, err := bd.AuctionRepository.FindAuctionById(ctx, bidValue.AuctionId)
			// se o leilão não for encontrado, o lance será ignorado
			if err != nil {
				logger.Error("Error trying to find auction by id", err)
				return
			}

			// Se o status do leilão não for ativo, o lance será ignorado
			if auctionEntity.Status != auction_entity.Active {
				return
			}

			bidEntityMongo := &BidEntityMongo{
				Id:        bidValue.Id,
				UserId:    bidValue.UserId,
				AuctionId: bidValue.AuctionId,
				Amount:    bid.Amount,
				Timestamp: bidValue.Timestamp.Unix(),
			}

			// Se não conseguir inserir, vai ser ignorado
			if _, err := bd.Collection.InsertOne(ctx, bidEntityMongo); err != nil {
				logger.Error("Error trying to insert bid", err)
				return
			}

		}(bid)
	}

	wg.Wait()
	return
}
