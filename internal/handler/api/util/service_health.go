package util

import (
	"net/http"

	"github.com/go-chi/chi"
	utilDomain "github.com/rudianto-dev/gotemp-api-gw/internal/domain/util"
	res "github.com/rudianto-dev/gotemp-sdk/pkg/response"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UtilHandler) GetServiceHealth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	request := utilDomain.GetServiceHealthRequest{
		Service: chi.URLParam(r, "service"),
	}
	err := utils.ValidateStruct(request)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_HANDLER_STAGE, err.Error())
		res.Nay(w, r, err)
		return
	}

	resp, err := s.utilUseCase.GetServiceHealth(ctx, request)
	if err != nil {
		s.logger.InfoWithContext(ctx, utils.ERROR_HANDLER_STAGE, err.Error())
		res.Nay(w, r, err)
		return
	}
	res.Yay(w, r, http.StatusOK, resp)
}
