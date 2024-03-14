package user

import (
	"net/http"

	"github.com/go-chi/chi"
	userDomain "github.com/rudianto-dev/gotemp-api-gw/internal/domain/user"
	res "github.com/rudianto-dev/gotemp-sdk/pkg/response"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

func (s *UserHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	request := userDomain.DeleteAccountRequest{
		ID: chi.URLParam(r, "id"),
	}
	err := utils.ValidateStruct(request)
	if err != nil {
		s.logger.ErrorWithContext(ctx, utils.ERROR_HANDLER_STAGE, err.Error())
		res.Nay(w, r, err)
		return
	}

	resp, err := s.userUseCase.DeleteAccount(ctx, request)
	if err != nil {
		s.logger.InfoWithContext(ctx, utils.ERROR_HANDLER_STAGE, err.Error())
		res.Nay(w, r, err)
		return
	}
	res.Yay(w, r, http.StatusOK, resp)
}
