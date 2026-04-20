package user

import (
	"shopify-user-command-module/contract/protogen/user"
	"shopify-user-command-module/internal/application/command/register_user"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToRegisterCommand(req *user.RegisterRequest) register_user.Command {
	return register_user.Command{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		FullName: req.GetFullName(),
		Gender:   req.GetGender(),
	}
}

func ToRegisterResponse(result *register_user.Result) *user.RegisterResponse {
	return &user.RegisterResponse{
		SessionToken: result.SessionToken,
		ExpiresAt:    timestamppb.New(result.ExpiresAt),
	}
}
