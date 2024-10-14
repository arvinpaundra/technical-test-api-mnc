package response

type (
	UpdateProfile struct {
		UserId      string `json:"user_id"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Address     string `json:"address"`
		UpdatedDate string `json:"updated_date"`
	}
)
