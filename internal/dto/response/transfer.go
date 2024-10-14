package response

type (
	Topup struct {
		TopupId       string  `json:"top_up_id"`
		AmountTopup   float64 `json:"amount_top_up"`
		BalanceBefore float64 `json:"balance_before"`
		BalanceAfter  float64 `json:"balance_after"`
		CreatedDate   string  `json:"created_date"`
	}

	Payment struct {
		PaymentId     string  `json:"payment_id"`
		Amount        float64 `json:"amount"`
		Remarks       string  `json:"remarks"`
		BalanceBefore float64 `json:"balance_before"`
		BalanceAfter  float64 `json:"balance_after"`
		CreatedDate   string  `json:"created_date"`
	}

	Transfer struct {
		TransferId    string  `json:"transfer_id"`
		Amount        float64 `json:"amount"`
		Remarks       string  `json:"remarks"`
		BalanceBefore float64 `json:"balance_before"`
		BalanceAfter  float64 `json:"balance_after"`
		CreatedDate   string  `json:"created_date"`
	}

	Transaction struct {
		TransferId      string  `json:"transfer_id"`
		Status          string  `json:"status"`
		UserId          string  `json:"user_id"`
		TransactionType string  `json:"transaction_type"`
		Amount          float64 `json:"amount"`
		Remarks         string  `json:"remarks"`
		BalanceBefore   float64 `json:"balance_before"`
		BalanceAfter    float64 `json:"balance_after"`
		CreatedDate     string  `json:"created_date"`
	}
)
