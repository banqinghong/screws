package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

func ConsumerTest(partition int32) {
	fmt.Printf("consumer_test\n")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// consumer
	consumer, err := sarama.NewConsumer([]string{"10.40.66.10:9092"}, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}

	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("jarvis.eventbus-vm", partition, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			//fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
			//	msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
			fmt.Printf("\n")
			fmt.Printf("%s\n", string(msg.Value))
			fmt.Printf("\n")
		case err := <-partitionConsumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}
	}

}

// 假设有12个分区
func main() {
	for n := 0; n < 12; n++ {
		go ConsumerTest(int32(n))
	}
	time.Sleep(6*time.Hour)
}
