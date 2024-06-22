package middlewaresUsecases

import _pkgMiddlewaresMiddlewaresRepositories "github.com/MarkTBSS/072_JWT_Authenthication/modules/middlewares/middlewaresRepositories"

type IMiddlewaresUsecase interface {
	FindAccessToken(userId, accessToken string) bool
}

type middlewaresUsecase struct {
	middlewaresRepository _pkgMiddlewaresMiddlewaresRepositories.IMiddlewaresRepository
}

func MiddlewaresUsecase(middlewaresRepository _pkgMiddlewaresMiddlewaresRepositories.IMiddlewaresRepository) IMiddlewaresUsecase {
	return &middlewaresUsecase{
		middlewaresRepository: middlewaresRepository,
	}
}

func (u *middlewaresUsecase) FindAccessToken(userId, accessToken string) bool {
	return u.middlewaresRepository.FindAccessToken(userId, accessToken)
}
