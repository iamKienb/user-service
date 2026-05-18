package outbox

import "github.com/iamKienb/shopify-go-platform/app_error"

func (s *outboxService) wrapError(err error) error {
	return app_error.WrapError(err, nil)
}
