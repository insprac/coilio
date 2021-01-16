package portfolio

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Portfolio struct {
	Coins []PortfolioCoin `yaml:"coins"`
}

type PortfolioCoin struct {
	Symbol   string  `yaml:"symbol"`
	Colour   string  `yaml:"colour"`
	Quantity float64 `yaml:"quantity"`
}

func FromFile(path string) (Portfolio, error) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return Portfolio{}, err
	}

	return FromBytes(data)
}

func FromBytes(data []byte) (Portfolio, error) {
	var portfolio Portfolio
	err := yaml.Unmarshal(data, &portfolio)
	return portfolio, err
}
