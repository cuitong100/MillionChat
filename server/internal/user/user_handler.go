package user
import "github.com/gin-gonic/gin"

type Handler struct {
	Service
}

// Constructor of handler
func NewHandler(s Service) *Handle{
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreatedUser(c *gin.Context) {
	var u CreatedUserReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error"; err.Error})
		return
	}

	res, err := h.Service.CreatedUser(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)


}