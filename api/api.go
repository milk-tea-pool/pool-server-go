package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"milkteapool.com/pool-server/pool"
)

type Controller struct {
	Pool *pool.Pool
}

type PoolInfoResponse struct {
	Name               string  `json:"name"`
	LogoURL            string  `json:"logoURL"`
	MinimumDifficulty  uint64  `json:"minimumDifficulty"`
	RelativeLockHeight uint32  `json:"relativeLockHeight"`
	ProtocolVersion    string  `json:"protocolVersion"`
	Dee                float32 `json:"fee"`
	Description        string  `json:"description"`
	TargetPuzzleHash   []byte  `json:"targetPuzzleHash"`
}

func (c *Controller) PoolInfo(ctx *gin.Context) {

	r := PoolInfoResponse{
		"Milk Tea Pool",
		"https://milkteapool.com/assets/img/miningstats-banner.png",
		c.Pool.MinDifficulty,
		c.Pool.RelativeLockHeight,
		"1.0.0",
		c.Pool.PoolFee,
		"(example) The Reference Pool allows you to pool with low fees, paying out daily using Chia.",
		c.Pool.DefaultTargetPuzzleHash,
	}

	ctx.JSON(http.StatusOK, r)
}

func (c *Controller) SubmitPartial(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "",
	})
}
