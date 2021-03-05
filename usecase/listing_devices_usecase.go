package usecase

import "github.com/rinsyan0518/adbkitting/domain/repository"

// ListingDevicesInputPort is input port
type ListingDevicesInputPort interface {
	Handle() error
}

// ListingDevicesOutputPort is output port
type ListingDevicesOutputPort interface {
	Printfln(format string, i ...interface{})
	Errorfln(format string, i ...interface{})
}

type listingDevicesUsecase struct {
	output  ListingDevicesOutputPort
	adbRepo repository.AdbRepository
}

// NewListingDevicesUsecase returns usecase
func NewListingDevicesUsecase(output ListingDevicesOutputPort, adbRepo repository.AdbRepository) ListingDevicesInputPort {
	return &listingDevicesUsecase{
		output:  output,
		adbRepo: adbRepo,
	}
}

func (uc *listingDevicesUsecase) Handle() error {
	devices, err := uc.adbRepo.FindAll()
	if err != nil {
		uc.output.Errorfln("failed collect android device: %s", err)
		return err
	}

	if len(devices) == 0 {
		uc.output.Printfln("No connected")
		return nil
	}

	for _, d := range devices {
		uc.output.Printfln("Serial(%s) Model(%s) Product(%s)", d.GetSerial(), d.GetModel(), d.GetProduct())
	}

	return nil
}
