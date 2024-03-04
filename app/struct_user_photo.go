package app

type UserPhotoData struct {
	Title    string `json:"title" valid:"required"`
	Caption  string `json:"caption" valid:"required"`
	PhotoURL string `json:"photo_url" valid:"required"`
}
