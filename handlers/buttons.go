package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	comp "github.com/MathisVerstrepen/templ_ui_library/components"
	"github.com/MathisVerstrepen/templ_ui_library/models"
	"github.com/MathisVerstrepen/templ_ui_library/ui"
)

func Buttons(c echo.Context) error {
	return Render(c, http.StatusOK, comp.Root(comp.Buttons(), "Buttons"))
}

func ButtonsChange(c echo.Context) error {
	exampleDecorator := ui.ButtonExampleDecorator()

	buttonData := models.Button{
		Text:              "Button",
		Variant:           c.FormValue("variant"),
		Color:             c.FormValue("color"),
		Size:              c.FormValue("size"),
		Decorator:         exampleDecorator,
		DecoratorPosition: c.FormValue("decorator"),
	}

	Render(c, http.StatusOK, comp.CodeBlockWithCopy("templ", comp.ButtonGetCode(buttonData)))

	return Render(c, http.StatusOK, ui.Button(buttonData))
}
