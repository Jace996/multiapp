package biz

import (
	"context"
	"github.com/jace996/multiapp/pkg/data"
	"github.com/jace996/multiapp/pkg/gorm"
	v1 "github.com/jace996/multiapp/user/api/account/v1"
)

// UserSetting contains key/value pairs of user settings
type UserSetting struct {
	gorm.UIDBase
	UserId string     `json:"user_id" gorm:"index"`
	Key    string     `json:"key" gorm:"index"`
	Value  data.Value `gorm:"embedded"`
}

type UpdateUserSetting struct {
	Key string
	//Value new value
	Value  *data.Value
	Delete bool
}

type UserSettingRepo interface {
	data.Repo[UserSetting, string, *v1.GetSettingsRequest]
	FindByUser(ctx context.Context, userId string, query *v1.GetSettingsRequest) ([]*UserSetting, error)
	UpdateByUser(ctx context.Context, userId string, updateBatch []UpdateUserSetting) error
}
