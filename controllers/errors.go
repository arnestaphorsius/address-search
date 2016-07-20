package controllers

import (
    "github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {
    c.IndentedJSON(404, gin.H{"status": "404", "message": "not found"})
}

func BadRequest(c *gin.Context) {
    c.IndentedJSON(400, gin.H{"status": "400", "message": "bad request"})
}
