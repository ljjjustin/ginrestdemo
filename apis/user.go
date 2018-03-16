package apis

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ljjjustin/ginrestdemo/models"
)

func init() {
	router.GET("/users", UserGetAll)
	router.POST("/users", UserCreate)
	router.GET("/users/:id", UserDetail)
	router.PUT("/users/:id", UserUpdate)
	router.DELETE("/users/:id", UserDelete)
}

func getId(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func UserGetAll(c *gin.Context) {
	var users []models.User

	errs := models.UserGetAll(&users)
	if len(errs) > 0 {
		c.JSON(500, NewError("failed to get all users."))
	} else {
		c.JSON(200, users)
	}
}

func UserDetail(c *gin.Context) {
	id, err := getId(c)
	if err != nil {
		c.JSON(400, NewError("invalid id sent"))
		return
	}

	var user models.User

	if models.UserFind(&user, id) {
		c.JSON(404, gin.H{"error": "not found"})
	} else {
		c.JSON(200, user)
	}
}

func UserCreate(c *gin.Context) {
	var user models.User

	if c.Bind(&user) != nil {
		c.JSON(400, NewError("invalid body"))
		return
	}
	models.UserSave(&user)

	c.JSON(201, user)
}

func UserUpdate(c *gin.Context) {
	var user models.User

	id, err := getId(c)
	if err != nil {
		c.JSON(400, NewError("invalid id sent"))
		return
	}

	if models.UserFind(&user, id) {
		c.JSON(404, gin.H{"error": "not found"})
	}

	if c.BindJSON(&user) != nil {
		c.JSON(400, NewError("invalid body"))
		return
	}

	models.UserSave(&user)
	c.JSON(200, user)
}

func UserDelete(c *gin.Context) {
	id, err := getId(c)
	if err != nil {
		c.JSON(400, NewError("invalid id sent"))
		return
	}

	var user models.User

	if models.UserFind(&user, id) {
		c.JSON(404, gin.H{"error": "not found"})
	} else {
		models.UserDelete(&user)
		c.Data(204, "application/json", make([]byte, 0))
	}
}
