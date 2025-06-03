package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"people-matching-api/model"
	"people-matching-api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// GetUser godoc
// @Summary Get list user
// @Description Return get list user
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} model.UserResponse "List of users"
// @Router /api/users [get]
func (userHandler *UserHandler) GetUser(ctx *gin.Context) {
	response, err := userHandler.UserService.GetUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed get data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success get data",
		"data":    response,
	})
}

// RegisterUser godoc
// @Summary Register user
// @Description Register user by providing user data
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.UserRequest true "User Data"
// @Success 200 {object} map[string]interface{} "Register user"
// @Failure 400 {object} map[string]interface{} "Error Request Body"
// @Router /api/users [post]
func (userHandler *UserHandler) RegisterUser(ctx *gin.Context) {
	var user model.UserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error request body",
		})
		return
	}

	err := userHandler.UserService.RegisterUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error register user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success register user",
	})
}

// MatchUser godoc
// @Summary Match user
// @Description Return match user
// @Tags users
// @Param age query int false "Age of the user"
// @Param friends query string false "Friends of the user"
// @Accept json
// @Produce json
// @Success 200 {object} model.MatchUserResponse "Match user"
// @Router /api/match_user [get]
func (userHandler *UserHandler) MatchUser(ctx *gin.Context) {
	age := ctx.Query("age")
	friends := ctx.Query("friends")

	ageInt, err := strconv.Atoi(age)
	if err != nil {
		log.Fatal(err)
	}

	response, err := userHandler.UserService.MatchUser(ageInt, friends)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "user not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success match user",
		"data":    response,
	})
}
