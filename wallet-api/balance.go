package walletapi

import "encoding/json"

// Balance - represents a wallet balance
type Balance struct {
	Unlocked float64 `json:"unlocked"`
	Locked   float64 `json:"locked"`
	Address  string `json:"address"`
}

// GetBalance - gets the balance for the entire wallet container
func (wAPI WalletAPI) GetBalance() (unlocked, locked float64, err error) {
	resp, _, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/balance",
		"",
	)
	if err == nil {
		unlocked = float64((*resp)["unlocked"].(float64))
		locked = float64((*resp)["locked"].(float64))
	}

	return unlocked, locked, err
}

// GetAddressBalance - gets the balance for the specified wallet address
func (wAPI WalletAPI) GetAddressBalance(address string) (bal *Balance, err error) {
	bal = &Balance{}
	resp, _, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/balance/"+address,
		"",
	)
	if err == nil {
		(*bal).Unlocked = float64((*resp)["unlocked"].(float64))
		(*bal).Locked = float64((*resp)["locked"].(float64))
		(*bal).Address = address
	}

	return bal, err
}

// GetBalances - gets the balance for the every wallet address
func (wAPI WalletAPI) GetBalances() ([]Balance, error) {
	balances := []Balance{}
	_, raw, err := wAPI.sendRequest(
		"GET",
		wAPI.Host+":"+wAPI.Port+"/balances",
		"",
	)
	if err == nil {
		json.Unmarshal(*raw, &balances)
	}

	return balances, err
}
