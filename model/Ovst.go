package model

type Ovst struct {
	HosGuid string `json:"hos_guid"`
	Vn 		string `json:"vn"`
	An      string `json:"an" binding:"required"`
	Hn      string `json:"hn"`
	Vstdate string `json:"vstdate"`
	Vsttime string
	Pttype string
	Hospmain string
	Hospsub string
	Oqueue string
	//Regdate string
	Ovstist string
	Pttypeno string
	Spclty string
	Hcode string
	Ovstost string
	Cur_dep string
	PtSubtype string
	//Regtime time.Time

}