package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	avro "github.com/elodina/go-avro"
	avroKafka "github.com/elodina/go-kafka-avro"
)

var schemaFlag = flag.String("schema", "http://schema:8081", "Schema Registry URL")

func main() {
	flag.Parse()

	decoder := avroKafka.NewKafkaAvroDecoder(*schemaFlag)

	in, err := ioutil.ReadAll(os.Stdin)
	check(err)

	if len(in) > 0 {
		decoded, err := decoder.Decode(in)
		check(err)

		rec := decoded.(*avro.GenericRecord)
		fmt.Println(rec)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
