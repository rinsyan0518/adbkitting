package gateway

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/rinsyan0518/adbkitting/domain/model"
)

// AdbGateway for Android Device Bridge
type AdbGateway struct {
	adb string
}

// NewAdbGateway returns new AdbGateway
func NewAdbGateway() *AdbGateway {
	return &AdbGateway{
		adb: "adb",
	}
}

// FindAll returns all connected Android devices.
func (a *AdbGateway) FindAll() ([]*model.AndroidDevice, error) {
	cmd := exec.Command(a.adb, "devices", "-l")
	r, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err = cmd.Start(); err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	devices := []*model.AndroidDevice{}
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if s == "" || strings.HasPrefix(s, "* ") || strings.HasPrefix(s, "List of devices") {
			continue
		}
		fields := strings.Fields(s)
		// 0: serial
		// 1; device
		// 2: usb:XXXXX
		// 3: product:XXXXX
		// 4: model:XXXXX
		// 5: device:XXXXX
		// 6: transport_id:[0-9]+
		d := model.NewAndroidDevice(
			fields[0],
			strings.Split(fields[3], ":")[1],
			strings.Split(fields[4], ":")[1],
		)
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

// InstallApk install apk into device
func (a *AdbGateway) InstallApk(device *model.AndroidDevice, apk string) error {
	cmd := exec.Command("adb", "-s", device.GetSerial(), "install", "-r", apk)
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}

// Shutdown run reboot device
func (a *AdbGateway) Shutdown(device *model.AndroidDevice) error {
	cmd := exec.Command("adb", "-s", device.GetSerial(), "reboot", "-p")
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}

// Reboot run reboot device
func (a *AdbGateway) Reboot(device *model.AndroidDevice) error {
	cmd := exec.Command("adb", "-s", device.GetSerial(), "reboot")
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}
