package main

import (
	pb "github.com/iochti/thing-service/proto"
	"golang.org/x/net/context"
)

// ThingSvc implements the RPCs from ThingSvc
type ThingSvc struct {
	Db DataLayerInterface
}

// GetThing gets a thing by its ID and returns it
func (t *ThingSvc) GetThing(ctx context.Context, in *pb.ThingIDRequest) (*pb.Thing, error) {
	return nil, nil
}

// CreateThing takes a thing as parameter, creates it and returns it with its ID
func (t *ThingSvc) CreateThing(ctx context.Context, in *pb.Thing) (*pb.Thing, error) {
	return nil, nil
}

// UpdateThing takes a thing as parameter, updates the corresponding data
// and returns the updated element
func (t *ThingSvc) UpdateThing(ctx context.Context, in *pb.Thing) (*pb.Thing, error) {
	return nil, nil
}

// DeleteThing takes a Thing ID as parameter, deletes it and returns the id of the deleted value
func (t *ThingSvc) DeleteThing(ctx context.Context, in *pb.ThingIDRequest) (*pb.ThingDeletion, error) {
	return nil, nil
}

// BulkDeleteThing takes a Thing array as parametes
func (t *ThingSvc) BulkDeleteThing(ctx context.Context, in *pb.ThingIDArray) (*pb.ThingDeletion, error) {
	return nil, nil
}
