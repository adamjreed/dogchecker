package handlers

import (
	"context"
	"dogchecker/internal/app"
	"dogchecker/pkg/mailer"
	"dogchecker/pkg/petfinder"
	"encoding/json"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/go-redis/redis/v8"
	"log"
)

func CheckForDogs(ctx context.Context) error {
	appContext := ctx.Value(app.AppContext).(*app.Context)

	err := appContext.Client.Authenticate()
	if err != nil {
		return errors.Wrap(err, "authenticating with petfinder")
	}

	dogs, err := appContext.Client.GetDogs()
	if err != nil {
		return errors.Wrap(err, "getting dogs from petfinder")
	}

	var dogTemplates []*mailer.DogTemplate
	for _, dog := range dogs {
		jsonDog, err := json.Marshal(dog)
		if err != nil {
			return errors.Wrap(err, "could not marshal dog to json")
		}

		res, err := appContext.Cache.SetArgs(ctx, fmt.Sprintf("%s%d", petfinder.DogsPrefix, dog.Id), jsonDog, redis.SetArgs{
			Get: true,
		}).Result()
		if err == redis.Nil {
			log.Println(fmt.Sprintf("new dog found: %d", dog.Id))
			dogTemplate := &mailer.DogTemplate{
				Name: dog.Name,
				Link: dog.URL,
			}

			if len(dog.Photos) > 0 {
				dogTemplate.Photo = dog.Photos[0].Medium
			}

			dogTemplates = append(dogTemplates, dogTemplate)
		} else if err != nil {
			return errors.Wrap(err, "redis error")
		} else {
			log.Println(fmt.Sprintf("existing dog found. id: %d, cached record: %s", dog.Id, res))
		}
	}

	if len(dogTemplates) > 0 {
		content, err := mailer.BuildHTML(dogTemplates)
		if err != nil {
			return errors.Wrap(err, "generating dogs email content")
		}

		err = appContext.Mailer.Send("New Dogs Listed", content)
		if err != nil {
			return errors.Wrap(err, "sending email")
		}
	} else {
		log.Println("no new dogs found")
	}

	return nil
}
