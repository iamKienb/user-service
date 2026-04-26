package otp

import (
	"shopify-user-command-module/internal/application/command/resend_otp"
	"shopify-user-command-module/internal/application/command/verify_otp"

	"github.com/iamKienb/shopify-go-api/gen/otp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToVerifyCommand(req *otp.VerifyRequest) verify_otp.Command {
	return verify_otp.Command{
		OTP:          req.GetOtp(),
		SessionToken: req.GetSessionToken(),
	}
}

func ToVerifyResponse(result *verify_otp.Result) *otp.VerifyResponse {
	return &otp.VerifyResponse{
		AccessToken:      result.AccessToken,
		RefreshToken:     result.RefreshToken,
		AccessExpiresAt:  timestamppb.New(result.AccessTokenExpiresAt),
		RefreshExpiresAt: timestamppb.New(result.RefreshTokenExpiresAt),
	}
}

func ToResendCommand(req *otp.ResendRequest) resend_otp.Command {
	return resend_otp.Command{
		SessionToken: req.GetSessionToken(),
	}
}

func ToResendResponse(result *resend_otp.Result) *otp.ResendResponse {
	return &otp.ResendResponse{
		ExpiresAt: timestamppb.New(result.ExpiresAt),
	}
}
