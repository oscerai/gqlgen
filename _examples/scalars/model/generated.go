// Code generated by github.com/oscerai/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/oscerai/gqlgen/_examples/scalars/external"
)

type Address struct {
	ID       external.ObjectID `json:"id"`
	Location *Point            `json:"location"`
}
