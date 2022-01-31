# Apache Pulsar Development with Golang

A quick and simple example for 


# Tasks

[] Setup
 [x] Run Pulsar Brokers
  [x] create message  topics
 [x] Run producer code (go)
  [x] produce messages ok
 [x] Run consumer code (go)
  [x] consume messages

[] Configure shared consumer pooling with read from "latest unread message" 
[] Implement as containerized services
 [] consumer
 [] producer

# Test Results


- running the producer

```
02:56:04.560 [pulsar-io-29-6] INFO  org.apache.pulsar.broker.service.ServerCnx - New connection from /172.17.0.1:62498
02:56:04.601 [pulsar-io-29-6] INFO  org.apache.pulsar.broker.service.ServerCnx - [/172.17.0.1:62498][persistent://ragnarok/requests/transactions] Creating producer. producerId=1
02:56:04.601 [ForkJoinPool.commonPool-worker-15] INFO  org.apache.pulsar.broker.service.ServerCnx - [/172.17.0.1:62498] persistent://ragnarok/requests/transactions configured with schema false
02:56:04.602 [ForkJoinPool.commonPool-worker-15] INFO  org.apache.pulsar.broker.service.ServerCnx - [/172.17.0.1:62498] Created new producer: Producer{topic=PersistentTopic{topic=persistent://ragnarok/requests/transactions}, client=/172.17.0.1:62498, producerName=standalone-0-1, producerId=1}
02:56:04.650 [pulsar-io-29-6] INFO  org.apache.pulsar.broker.service.ServerCnx - [PersistentTopic{topic=persistent://ragnarok/requests/transactions}][standalone-0-1] Closing producer on cnx /172.17.0.1:62498. producerId=1
02:56:04.650 [pulsar-io-29-6] INFO  org.apache.pulsar.broker.service.ServerCnx - [PersistentTopic{topic=persistent://ragnarok/requests/transactions}][standalone-0-1] Closed producer on cnx /172.17.0.1:62498. producerId=1
02:56:04.651 [pulsar-io-29-6] INFO  org.apache.pulsar.broker.service.ServerCnx - Closed connection from /172.17.0.1:62498

```

```
(base) welcome@Traianos-MacBook-Pro producer % ./producer 
INFO[0000] [Connecting to broker]                        remote_addr="pulsar://localhost:6650"
INFO[0000] [TCP connection established]                  local_addr="[::1]:50463" remote_addr="pulsar://localhost:6650"
INFO[0000] [Connection is ready]                         local_addr="[::1]:50463" remote_addr="pulsar://localhost:6650"
INFO[0000] Found connection in pool key=localhost:665045 0 logical_addr=pulsar://localhost:6650 physical_addr=pulsar://localhost:6650 
INFO[0000] Found connection in pool key=localhost:665045 0 logical_addr=pulsar://localhost:6650 physical_addr=pulsar://localhost:6650 
INFO[0000] [Created producer]                            cnx="[::1]:50463 -> [::1]:6650" producerID=1 producer_name=standalone-0-0 topic="persistent://ragnarok/requests/transactions"
2022/01/31 10:54:54 Published message:  11:0:0
2022/01/31 10:54:54 Published message:  11:1:0
2022/01/31 10:54:54 Published message:  11:2:0
2022/01/31 10:54:54 Published message:  11:3:0
2022/01/31 10:54:54 Published message:  11:4:0
2022/01/31 10:54:54 Published message:  11:5:0
2022/01/31 10:54:54 Published message:  11:6:0
2022/01/31 10:54:54 Published message:  11:7:0
2022/01/31 10:54:54 Published message:  11:8:0
2022/01/31 10:54:54 Published message:  11:9:0
INFO[0000] [Closing producer]                            producerID=1 producer_name=standalone-0-0 topic="persistent://ragnarok/requests/transactions"
INFO[0000] [Closed producer]                             producerID=1 producer_name=standalone-0-0 topic="persistent://ragnarok/requests/transactions"
(base) welcome@Traianos-MacBook-Pro producer % 
```

- running consumer


```
(base) welcome@Traianos-MacBook-Pro consumer % ./consumer 
INFO[0000] [Connecting to broker]                        remote_addr="pulsar://localhost:6650"
INFO[0000] [TCP connection established]                  local_addr="[::1]:50651" remote_addr="pulsar://localhost:6650"
INFO[0000] [Connection is ready]                         local_addr="[::1]:50651" remote_addr="pulsar://localhost:6650"
INFO[0000] Found connection in pool key=localhost:665045 0 logical_addr=pulsar://localhost:6650 physical_addr=pulsar://localhost:6650 
INFO[0000] Found connection in pool key=localhost:665045 0 logical_addr=pulsar://localhost:6650 physical_addr=pulsar://localhost:6650 
INFO[0000] [Connected consumer]                          consumerID=1 name=cyddm subscription=sub001 topic="persistent://ragnarok/requests/transactions"
INFO[0000] [Created consumer]                            consumerID=1 name=cyddm subscription=sub001 topic="persistent://ragnarok/requests/transactions"
Received message msgId: pulsar.trackingMessageID{messageID:pulsar.messageID{ledgerID:70, entryID:0, batchIdx:0, partitionIdx:0}, tracker:(*pulsar.ackTracker)(nil), consumer:(*pulsar.partitionConsumer)(0xc0001c7500), receivedTime:time.Time{wall:0xc075f32ed06f8118, ext:50861666194, loc:(*time.Location)(0x4c67960)}} -- content: 'hello-0'
Received message msgId: pulsar.trackingMessageID{messageID:pulsar.messageID{ledgerID:70, entryID:1, batchIdx:0, partitionIdx:0}, tracker:(*pulsar.ackTracker)(nil), consumer:(*pulsar.partitionConsumer)(0xc0001c7500), receivedTime:time.Time{wall:0xc075f32ed0e8a4d0, ext:50869604662, loc:(*time.Location)(0x4c67960)}} -- content: 'hello-1'
Received message msgId: pulsar.trackingMessageID{messageID:pulsar.messageID{ledgerID:70, entryID:2, batchIdx:0, partitionIdx:0}, tracker:(*pulsar.ackTracker)(nil), consumer:(*pulsar.partitionConsumer)(0xc0001c7500), receivedTime:time.Time{wall:0xc075f32ed0e8b858, ext:50869609748, loc:(*time.Location)(0x4c67960)}} -- content: 'hello-2'
Received message msgId: pulsar.trackingMessageID{messageID:pulsar.messageID{ledgerID:70, entryID:3, batchIdx:0, partitionIdx:0}, tracker:(*pulsar.ackTracker)(nil), consumer:(*pulsar.partitionConsumer)(0xc0001c7500), receivedTime:time.Time{wall:0xc075f32ed190d768, ext:50880628095, loc:(*time.Location)(0x4c67960)}} -- content: 'hello-3'
Received message msgId: pulsar.trackingMessageID{messageID:pulsar.messageID{ledgerID:70, entryID:4, batchIdx:0, partitionIdx:0}, tracker:(*pulsar.ackTracker)(nil), consumer:(*pulsar.partitionConsumer)(0xc0001c7500), receivedTime:time.Time{wall:0xc075f32ed190f6a8, ext:50880635610, loc:(*time.Location)(0x4c67960)}} -- content: 'hello-4'
Received message msgId: pulsar.trackingMessageID{messageID:pulsar.messageID{ledgerID:70, entryID:5, batchIdx:0, partitionIdx:0}, tracker:(*pulsar.ackTracker)(nil), consumer:(*pulsar.partitionConsumer)(0xc0001c7500), receivedTime:time.Time{wall:0xc075f32ed190fe78, ext:50880637480, loc:(*time.Location)(0x4c67960)}} -- content: 'hello-5'
Received message msgId: pulsar.trackingMessageID{messageID:pulsar.messageID{ledgerID:70, entryID:6, batchIdx:0, partitionIdx:0}, tracker:(*pulsar.ackTracker)(nil), consumer:(*pulsar.partitionConsumer)(0xc0001c7500), receivedTime:time.Time{wall:0xc075f32ed2361818, ext:50891458175, loc:(*time.Location)(0x4c67960)}} -- content: 'hello-6'
Received message msgId: pulsar.trackingMessageID{messageID:pulsar.messageID{ledgerID:70, entryID:7, batchIdx:0, partitionIdx:0}, tracker:(*pulsar.ackTracker)(nil), consumer:(*pulsar.partitionConsumer)(0xc0001c7500), receivedTime:time.Time{wall:0xc075f32ed23671f0, ext:50891480536, loc:(*time.Location)(0x4c67960)}} -- content: 'hello-7'
Received message msgId: pulsar.trackingMessageID{messageID:pulsar.messageID{ledgerID:70, entryID:8, batchIdx:0, partitionIdx:0}, tracker:(*pulsar.ackTracker)(nil), consumer:(*pulsar.partitionConsumer)(0xc0001c7500), receivedTime:time.Time{wall:0xc075f32ed2367da8, ext:50891483335, loc:(*time.Location)(0x4c67960)}} -- content: 'hello-8'
Received message msgId: pulsar.trackingMessageID{messageID:pulsar.messageID{ledgerID:70, entryID:9, batchIdx:0, partitionIdx:0}, tracker:(*pulsar.ackTracker)(nil), consumer:(*pulsar.partitionConsumer)(0xc0001c7500), receivedTime:time.Time{wall:0xc075f32ed2d73e00, ext:50902019071, loc:(*time.Location)(0x4c67960)}} -- content: 'hello-9'
INFO[0060] Found connection in pool key=localhost:665045 0 logical_addr=pulsar://localhost:6650 physical_addr=pulsar://localhost:6650 
```

# Notes


- Create a client

```
import (
    "log"
    "time"
    "github.com/apache/pulsar-client-go/pulsar"
)
 
func main() {
    client, err := pulsar.NewClient(
    pulsar.ClientOptions{
        URL:               "pulsar://localhost:6650",
        OperationTimeout:  30 * time.Second,
        ConnectionTimeout: 30 * time.Second,
    })
 
    if err != nil {
        log.Fatalf("Could not instantiate Pulsar client: %v", err)
    }
 
    defer client.Close()
}
```

- run local broker for dev

```
docker run -it -p 6650:6650 -p 8080:8080-v $PWD/data:/pulsar/data --name pulsar apachepulsar/pulsar-standalone
```

- setup topics:

```
docker exec -it pulsar /pulsar/bin/pulsar-admin tenants create ragnarok
docker exec -it pulsar /pulsar/bin/pulsar-admin namespaces create ragnarok/requests
docker exec -it pulsar /pulsar/bin/pulsar-admin topics create persistent://ragnarok/requests/transactions
```

# Refs:

- https://livebook.manning.com/book/pulsar-in-action/chapter-3/97
- https://pulsar.apache.org/docs/en/client-libraries-go/
