package bid_usecase

import (
	"time"

	"github.com/wandermaia/aulas-pos-golang/lab-leilao/internal/entity/bid_entity"
)

type BidOutputDTO struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	AuctionId string    `json:"auction_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type BidUseCase struct {
	BidRepository bid_entity.BidEntityRepository
}
