package user

import (
	"user-query-module/internal/application/service/models"

	api "github.com/iamKienb/api-contract/gen/user"
)

func ToUserView(user *models.User) *api.UserDetail {
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

func ToUserViews(users []models.User) []*api.UserDetail {
	views := make([]*api.UserDetail, 0, len(users))
	for i := range users {
		views = append(views, ToUserView(&users[i]))
	}
	return views
}

func ToUserProfileView(profile *models.UserProfile) *api.ProfileDetail {
	if profile == nil {
		return nil
	}
	return &api.ProfileDetail{
		FullName: profile.FullName,
		Gender:   profile.Gender,
	}
}

func ToUserAddressViews(addresses []models.UserAddress) []*api.AddressDetail {
	views := make([]*api.AddressDetail, 0, len(addresses))
	for _, address := range addresses {
		views = append(views, &api.AddressDetail{
			AddressId:    address.ID,
			ReceiverName: address.ReceiverName,
			PhoneNumber:  address.PhoneNumber,
			ProvinceId:   address.Province.ID,
			ProvinceName: address.Province.Name,
			WardId:       address.Ward.ID,
			WardName:     address.Ward.Name,
			AddressLine:  address.AddressLine,
			Label:        address.Label,
			IsDefault:    address.IsDefault,
		})
	}
	return views
}
