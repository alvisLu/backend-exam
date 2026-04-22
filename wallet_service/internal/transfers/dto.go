package transfers

import "github.com/alvis/wallet_service/internal/accounts"

type TransferRequest struct {
	FromID string `json:"from_id"`
	ToID   string `json:"to_id"`
	Amount string `json:"amount"`
}

type accountSummary struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Balance string `json:"balance"`
}

type TransferResponse struct {
	From   accountSummary `json:"from"`
	To     accountSummary `json:"to"`
	Amount string         `json:"amount"`
}

func toSummary(a *accounts.Account) accountSummary {
	return accountSummary{
		ID:      a.ID.String(),
		Name:    a.Name,
		Balance: a.Balance.StringFixed(8),
	}
}

func toTransferResponse(r *Result) TransferResponse {
	return TransferResponse{
		From:   toSummary(r.From),
		To:     toSummary(r.To),
		Amount: r.Amount.String(),
	}
}
