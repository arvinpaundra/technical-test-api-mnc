package request

type (
	UpdateProfile struct {
		FirstName string `json:"first_name" validate:"required,min=3"`
		LastName  string `json:"last_name"`
		Address   string `json:"address" validate:"required,max=100"`
	}
)
