package entity

type AuthConfig struct {
	Secret            string
	MinutesExp        int
	RefreshMinutesExp int
}
