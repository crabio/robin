package main

import (
	// External
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"

	// Internal
	"github.com/iakrevetkho/robin/internal/config"
)

func RegisterRoutes(config config.Config, nc *nats.Conn, r *gin.Engine) {
	r.GET("/login", func(c *gin.Context) {
		LoginHandler(config, nc, c)
	})
	r.GET("/auth", func(c *gin.Context) {
		AuthHandler(config, nc, c)
	})
}
