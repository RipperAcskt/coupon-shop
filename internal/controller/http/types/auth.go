package types

import (
	"fmt"
	"strings"
)

type (
	AuthUserRequest struct {
		Channel  string `json:"channel" validate:"required"`
		Resource string `json:"resource" validate:"required"`
	}

	TokenResponse struct {
		Token string `json:"token"`
	}

	Channel string
)

const (
	Phone Channel = "phone"
	Email Channel = "email"
	Code  Channel = "code"
)

func (c Channel) String() string {
	return string(c)
}

func (c Channel) IsEmail() bool {
	return c == Email
}

func (c Channel) IsPhone() bool {
	return c == Phone
}

func (c Channel) IsCode() bool {
	return c == Code
}

var channels = []Channel{
	Phone,
	Email,
	Code,
}

func ResolveByChannel(ch string) (*Channel, error) {
	for i, channel := range channels {
		if strings.EqualFold(channel.String(), ch) {
			return &channels[i], nil
		}
	}

	return nil, fmt.Errorf("unknown channel: %s", ch)
}
