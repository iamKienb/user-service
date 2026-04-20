package otp

import (
	"context"
	"shopify-user-command-module/contract/protogen/otp"
	"shopify-user-command-module/contract/protogen/otp/otpconnect"
	"shopify-user-command-module/internal/application/command/resend_otp"
	"shopify-user-command-module/internal/application/command/verify_otp"

	"connectrpc.com/connect"
)

type OTPServer struct {
	verifyOTPExecutor verify_otp.Executor
	resendOTPExecutor resend_otp.Executor
}

func NewOTPServer(verifyOTPExecutor verify_otp.Executor, resendOTPExecutor resend_otp.Executor) *OTPServer {
	return &OTPServer{
		verifyOTPExecutor: verifyOTPExecutor,
		resendOTPExecutor: resendOTPExecutor,
	}
}

func (s *OTPServer) Verify(ctx context.Context, req *connect.Request[otp.VerifyRequest]) (*connect.Response[otp.VerifyResponse], error) {
	cmd := ToVerifyCommand(req.Msg)

	result, err := s.verifyOTPExecutor.Execute(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(ToVerifyResponse(result)), nil
}

func (s *OTPServer) Resend(ctx context.Context, req *connect.Request[otp.ResendRequest]) (*connect.Response[otp.ResendResponse], error) {
	cmd := ToResendCommand(req.Msg)

	result, err := s.resendOTPExecutor.Execute(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(ToResendResponse(result)), nil
}

var _ otpconnect.OTPCommandServiceHandler = (*OTPServer)(nil)
