package interceptor

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	apperror "shopify-user-common-module/error"

// 	"connectrpc.com/connect"
// )

// func ErrorResponseInterceptor() connect.UnaryInterceptorFunc {
// 	return func(next connect.UnaryFunc) connect.UnaryFunc {
// 		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
// 			resp, err := next(ctx, req)
// 			if err == nil {
// 				return resp, nil
// 			}
// 			fmt.Printf("[ERROR] Path: %s | Error: %v\n", req.Spec().Procedure, err)

// 			var connectErr *connect.Error
// 			if errors.As(err, &connectErr) {
// 				return nil, connectErr
// 			}

// 			var appErr *apperror.AppError
// 			if errors.As(err, &appErr) {
// 				if appErr.Kind == apperror.KindInternal {
// 					fmt.Printf("[SERVER-ERROR] Detail: %v\n", appErr)
// 				}
// 				return nil, connect.NewError(appErr.Kind.ConnectCode(), errors.New(appErr.Message))
// 			}

// 			fmt.Printf("[UNHANDLED-ERROR] %v\n", err)
// 			return nil, connect.NewError(connect.CodeInternal, errors.New("internal server error"))
// 		}
// 	}
// }
