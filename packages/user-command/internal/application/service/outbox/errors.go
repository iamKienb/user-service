package outbox

import "github.com/iamKienb/go-core/app_error"

func (s *outboxService) wrapError(err error) error {
	return app_error.WrapError(err, nil)
}
