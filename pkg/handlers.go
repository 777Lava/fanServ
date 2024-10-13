package pkg

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}
type RequestData struct {
	Param1 string `json:"param1" binding:"required"`
	Param2 string `json:"param2" binding:"required"`
}

func NewHandler() *Handler {
	return &Handler{}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Static("/static", "./static")
	router.LoadHTMLGlob("./templates/*")

	router.GET("/", h.middleware)
	router.POST("/", h.resp)
	return router
}

func (h *Handler) resp(c *gin.Context) {
	var req RequestData
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	resp := fmt.Sprintf("Пусть тебе сегодня приснятся %s и %s 💗💗💗💗", req.Param1, req.Param2)
	c.JSON(http.StatusOK, gin.H{"result": resp})
}

func (h *Handler) middleware(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
