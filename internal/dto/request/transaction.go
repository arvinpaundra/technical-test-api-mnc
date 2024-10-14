package request

type (
	Topup struct {
		Amount float64 `json:"amount" validate:"required"`
	}

	Payment struct {
		Amount  float64 `json:"amount" validate:"required"`
		Remarks string  `json:"remarks" validate:"required"`
	}

	Transfer struct {
		TargetUser string  `json:"target_user" validate:"required"`
		Amount     float64 `json:"amount" validate:"required"`
		Remarks    string  `json:"remarks" validate:"required"`
	}
)
