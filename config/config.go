package config

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewAWSSession() (*session.Session, error) {
	return session.NewSession()
}
