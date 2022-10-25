package api

type HttpAPIResponse struct {
	Code  int         `json:"code"`
	Error *string     `json:"error"`
	Data  interface{} `json:"_data"`
}
