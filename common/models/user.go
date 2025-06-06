package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	AutoID          uint    `gorm:"primaryKey;autoIncrement"`
	ID              string  `json:"id" gorm:"uniqueIndex;not null"`
	Email           string  `json:"email"`
	Username        string  `json:"username"`
	AvatarURL       string  `json:"avatar_url"`
	Bio             string  `json:"bio"`
	TotalWorkouts   int     `json:"total_workouts"`
	TotalDistanceKm float64 `json:"total_distance_km"`
	TotalCalories   int     `json:"total_calories"`
	IsMetric        bool    `json:"is_metric"`
	IsPublic        bool    `json:"is_public"`
	City            string  `json:"city"`
}
type Training struct {
	gorm.Model
	UserID      string      `json:"user_id"`
	StartTime   time.Time   `json:"start_time"`
	EndTime     time.Time   `json:"end_time"`
	WorkoutData WorkoutData `json:"workout_data" gorm:"type:json"`
}
type WorkoutData struct {
	DeviceType     string     `json:"device_type"`
	DurationSec    int        `json:"duration_sec"`
	DistanceMeters *float64   `json:"distance_meters,omitempty"`
	SpeedAvgKmh    float64    `json:"speed_avg_kmh"`
	SpeedMaxKmh    float64    `json:"speed_max_kmh"`
	Calories       *int       `json:"calories,omitempty"`
	HeartRateAvg   *int       `json:"heart_rate_avg,omitempty"`
	HeartRateMax   *int       `json:"heart_rate_max,omitempty"`
	PowerAvgWatts  *float64   `json:"power_avg_watts,omitempty"`
	GpsPoints      []GpsPoint `json:"-" gorm:"-"`
}

type GpsPoint struct {
	gorm.Model
	WorkoutDataID uint      `json:"workout_data_id" gorm:"index"`
	Lat           float64   `json:"lat"`
	Lng           float64   `json:"lng"`
	Time          time.Time `json:"time"`
}
