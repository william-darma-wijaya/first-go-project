package model

type UserEvent struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *UserEvent) GetId() string {
	return u.Id
}