package app

import (
	"dogchecker/internal/config"
	"dogchecker/pkg/mailer"
	"dogchecker/pkg/petfinder"
	"dogchecker/pkg/sendgrid"
	"github.com/go-redis/redis/v8"
)

const AppContext string = "appContext"

type Context struct {
	Config *config.Config
	Client *petfinder.Client
	Cache  *redis.Client
	Mailer mailer.Mailer
}

func CreateContext() (*Context, error) {
	conf, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	var mailClient mailer.Mailer
	if conf.Environment == config.EnvProduction {
		mailClient = sendgrid.NewMailer(conf.Mailer)
	} else {
		mailClient = mailer.NewDevMailer(conf.Mailer)
	}

	return &Context{
		Config: conf,
		Client: petfinder.NewClient(conf.Petfinder.BaseUrl, conf.Petfinder.ClientId, conf.Petfinder.ClientSecret),
		Cache: redis.NewClient(&redis.Options{
			Addr:     conf.Redis.BaseUrl,
			Password: conf.Redis.Password,
			DB:       0,
		}),
		Mailer: mailClient,
	}, nil
}
