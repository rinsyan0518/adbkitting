package cmd

import (
	"bufio"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

type options struct {
	reboot bool
}

type device struct {
	serial string
	name   string
}

var (
	o = &options{}
)

// NewCmdInstallAll is command
func NewCmdInstallAll() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "installall",
		Short:         "install apk to all connected devices.",
		Run:           run,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	cmd.Flags().BoolVarP(&o.reboot, "reboot", "r", false, "reboot device after install APK")

	return cmd
}

func init() {
}

func run(cmd *cobra.Command, args []string) {
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

	devices, err := collectAndroidDevice()
	if err != nil {
		cmd.PrintErrln("failed collect android device: %s", err)
		os.Exit(1)
	}

	if len(devices) == 0 {
		cmd.PrintErrln("no connected devices")
		os.Exit(1)
	}

	for _, d := range devices {
		cmd.Printf("install %s to device(%s)\n", apk, d.serial)

		err = installApk(d.serial, apk)
		if err != nil {
			cmd.PrintErrf("failed install %s to device(%s)\n%s\n", apk, d.serial, err)
			continue
		}
		cmd.Println("install success")

		if o.reboot {
			cmd.Println("reboot device(%s)", d.serial)
			err = rebootDevice(d.serial)
			if err != nil {
				cmd.PrintErrf("failed reboot device(%s)\n%s\n", d.serial, err)
				continue
			}
		}
	}
}

func collectAndroidDevice() ([]device, error) {
	cmd := exec.Command("adb", "devices")
	r, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err = cmd.Start(); err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	devices := []device{}
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if s == "" || strings.HasPrefix(s, "* ") || strings.HasPrefix(s, "List of devices") {
			continue
		}
		fields := strings.Fields(s)
		d := device{
			serial: fields[0],
			name:   fields[1],
		}
		devices = append(devices, d)
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	if err = cmd.Wait(); err != nil {
		return nil, err
	}

	return devices, nil
}

func installApk(serial string, apk string) error {
	cmd := exec.Command("adb", "-s", serial, "install", "-r", apk)
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}

func rebootDevice(serial string) error {
	cmd := exec.Command("adb", "-s", serial, "reboot")
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}
