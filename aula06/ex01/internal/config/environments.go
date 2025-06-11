package config

import "os"

type app struct {
	SecretKey string
}

var bootcampApp *app

func New() {
	if bootcampApp == nil {
		bootcampApp = &app{
			SecretKey: os.Getenv("SECRET_KEY"),
		}
	}
}

func GetSecretKey() string {
	return bootcampApp.SecretKey
}
