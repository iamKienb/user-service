package user

import (
	"strconv"

	"user-query-module/internal/application/port"

	api "github.com/iamKienb/api-contract/gen/user"
)

func ToUserView(user *port.User) *api.UserDetail {
	if user == nil {
		return nil
	}
	return &api.UserDetail{
		Id:        user.ID,
		Email:     user.Email,
		Status:    user.Status,
		Roles:     user.Roles,
		Profile:   ToUserProfileView(user.Profile),
		Addresses: ToUserAddressViews(user.Addresses),
	}
}

func ToUserViews(users []port.User) []*api.UserDetail {
	views := make([]*api.UserDetail, 0, len(users))
	for i := range users {
		views = append(views, ToUserView(&users[i]))
	}
	return views
}

func ToUserProfileView(profile *port.UserProfile) *api.ProfileDetail {
	if profile == nil {
		return nil
	}
	return &api.ProfileDetail{
		FullName: profile.FullName,
		Gender:   profile.Gender,
	}
}

func ToUserAddressViews(addresses []port.UserAddress) []*api.AddressDetail {
	views := make([]*api.AddressDetail, 0, len(addresses))
	for _, address := range addresses {
		views = append(views, &api.AddressDetail{
			AddressId:    address.ID,
			ReceiverName: address.ReceiverName,
			PhoneNumber:  address.PhoneNumber,
			ProvinceId:   extraInt32(address.Extra, "province_id"),
			ProvinceName: extraString(address.Extra, "province_name"),
			WardId:       extraInt32(address.Extra, "ward_id"),
			WardName:     extraString(address.Extra, "ward_name"),
			AddressLine:  address.AddressLine,
			Label:        address.Label,
			IsDefault:    address.IsDefault,
		})
	}
	return views
}

func extraString(extra map[string]any, key string) string {
	value, ok := extra[key]
	if !ok || value == nil {
		return ""
	}
	switch typed := value.(type) {
	case string:
		return typed
	default:
		return ""
	}
}

func extraInt32(extra map[string]any, key string) int32 {
	value, ok := extra[key]
	if !ok || value == nil {
		return 0
	}
	switch typed := value.(type) {
	case int:
		return int32(typed)
	case int64:
		return int32(typed)
	case float64:
		return int32(typed)
	case string:
		parsed, err := strconv.ParseInt(typed, 10, 32)
		if err != nil {
			return 0
		}
		return int32(parsed)
	default:
		return 0
	}
}
