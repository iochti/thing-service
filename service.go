package main

import (
	"encoding/json"
	"fmt"

	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/iochti/thing-service/models"
	pb "github.com/iochti/thing-service/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// ThingSvc implements the RPCs from ThingSvc
type ThingSvc struct {
	Db DataLayerInterface
}

// GetThing gets a thing by its ID and returns it
func (t *ThingSvc) GetThing(ctx context.Context, in *pb.ThingIDRequest) (*pb.Thing, error) {
	if in.GetID() == "" {
		return nil, grpc.Errorf(codes.Internal, "Error: missing ID in request")
	}
	tng := new(models.Thing)
	if err := t.Db.GetThingByID(in.GetID(), tng); err != nil {
		return nil, grpc.Errorf(codes.NotFound, err.Error())
	}
	resp, err := json.Marshal(tng)
	if err != nil {
		return nil, err
	}
	return &pb.Thing{Item: resp}, nil
}

// CreateThing takes a thing as parameter, creates it and returns it with its ID
func (t *ThingSvc) CreateThing(ctx context.Context, in *pb.Thing) (*pb.Thing, error) {
	thing := models.Thing{}
	if err := json.Unmarshal(in.GetItem(), &thing); err != nil {
		return nil, err
	}

	if err := t.Db.CreateThing(&thing); err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}

	if thing.ID == "" {
		return nil, grpc.Errorf(codes.Internal, "Error: no userid after insertion")
	}
	bytesResp, err := json.Marshal(thing)
	if err != nil {
		return nil, err
	}
	return &pb.Thing{Item: bytesResp}, nil
}

// UpdateThing takes a thing as parameter, updates the corresponding data
// and returns the updated element
func (t *ThingSvc) UpdateThing(ctx context.Context, in *pb.Thing) (*pb.Thing, error) {
	tng := models.Thing{}
	if err := json.Unmarshal(in.GetItem(), &tng); err != nil {
		return nil, err
	}

	if err := t.Db.UpdateThing(&tng); err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	byteResp, err := json.Marshal(tng)
	if err != nil {
		return nil, err
	}
	return &pb.Thing{Item: byteResp}, nil
}

// DeleteThing takes a Thing ID as parameter, deletes it and returns the id of the deleted value
func (t *ThingSvc) DeleteThing(ctx context.Context, in *pb.ThingIDRequest) (*empty.Empty, error) {
	if in.GetID() == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "Error: invalid id as parameter")
	}
	if err := t.Db.DeleteThing(in.GetID()); err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	return new(empty.Empty), nil
}

// BulkDeleteThing takes a Thing array as parametes
func (t *ThingSvc) BulkDeleteThing(ctx context.Context, in *pb.ThingIDArray) (*empty.Empty, error) {
	if len(in.GetThings()) == 0 {
		return nil, grpc.Errorf(codes.Internal, "Error: empty ids array")
	}
	fmt.Println(in.GetThings())
	if err := t.Db.DeleteThingArray(in.GetThings()); err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	return new(empty.Empty), nil
}

// ListGroupThings sets a stream that sends a list of things fetched by their group ID
func (t *ThingSvc) ListGroupThings(in *pb.GroupRequest, stream pb.ThingSvc_ListGroupThingsServer) error {
	return nil
}
