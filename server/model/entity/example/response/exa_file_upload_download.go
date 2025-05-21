package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/entity/example"
)

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
