package cmd

import (
	"github.com/rinsyan0518/adbkitting/interface/gateway"
	"github.com/rinsyan0518/adbkitting/interface/presenter"
	"github.com/rinsyan0518/adbkitting/usecase"
	"github.com/spf13/cobra"
)

func init() {
}

// NewCmdShutdownAll is command
func NewCmdShutdownAll() *cobra.Command {
	var uc usecase.ShuttingDownAllDevicesInputPort

	cmd := &cobra.Command{
		Use:   "shutdownall",
		Short: "shut down all connected devices",
		Run: func(cmd *cobra.Command, args []string) {
			uc.Handle()
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	uc = usecase.NewShuttingDownAllDevicesUsecase(
		presenter.NewConsolePresenter(cmd),
		gateway.NewAdbGateway())

	return cmd
}
