package model

import (
	"gorm.io/gorm"
	"time"
)

type HeroSection struct {
	gorm.Model
	Title               string
	Subtitle            string
	HeroSectionCarousel []*HeroSectionCarousel `gorm:"many2many:hero_section_carousel"`
}
type HeroSectionCarousel struct {
	gorm.Model
	Topic string
	Title string
}

type ImageCarousel struct {
	gorm.Model
	Title              string
	Subtitle           string
	ImageCarouselImage []*ImageCarouselImage `gorm:"many2many:image_section_carousel"`
}

type ImageCarouselImage struct {
	gorm.Model
	Url string
}

type VideoSection struct {
	gorm.Model
	Url string
}

type ServiceSection struct {
	gorm.Model
	Title    string
	Subtitle string
}

type ImpactSection struct {
	gorm.Model
	Title           string
	Subtitle        string
	HoursTutored    string
	StudentsTutored string
}

type HomePage struct {
	gorm.Model
	PublishDate      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	IsActive         bool
	IsDraft          bool
	HeroSectionID    uint
	ImageCarouselID  uint
	ServiceSectionID uint
	VideoSectionID   uint
	ImpactSectionID  uint
	HeroSection      HeroSection
	ImageCarousel    ImageCarousel
	ServiceSection   ServiceSection
	ImpactSection    ImpactSection
	VideoSection     VideoSection
}
