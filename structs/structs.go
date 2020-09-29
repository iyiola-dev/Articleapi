package structs

type Response struct {
	code int         `json:"code"`
	body interface{} `json:"body"`
}
