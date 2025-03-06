package exthost

import (
	"github.com/rs/zerolog/log"
	"golang.org/x/sys/windows/svc"
	"time"
)

type ExtensionService struct {
	stopHandler func()
}

func (s *ExtensionService) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown
	changes <- svc.Status{State: svc.StartPending}
	tick := time.Tick(500 * time.Millisecond)
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
	for {
		select {
		case <-tick:
			log.Trace().Msg("tick")
		case c := <-r:
			switch c.Cmd {
			case svc.Stop, svc.Shutdown:
				changes <- svc.Status{State: svc.Stopped, Accepts: cmdsAccepted}
				log.Fatal().Msg("Received Windows service stop command, exiting")
			default:
				log.Info().Msgf("unexpected control request #%d", c)
			}
		}
	}
}

func NewExtensionService(stopHandler func()) error {
	ext := &ExtensionService{
		stopHandler: stopHandler,
	}
	return svc.Run("SteadybitExtensionHostWindows", ext)
}
