package usecases

import (
	"back/src/core/domain"
	"github.com/google/uuid"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type SetupUsecases interface {
	IsSetupDone() (bool, string)
}

type Usecases interface {
	SetupUsecases
}

type SettingsRepository interface {
	GetSettingsValue(key string) string
	SetSettingsValue(key string, value string)
}

type UserRepository interface {
	CountUsers() int64

	FindUserByID(id uuid.UUID) (domain.User, error)
}

type ServiceRepository interface {
}

type interactor struct {
	settingsRepo SettingsRepository
	userRepo     UserRepository
	serviceRepo  ServiceRepository

	translator          *i18n.Bundle
	remainingSetupSteps []string
}

func NewInteractor(sR SettingsRepository, uR UserRepository, servR ServiceRepository, i18n *i18n.Bundle) *interactor {
	inter := &interactor{
		settingsRepo: sR,
		userRepo:     uR,
		serviceRepo:  servR,
		translator:   i18n,
	}

	inter.prepareSetupSteps()

	return inter
}
