// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	context "context"
	errors "github.com/go-kratos/kratos/v2/errors"
	i18n "github.com/jace996/go-i18n/v2/i18n"
	localize "github.com/jace996/multiapp/pkg/localize"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsConfirmPasswordMismatch(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CONFIRM_PASSWORD_MISMATCH.String() && e.Code == 400
}

func ErrorConfirmPasswordMismatchLocalized(ctx context.Context, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	localizer := localize.FromContext(ctx)
	if localizer == nil {
		return errors.New(400, ErrorReason_CONFIRM_PASSWORD_MISMATCH.String(), "")
	}
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "ConfirmPasswordMismatch",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(400, ErrorReason_CONFIRM_PASSWORD_MISMATCH.String(), msg)
	} else {
		return errors.New(400, ErrorReason_CONFIRM_PASSWORD_MISMATCH.String(), "")
	}
}

func IsPasswordInsufficientStrength(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_PASSWORD_INSUFFICIENT_STRENGTH.String() && e.Code == 400
}

func ErrorPasswordInsufficientStrengthLocalized(ctx context.Context, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	localizer := localize.FromContext(ctx)
	if localizer == nil {
		return errors.New(400, ErrorReason_PASSWORD_INSUFFICIENT_STRENGTH.String(), "")
	}
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "PasswordInsufficientStrength",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(400, ErrorReason_PASSWORD_INSUFFICIENT_STRENGTH.String(), msg)
	} else {
		return errors.New(400, ErrorReason_PASSWORD_INSUFFICIENT_STRENGTH.String(), "")
	}
}

func IsInvalidPassword(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_INVALID_PASSWORD.String() && e.Code == 400
}

func ErrorInvalidPasswordLocalized(ctx context.Context, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	localizer := localize.FromContext(ctx)
	if localizer == nil {
		return errors.New(400, ErrorReason_INVALID_PASSWORD.String(), "")
	}
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "InvalidPassword",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(400, ErrorReason_INVALID_PASSWORD.String(), msg)
	} else {
		return errors.New(400, ErrorReason_INVALID_PASSWORD.String(), "")
	}
}

func IsDuplicateUsername(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DUPLICATE_USERNAME.String() && e.Code == 400
}

func ErrorDuplicateUsernameLocalized(ctx context.Context, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	localizer := localize.FromContext(ctx)
	if localizer == nil {
		return errors.New(400, ErrorReason_DUPLICATE_USERNAME.String(), "")
	}
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "DuplicateUsername",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(400, ErrorReason_DUPLICATE_USERNAME.String(), msg)
	} else {
		return errors.New(400, ErrorReason_DUPLICATE_USERNAME.String(), "")
	}
}

func IsDuplicateEmail(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DUPLICATE_EMAIL.String() && e.Code == 400
}

func ErrorDuplicateEmailLocalized(ctx context.Context, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	localizer := localize.FromContext(ctx)
	if localizer == nil {
		return errors.New(400, ErrorReason_DUPLICATE_EMAIL.String(), "")
	}
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "DuplicateEmail",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(400, ErrorReason_DUPLICATE_EMAIL.String(), msg)
	} else {
		return errors.New(400, ErrorReason_DUPLICATE_EMAIL.String(), "")
	}
}

func IsDuplicatePhone(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DUPLICATE_PHONE.String() && e.Code == 400
}

func ErrorDuplicatePhoneLocalized(ctx context.Context, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	localizer := localize.FromContext(ctx)
	if localizer == nil {
		return errors.New(400, ErrorReason_DUPLICATE_PHONE.String(), "")
	}
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "DuplicatePhone",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(400, ErrorReason_DUPLICATE_PHONE.String(), msg)
	} else {
		return errors.New(400, ErrorReason_DUPLICATE_PHONE.String(), "")
	}
}

func IsInvalidEmail(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_INVALID_EMAIL.String() && e.Code == 400
}

func ErrorInvalidEmailLocalized(ctx context.Context, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	localizer := localize.FromContext(ctx)
	if localizer == nil {
		return errors.New(400, ErrorReason_INVALID_EMAIL.String(), "")
	}
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "InvalidEmail",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(400, ErrorReason_INVALID_EMAIL.String(), msg)
	} else {
		return errors.New(400, ErrorReason_INVALID_EMAIL.String(), "")
	}
}

func IsInvalidPhone(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_INVALID_PHONE.String() && e.Code == 400
}

func ErrorInvalidPhoneLocalized(ctx context.Context, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	localizer := localize.FromContext(ctx)
	if localizer == nil {
		return errors.New(400, ErrorReason_INVALID_PHONE.String(), "")
	}
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "InvalidPhone",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(400, ErrorReason_INVALID_PHONE.String(), msg)
	} else {
		return errors.New(400, ErrorReason_INVALID_PHONE.String(), "")
	}
}

func IsInvalidUsername(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_INVALID_USERNAME.String() && e.Code == 400
}

func ErrorInvalidUsernameLocalized(ctx context.Context, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	localizer := localize.FromContext(ctx)
	if localizer == nil {
		return errors.New(400, ErrorReason_INVALID_USERNAME.String(), "")
	}
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "InvalidUsername",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(400, ErrorReason_INVALID_USERNAME.String(), msg)
	} else {
		return errors.New(400, ErrorReason_INVALID_USERNAME.String(), "")
	}
}

func IsUserNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_NOT_FOUND.String() && e.Code == 404
}

func ErrorUserNotFoundLocalized(ctx context.Context, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	localizer := localize.FromContext(ctx)
	if localizer == nil {
		return errors.New(404, ErrorReason_USER_NOT_FOUND.String(), "")
	}
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "UserNotFound",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(404, ErrorReason_USER_NOT_FOUND.String(), msg)
	} else {
		return errors.New(404, ErrorReason_USER_NOT_FOUND.String(), "")
	}
}
