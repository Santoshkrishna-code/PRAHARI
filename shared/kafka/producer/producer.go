package producer

import (
	"github.com/segmentio/kafka-go"
	prahariKafka "prahari/shared/kafka"
)

// Producer wraps segmentio/kafka-go Writer to publish messages.
type Producer struct {
	writer *kafka.Writer
}

// NewProducer constructs a Kafka message publisher.
func NewProducer(cfg prahariKafka.Config) (*Producer, error) {
	dialer, err := prahariKafka.GetDialer(cfg)
	if err != nil {
		return nil, err
	}

	transport := &kafka.Transport{
		Dial:   dialer.DialFunc,
		SASL:   dialer.SASLMechanism,
		TLS:    dialer.TLS,
	}

	writer := &kafka.Writer{
		Addr:         kafka.TCP(cfg.Brokers...),
		Balancer:     &kafka.LeastBytes{},
		Transport:    transport,
		MaxAttempts:  3,
		RequiredAcks: kafka.RequireAll, // Enforce strong consistency in production
	}

	return &Producer{
		writer: writer,
	}, nil
}

// Close releases writer connections.
func (p *Producer) Close() error {
	if p.writer == nil {
		return nil
	}
	return p.writer.Close()
}
