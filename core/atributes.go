package core

import (
	utility "milk/core/utility"
)

type Atributes struct {
	ImageSrc    string
	ImageAlt    string
	ImageWidth  int
	ImageHeight int
}

func (c *Component) Src(src string) *Component {
	if c.Tag == "img" {
		c.Atributes.ImageSrc = src

	} else {
		utility.Error("Src can only be set on img components")
	}
	return c
}

func (c *Component) Alt(alt string) *Component {
	if c.Tag == "img" {
		c.Atributes.ImageAlt = alt

	} else {
		utility.Error("Alt can only be set on img components")
	}
	return c
}

func (c *Component) Width(width int) *Component {
	if c.Tag == "img" {
		c.Atributes.ImageWidth = width

	} else {
		utility.Error("Width can only be set on img components")
	}
	return c
}

func (c *Component) Height(height int) *Component {
	if c.Tag == "img" {
		c.Atributes.ImageHeight = height

	} else {
		utility.Error("Height can only be set on img components")
	}
	return c
}
