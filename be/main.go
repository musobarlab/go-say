package main

import (
	//"os"
	//"os/exec"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/wuriyanto48/go-say/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

  log.Println("your server is running")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTextToSpeechServer(grpcServer, server{})
	err = grpcServer.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}

type server struct {
}

func (s server) Say(ctx context.Context, text *pb.Text) (*pb.Speech, error) {
	return nil, fmt.Errorf("not implemented yet fuck")
}

// cmd := exec.Command("flite", "-t", os.Args[1], "-o", "output.wav")
// cmd.Stdout = os.Stdout
// cmd.Stderr = os.Stderr
// if err := cmd.Run(); err != nil {
//   log.Fatal(err)
// }
