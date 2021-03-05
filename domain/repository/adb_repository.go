package repository

import "github.com/rinsyan0518/adbkitting/domain/model"

// AdbRepository connect to Android Debice Bridge
type AdbRepository interface {
	FindAll() ([]*model.AndroidDevice, error)
	InstallApk(*model.AndroidDevice, string) error
	Shutdown(*model.AndroidDevice) error
	Reboot(*model.AndroidDevice) error
}
