#!/bin/bash
#


#run the pulsar standalone container
docker run  -p 6650:6650 -p 8080:8080 -v $PWD/data:/pulsar/data --name pulsar apachepulsar/pulsar-standalone

#build and run the producer
cd producer
go build producer.go
./producer

#build and run the consumer
cd consumer/
go build consumer.go
./consumer
