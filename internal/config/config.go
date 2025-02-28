package config

import "errors"

type Config struct {
	DbURL           string `json:"db_URL"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(user string) error {
	c.CurrentUserName = user
	err := write(c)
	if err != nil {
		return errors.New("failed to write changes to config file")
	}
	return nil
}
