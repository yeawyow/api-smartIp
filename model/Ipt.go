package model

type Ipt struct {
	An      string `json:"an" binding:"required"`
	Hn      string
	//Dchdate time.Time
	//Dchtype string
	Pttype string
	//Wardname string
	//Regdate string
	Fullname string
	//Regtime time.Time

}