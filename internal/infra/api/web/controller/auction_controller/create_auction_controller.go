package auction_controller

import (
	"context"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcelockdata/go-leilao/configuration/rest_error"
	"github.com/marcelockdata/go-leilao/internal/infra/api/web/validation"
	"github.com/marcelockdata/go-leilao/internal/usecase/auction_usecase"
)

type AuctionController struct {
	auctionUseCase auction_usecase.AuctionUseCaseInterface
}

func NewAuctionController(auctionUseCase auction_usecase.AuctionUseCaseInterface) *AuctionController {
	return &AuctionController{
		auctionUseCase: auctionUseCase,
	}
}

func (u *AuctionController) CreateAuction(c *gin.Context) {
	var auctionInputDTO auction_usecase.AuctionInputDTO

	if err := c.ShouldBindJSON(&auctionInputDTO); err != nil {
		restErr := validation.ValidateErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	err := u.auctionUseCase.CreateAuction(context.Background(), auctionInputDTO)
	if err != nil {
		restErr := rest_error.ConvertError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusCreated)
}
