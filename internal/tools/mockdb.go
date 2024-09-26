package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"darius": {
		AuthToken: "123ABC",
		Username:  "darius",
	}, "jack": {
		AuthToken: "123ABC",
		Username:  "jack",
	}, "beniamin": {
		AuthToken: "123ABC",
		Username:  "beniamin",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"darius": {
		Coins:    500,
		Username: "darius",
	}, "jack": {
		Coins:    500,
		Username: "jack",
	}, "beniamin": {
		Coins:    500,
		Username: "beniamin",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate db call
	time.Sleep(time.Second * 1)

	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}
func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate db call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
