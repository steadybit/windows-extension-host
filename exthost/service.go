package exthost

import (
	"github.com/kardianos/service"
	"github.com/rs/zerolog/log"
)

type ExtensionService struct {
}

func NewExtensionService() (*ExtensionService, error) {
	svcConfig := &service.Config{
		Name:        "SteadybitExtensionHostWindows",
		DisplayName: "Steadybit Extension Host Windows",
		Description: "The Steadybit Extension Host for Windows hosts.",
	}

	prg := &ExtensionService{}

	_, err := service.New(prg, svcConfig)
	if err != nil {
		return nil, err
	}
	return prg, nil
}

func (p *ExtensionService) Start(s service.Service) error {
	log.Info().Msgf("Service start call received")
	return nil
}

func (p *ExtensionService) Stop(s service.Service) error {
	return nil
}
