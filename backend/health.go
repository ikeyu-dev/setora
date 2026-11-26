package main

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type HealthResponse struct {
	Status string `json:"status"`
}

// health はサーバーの稼働状態を返します。
// @Summary ヘルスチェック
// @Tags health
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func health(c *echo.Context) error {
	return c.JSON(http.StatusOK, HealthResponse{Status: "ok"})
}
