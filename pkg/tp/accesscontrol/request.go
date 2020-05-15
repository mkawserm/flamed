package accesscontrol

type Request struct {
	Username  string `json:"name"`
	Namespace []byte `json:"namespace"`
}
