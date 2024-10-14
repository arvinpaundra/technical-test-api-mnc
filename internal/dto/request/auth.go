package request

type (
	Register struct {
		FirstName   string `json:"first_name" validate:"required,min=3"`
		LastName    string `json:"last_name"`
		PhoneNumber string `json:"phone_number" validate:"required,min=10,max=15"`
		Address     string `json:"address" validate:"required,max=100"`
		PIN         string `json:"pin" validate:"required,min=6,max=6"`
	}

	Login struct {
		PhoneNumber string `json:"phone_number" validate:"required"`
		PIN         string `json:"pin" validate:"required"`
	}
)
