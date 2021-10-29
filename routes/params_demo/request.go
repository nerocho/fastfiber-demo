package params_demo

type PersonGetRequest struct {
	Name     string   `query:"name" json:"name"`
	Pass     string   `query:"pass" json:"pass"`
	Products []string `query:"products" json:"products"`
}

type PersonPostRequest struct {
	Name string `json:"name" xml:"name" form:"name"`
	Pass string `json:"pass" xml:"pass" form:"pass"`
}
