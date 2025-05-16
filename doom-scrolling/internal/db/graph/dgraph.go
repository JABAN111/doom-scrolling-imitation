package graph

import (
	"github.com/dgraph-io/dgo/v240"
	"github.com/dgraph-io/dgo/v240/protos/api"

	"log/slog"
	"rshd/lab1/v2/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DgraphClient struct {
	Conn *grpc.ClientConn
	Dg   *dgo.Dgraph
}

func InitClient(cfg config.Config, log *slog.Logger) (*DgraphClient, error) {
	conn, err := grpc.NewClient("not used", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("Failed to dial gRPC", "error", err)
		return nil, err
	}

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	return &DgraphClient{Conn: conn, Dg: dg}, nil
}

func (c *DgraphClient) Close() {
	if c.Conn != nil {
		_ = c.Conn.Close()
	}
}
