package main

import (
	"database/sql"
	"fmt"

	pb "github.com/iochti/thing-service/proto"
	_ "github.com/lib/pq"
	"github.com/namsral/flag"
)

// DataLayerInterface is an interface to abstract DB queries
type DataLayerInterface interface {
	GetThingByID(id int32, thing *pb.Thing) error
	CreateThing(thing *pb.Thing) error
	UpdateThing(thing *pb.Thing) error
	DeleteThing(id int32) error
	DeleteThingArray(id []int32) error
}

// PostgresDL implements DataLayerInterface and CRUD methods
type PostgresDL struct {
	DBName string
	Host   string
	Db     *sql.DB
}

// Init the DB host,user,etc...
func (p *PostgresDL) Init() error {
	host := flag.String("pq-host", "localhost", "PostgresSQL database host")
	user := flag.String("pq-user", "postgres", "PostgresSQL user")
	dbName := flag.String("pq-db", "iochti", "PostgresSQL DBName")
	password := flag.String("pq-pwd", "", "PostgresSQL user password")
	flag.Parse()

	// Create a db connection
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", *user, *password, *host, *dbName))
	if err != nil {
		return err
	}
	p.Db = db
	return nil
}

// GetThingByID gets a thing by its ID and modify the thing passed as parameter
func (p *PostgresDL) GetThingByID(id int32, thing *pb.Thing) error {
	return nil
}

// CreateThing creates a thing and modify the thing's ID passed as parameter
func (p *PostgresDL) CreateThing(thing *pb.Thing) error {
	return nil
}

// UpdateThing updates a thing and updates the thing passed as parameter
func (p *PostgresDL) UpdateThing(thing *pb.Thing) error {
	return nil
}

// DeleteThing deletes a thing identified by its ID
func (p *PostgresDL) DeleteThing(id int32) error {
	return nil
}

// DeleteThingArray deletes a list of things identifed by their ID
func (p *PostgresDL) DeleteThingArray(ids []int32) error {
	return nil
}
