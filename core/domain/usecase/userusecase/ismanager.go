package userusecase

import "github.com/ferreiraHenrique/go-auth/core/domain"

func (usecase usecase) IsManager(signedToken string) (bool, *domain.Manager) {
	token, _ := domain.NewTokenFromString(signedToken)
	username := token.GetClaim("username").(string)

	user, error := usecase.repository.FindByUsername(username)
	if error != nil {
		return false, nil
	}

	isManager, manager := usecase.repository.IsManager(user.ID)
	if !isManager {
		return false, nil
	}

	return true, manager
}
