package user

// User is representing the user data struct
type User struct {
	ID   string
	Name string
}

type GetProfileRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetProfileResponse struct {
	User User `json:"user"`
}

type UpdateProfileRequest struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UpdateProfileResponse struct {
	User User `json:"user"`
}

type DeleteAccountRequest struct {
	ID string `json:"id" validate:"required"`
}

type DeleteAccountResponse struct {
	User User `json:"user"`
}
