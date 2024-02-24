package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"anshu": {
		AuthToken: "meena",
		Username:  "Anshu",
	},
	"aman": {
		AuthToken: "meena",
		Username:  "Aman",
	},
	"anuj": {
		AuthToken: "gupta",
		Username:  "Anuj",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"anshu": {
		Coins:    100,
		Username: "Anshu",
	},
	"aman": {
		Coins:    200,
		Username: "Aman",
	},
	"anuj": {
		Coins:    300,
		Username: "Anuj",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
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
