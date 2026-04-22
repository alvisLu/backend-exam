package accounts

type CreateAccountRequest struct {
	Name string `json:"name"`
}

type AccountResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Balance string `json:"balance"`
}

func toAccountResponse(a *Account) AccountResponse {
	return AccountResponse{
		ID:      a.ID.String(),
		Name:    a.Name,
		Balance: a.Balance.StringFixed(8),
	}
}
