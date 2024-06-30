package bid

import (
	"context"
	"fmt"
	"time"

	"github.com/wandermaia/aulas-pos-golang/lab-leilao/configuration/logger"
	"github.com/wandermaia/aulas-pos-golang/lab-leilao/internal/entity/bid_entity"
	"github.com/wandermaia/aulas-pos-golang/lab-leilao/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Parte do objeto BidRepository
func (bd *BidRepository) FindBidByAuctionId(
	ctx context.Context, auctionId string) ([]bid_entity.Bid, *internal_error.InternalError) {

	filter := bson.M{"auctionId": auctionId}
	cursor, err := bd.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error(fmt.Sprintf("Error trying to find bids by auctionId %s", auctionId), err)
		return nil, internal_error.NewInternalServerError(fmt.Sprintf("Error trying to find bids by auctionId %s", auctionId))
	}

	var bidEntitiesMongo []BidEntityMongo
	if err := cursor.All(ctx, &bidEntitiesMongo); err != nil {
		logger.Error(fmt.Sprintf("Error trying to find bids by auctionId %s", auctionId), err)
		return nil, internal_error.NewInternalServerError(fmt.Sprintf("Error trying to find bids by auctionId %s", auctionId))

	}
	var bidEntities []bid_entity.Bid
	for _, bidEntityMongo := range bidEntitiesMongo {
		bidEntities = append(bidEntities, bid_entity.Bid{
			Id:        bidEntityMongo.Id,
			UserId:    bidEntityMongo.UserId,
			AuctionId: bidEntityMongo.AuctionId,
			Amount:    bidEntityMongo.Amount,
			Timestamp: time.Unix(bidEntityMongo.Timestamp, 0),
		})
	}
	return bidEntities, nil
}

// Parte do objeto BidRepository
func (bd *BidRepository) FindWinningBidByAuctionId(
	ctx context.Context, auctionId string) (*bid_entity.Bid, *internal_error.InternalError) {

	filter := bson.M{"auctionId": auctionId}

	var bidEntityMongo BidEntityMongo

	// Para retornar o maior valor
	opts := options.FindOne().SetSort(bson.D{{Key: "amount", Value: -1}})

	if err := bd.Collection.FindOne(ctx, filter, opts).Decode(&bidEntityMongo); err != nil {
		logger.Error("Error trying to find the auction winner", err)
		return nil, internal_error.NewInternalServerError("Error trying to find the auction winner")
	}

	return &bid_entity.Bid{
		Id:        bidEntityMongo.Id,
		UserId:    bidEntityMongo.UserId,
		AuctionId: bidEntityMongo.AuctionId,
		Amount:    bidEntityMongo.Amount,
		Timestamp: time.Unix(bidEntityMongo.Timestamp, 0),
	}, nil

}
