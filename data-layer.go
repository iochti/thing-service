package main

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/iochti/thing-service/models"
	_ "github.com/lib/pq"
	"github.com/namsral/flag"
)

const THING_COLLECTION = "thing"

// DataLayerInterface is an interface to abstract DB queries
type DataLayerInterface interface {
	GetThingByID(id string, thing *models.Thing) error
	CreateThing(thing *models.Thing) error
	UpdateThing(thing *models.Thing) error
	DeleteThing(id string) error
	DeleteThingArray(id []string) error
}

// MgoDL implements DataLayerInterface
type MgoDL struct {
	DBName  string
	Session *mgo.Session
}

var (
	mainSession *mgo.Session
	mainDB      *mgo.Database
	DBName      string
)

// Init inits the DB
func (m *MgoDL) Init() error {
	mHost := flag.String("mhost", "localhost", "MongoDB database host")
	mPort := flag.String("mport", "27017", "MongoDB's port")
	mName := flag.String("mname", "crm", "MongoDB's name")
	flag.Parse()
	mainSession, err := mgo.Dial(fmt.Sprintf("mongodb://%s:%s", *mHost, *mPort))
	if err != nil {
		panic(err)
	}
	mainDB = mainSession.DB(*mName)

	m.Session = mainSession.Copy()

	s := m.Session.Copy()
	defer s.Close()
	return nil
}

// GetThingByID gets a thing by its ID and modify the thing passed as parameter
func (m *MgoDL) GetThingByID(id string, thing *models.Thing) error {
	if id == "" {
		return fmt.Errorf("Error: invalid id")
	}

	if thing == nil {
		return fmt.Errorf("Error: nil thing parameter")
	}
	fmt.Println(id)
	sess := m.Session.Copy()
	defer sess.Close()
	if err := sess.DB(DBName).C(THING_COLLECTION).FindId(bson.ObjectIdHex(id)).One(thing); err != nil {
		return err
	}
	return nil
}

// CreateThing creates a thing and modify the thing's ID passed as parameter
func (m *MgoDL) CreateThing(thing *models.Thing) error {
	timeNow := time.Now()
	thing.CreatedAt = timeNow
	thing.UpdatedAt = timeNow
	thing.ID = bson.NewObjectId()
	sess := m.Session.Copy()
	defer sess.Close()
	if err := sess.DB(DBName).C(THING_COLLECTION).Insert(&thing); err != nil {
		return err
	}

	return nil
}

// UpdateThing updates a thing and updates the thing passed as parameter
func (m *MgoDL) UpdateThing(thing *models.Thing) error {
	updateTime := time.Now()
	thing.UpdatedAt = updateTime
	sess := m.Session.Copy()
	defer sess.Close()
	if err := sess.DB(DBName).C(THING_COLLECTION).UpdateId(thing.ID, thing); err != nil {
		return err
	}
	return nil
}

// DeleteThing deletes a thing identified by its ID
func (m *MgoDL) DeleteThing(id string) error {
	if id == "" {
		return fmt.Errorf("Error: invalid ID")
	}
	sess := m.Session.Copy()
	defer sess.Close()
	if err := sess.DB(DBName).C(THING_COLLECTION).RemoveId(bson.ObjectIdHex(id)); err != nil {
		return err
	}
	return nil
}

// DeleteThingArray deletes a list of things identifed by their ID
func (m *MgoDL) DeleteThingArray(ids []string) error {
	if len(ids) == 0 {
		return fmt.Errorf("Error: empty IDs array")
	}
	sess := m.Session.Copy()
	defer sess.Close()
	for _, o := range ids {
		if err := sess.DB(DBName).C(THING_COLLECTION).Remove(bson.M{"_id": bson.M{"$eq": bson.ObjectIdHex(o)}}); err != nil {
			return err
		}
	}
	return nil
}
