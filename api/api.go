package api

import (
	"encoding/json"
	"net/http"
)

// Coin balance parameters
type CoinBalanceParams struct {
	Username string
}

// Coin balance response
type CoinBalanceResponse struct {
	// Status code
	Code int

	// Current balance
	Ballance int64
}

type Error struct {
	// Error code
	Code int

	// Error message
	Message string
}
