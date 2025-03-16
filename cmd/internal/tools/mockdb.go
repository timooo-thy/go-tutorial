package tools

import "time"

type mockDB struct {}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "1234",
		Username: "alex",
	},
	"bob": {
		AuthToken: "5678",
		Username: "bob",
	},
	"charlie": {
		AuthToken: "9012",
		Username: "charlie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins: 100,
		Username: "alex",
	},
	"bob": {
		Coins: 200,
		Username: "bob",
	},
	"charlie": {
		Coins: 300,
		Username: "charlie",
	},
}

func (db *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep((time.Second * 2))

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep((time.Second * 2))

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDB) SetupDatabase() error {
	return nil
}