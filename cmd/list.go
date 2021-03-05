package cmd

import (
	"github.com/rinsyan0518/adbkitting/interface/gateway"
	"github.com/rinsyan0518/adbkitting/interface/presenter"
	"github.com/rinsyan0518/adbkitting/usecase"
	"github.com/spf13/cobra"
)

var (
	uc usecase.ListingDevicesInputPort
)

func init() {
}

// NewCmdList is command
func NewCmdList() *cobra.Command {
	var uc usecase.ListingDevicesInputPort

	cmd := &cobra.Command{
		Use:   "list",
		Short: "list connected devices",
		Run: func(cmd *cobra.Command, args []string) {
			uc.Handle()
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	uc = usecase.NewListingDevicesUsecase(
		presenter.NewConsolePresenter(cmd),
		gateway.NewAdbGateway())

	return cmd
}
