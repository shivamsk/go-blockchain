package controllers

import (
	"go-blockchain/internal/constants"
	"go-blockchain/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HttpResponseInjector() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Next()

		statusCode, statusCodeExists := c.Get(constants.StatusCode)

		if errorMsg, errMsgExists := c.Get(constants.ErrorMsg); errMsgExists {
			httpStatusCode := http.StatusInternalServerError

			if statusCodeExists {
				httpStatusCode = statusCode.(int)
			}

			errorResponse := response.ErrorResponse{
				Data:       errorMsg,
				StatusCode: httpStatusCode,
			}
			c.JSON(httpStatusCode, &errorResponse)
		} else {

			// 410(gone) if router doesn't have a handler

			httpStatusCode := http.StatusGone

			if statusCodeExists {
				httpStatusCode = statusCode.(int)
			}

			obj, responseBodyExists := c.Get(constants.ResponseValue)

			res := response.APIResponse{
				Data:       obj,
				StatusCode: httpStatusCode,
			}

			if responseBodyExists {
				c.JSON(httpStatusCode, res)
			} else {
				c.Status(httpStatusCode)
			}
		}
	}
}
