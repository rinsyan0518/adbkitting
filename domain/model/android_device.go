package model

// AndroidDevice is data
type AndroidDevice struct {
	serial  string
	product string
	model   string
}

// NewAndroidDevice create androidDevice instance
func NewAndroidDevice(serial string, product string, model string) *AndroidDevice {
	return &AndroidDevice{
		serial:  serial,
		product: product,
		model:   model,
	}
}

// GetSerial returns Android Device Serial
func (ad AndroidDevice) GetSerial() string {
	return ad.serial
}

// GetProduct returns Android Device Product Name
func (ad AndroidDevice) GetProduct() string {
	return ad.product
}

// GetModel returns Android Device Model Name
func (ad AndroidDevice) GetModel() string {
	return ad.model
}
