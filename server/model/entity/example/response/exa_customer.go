package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/entity/example"
)

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
