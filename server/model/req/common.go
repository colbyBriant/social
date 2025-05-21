package req

type StatusReq struct {
	ID    uint   `json:"id"`
	Value int8   `json:"value"`
	Field string `json:"field"`
}
