package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
)

var (
	lastReq = uint64(time.Now().UnixMilli())
	cPeriod = uint64(60 * 60 * 1000)

	reqList  = make(map[string]uint64)
	rateList = make(map[string]uint64)
)

func RateLimit(route string, delay uint64) gin.HandlerFunc {
	rateList[route] = delay

	return func(ctx *gin.Context) {
		reqUrl := (ctx.Request.Method + ":" + ctx.Request.RequestURI)
		dateNow := uint64(time.Now().UnixMilli())

		// clear memory.
		if (dateNow - lastReq) > cPeriod {
			reqList = make(map[string]uint64)

		}

		lastReq = dateNow

		if value, key := reqList[reqUrl+":"+ctx.ClientIP()]; key && dateNow < value {
			ctx.AbortWithStatusJSON(429, gin.H{
				"message": "429 - Too Many Requests.",
			})

			return
		}

		reqList[reqUrl+":"+ctx.ClientIP()] = (dateNow + rateList[reqUrl])
		ctx.Next()
	}

}
