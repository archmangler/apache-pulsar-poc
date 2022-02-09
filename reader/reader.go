package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func writeBytesToFile(f string, byteSlice []byte) int {
	// Open a new file for writing only

	f = "./data/" + f

	file, err := os.OpenFile(
		f,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write bytes to file
	bytesWritten, err := file.Write(byteSlice)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Wrote %d bytes.\n", bytesWritten)

	return bytesWritten
}

func readBackByEntryId(msgDir string, msgIndex string) (yourBytes []byte) {

	//We know the file name by convention
	fname := msgDir + "/" + msgIndex + ".dat"

	yourBytes, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Printf("error reading %s", fname)
		return nil
	}

	return yourBytes
}

func getFiles(aDir string) []string {

	var theFiles []string

	files, err := ioutil.ReadDir("./data/")

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {

		theFiles = append(theFiles, f.Name())

	}

	return theFiles
}

func streamAll(reader pulsar.Reader, startMsgIndex int64, stopMsgIndex int64) {

	read := false

	for reader.HasNext() {

		msg, err := reader.Next(context.Background())

		if err != nil {
			log.Fatal(err)
		}

		//can I access the details of the message ? yes
		fmt.Printf("%v -> %#v\n", msg.ID().EntryID(), msg.ID())

		//Can i serialize into bytes? Yes
		myBytes := msg.ID().Serialize()

		//Can I store it somewhere? Perhaps a map ? or even on disk in a file ?
		//In other words: Can I write a byte[] slice to a file? Yes!
		msgIndex := msg.ID().EntryID()

		if msgIndex == startMsgIndex {
			fmt.Println("start read: ", msgIndex)
			read = true
		}

		if msgIndex > stopMsgIndex {
			fmt.Println("stop reading: ", msgIndex)
			read = false
		}

		if read == false {

			fmt.Println("skipping ", msgIndex)

		} else {

			fname := strconv.FormatInt(msgIndex, 10) + ".dat"

			fmt.Println("written bytes: ", writeBytesToFile(fname, myBytes))

			fmt.Printf("Received message msgId: %#v -- content: '%s' published at %v\n",
				msg.ID(), string(msg.Payload()), msg.PublishTime())

		}

		/*
			//FYI - to save and reread a msgId from store: https://githubmemory.com/@storm-5
			msgId := msg.ID()
			msgIdBytes := msgId.Serialize()
			idNew, _ := pulsar.DeserializeMessageID(msgIdBytes)

			readerInclusive, err := client.CreateReader(pulsar.ReaderOptions{
				Topic:                   "ragnarok/transactions/requests",
				StartMessageID:          idNew,
				StartMessageIDInclusive: true,
			})
		*/
	}

}

func retrieveRange(client pulsar.Client) {

	someFiles := getFiles("./data/")

	for _, f := range someFiles {

		fIndex := strings.Split(f, ".")[0]

		fmt.Println("re-reading message index -> ", fIndex)

		msgIdBytes := readBackByEntryId("./data", fIndex)

		fmt.Printf("boom -> %#v\n", msgIdBytes)

		idNew, err := pulsar.DeserializeMessageID(msgIdBytes)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Got message entry id => ", idNew.EntryID())

		readerInclusive, err := client.CreateReader(pulsar.ReaderOptions{
			Topic:                   "ragnarok/transactions/requests",
			StartMessageID:          idNew,
			StartMessageIDInclusive: true,
		})

		if err != nil {
			log.Fatal(err)
		}

		defer readerInclusive.Close()

		//defer readerInclusive.Close()
		fmt.Println("bleep!")

		msg, err := readerInclusive.Next(context.Background())

		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println("retrieved message -> ", string(msg.Payload()))
		fmt.Printf("Retrieved message ID message msgId: %#v -- content: '%s' published at %v\n",
			msg.ID(), string(msg.Payload()), msg.PublishTime())

	}
}

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

	reader, err := client.CreateReader(pulsar.ReaderOptions{
		Topic:          "ragnarok/transactions/requests",
		StartMessageID: pulsar.EarliestMessageID(),
	})

	if err != nil {
		log.Fatal(err)
	}

	defer reader.Close()

	if err != nil {
		log.Fatal(err)
	}

	var startMsgId int64 = 55
	var stopMsgId int64 = 66

	//stream all the messages from the earliest to latest
	//pick a subset between a start and stop id
	streamAll(reader, startMsgId, stopMsgId)

	//retrieve the picked range
	retrieveRange(client)

	//Can I read back the data associated with a given EntryID, from the external data source (e.g file/redis)?
	//1. [x] Read all files in ./data/
	//2. [x] Iterate through the filename list
	//3. [x] Get the contents of each file for a msgId object
	//4. [x] Open a reader and
	//5. [x] read with the msgId
	//6. [x] Print data to output

	//8. [] Input startRead and stopRead parameters to constrain the sequence
	//7. [] put data in format that redis can store ??

}
