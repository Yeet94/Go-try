package tools

type LoginDetails struct {
	AuthToken string
	Username  string
}

type CoinDetails struct {
	Coins    float64
	Username string
}

var mockLoginDetails = map[string]LoginDetails{
	"John": {AuthToken: "token123", Username: "john_doe"},
	"Jane": {AuthToken: "token456", Username: "jane_doe"},
	"Mike": {AuthToken: "token789", Username: "mike_smith"},
}

var mockCoinDetails = map[string]CoinDetails{
	"BTC": {Coins: 1.5, Username: "john_doe"},
	"ETH": {Coins: 10.0, Username: "jane_doe"},
	"ADA": {Coins: 100.0, Username: "mike_smith"},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second *1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoin(username string) *CoinDetails {
	time.Sleep(time.Second *1)

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