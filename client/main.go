package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"

	pb "github.com/wuriyanto48/go-say/api"
	"google.golang.org/grpc"
)

func main() {
	backend := flag.String("b", "localhost:8080", "backend address")
	output := flag.String("o", "output.wav", "wav output")
	flag.Parse()

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewTextToSpeechClient(conn)
	text := &pb.Text{Text: "hello, fuck"}
	res, err := client.Say(context.Background(), text)
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(*output, res.Audio, 0666); err != nil {
		log.Fatal(err)
	}

}
