package otp

import (
	"context"
	"user-command-module/internal/application/commands/resend_otp"
	"user-command-module/internal/application/commands/verify_otp"

	"connectrpc.com/connect"
	otpv1 "github.com/iamKienb/api-contract/gen/otp"
	"github.com/iamKienb/api-contract/gen/otp/otpconnect"
)

type otpServer struct {
	verifyOTPExecutor verify_otp.Executor
	resendOTPExecutor resend_otp.Executor
}

func NewOTPServer(verifyOTPExecutor verify_otp.Executor, resendOTPExecutor resend_otp.Executor) *otpServer {
	return &otpServer{
		verifyOTPExecutor: verifyOTPExecutor,
		resendOTPExecutor: resendOTPExecutor,
	}
}

func (s *otpServer) Verify(ctx context.Context, req *connect.Request[otpv1.VerifyRequest]) (*connect.Response[otpv1.VerifyResponse], error) {
	cmd := ToVerifyCommand(req.Msg)
	result, err := s.verifyOTPExecutor.Execute(ctx, cmd)
	if err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(ToVerifyResponse(result)), nil
}

func (s *otpServer) Resend(ctx context.Context, req *connect.Request[otpv1.ResendRequest]) (*connect.Response[otpv1.ResendResponse], error) {
	cmd := ToResendCommand(req.Msg)
	result, err := s.resendOTPExecutor.Execute(ctx, cmd)
	if err != nil {
		return nil, mapError(err)
	}

	return connect.NewResponse(ToResendResponse(result)), nil
}

var _ otpconnect.OTPCommandServiceHandler = (*otpServer)(nil)
