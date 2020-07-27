# coding=utf-8
from kafka import KafkaConsumer
import json

# 一个很简单的kafka消费者,主要是用于调试

consumer = KafkaConsumer('alert', group_id='cmdb-test-ly',
                         bootstrap_servers='ops-kafka-01.domain.com:9092,ops-kafka-02.domain.com:9092,ops-kafka-03.domain.com:9092')
for msg in consumer:
    # recv = "%s:%d:%d: key=%s value=%s" % (msg.topic, msg.partition, msg.offset, msg.key, msg.value)
    # recv = "%s" % (msg.value)
    metric = json.loads(msg.value)

    print (msg.value)
    # if metric["metric_name"] != "machine":
    #     continue
    # else:
    #     print (metric["bytes_uuid"])
