package mxvd

import (
	"context"

	pbV0 "github.com/sjy-dv/mind-x/processor/protobuf/protocol/v0"
)

func (mxvd *MXVDProxy) CreateDatabase(
	ctx context.Context, dim, partitionCount, replicationFactor uint32,
) ([]byte, error) {

	dataset, err := mxvd.datasetManager().Create(ctx, &pbV0.Dataset{
		Dimension:         dim,
		PartitionCount:    partitionCount,
		ReplicationFactor: replicationFactor,
	})
	if err != nil {
		return nil, err
	}
	return dataset.GetId(), nil
}

func (mxvd *MXVDProxy) DatabaseStatus(
	ctx context.Context, databaseID []byte,
) (*pbV0.Dataset, error) {
	status, err := mxvd.datasetManager().Get(ctx, &pbV0.GetDatasetRequest{
		DatasetId: databaseID,
		WithSize:  true,
	})
	return status, err
}

func (mxvd *MXVDProxy) DeleteDatabase(
	ctx context.Context, databaseID []byte,
) error {
	_, err := mxvd.datasetManager().Delete(ctx, &pbV0.UUIDRequest{
		Id: databaseID,
	})
	return err
}
