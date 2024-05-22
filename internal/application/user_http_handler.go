package application

import (
	"github.com/gin-gonic/gin"
	"github.com/techerjeansebastienpro/go-hexa-example/internal/domain"
)

type UserHttpHandler struct {
	userInput domain.UsersInput
	app       *gin.Engine
}

func NewUserHttpHandler(app *gin.Engine, userInput domain.UsersInput) *UserHttpHandler {
	return &UserHttpHandler{
		userInput: userInput,
		app:       app,
	}
}

func (u *UserHttpHandler) RegisterRoutes() {
	u.app.GET("/users/:id", u.GetByID)
	u.app.POST("/users", u.Create)
}

func (u *UserHttpHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	user, err := u.userInput.GetByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func (u *UserHttpHandler) Create(c *gin.Context) {
	var createUser domain.CreateUser
	if err := c.ShouldBindJSON(&createUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := u.userInput.CreateOne(createUser)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, &UserDTO{
		ID:    user.ID,
		Email: user.Email,
	})
}
