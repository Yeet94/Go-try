package tools

import (
	"strings"
	"time"
	"log"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"John": {AuthToken: "token123", Username: "john_doe"},
	"Jane": {AuthToken: "token456", Username: "jane_doe"},
	"Mike": {AuthToken: "token789", Username: "mike_smith"},
}

var mockCoinDetails = map[string]CoinDetails{
	"BTC": {Coins: 15.0, Username: "john_doe"},
	"ETH": {Coins: 10.0, Username: "jane_doe"},
	"ADA": {Coins: 100.0, Username: "mike_smith"},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	log.Printf("Searching for username: %s", username) // Log the username being searched
	for key, clientData := range mockLoginDetails {
		if strings.EqualFold(key, username) { // Perform case-insensitive comparison
			return &clientData
		}
	}

	log.Printf("Username not found: %s", username) // Log if username is not found
	return nil
}

func (d *mockDB) GetUserCoin(username string) *CoinDetails {
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
