package users

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/controller/make_request"
	"focus-finance/src/models"
	"focus-finance/src/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestFindUserController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserServiceInterface(ctrl)
	Controller := NewUserController(service)

	t.Run("id_is_null_return_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := make_request.GetTestGinContext(recorder)
		params := []gin.Param{
			{
				"id",
				"",
			},
		}

		make_request.MakeRequest(context, params, url.Values{}, "GET", nil)
		Controller.GetUser(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("id_not_null_returning_service_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := make_request.GetTestGinContext(recorder)
		params := []gin.Param{
			{
				"id",
				"1",
			},
		}

		service.EXPECT().GetUserByID(1).Return(nil, rest_err.NewInternalServerError("error"))
		make_request.MakeRequest(context, params, url.Values{}, "GET", nil)

		Controller.GetUser(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)

	})

	t.Run("id_not_null_returning_service_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := make_request.GetTestGinContext(recorder)
		params := []gin.Param{
			{
				"id",
				"1",
			},
		}

		service.EXPECT().GetUserByID(1).Return(models.NewUserDomain("TESTE", "TESTE@TESTE.COM", "123456"), nil)
		make_request.MakeRequest(context, params, url.Values{}, "GET", nil)

		Controller.GetUser(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)

	})
}
