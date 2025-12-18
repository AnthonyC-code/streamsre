// Package kafka provides topic management utilities.
//
// OPTIONAL: You can skip this and create topics manually with rpk.
// But if you want, implement topic creation from Go.
package kafka

// TODO: Import required packages:
// - "context"
// - "github.com/segmentio/kafka-go"

// Topic names as constants
const (
	MainTopic = "events.main"
	DLQTopic  = "events.dlq"
)

// CreateTopic creates a Kafka topic if it doesn't exist.
//
// TODO: Implement (OPTIONAL):
// 1. Connect to broker: conn, _ := kafka.Dial("tcp", broker)
// 2. Get controller: controller, _ := conn.Controller()
// 3. Connect to controller
// 4. Call controllerConn.CreateTopics(kafka.TopicConfig{
//       Topic:             topic,
//       NumPartitions:     partitions,
//       ReplicationFactor: 1,
//    })
// func CreateTopic(ctx context.Context, broker, topic string, partitions int) error

// Alternative: Just use rpk from scripts/seed_topics.sh
// docker exec streamsre-redpanda rpk topic create events.main -p 6
