package models

import "github.com/a-h/templ"

const (
	DecoratorPositionNone   = "none"
	DecoratorPositionBefore = "before"
	DecoratorPositionAfter  = "after"
	DecoratorPositionOnly   = "only"
)

type Attributes map[string]string

type Button struct {
	Text              string
	Variant           string
	Color             string
	Size              string
	Decorator         templ.Component
	DecoratorPosition string
	HTMX              Attributes
	HTML              Attributes
}

type Chip struct {
	Text    string
	Variant string
	Color   string
	Size    string
	HTMX    Attributes
	HTML    Attributes
}
