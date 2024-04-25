package mxvd

import (
	"context"

	"github.com/sjy-dv/mind-x/processor/entity"
	"github.com/sjy-dv/mind-x/processor/protobuf/protocol/v0"
)

func (mxvd *MXVDProxy) Insert(
	ctx context.Context, obj *entity.InsertObject,
) error {

	_, err := mxvd.dataManager().Insert(ctx, &protocol.InsertRequest{
		DatasetId: mxvd.PersonalID,
		Id:        obj.ID,
		Value:     obj.Vector,
		Metadata:  obj.Metadata,
	})
	return err
}
