package shared

import (
	"github.com/google/uuid"
)

type AddressID int
type RoleID int

type UserID uuid.UUID
type UserAddressID uuid.UUID

type ShopID uuid.UUID
type ShopAddressID uuid.UUID

func NewID[T ~[16]byte]() T {
	return T(uuid.Must(uuid.NewV7()))
}

func (id UserID) String() string {
	return "user_" + uuid.UUID(id).String()
}
func (id UserAddressID) String() string {
	return "addr_" + uuid.UUID(id).String()
}

func (id ShopID) String() string {
	return "shop_" + uuid.UUID(id).String()
}
func (id ShopAddressID) String() string {
	return "addr_" + uuid.UUID(id).String()
}

func (id UserID) RawID() uuid.UUID {
	return uuid.UUID(id)
}
func (id UserAddressID) RawID() uuid.UUID {
	return uuid.UUID(id)
}

func (id ShopID) RawID() uuid.UUID {
	return uuid.UUID(id)
}
func (id ShopAddressID) RawID() uuid.UUID {
	return uuid.UUID(id)
}
