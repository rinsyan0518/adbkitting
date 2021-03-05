package usecase

import "github.com/rinsyan0518/adbkitting/domain/repository"

// ShuttingDownAllDevicesInputPort is input port
type ShuttingDownAllDevicesInputPort interface {
	Handle() error
}

// ShuttingDownAllDevicesOutputPort is output port
type ShuttingDownAllDevicesOutputPort interface {
	Printfln(format string, i ...interface{})
	Errorfln(format string, i ...interface{})
}

type shuttingDownAllDevicesUsecase struct {
	output  ShuttingDownAllDevicesOutputPort
	adbRepo repository.AdbRepository
}

// NewShuttingDownAllDevicesUsecase returns usecase
func NewShuttingDownAllDevicesUsecase(output ShuttingDownAllDevicesOutputPort, adbRepo repository.AdbRepository) ShuttingDownAllDevicesInputPort {
	return &shuttingDownAllDevicesUsecase{
		output:  output,
		adbRepo: adbRepo,
	}
}

func (uc *shuttingDownAllDevicesUsecase) Handle() error {
	devices, err := uc.adbRepo.FindAll()

	if err != nil {
		uc.output.Errorfln("Failed to collect android device: %s", err)
		return err
	}

	if len(devices) == 0 {
		uc.output.Printfln("No connected device")
		return nil
	}

	for _, d := range devices {
		uc.output.Printfln("Shutdown %s device(%s)", d.GetSerial(), d.GetModel())
		err = uc.adbRepo.Shutdown(d)
		if err != nil {
			uc.output.Errorfln("Failed to shut down device(%s): %s", d.GetSerial(), err)
			continue
		}
	}

	return nil
}
