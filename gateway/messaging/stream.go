package messaging

import (
	"context"
	"encoding/json"

	"github.com/ahyalfan/go-toolbox-icems/gateway/model"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type Stream[T model.Event] struct {
	client     *redis.Client
	StreamName string
	Log        *logrus.Logger
}

func (s *Stream[T]) GetStream() string {
	return s.StreamName
}

func (s *Stream[T]) Send(ctx context.Context, event T) error {
	value, err := json.Marshal(event)
	if err != nil {
		s.Log.WithError(err).Error("failed to marshal event")
		return err
	}

	message := &redis.XAddArgs{
		Stream: s.GetStream(),
		Values: map[string]any{
			"event_struct": string(value),
		},
	}

	err = s.client.XAdd(ctx, message).Err()
	if err != nil {
		s.Log.WithError(err).Error("failed to produce message")
		return err
	}
	return nil
}
