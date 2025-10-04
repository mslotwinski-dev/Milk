package examples

import (
	c "milk/core"
	"milk/core/web"
)

func HelloWorld() {
	page := c.Page(
		c.H1(
			c.Div("Witaj w Milk!").Style(c.FontSize("24px"), c.Color("#00647d")),
		),
		c.P(
			"To jest akapit tekstu wygenerowany w Go przy użyciu Milk.",
		).Style(c.Color("white"), c.BackgroundColor("black"), c.FontSize("18px")),
	)

	page.SetTitle("Moja Pierwsza Strona w Milk")

	web.Run(page, 7250)

}
