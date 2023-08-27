package di

import (
	"github.com/ferreiraHenrique/go-auth/adapter/http/clientservice"
	"github.com/ferreiraHenrique/go-auth/adapter/sqlite"
	"github.com/ferreiraHenrique/go-auth/adapter/sqlite/clientrepository"
	"github.com/ferreiraHenrique/go-auth/core/domain"
	"github.com/ferreiraHenrique/go-auth/core/domain/usecase/clientusecase"
)

func ConfigClientDI(conn sqlite.PoolInterface) domain.ClientService {
	clientRepository := clientrepository.New(conn)
	clientUseCase := clientusecase.New(clientRepository)
	clientService := clientservice.New(clientUseCase)

	return clientService
}
