package biz

import (
	"context"
	"github.com/jace996/multiapp/pkg/data"
	"github.com/jace996/multiapp/pkg/gorm"
	"github.com/jace996/multiapp/pkg/idgen"
	"github.com/jace996/multiapp/pkg/utils"
	v1 "github.com/jace996/multiapp/realtime/api/notification/v1"
	v12 "github.com/jace996/multiapp/realtime/event/v1"
	gorm2 "github.com/jace996/saas/gorm"
	"time"
)

type Notification struct {
	ID string `gorm:"type:char(36)" json:"id"`
	gorm.AuditedModel
	gorm2.MultiTenancy
	Group    string
	Title    string
	Desc     string
	SendTime *time.Time
	Image    string
	Link     string
	Source   string
	UserId   string `json:"user_id" gorm:"index:,type:char(36)"`
	HasRead  bool
	Extra    data.JSONMap
	Level    int32
}

func NewNotification() *Notification {
	id, _ := (&idgen.Ksuid{}).Gen(context.Background())
	ret := &Notification{
		ID: id,
	}
	return ret
}

func FromNotificationEvents(event *v12.NotificationEvent) []*Notification {
	ret := make([]*Notification, len(event.UserIds))
	for i, id := range event.UserIds {
		r := NewNotification()
		ret[i] = r
		r.UserId = id
		r.MultiTenancy.TenantId = gorm2.NewTenantId(event.TenantId)
		r.Group = event.Group
		r.Title = event.Title
		r.Desc = event.Desc
		r.Image = event.Image
		r.Link = event.Link
		r.Source = event.Source
		r.Extra = utils.Structpb2Map(event.Extra)
		r.Level = int32(event.Level)
	}
	return ret
}

type NotificationRepo interface {
	data.Repo[Notification, string, *v1.ListNotificationRequest]
	MyUnreadCount(ctx context.Context) (int32, error)
	SetMyRead(ctx context.Context, idFilter string) error
	DeleteMy(ctx context.Context, id string) error
}
