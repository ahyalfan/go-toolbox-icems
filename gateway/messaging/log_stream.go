package messaging

import (
	"github.com/ahyalfan/go-toolbox-icems/gateway/model"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type LogStream struct {
	Stream[*model.LogAuditTrailEvent]
}

func NewLogStream(
	client *redis.Client,
	log *logrus.Logger,
	streamName string,
) *LogStream {
	return &LogStream{
		Stream: Stream[*model.LogAuditTrailEvent]{
			client:     client,
			StreamName: streamName,
			Log:        log,
		},
	}
}
