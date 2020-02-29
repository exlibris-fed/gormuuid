// Package gormuuid provides a type that can be embedded in a GORM model (http://gorm.io) to provide support for UUID primary keys.
//
// While some databases (ie postgres) have native support for UUIDs, for ease of use and interoperability the column type is a byte slice.
//
// Usage:
//
//    package foo
//
//    import (
//        "github.com/exlibris-fed/gormuuid"
//    )
//
//    type Person struct {
//        gormuuid.UUID
//        Name string
//        // etc
//    }
package gormuuid

import (
	"errors"

	"github.com/google/uuid"
)

var (
	// ErrorNoUUID is the error returned when you attempt to get a UUID when it hasn't been generated.
	ErrorNoUUID = errors.New("UUID has not been created")
)

// A UUID is a struct that can be embedded to add UUID primary key support to a GORM model.
type UUID struct {
	ID []byte `gorm:"primary key"`
}

// BeforeCreate ensures that a model has a valid UUID before insertion into the database. If one exists already (ie your implementation needed to specify one) it will be respected.
func (u *UUID) BeforeCreate() (err error) {
	if len(u.ID) == 16 {
		return
	}

	uuid, err := uuid.New().MarshalBinary()
	if err != nil {
		return
	}

	u.ID = uuid
	return
}

// UUID returns the UUID of the model's ID.
//
// As it returns a uuid.UUID object, you can then call any of that package's methods. (see https://godoc.org/github.com/google/uuid). For example, to use as a string:
//
//    c := Person{
//        Name: "Frank",
//    }
//    uuid, _ := p.UUID() // don't ignore errors!
//    id := uuid.String()
func (u *UUID) UUID() (uuid.UUID, error) {
	return uuid.FromBytes(u.ID)
}
