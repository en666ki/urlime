package result

import (
	"github.com/en666ki/urlime/internal/shortener/viewmodels"
)

type Result struct {
	Message string            `json:"msg,omitempty"`
	Code    int               `json:"code"`
	Data    *viewmodels.UrlVM `json:"data,omitempty"`
}
