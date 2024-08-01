package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	comp "diikstra.fr/template/components"
)

func HomeHandler(c echo.Context) error {
	return Render(c, http.StatusOK, comp.Root(comp.Home(), "Home"))
}
