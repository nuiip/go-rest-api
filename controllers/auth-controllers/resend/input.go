package resend

type InputResend struct {
	Email string `json:"email" validate:"required,email"`
}
