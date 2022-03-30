package mongo

import (
	"github.com/joaopandolfi/blackwhale/remotes/mongo"
	"golang.org/x/xerrors"
	"gopkg.in/mgo.v2"
)

type Dao struct {
	Collection string
}

// Save data
func (d *Dao) Save(data interface{}) error {
	session, err := mongo.NewSession()
	if err != nil {
		return xerrors.Errorf("(%s): connecting on mongo: %w", d.Collection, err)
	}

	err = session.GetCollection(d.Collection).Insert(&data)
	if err != nil {
		return xerrors.Errorf("Insert Log error %v", err)
	}
	return nil
}

// GetCollection return collection pointer
func (d *Dao) GetCollection() (*mgo.Collection, error) {
	session, err := mongo.NewSession()
	if err != nil {
		err = xerrors.Errorf("Unable to connect to mongo: %v", err)
		return nil, err
	}

	return session.GetCollection(d.Collection), nil
}

func GracefullShutdown() {
	mongo.Close()
}

func CreateIndex(collection string, keys ...string) error {
	return mongo.CreateIndex(collection, keys...)
}
