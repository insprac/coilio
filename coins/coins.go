package main

import (
	"errors"
	"fmt"
	gecko "github.com/insprac/go-coingecko"
)

var coinsCache []gecko.ListedCoin

func CurrentPrice(symbol string) (float64, err) {
	id, err := IDFromSymbol(symbol)

	if err != nil {
		return 0, err
	}

	params = gecko.GetCoinParams{
		Localization:  false,
		Tickers:       false,
		MarketData:    true,
		CommunityData: false,
		DeveloperData: false,
		Sparkline:     false,
	}

	coin, err := gecko.GetCoin(params)

	if err != nil {
		return 0, err
	}

	return coin.MarketData.CurrentPrice["usd"], nil
}

func IDFromSymbol(symbol string) (string, err) {
	coins, err := ListCoins()

	if err != nil {
		return "", err
	}

	for coin := range coins {
		if coin.Symbol == symbol {
			return coin.ID, nil
		}
	}

	return "", newError("coin doesn't exist with symbol '%v'", symbol)
}

func ListCoins() ([]gecko.ListedCoin, error) {
	if len(coinsCache) > 0 {
		return coinsCache, nil
	}

	coins, err := gecko.ListCoins()

	if err != nil {
		return coinsCache, nil
	}

	coinsCache = coins
	return coinsCache, nil
}

func newError(message string, args ...interface{}) error {
	return errors.New("coins: " + fmt.Sprintf(message, args...))
}
