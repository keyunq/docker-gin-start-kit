package docs

// User Info
//
// swagger:response AuthResponse
type AuthResponseWapper struct {
	// in: body
	Body AuthResponse
}
type AuthResponse struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}
