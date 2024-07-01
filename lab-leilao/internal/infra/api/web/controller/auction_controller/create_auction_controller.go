package auction_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wandermaia/aulas-pos-golang/lab-leilao/configuration/rest_err.go"
	"github.com/wandermaia/aulas-pos-golang/lab-leilao/internal/infra/api/web/validation"
	"github.com/wandermaia/aulas-pos-golang/lab-leilao/internal/usecase/auction_usecase"
)

type AuctionController struct {
	auctionUseCase auction_usecase.AuctionUseCaseInterface
}

// Função Construtora
func NewAuctionController(auctionUseCase auction_usecase.AuctionUseCaseInterface) *AuctionController {

	return &AuctionController{
		auctionUseCase: auctionUseCase,
	}
}

func (u *AuctionController) CreateAuction(c *gin.Context) {
	var auctionInputDTO auction_usecase.AuctionIntputDTO

	if err := c.ShouldBindJSON(&auctionInputDTO); err != nil {
		restErr := validation.ValidateErr(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	err := u.auctionUseCase.CreateAuction(context.Background(), auctionInputDTO)
	if err != nil {
		restErr := rest_err.ConvertError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusCreated)
}
