package service

import (
	"context"
	"github.com/jace996/multiapp/pkg/authz/authz"
	"github.com/jace996/multiapp/user/private/biz"
	"github.com/jace996/saas/data"
)

type UserRoleContrib struct {
	um *biz.UserManager
}

func NewUserRoleContrib(um *biz.UserManager) *UserRoleContrib {
	return &UserRoleContrib{um: um}
}

func (u *UserRoleContrib) Process(ctx context.Context, subject authz.Subject) ([]authz.Subject, error) {
	if us, ok := authz.ParseUserSubject(subject); ok {
		if us.GetUserId() != "" {
			//TODO ?
			ctx = data.NewMultiTenancyDataFilter(ctx, false)
			roles, err := u.um.GetUserRoleIds(ctx, us.GetUserId(), false)
			if err != nil {
				return nil, err
			}

			roleSubjects := make([]authz.Subject, len(roles))
			for i, r := range roles {
				roleSubjects[i] = authz.NewRoleSubject(r.RoleId)
			}
			return roleSubjects, nil
		}
	}
	return nil, nil
}

var _ authz.SubjectContrib = (*UserRoleContrib)(nil)
