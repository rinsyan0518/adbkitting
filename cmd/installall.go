package cmd

import (
	"os"

	"github.com/rinsyan0518/adbkitting/interface/gateway"
	"github.com/rinsyan0518/adbkitting/interface/presenter"
	"github.com/rinsyan0518/adbkitting/usecase"
	"github.com/spf13/cobra"
)

type installAllOptions struct {
	reboot bool
}

func init() {
}

// NewCmdInstallAll is command
func NewCmdInstallAll() *cobra.Command {
	var uc usecase.InstallingApkAllDevicesInputPort
	o := &installAllOptions{}

	cmd := &cobra.Command{
		Use:   "installall",
		Short: "install apk to all connected devices.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.PrintErrln("need one command line argument")
				cmd.PrintErrln("Usage: ")
				cmd.PrintErrln("    adbkitting installall APK_PATH")
				os.Exit(1)
			}

			apk := args[0]
			if f, err := os.Stat(apk); os.IsNotExist(err) || f.IsDir() {
				cmd.PrintErrf("%s is not found\n", apk)
				os.Exit(1)
			}

			uc.Handle(apk, o.reboot)
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	uc = usecase.NewInstallingApkAllDeviceUsecase(
		presenter.NewConsolePresenter(cmd),
		gateway.NewAdbGateway())
	cmd.Flags().BoolVarP(&o.reboot, "reboot", "r", false, "reboot device after install APK")

	return cmd
}
