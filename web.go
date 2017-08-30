package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	r = gin.Default()
)

func pong(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
