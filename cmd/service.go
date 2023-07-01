package cmd

import "github.com/meivaldi/gaplek/internal/service"

func SetupService() *service.Service {
	jitterSvc := service.NewJitterService()

	return &service.Service{
		JitterService: jitterSvc,
	}
}
