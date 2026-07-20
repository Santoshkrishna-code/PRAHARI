package adapter

import "time"

// Adapter represents translation driver configuration metadata.
type Adapter struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Protocol  string    `json:"protocol"` // MQTT, AMQP, OPCUA, SOAP, REST
	Version   string    `json:"version"`
	CreatedAt time.Time `json:"created_at"`
}
