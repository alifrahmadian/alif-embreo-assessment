package handlers

import (
	"net/http"

	"github.com/alifrahmadian/alif-embreo-assessment/internal/handlers/dtos"
	"github.com/alifrahmadian/alif-embreo-assessment/internal/handlers/responses"
	"github.com/alifrahmadian/alif-embreo-assessment/internal/services"
	"github.com/gin-gonic/gin"
)

type VendorHandler struct {
	VendorService services.VendorService
}

func NewVendorHandler(vendorService *services.VendorService) *VendorHandler {
	return &VendorHandler{
		VendorService: *vendorService,
	}
}

func (h *VendorHandler) GetVendors(c *gin.Context) {
	vendors, err := h.VendorService.GetVendors()
	if err != nil {
		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	responseData := make([]*dtos.GetVendorResponse, len(vendors))

	for i, vendor := range vendors {
		responseData[i] = &dtos.GetVendorResponse{
			ID:   *vendor.ID,
			Name: *vendor.Name,
		}
	}

	responses.SuccessResponse(c, "get vendors successful", responseData)
}
