package model

import "time"

type FeedingProgram struct {
	BaseModel
	Frequency int8  `gorm:"type:tinyint unsigned;index;not null" json:"frequency"`
	Food      int8  `gorm:"type:tinyint unsigned;not null" json:"food"`
	Amount    int16 `gorm:"type:smallint unsigned;not null" json:"amount"`
	DeviceID  int32 `gorm:"type:int unsigned;not null" json:"device_id"`
	// differ from FeedAt which is of type datetime in database
	FeedTime time.Time `gorm:"type:time;not null" json:"feed_time"`
}

type FeederStatus struct {
	DeviceId      int32 `gorm:"type:int unsigned;not null" json:"device_id"`
	FoodSurplus   int16 `gorm:"type:smallint unsigned;not null" json:"food_surplus"`
	BatteryCharge int8  `gorm:"type:tinyint unsigned;not null" json:"battery_charge"`
}

type FeedData struct {
	Success  bool  `gorm:"type:bool;not null" json:"success"`
	DeviceID int32 `gorm:"type:int unsigned;not null;index" json:"device_id"`
	Amount   int16 `gorm:"type:smallint unsigned;not null" json:"amount"`
	Food     int8  `gorm:"type:tinyint unsigned;not null" json:"food"`
	// Consider adding an index to the field FeedAt
	FeedAt time.Time `gorm:"type:datetime;not null" json:"feed_at"`
	Remark string    `gorm:"type:varchar(255);not null;comment:Provides additional details about the feeding process, including failure reasons, adjustments, and other observations." json:"remark"`
}

type TimedFeedingRecord struct {
	BaseModel
	FeedingProgramID int32 `gorm:"type:int unsigned;index" json:"feeding_program_id"`
	FeedData
}

type InstantFeedingRecord struct {
	BaseModel
	FeedData
}
