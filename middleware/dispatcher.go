package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"go-learn/entities"
	"go-learn/repositories"
	"io/ioutil"
	"log"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

type _Dispatcher struct {
	repo repositories.Repo
}

func NewDispatcher(repo repositories.Repo) *_Dispatcher {
	return &_Dispatcher{repo: repo}
}

func (d *_Dispatcher) StartDispatcher() {

	ctx := context.Background()

	key, err := ioutil.ReadFile(os.Getenv("ACCOUNT_PATH"))
	if err != nil {
		log.Println("Error readfile", err)
		return
	}

	client, err := pubsub.NewClient(ctx, os.Getenv("PROJECTID"), option.WithCredentialsJSON(key))
	if err != nil {
		log.Println("pubsub.NewClient:", err)
		return
	}
	defer client.Close()

	sub := client.Subscription(os.Getenv("SUBSID"))

	for {
		cctx, cancel := context.WithCancel(ctx)

		log.Println("Pull Message From Pubsub.........")

		err = sub.Receive(cctx, func(_ context.Context, msg *pubsub.Message) {
			log.Println("Receivingggg.........")
			var (
				id      = uuid.NewString()
				timeNow = time.Now().Local()
			)

			formatedTime := timeNow.Format("2006-01-02")
			newId := fmt.Sprintf("%s-%s", formatedTime, id)

			var p entities.LogsPayload
			err := json.Unmarshal(msg.Data, &p)
			if err != nil {
				log.Println("error marshaling", err)
				return
			}
			log.Println("Pull Message From Pubsub Success.........")

			p.ID = newId

			msg.Ack()

			d.repo.Logs.Create(&p)
		})
		if err != nil {
			log.Println("err receive", err)
		}

		cancel()

	}
}
