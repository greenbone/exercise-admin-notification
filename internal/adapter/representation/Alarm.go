package representation

type Alarm struct {
	Level                string `json:"level"`
	EmployeeAbbreviation string `json:"employeeAbbreviation"`
	Message              string `json:"message"`
}

func (u *Alarm) OK() error {
	if len(u.Level) == 0 {
		return MissingFieldError("level")
	}
	if len(u.EmployeeAbbreviation) == 0 {
		return MissingFieldError("employeeAbbreviation")
	}
	if len(u.Message) == 0 {
		return MissingFieldError("message")
	}
	return nil
}

func (e MissingFieldError) Error() string {
	return string(e) + " is required"
}
