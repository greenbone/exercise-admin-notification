package representation

type Alarm struct {
	AlarmType string `json:"type"`
	Message   string `json:"message"`
	User      string `json:"user"`
}

func (u *Alarm) OK() error {
	if len(u.AlarmType) == 0 {
		return MissingFieldError("type")
	}
	if len(u.Message) == 0 {
		return MissingFieldError("message")
	}
	if len(u.User) == 0 {
		return MissingFieldError("user")
	}
	return nil
}

func (e MissingFieldError) Error() string {
	return string(e) + " is required"
}
