package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	comp "diikstra.fr/template/components"
	"diikstra.fr/template/models"
	"diikstra.fr/template/ui"
)

func Buttons(c echo.Context) error {
	return Render(c, http.StatusOK, comp.Root(comp.Buttons(), "Home"))
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
	fmt.Printf("Button Data: %+v\n", buttonData)

	Render(c, http.StatusOK, comp.CodeBlockWithCopy("templ", comp.GetCode(buttonData)))

	return Render(c, http.StatusOK, ui.Button(buttonData))
}
