package user
import "github.com/gin-gonic/gin"

type Handler struct {
	Service
}
func NewHandler(s Service) *Handle{
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreatedUser(c *gin.Context) {
	
}