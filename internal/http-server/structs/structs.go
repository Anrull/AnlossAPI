package structs

type Record struct {
	Name    string `json:"name" form:"name" query:"name"`
	Class   string `json:"class" form:"class" query:"class"`
	Olimp   string `json:"olimp" form:"olimp" query:"olimp"`
	Sub     string `json:"sub" form:"sub" query:"sub"`
	Teacher string `json:"teacher" form:"teacher" query:"teacher"`
	Stage   string `json:"stage" form:"stage" query:"stage"`
	Date    string `json:"date" form:"date" query:"date"`
}

type GetRecord struct {
	Name    string `json:"name" form:"name" query:"name"`
	Sub     string `json:"sub" form:"sub" query:"sub"`
	Olimp   string `json:"olimp" form:"olimp" query:"olimp"`
	Stage   string `json:"stage" form:"stage" query:"stage"`
	Teacher string `json:"teacher" form:"teacher" query:"teacher"`
}

type Admin struct {
	Password string `json:"password" form:"password" query:"password"`
}

type CheckSnils struct {
	Snils string `json:"snils" form:"snils" query:"snils"`
}

type AddStudent struct {
	Name  string `json:"name" form:"name" query:"name"`
	Class string `json:"class" form:"class" query:"class"`
	Snils string `json:"snils" form:"snils" query:"snils"`
}
