package port

import (
	postgresx "github.com/iamKienb/shopify-go-platform/postgres"
)

type TxManager interface {
	postgresx.TxManager
}
