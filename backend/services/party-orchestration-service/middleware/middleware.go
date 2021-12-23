package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/toeydevelopment/microservices-example/party-orchestration-service/constant"
)

func NewAuthMiddleware(authHost string, hc *http.Client) gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		splited := strings.Split(token, " ")

		if len(splited) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "bearer token incorrect format",
			})
			return
		}

		if splited[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "bearer token incorrect format should start with Bearer",
			})
			return
		}

		type verify struct {
			Token string `json:"token"`
		}

		v := &verify{
			Token: splited[1],
		}

		by, _ := json.Marshal(v)

		req, err := http.NewRequestWithContext(
			c.Request.Context(),
			http.MethodPost,
			fmt.Sprintf("%s/auth/verify", authHost),
			bytes.NewBuffer(by),
		)

		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		res, err := hc.Do(req)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if res.StatusCode != 200 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		defer res.Body.Close()

		type Response struct {
			UserEmail string `json:"user_email"`
		}

		var r Response

		b, err := ioutil.ReadAll(res.Body)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := json.Unmarshal(b, &r); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Set(constant.UserEmail, r.UserEmail)

		c.Next()
	}
}
