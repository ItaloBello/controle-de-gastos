package category_handler

import (
	"controle-de-gastos/src/model"
	"controle-de-gastos/src/service/category_service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	service category_service.CategoryService
}

type CategoryHandler interface {
	GetAllCategories(c *gin.Context)
	GetCategoryById(c *gin.Context)

	CreateCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
}

func NewCategoryHandler(serv category_service.CategoryService) CategoryHandler {
	return &categoryHandler{service: serv}
}

func (h *categoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "error during get all categories"})
		fmt.Println("error during get all categories: " + err.Error())
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (h *categoryHandler) GetCategoryById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "id must be an integer"})
		fmt.Println("id must be an integer: " + err.Error())
		return
	}
	category, err := h.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "error during get category by id"})
		fmt.Println("error during get category by id: " + err.Error())
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *categoryHandler) CreateCategory(c *gin.Context){
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "invalid request body"})
		fmt.Println("invalid request body: " + err.Error())
		return
	}
	id, err := h.service.Create(category)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "error during create category"})
		fmt.Println("error during create category: " + err.Error())
		return
	}
	c.JSON(http.StatusOK, id)
}

func (h *categoryHandler) UpdateCategory(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "id must be an integer"})
		fmt.Println("id must be an integer: " + err.Error())
		return
	}
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "invalid request body"})
		fmt.Println("invalid request body: " + err.Error())
		return
	}
	category.Id = id
	err = h.service.Update(category)
		if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "error during update category"})
		fmt.Println("error during update category: " + err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category updated successfully"})
}

func (h *categoryHandler) DeleteCategory(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "id must be an integer"})
		fmt.Println("id must be an integer: " + err.Error())
		return
	}
	err = h.service.Delete(id)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "error during delete category"})
		fmt.Println("error during delete category: " + err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category deleted successfully"})
}