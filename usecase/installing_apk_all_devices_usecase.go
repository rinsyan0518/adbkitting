package usecase

import (
	"errors"

	"github.com/rinsyan0518/adbkitting/domain/repository"
)

// InstallingApkAllDevicesInputPort is input port
type InstallingApkAllDevicesInputPort interface {
	Handle(apk string, afterReboot bool) error
}

// InstallingApkAllDevicesOutputPort is output port
type InstallingApkAllDevicesOutputPort interface {
	Printfln(format string, i ...interface{})
	Errorfln(format string, i ...interface{})
}

type installingApkAllDevicesUsecase struct {
	output  InstallingApkAllDevicesOutputPort
	adbRepo repository.AdbRepository
}

// NewInstallingApkAllDeviceUsecase returns usecase
func NewInstallingApkAllDeviceUsecase(output InstallingApkAllDevicesOutputPort, adbRepo repository.AdbRepository) InstallingApkAllDevicesInputPort {
	return &installingApkAllDevicesUsecase{
		output:  output,
		adbRepo: adbRepo,
	}
}

func (uc *installingApkAllDevicesUsecase) Handle(apk string, afterReboot bool) error {
	devices, err := uc.adbRepo.FindAll()

	if err != nil {
		uc.output.Errorfln("Failed to collect android device: %s", err)
		return err
	}

	if len(devices) == 0 {
		uc.output.Errorfln("No connected device")
		return errors.New("NotConnected")
	}

	for _, d := range devices {
		uc.output.Printfln("Install %s to %s device(%s)", apk, d.GetModel(), d.GetSerial())

		err = uc.adbRepo.InstallApk(d, apk)
		if err != nil {
			uc.output.Errorfln("Failed to install APK: %s", err)
			continue
		}
		uc.output.Printfln("Success to install APK")

		if afterReboot {
			uc.output.Printfln("Reboot device")
			err = uc.adbRepo.Reboot(d)
			if err != nil {
				uc.output.Errorfln("Failed to reboot device(%s): %s", d.GetSerial(), err)
				continue
			}
		}
	}

	return nil
}
