package response

type (
	Register struct {
		UserId      string `json:"user_id"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		PhoneNumber string `json:"phone_number"`
		Address     string `json:"address"`
		CreatedDate string `json:"created_date"`
	}

	Login struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	Authenticate struct {
		UserId string `json:"user_id"`
	}
)
