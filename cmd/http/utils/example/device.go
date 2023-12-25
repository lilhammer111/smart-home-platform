package example

type AlertData struct {
	Id         *int32 `example:"1" json:"id,omitempty"`
	DeviceId   int32  `example:"1" json:"device_id,omitempty"`
	Count      int8   `example:"50" json:"count,omitempty"`
	Level      int8   `example:"1" json:"level,omitempty"`
	Desc       string `example:"hello,alert" json:"desc,omitempty"`
	FirstAlarm string `example:"2023-12-22 15:16:00" json:"first_alarm,omitempty"`
	LastAlarm  string `example:"2023-12-30 15:16:00" json:"last_alarm,omitempty"`
	IsOngoing  bool   `example:"true" json:"is_ongoing,omitempty"`
}

type DeviceData struct {
	Id              *int32 `example:"1" json:"id,omitempty"`
	OwnerId         int32  `example:"1" json:"owner_id,omitempty"`
	ProductId       int32  `example:"1" json:"product_id,omitempty"`
	SerialNo        string `example:"1234567890123456" json:"serial_no,omitempty"`
	Name            string `example:"test" json:"name,omitempty"`
	State           int8   `example:"1" json:"state,omitempty"`
	LocationId      string `example:"1" json:"location_id,omitempty"`
	HardwareVersion string `example:"1.0.0" json:"hardware_version,omitempty"`
	SoftwareVersion string `example:"1.0.0" json:"software_version,omitempty"`
	Desc            string `example:"test" json:"desc,omitempty"`
}
