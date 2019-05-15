package models

type BannerModel struct {
	ID        int
	BannerUrl string
}

func (BannerModel) TableName() string {
	return "home_banner"
}
