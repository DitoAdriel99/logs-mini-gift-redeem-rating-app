package pubsubx

import (
	"context"
	"encoding/json"
	"fmt"
	"go-learn/entities"
	"log"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
)

func PullMsgs() (*entities.LogsPayload, error) {
	log.Println("Pull Message From Pubsub.........")

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, os.Getenv("PROJECTID"))
	if err != nil {
		return nil, fmt.Errorf("pubsub.NewClient: %w", err)
	}
	defer client.Close()

	sub := client.Subscription(os.Getenv("SUBSID"))
	cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = sub.Receive(cctx, func(_ context.Context, msg *pubsub.Message) {
		log.Println("Receivingggg.........")

		var p entities.LogsPayload
		err := json.Unmarshal(msg.Data, &p)
		if err != nil {
			log.Println("error marshaling", err)
			return
		}
		log.Println("Pull Message From Pubsub Success.........")

		msg.Ack()

	})
	if err != nil {
		return nil, fmt.Errorf("sub.Receive: %w", err)
	}

	return nil, nil
}
