package tests

import (
	"focus-finance/src/controller/make_request"
	"focus-finance/src/controller/movements"
	"focus-finance/src/models/repository"
	"focus-finance/src/models/repository/entity"
	"focus-finance/src/models/service/movement_service"
	"focus-finance/src/tests/connection"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

var (
	db                 *gorm.DB
	MovementController movements.MovementControllerInterface
)

func TestMain(m *testing.M) {
	closeConnection := func() {}
	db, closeConnection = connection.OpenConnection()

	repo := repository.NewMovementRepository(db)
	service := movement_service.NewMovementService(repo)
	MovementController = movements.NewMovementController(service)

	defer closeConnection()

	err := db.AutoMigrate(&entity.Movement{}, &entity.User{}, &entity.Type{}, &entity.Operation{})
	if err != nil {
		log.Fatalf("Error trying to create migrate")
	}

	os.Exit(m.Run())
}

func TestFindMovementController(t *testing.T) {
	movementsDomain := entity.Movement{
		Value:       15,
		Description: "TESTE",
		UserID:      1,
		TypeID:      1,
		OperationID: 1,
	}

	user := entity.User{
		Name:     "TESTE",
		Email:    "TESTE@GMAI.COM",
		Password: "123456",
	}

	typeEntity := entity.Type{
		Name: "TESTE",
	}

	operation := entity.Operation{
		Name: "TESTE",
	}

	err := db.Create(&typeEntity).Error
	if err != nil {
		log.Fatal("error ao criar type")
	}

	err = db.Create(&operation).Error
	if err != nil {
		log.Fatal("error ao criar operation")
	}

	err = db.Create(&user).Error
	if err != nil {
		log.Fatal("error ao criar user")
	}

	err = db.Create(&movementsDomain).Error
	if err != nil {
		log.Fatal("error ao criar movement")
	}

	t.Run("find_movement_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := make_request.GetTestGinContext(recorder)
		param := []gin.Param{
			{"userID",
				"1",
			},
		}

		make_request.MakeRequest(ctx, param, url.Values{}, "GET", nil)
		MovementController.GetMovement(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})

}
