package userusecase

import "github.com/ferreiraHenrique/go-auth/core/domain"

func (usecase usecase) IsAdmin(signedToken string) bool {
	token, _ := domain.NewTokenFromString(signedToken)
	username := token.GetClaim("username").(string)

	user, error := usecase.repository.FindByUsername(username)
	if error != nil {
		return false
	}

	return usecase.repository.IsAdmin(user.ID)
}
