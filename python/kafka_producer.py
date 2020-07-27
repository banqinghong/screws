# coding=utf-8
import json
from kafka import KafkaProducer
# from kafka import KafkaConsumer

# 简单的kafka生产者，主要是用于调试

producer = KafkaProducer(bootstrap_servers='ops-kafka-01.domain.com:9092,ops-kafka-02.domain.com:9092,ops-kafka-03.domain.com:9092')

msg_dict = {
    "client_request_id": "1111111",
    "app_name": "测试",
    "event_id": "22222222222",
    "receiver": ["liyang"],
    "notify_type": ["Sms", "DingTalk"],
    "message_content": "测试内容"
}
msg = json.dumps(msg_dict)
producer.send("alert", msg)
producer.close()


