package auction_entity

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/wandermaia/aulas-pos-golang/lab-leilao/internal/internal_error"
)

// Possui o método Validate
type Auction struct {
	Id          string
	ProductName string
	Category    string
	Description string
	Condition   ProductCondition
	Status      AuctionStatus
	TimeStamp   time.Time
}

type ProductCondition int
type AuctionStatus int

// Ao utilizar o iota, é como se recebesse um valor sequencial
// A Variável Active=0 e O Completed=1
const (
	Active AuctionStatus = iota
	Completed
)

// Semelhante ao caso anterior, mas com três variáveis
// A Variável New=0 , Used=1 e Refurbished=2
const (
	New ProductCondition = iota
	Used
	Refurbished
)

type AuctionRepositoryInterface interface {
	CreateAuction(
		ctx context.Context,
		auctionEntity Auction) *internal_error.InternalError

	FindAuctions(
		ctx context.Context,
		status AuctionStatus,
		category, productName string) ([]Auction, *internal_error.InternalError)

	FindAuctionById(
		ctx context.Context, id string) (*Auction, *internal_error.InternalError)
}

func CreateAuction(
	productName, category, description string,
	condition ProductCondition) (*Auction, *internal_error.InternalError) {
	auction := &Auction{
		Id:          uuid.New().String(),
		ProductName: productName,
		Category:    category,
		Description: description,
		Condition:   condition,
		Status:      Active,
		TimeStamp:   time.Now(),
	}

	if err := auction.Validate(); err != nil {
		return nil, err
	}

	return auction, nil
}

// Parte do objeto Auction
func (au *Auction) Validate() *internal_error.InternalError {

	if len(au.ProductName) <= 1 ||
		len(au.Category) <= 2 ||
		len(au.Description) <= 10 && (au.Condition != New &&
			au.Condition != Refurbished &&
			au.Condition != Used) {
		internal_error.NewBadRequestError("Invalid object Auction")
	}
	return nil
}
