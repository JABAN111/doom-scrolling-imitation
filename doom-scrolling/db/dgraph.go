package db

import (
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"

	//_ "github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"rshd/lab1/v2/config"
)

type DgraphClient struct {
	Conn *grpc.ClientConn
	Dg   *dgo.Dgraph
}

func InitClient(cfg config.Config, log *slog.Logger) (*DgraphClient, error) {
	conn, err := grpc.NewClient(cfg.DgraphCfg.URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
