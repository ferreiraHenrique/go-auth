package clientusecase

import "github.com/ferreiraHenrique/go-auth/core/dto"

func (usecase usecase) AttachPermission(clientRequest *dto.AttachClientPermissionRequest) error {
	client, err := usecase.clientRepository.FindByUUID(clientRequest.ClientUUID)
	if err != nil {
		return err
	}

	permission, err := usecase.permissionRepository.FindByUUID(clientRequest.PermissionUUID)
	if err != nil {
		return err
	}

	return usecase.clientRepository.AttachPermission(client, permission)
}
