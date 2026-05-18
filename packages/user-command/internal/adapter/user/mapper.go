package user

import (
	"user-command-module/internal/application/command/add_user_address"
	"user-command-module/internal/application/command/login_user"
	"user-command-module/internal/application/command/register_user"
	"user-command-module/internal/domain/account"
	"user-command-module/internal/domain/shared"

	"github.com/iamKienb/api-contract/gen/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToRegisterCommand(req *user.RegisterUserRequest) register_user.Command {
	profile := req.GetProfile()

	return register_user.Command{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Profile: register_user.UserProfile{
			FullName: profile.GetFullName(),
			Gender:   profile.GetGender(),
		},
	}
}

func ToRegisterResponse(result *register_user.Result) *user.RegisterUserResponse {
	return &user.RegisterUserResponse{
		SessionToken: result.SessionToken,
		ExpiresAt:    timestamppb.New(result.ExpiresAt),
	}
}

func ToLoginCommand(req *user.LoginUserRequest) login_user.Command {
	return login_user.Command{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
}

func ToLoginResponse(result *login_user.Result) *user.LoginUserResponse {
	return &user.LoginUserResponse{
		AccessToken:      result.AccessToken,
		RefreshToken:     result.RefreshToken,
		AccessExpiresAt:  timestamppb.New(result.AccessTokenExpiresAt),
		RefreshExpiresAt: timestamppb.New(result.RefreshTokenExpiresAt),
	}
}

func ToAddAddressCommand(userID string, req *user.AddUserAddressRequest) (add_user_address.Command, error) {
	parsedUserID, err := shared.ParseToRawID[shared.UserID](userID)
	if err != nil {
		return add_user_address.Command{}, account.ErrUserInvalid

	}
	return add_user_address.Command{
		UserID: parsedUserID,

		Country:  toUserLocationInfo(req.GetCountry()),
		City:     toUserLocationInfo(req.GetCity()),
		District: toUserLocationInfo(req.GetDistrict()),
		Ward:     toUserLocationInfo(req.GetWard()),

		AddressLine:  req.GetAddressLine(),
		ReceiverName: req.GetReceiverName(),
		PhoneNumber:  req.GetPhoneNumber(),
		Label:        req.GetLabel(),
		IsDefault:    req.GetIsDefault(),
	}, nil
}

func ToAddAddressResponse(result *add_user_address.Result) *user.AddUserAddressResponse {
	return &user.AddUserAddressResponse{
		AddressId: result.UserAddressID,
	}
}

type userLocationSource interface {
	GetId() int64
	GetName() string
}

func toUserLocationInfo(src userLocationSource) add_user_address.LocationInfo {
	if src == nil {
		return add_user_address.LocationInfo{}
	}

	return add_user_address.LocationInfo{
		ID:   int(src.GetId()),
		Name: src.GetName(),
	}
}
