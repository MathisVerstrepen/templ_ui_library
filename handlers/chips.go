package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	comp "github.com/MathisVerstrepen/templ_ui_library/components"
	"github.com/MathisVerstrepen/templ_ui_library/models"
	"github.com/MathisVerstrepen/templ_ui_library/ui"
)

func Chips(c echo.Context) error {
	return Render(c, http.StatusOK, comp.Root(comp.Chips(), "Chips"))
}

func ChipsChange(c echo.Context) error {
	chipData := models.Chip{
		Text:    "Chip",
		Variant: c.FormValue("variant"),
		Color:   c.FormValue("color"),
		Size:    c.FormValue("size"),
	}

	Render(c, http.StatusOK, comp.CodeBlockWithCopy("templ", comp.ChipGetCode(chipData)))

	return Render(c, http.StatusOK, ui.Chip(chipData))
}
