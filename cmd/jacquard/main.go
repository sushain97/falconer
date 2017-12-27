package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/gphat/jacquard"
	"google.golang.org/grpc"

	"github.com/stripe/veneur/ssf"
)

type JacquardServer struct {
	workerCount int64
	workers     []*jacquard.Worker
}

func NewServer(workerCount int64) *JacquardServer {
	workers := make([]*jacquard.Worker, workerCount)
	for i := range workers {
		workers[i] = jacquard.NewWorker()
		go workers[i].Work()
	}
	return &JacquardServer{
		workerCount: workerCount,
		workers:     workers,
	}
}

func (s *JacquardServer) DispatchSpan(span *ssf.SSFSpan) {
	s.workers[span.TraceId%s.workerCount].SpanChan <- span
}

func (s *JacquardServer) SendSpans(stream jacquard.Jacquard_SendSpansServer) error {
	count := 0
	start := time.Now()
	for {
		batch, err := stream.Recv()
		if err == io.EOF {
			d := time.Since(start)
			fmt.Printf("Sending response: %d in %f @ %f/sec\n", count, d.Seconds(), float64(count)/d.Seconds())
			return stream.SendMsg(&jacquard.SpanResponse{
				Greeting: "fart",
			})
		}
		if err != nil {
			return err
		}
		for _, span := range batch.Spans {
			s.DispatchSpan(span)
		}
		count = count + len(batch.Spans)
	}
}

func (s *JacquardServer) GetTrace(req *jacquard.TraceRequest, stream jacquard.Jacquard_GetTraceServer) error {
	log.Printf("Looking for %v\n", req.GetTraceID())
	worker := s.workers[req.GetTraceID()%s.workerCount]
	spans := worker.GetTrace(req.GetTraceID())
	log.Printf("Found %v", spans)
	for _, span := range spans {
		stream.Send(span)
	}
	fmt.Println("Done with GetTrace")
	return nil
}

func (s *JacquardServer) FindSpans(req *jacquard.FindSpanRequest, stream jacquard.Jacquard_FindSpansServer) error {
	log.Printf("Looking for %v", req)

	tagsToFind := req.GetTags()

	start := time.Now()
	scanned := 0
	var foundSpans []*ssf.SSFSpan
	for _, worker := range s.workers {
		for _, spans := range worker.Spans {
			for _, span := range spans {
				scanned++
				for fk, fv := range tagsToFind {
					if v, ok := span.Tags[fk]; ok {
						if v == fv {
							foundSpans = append(foundSpans, span)
							continue
						}
					}
				}
			}
		}
	}

	duration := time.Since(start)
	log.Printf("Scanned %d in %f seconds @ %f/second", scanned, duration.Seconds(), float64(scanned)/duration.Seconds())

	for _, span := range foundSpans {
		stream.Send(span)
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:3000"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	jacquard.RegisterJacquardServer(grpcServer, NewServer(256))
	grpcServer.Serve(lis)
}
