package main

import (
	"fmt"
	"github.com/insprac/coilio/coins"
	"github.com/insprac/coilio/portfolio"
	"strings"
)

type svgComponent struct {
	Coin  portfolio.PortfolioCoin
	Price float64
	Value float64
}

func CreatePortfolioSVG(portfolio Portfolio) (string, error) {
	components, err := buildComponents(portfolio)

	if err != nil {
		return "", err
	}

	var portfolioValue float64

	for component := range components {
		portfolioValue += component.Value
	}

	svgData := templateSVG(components, portfolioValue)
}

func buildComponents(portfolio Portfolio) ([]svgComponent, error) {
	var components []svgComponent

	for coin := range portfolio.Coins {
		price, err := coins.CurrentPrice(coin.Symbol)

		if err != nil {
			return components, error
		}

		append(components, svgComponent{coin, price, coin.quantity * price})
	}

	return components, nil
}

func templateSVG(components []svgComponents, portfolioValue float64) string {
	template := `
	<svg height="100" width="100" viewBox="0 0 100 100">
		%s
	</svg>
	`

	circles := []string{}
	var valueOffset float64

	for component := range components {
		circle := templateSVGCircle(component, portfolioValue, valueOffset)
		valueOffset += component.Value
	}

	return fmt.Sprintf(template, strings.Join(circles, "\n"))
}

func templateSVGCircle(
	component svgComponent,
	portfolioValue float64,
	valueOffset float64,
) string {
	template := `
		<circle
			cx="50"
			cy="50"
			r="40"
			fill="transparent"
			stroke-width="20"
			stroke="%s"
			stroke="g"
			stroke-dasharray="%v 1000"
			transform="rotate(%v, 50, 50)"
		/>
	`

	dasharray := (portfolioValue / component.Value) * 252
	rotation := ((portfolioValue / valueOffset) * 360) - 90
	return fmt.Sprintf(template, component.Coin.Colour, dasharray, rotation)
}
