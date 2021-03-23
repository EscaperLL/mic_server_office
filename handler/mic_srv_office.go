package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	mic_srv_office "mic_srv_office/proto/mic_srv_office"
)

type Mic_srv_office struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Mic_srv_office) Call(ctx context.Context, req *mic_srv_office.Request, rsp *mic_srv_office.Response) error {
	log.Info("Received Mic_srv_office.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Mic_srv_office) Stream(ctx context.Context, req *mic_srv_office.StreamingRequest, stream mic_srv_office.Mic_srv_office_StreamStream) error {
	log.Infof("Received Mic_srv_office.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&mic_srv_office.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Mic_srv_office) PingPong(ctx context.Context, stream mic_srv_office.Mic_srv_office_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&mic_srv_office.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
