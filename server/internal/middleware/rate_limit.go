package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	memory "github.com/ulule/limiter/v3/drivers/store/memory"
	ginLimiter "github.com/ulule/limiter/v3/drivers/middleware/gin"
)

func RateLimiter() gin.HandlerFunc {
	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit: 60,
	}

	store := memory.NewStore()

	instance := limiter.New(store,rate)

	return ginLimiter.NewMiddleware(instance)
}