package models

type Pagelimit struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type Listrequest struct {
	Pagelimit
}

type Detailsrequest struct {
	Id int `form:"id" validate:"min=1,max=801"`
}
