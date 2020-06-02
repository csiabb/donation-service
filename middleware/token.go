package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/csiabb/donation-service/common/rest"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) TokenCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(rest.AuthToken)
		if token == "" {
			c.JSON(http.StatusOK, rest.ErrorResponse(rest.TokenMissingCode, rest.TokenMsgMissing))
			c.Abort()
			return
		}

		id, err := m.ctx.Redis.Get(context.Background(), token).Result()
		if err != nil {
			e := fmt.Errorf("middleware check token error, %v", err)
			logger.Error(e)
			c.JSON(http.StatusBadRequest, rest.ErrorResponse(rest.TokenInvalidCode, e.Error()))
			return
		}

		if id == "" {
			c.JSON(http.StatusOK, rest.ErrorResponse(rest.TokenInvalidCode, rest.TokenMsgMissing))
			c.Abort()
			return
		}

		c.Request.Header.Set(rest.AuthUID, id)
	}
}
