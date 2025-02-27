package route

import (
	"time"

	"go-api-mock-test/api/controller"
	"go-api-mock-test/bootstrap"
	"go-api-mock-test/domain"
	"go-api-mock-test/mongo"
	"go-api-mock-test/repository"
	"go-api-mock-test/usecase"

	"github.com/gin-gonic/gin"
)

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
