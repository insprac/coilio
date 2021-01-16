package portfolio

import (
	"io/ioutil"
	"os"
	"testing"
)

const testYAMLData string = `
coins:
  - symbol: BTC
    colour: '#f6921a'
    quantity: 3.562
  - symbol: ADA
    colour: '#3bc6c6'
    quantity: 64000
`

func assertEquals(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("Equal assertion failed: %v != %v", a, b)
	}
}

func assertNilError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("Error wasn't nil: %v", err)
	}
}

func assertPortfolio(t *testing.T, portfolio Portfolio) {
	assertEquals(t, len(portfolio.Coins), 2)
	assertEquals(t, portfolio.Coins[0].Symbol, "BTC")
	assertEquals(t, portfolio.Coins[0].Colour, "#f6921a")
	assertEquals(t, portfolio.Coins[0].Quantity, 3.562)
	assertEquals(t, portfolio.Coins[1].Symbol, "ADA")
	assertEquals(t, portfolio.Coins[1].Colour, "#3bc6c6")
	assertEquals(t, portfolio.Coins[1].Quantity, 64000.0)
}

func TestFromBytes(t *testing.T) {
	portfolio, err := FromBytes([]byte(testYAMLData))

	assertNilError(t, err)
	assertPortfolio(t, portfolio)
}

func TestFromFile(t *testing.T) {
	file, err := ioutil.TempFile("", "portfolio.*.yaml")

	assertNilError(t, err)

	defer os.Remove(file.Name())

	_, err = file.Write([]byte(testYAMLData))

	assertNilError(t, err)

	portfolio, err := FromFile(file.Name())

	assertNilError(t, err)
	assertPortfolio(t, portfolio)
}
