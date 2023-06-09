package crudapi

import (
	"fmt"
	"github.com/multiheaded/go-WebserviceTemplate/pkg/datamodel"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GinRouteCRUD[T any](grp *gin.RouterGroup, hnd Handler[gin.HandlerFunc]) {
	var inst T
	typeName := strings.ToLower(reflect.TypeOf(inst).Name())

	typeRoute := fmt.Sprintf("/%s", typeName)
	idRoute := typeRoute + "/:id"

	grp.POST(typeRoute, hnd.Create()) // Create
	grp.GET(typeRoute, hnd.ReadAll()) // ReadAll
	grp.GET(idRoute, hnd.ReadOne())   // ReadOne
	grp.PUT(idRoute, hnd.Update())    // Update
	grp.DELETE(idRoute, hnd.Delete()) // Delete
}

// GinCRUDHandler implements type Handler[F any] interface
type GinCRUDHandler[T any] struct {
	Repository datamodel.CRUDRepository[T]
}

func NewGinCRUDHandler[T any](repo datamodel.CRUDRepository[T]) Handler[gin.HandlerFunc] {
	return GinCRUDHandler[T]{
		Repository: repo,
	}
}

func parseIdParameter(c *gin.Context) (uint64, error) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println("Could not convert id from string to int, ", err)
		return ^uint64(0), err
	} else {
		return id, nil
	}
}

func (hnd GinCRUDHandler[T]) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := new(T)

		err := c.BindJSON(&requestBody)
		if err != nil {
			fmt.Println("Could not parse JSON request, ", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		obj, err := hnd.Repository.Create(*requestBody)
		if err != nil {
			c.JSON(http.StatusInternalServerError, obj)
		} else {
			c.JSON(http.StatusOK, obj)
		}
	}
}

func (hnd GinCRUDHandler[T]) ReadAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		var objs []T
		err := hnd.Repository.ReadAll(&objs)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
		}

		c.JSON(http.StatusOK, objs)
	}
}

func (hnd GinCRUDHandler[T]) ReadOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parseIdParameter(c)
		if err != nil {
			c.AbortWithError(http.StatusUnprocessableEntity, err)
			return
		}

		var obj T
		err = hnd.Repository.ReadOne(id, &obj)

		if err != nil {
			c.JSON(http.StatusInternalServerError, obj)
		} else {
			c.JSON(http.StatusOK, obj)
		}
	}
}

func (hnd GinCRUDHandler[T]) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parseIdParameter(c)
		if err != nil {
			c.AbortWithError(http.StatusUnprocessableEntity, err)
			return
		}

		requestBody := new(T)

		err = c.BindJSON(requestBody)
		if err != nil {
			fmt.Println("Could not parse JSON request, ", err)
			return
		}

		obj, err := hnd.Repository.Update(id, *requestBody)
		fmt.Println(obj)
		if err != nil {
			c.JSON(http.StatusInternalServerError, obj)
		} else {
			c.JSON(http.StatusOK, obj)
		}
	}
}

func (hnd GinCRUDHandler[T]) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parseIdParameter(c)
		if err != nil {
			c.AbortWithError(http.StatusUnprocessableEntity, err)
			return
		}

		result := hnd.Repository.Delete(id)

		if result != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
		} else {
			c.JSON(http.StatusOK, gin.H{})
		}
	}
}
