package auction_entity

import "time"

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
