package controller

import (
	"cocodingding/gptbots/pkg/client/chatgpt"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetGptResponse(c *gin.Context) {
	var resp string
	var err error

	query := c.Query("script")
	if query == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Empty query"})
		return
	}

	resp, err = chatgpt.Ask(query)
	if err != nil {
		// go-openai sample error
		// 'error, status code: 404, message: bad request'

		// 문자열을 콤마로 분할하여 배열로 변환
		parts := strings.Split(err.Error(), ", ")

		// 에러 메시지 추출
		message := strings.Split(parts[2], ": ")[1]

		// 에러 코드 추출
		code := http.StatusBadRequest
		codeStr := parts[1]
		_, _ = fmt.Sscanf(codeStr, "%d", &code)

		// 에러 생성
		c.AbortWithStatusJSON(code, gin.H{"error": message})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": resp})
}
