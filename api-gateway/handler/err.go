package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

func HandleError(c *gin.Context, err error) {
	if grpcErr, ok := status.FromError(err); ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": grpcErr.Message()})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unknown error"})
	}
}
