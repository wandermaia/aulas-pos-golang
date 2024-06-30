package auction_usecase

import (
	"context"
	"time"

	"github.com/wandermaia/aulas-pos-golang/lab-leilao/internal/entity/auction_entity"
	"github.com/wandermaia/aulas-pos-golang/lab-leilao/internal/entity/bid_entity"
	"github.com/wandermaia/aulas-pos-golang/lab-leilao/internal/internal_error"
	"github.com/wandermaia/aulas-pos-golang/lab-leilao/internal/usecase/bid_usecase"
)

type ProductCondition int64
type AuctionStatus int64

type AuctionIntputDTO struct {
	ProductName string           `json:"product_name" binding:"required,min=1"` // essas tags são utilizadas no validation: internal/infra/api/web/validation
	Category    string           `json:"category" binding:"required,min=2"`
	Description string           `json:"description"  binding:"required,min=10,max=200"`
	Condition   ProductCondition `json:"condition"` // binding:"oneOf" Um dos pré-configurados
}

type AuctionOutputDTO struct {
	Id          string           `json:"id"`
	ProductName string           `json:"product_name"`
	Category    string           `json:"category"`
	Description string           `json:"description"`
	Condition   ProductCondition `json:"condition"`
	Status      AuctionStatus    `json:"status"`
	Timestamp   time.Time        `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type WinningInfoOutputDTO struct {
	Auction AuctionOutputDTO          `json:"auction"`
	Bid     *bid_usecase.BidOutputDTO `json:"bid,omitempty"`
}

// Possui os métodos CreateAuction, FindAuctionById e FindAuction
type AuctionUseCase struct {

	//Essa interface implementa os métodos de acesso dos repositórios
	auctionRepositoryInterface auction_entity.AuctionRepositoryInterface
	bidRepositoryInterface     bid_entity.BidEntityRepository
}

// Parte do objeto AuctionUseCase
func (au AuctionUseCase) CreateAuction(
	ctx context.Context,
	auctionIntput AuctionIntputDTO) *internal_error.InternalError {
	auction, err := auction_entity.CreateAuction(
		auctionIntput.ProductName,
		auctionIntput.Category,
		auctionIntput.Description,
		auction_entity.ProductCondition(auctionIntput.Condition),
	)
	if err != nil {
		return err
	}

	if err := au.auctionRepositoryInterface.CreateAuction(ctx, *auction); err != nil {
		return err
	}

	return nil

}
