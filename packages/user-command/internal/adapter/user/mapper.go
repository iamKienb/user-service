package user

import (
	"strconv"

	"user-command-module/internal/application/commands/add_user_address"
	"user-command-module/internal/application/commands/login_user"
	"user-command-module/internal/application/commands/register_user"
	get_user_address_by_id "user-command-module/internal/application/queries/get_address_by_id"
	"user-command-module/internal/domain/shared"

	"github.com/iamKienb/api-contract/gen/user"
	"github.com/iamKienb/go-core/app_error"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toRegisterCommand(req *user.RegisterUserRequest) register_user.Command {
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

func toRegisterResponse(result *register_user.Result) *user.RegisterUserResponse {
	return &user.RegisterUserResponse{
		SessionToken: result.SessionToken,
		ExpiresAt:    timestamppb.New(result.ExpiresAt),
	}
}

func toLoginCommand(req *user.LoginUserRequest) login_user.Command {
	return login_user.Command{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
}

func toLoginResponse(result *login_user.Result) *user.LoginUserResponse {
	return &user.LoginUserResponse{
		AccessToken:      result.AccessToken,
		RefreshToken:     result.RefreshToken,
		AccessExpiresAt:  timestamppb.New(result.AccessTokenExpiresAt),
		RefreshExpiresAt: timestamppb.New(result.RefreshTokenExpiresAt),
	}
}

func toAddAddressCommand(userID string, req *user.AddUserAddressRequest) (add_user_address.Command, error) {
	parsedUserID, err := parseUserID(userID)
	if err != nil {
		return add_user_address.Command{}, err

	}
	return add_user_address.Command{
		UserID: parsedUserID,

		Country:  toUserLocationInfo(req.GetCountry()),
		Province: toUserLocationInfo(req.GetProvince()),
		Ward:     toUserLocationInfo(req.GetWard()),

		AddressLine:  req.GetAddressLine(),
		ReceiverName: req.GetReceiverName(),
		PhoneNumber:  req.GetPhoneNumber(),
		Label:        req.GetLabel(),
		IsDefault:    req.GetIsDefault(),
	}, nil
}

func toAddAddressResponse(result *add_user_address.Result) *user.AddUserAddressResponse {
	return &user.AddUserAddressResponse{
		AddressId: result.UserAddressID.String(),
	}
}

func toGetUserAddressQuery(req *user.GetUserAddressByIDRequest) (get_user_address_by_id.Query, error) {
	parsedUserID, err := parseUserID(req.GetUserId())
	if err != nil {
		return get_user_address_by_id.Query{}, err
	}

	parseUserAddressID, err := parseUserAddressID(req.GetUserAddressId())
	if err != nil {
		return get_user_address_by_id.Query{}, err
	}

	return get_user_address_by_id.Query{
		UserID:        parsedUserID,
		UserAddressID: parseUserAddressID,
	}, nil
}

func toGetUserAddressResponse(result *get_user_address_by_id.Result) *user.GetUserAddressByIDResponse {
	return &user.GetUserAddressByIDResponse{
		Address: &user.AddressDetail{
			AddressId:    result.UserAddressID.String(),
			ReceiverName: result.ReceiverName,
			PhoneNumber:  result.PhoneNumber,

			ProvinceId:   toInt32(result.ProvinceID),
			ProvinceName: result.ProvinceName,

			WardId:   toInt32(result.WardID),
			WardName: result.WardName,

			AddressLine: result.AddressLine,
			Label:       result.Label,
			IsDefault:   result.IsDefault,
		},
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
		ID:   strconv.FormatInt(src.GetId(), 10),
		Name: src.GetName(),
	}
}

func toInt32(value string) int32 {
	parsed, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0
	}
	return int32(parsed)
}

func parseUserID(value string) (shared.UserID, error) {
	parsed, err := shared.ParseToRawID[shared.UserID](value)
	if err != nil {
		return parsed, app_error.New(app_error.KindValidation, "user_invalid", "invalid user id", err)
	}

	return parsed, nil
}

func parseUserAddressID(value string) (shared.UserAddressID, error) {
	parsed, err := shared.ParseToRawID[shared.UserAddressID](value)
	if err != nil {
		return parsed, app_error.New(app_error.KindValidation, "user_address_invalid", "invalid user address id", err)
	}

	return parsed, nil
}
