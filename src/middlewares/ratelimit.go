package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
)

var (
	lastReq = uint64(time.Now().UnixMilli())
	cPeriod = uint64(60 * 60 * 1000)

	reqList = make(map[string]uint64)
)

func RateLimit(delay uint64) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		reqUrl := ctx.Request.URL.String() + ctx.ClientIP()
		dateNow := uint64(time.Now().UnixMilli())

		// clear memory.
		if (dateNow - lastReq) > cPeriod {
			reqList = make(map[string]uint64)

		}

		lastReq = dateNow

		if value, key := reqList[reqUrl]; key && dateNow < value {
			ctx.AbortWithStatusJSON(429, gin.H{
				"message": "429 - Too Many Requests.",
			})

			return
		}

		reqList[reqUrl] = (dateNow + delay)
		ctx.Next()
	}

}