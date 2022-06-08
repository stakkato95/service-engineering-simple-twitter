package domain

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/stakkato95/service-engineering-go-lib/logger"
	"github.com/stakkato95/twitter-service-tweets/config"
	"github.com/stakkato95/twitter-service-tweets/dto"
)

type TweetsSink interface {
	AddTweet(dto.TweetDto) error
}

const partition = 0

type kafkaTweetsSink struct {
	conn *kafka.Conn
}

func NewTweetsSink() TweetsSink {
	conn, err := kafka.DialLeader(context.Background(), "tcp", config.AppConfig.KafkaService, config.AppConfig.KafkaTopic, partition)
	if err != nil {
		logger.Fatal("failed to dial leader: " + err.Error())
	}

	return &kafkaTweetsSink{conn}
}

func (s *kafkaTweetsSink) AddTweet(tweet dto.TweetDto) error {
	w := new(bytes.Buffer)

	if err := json.NewEncoder(w).Encode(tweet); err != nil {
		logger.Fatal("can not encode tweet struct: " + err.Error())
	}

	s.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if bytesWritten, err := s.conn.Write(w.Bytes()); err != nil {
		logger.Error("failed to write messages: " + err.Error())
		return err
	} else {
		logger.Info(fmt.Sprintf("writter user: %s, written bytes: %d", w.String(), bytesWritten))
		return nil
	}
}
