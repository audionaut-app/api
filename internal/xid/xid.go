package xid

import (
	"fmt"
	"github.com/rs/xid"
)

type ResourceType interface {
	Prefix() string
}

type Token struct{}

func (t Token) Prefix() string { return "tok" }

type ID[T ResourceType] xid.ID

func New[T ResourceType]() ID[T] {
	return ID[T](xid.New())
}

func (id ID[T]) String() string {
	var resourceType T
	return fmt.Sprintf("%s_%s", resourceType.Prefix(), xid.ID(id).String())
}
