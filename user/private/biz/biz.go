package biz

import (
	"github.com/jace996/multiapp/event"
	"github.com/jace996/multiapp/pkg/dal"
	kitdi "github.com/jace996/multiapp/pkg/di"
	"github.com/defval/di"
)

const ConnName dal.ConnName = "user"

// ProviderSet is biz providers.
var ProviderSet = kitdi.NewSet(
	NewUserManager,
	NewSignInManager,
	NewUserValidator,
	NewRoleManager,
	NewLookupNormalizer,

	kitdi.NewProvider(NewOtpTokenProvider, di.As(new(OtpTokenProvider))),
	NewEmailTokenProvider,
	NewPhoneTokenProvider,
	NewPasswordHasher,
	NewPasswordValidator,
	NewRoleSeed,
	NewUserSeed,
	NewPermissionSeeder,
	NewEmailSender,

	kitdi.NewProvider(NewUserRoleChangeEventHandler, di.As(new(event.ConsumerHandler))),
)

const UserAvatarPath = "user/avatar"
