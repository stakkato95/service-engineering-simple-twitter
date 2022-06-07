package domain

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/stakkato95/service-engineering-go-lib/logger"
	"github.com/stakkato95/twitter-service-analytics/config"
)

type TweetProcessor interface {
	GetTweetCount() int
	Destroy() error
}

const partition = 0
const msgBufferSize = 10e3 //10KB

type kafkaTweetProcessor struct {
	conn       *kafka.Conn
	tweetCount int
}

func NewTweetProcessor() TweetProcessor {
	kafkaService := config.AppConfig.KafkaService
	topic := config.AppConfig.KafkaTopic
	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaService, topic, partition)
	if err != nil {
		logger.Fatal("failed to dial leader: " + err.Error())
	}

	repo := kafkaTweetProcessor{conn: conn}

	go func() {
		for {
			//zwischen timeout und anderen Fehlern unterscheiden
			conn.SetReadDeadline(time.Now().Add(60 * time.Hour))
			if msg, err := conn.ReadMessage(msgBufferSize); err != nil {
				logger.Error("error when reading a msg from kafka: " + err.Error())
			} else {
				// logger.Info(fmt.Sprintf("topic: %s, offset: %d, value: %s", msg.Topic, msg.Offset, string(msg.Value[:])))
				repo.tweetCount += 1

				var tweet TweetDto
				json.NewDecoder(bytes.NewReader(msg.Value)).Decode(&tweet)
				logger.Info(fmt.Sprintf("count: %d, read tweet: %v", repo.tweetCount, tweet))
			}
		}
	}()

	return &repo
}

func (k *kafkaTweetProcessor) GetTweetCount() int {
	return k.tweetCount
}

func (k *kafkaTweetProcessor) Destroy() error {
	if err := k.conn.Close(); err != nil {
		logger.Fatal("failed to close writer: " + err.Error())
	}

	return nil
}
