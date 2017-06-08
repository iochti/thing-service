package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
	if id <= 0 {
		return fmt.Errorf("Error: invalid id")
	}

	if thing == nil {
		return fmt.Errorf("Error: nil thing parameter")
	}
	var createdAt time.Time
	var updatedAt time.Time

	if err := p.Db.QueryRow("SELECT id, name, description, created_at, updated_at FROM things WHERE id = $1;", id).Scan(
		&thing.ID,
		&thing.Name,
		&thing.Description,
		&createdAt,
		&updatedAt,
	); err != nil {
		return err
	}
	thing.CreatedAt = &timestamp.Timestamp{Seconds: createdAt.Unix()}
	thing.UpdatedAt = &timestamp.Timestamp{Seconds: updatedAt.Unix()}
	return nil
}

// CreateThing creates a thing and modify the thing's ID passed as parameter
func (p *PostgresDL) CreateThing(thing *pb.Thing) error {
	var thingID int
	timeNow := time.Now()
	err := p.Db.QueryRow(`INSERT INTO things(name, description, created_at, updated_at)
	VALUES($1, $2, $3, $3) RETURNING id;`, thing.GetName(), thing.GetDescription(), timeNow).Scan(&thingID)
	if err != nil {
		return err
	}
	thing.ID = int32(thingID)
	thing.CreatedAt = &timestamp.Timestamp{Seconds: int64(timeNow.Unix())}
	thing.UpdatedAt = &timestamp.Timestamp{Seconds: int64(timeNow.Unix())}

	return nil
}

// UpdateThing updates a thing and updates the thing passed as parameter
func (p *PostgresDL) UpdateThing(thing *pb.Thing) error {
	updateTime := time.Now()
	var createdAt time.Time
	err := p.Db.QueryRow(`UPDATE things SET
		name=$1,
		description=$2,
		updated_at=$3
		WHERE id=$4 RETURNING created_at;`, thing.GetName(), thing.GetDescription(), updateTime, thing.GetID()).Scan(&createdAt)
	if err != nil {
		return err
	}
	thing.UpdatedAt = &timestamp.Timestamp{Seconds: updateTime.Unix()}
	thing.CreatedAt = &timestamp.Timestamp{Seconds: createdAt.Unix()}
	return nil
}

// DeleteThing deletes a thing identified by its ID
func (p *PostgresDL) DeleteThing(id int32) error {
	if id <= 0 {
		return fmt.Errorf("Error: invalid ID")
	}
	_, err := p.Db.Exec("DELETE FROM things WHERE id=$1;", id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteThingArray deletes a list of things identifed by their ID
func (p *PostgresDL) DeleteThingArray(ids []int32) error {
	if len(ids) == 0 {
		return fmt.Errorf("Error: empty IDs array")
	}
	query := make([]string, len(ids))
	for i, id := range ids {
		query[i] = strconv.Itoa(int(id))
	}
	_, err := p.Db.Exec(fmt.Sprintf("DELETE FROM things WHERE id IN (%s)", strings.Join(query, ", ")))
	if err != nil {
		return err
	}
	return nil
}
