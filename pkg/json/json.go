package json

import (
	"encoding/json"
)

type Host struct {
	IP   string
	Name string
}

func Convert2Json(h *Host) (string, error) {
	b, err := json.Marshal(h)
	return string(b), err
}


