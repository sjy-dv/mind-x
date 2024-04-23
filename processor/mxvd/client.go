package mxvd

import (
	pbV0 "github.com/sjy-dv/mind-x/processor/protobuf/protocol/v0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MXVDProxy struct {
	Conn *grpc.ClientConn
}

func InitMXVD() (*MXVDProxy, error) {
	proxy := &MXVDProxy{}
	conn, err := grpc.Dial(":50003", grpc.WithTransportCredentials(
		insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	proxy.Conn = conn

	return proxy, nil
}

func (mxvd *MXVDProxy) datasetManager() pbV0.DatasetManagerClient {
	return pbV0.NewDatasetManagerClient(mxvd.Conn)
}

func (mxvd *MXVDProxy) dataManager() pbV0.DataManagerClient {
	return pbV0.NewDataManagerClient(mxvd.Conn)
}

func (mxvd *MXVDProxy) searchManager() pbV0.SearchClient {
	return pbV0.NewSearchClient(mxvd.Conn)
}
