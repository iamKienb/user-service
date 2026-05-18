package indexing

import (
	_ "embed"
)

//go:embed mappings/users.json
var UserMapping string

//go:embed mappings/shops.json
var ShopMapping string
