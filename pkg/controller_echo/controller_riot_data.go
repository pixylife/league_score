package controller_echo

import (
	"github.com/labstack/echo/v4"
	"league_score/pkg/api_echo/match"
	"net/http"
)

func GetMatches(c echo.Context) error {
	result, err := match.GetMatchData(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func APIControllerMatch(g *echo.Group) {
	g.GET("api/match", GetMatches)
}
