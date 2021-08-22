package representation

type Error struct {
	Message string `json:"message"`
}

type MissingFieldError string
