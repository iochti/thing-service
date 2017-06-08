package main

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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
	if in.GetID() == 0 {
		return nil, grpc.Errorf(codes.Internal, "Error: missing ID in request")
	}
	thingID := in.GetID()
	tng := new(pb.Thing)
	if err := t.Db.GetThingByID(thingID, tng); err != nil {
		return nil, grpc.Errorf(codes.NotFound, err.Error())
	}
	return tng, nil
}

// CreateThing takes a thing as parameter, creates it and returns it with its ID
func (t *ThingSvc) CreateThing(ctx context.Context, in *pb.Thing) (*pb.Thing, error) {
	if err := t.Db.CreateThing(in); err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}

	if in.GetID() == 0 {
		return nil, grpc.Errorf(codes.Internal, "Error: no userid after insertion")
	}
	return in, nil
}

// UpdateThing takes a thing as parameter, updates the corresponding data
// and returns the updated element
func (t *ThingSvc) UpdateThing(ctx context.Context, in *pb.Thing) (*pb.Thing, error) {
	if in.GetID() == 0 {
		return nil, grpc.Errorf(codes.InvalidArgument, "Error: can't update user with no ID")
	}
	if err := t.Db.UpdateThing(in); err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	return in, nil
}

// DeleteThing takes a Thing ID as parameter, deletes it and returns the id of the deleted value
func (t *ThingSvc) DeleteThing(ctx context.Context, in *pb.ThingIDRequest) (*empty.Empty, error) {
	if in.GetID() == 0 {
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

	if err := t.Db.DeleteThingArray(in.GetThings()); err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	return new(empty.Empty), nil
}
