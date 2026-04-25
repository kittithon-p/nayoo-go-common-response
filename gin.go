package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ============================================================
// Gin Helper Functions
// ============================================================

// ResponseSuccess sends a success response via Gin context
func ResponseSuccess(c *gin.Context, status int, data interface{}) {
	c.JSON(status, Success(data))
}

// ResponseSuccessWithMessage sends a success response with message via Gin context
func ResponseSuccessWithMessage(c *gin.Context, status int, data interface{}, message string) {
	c.JSON(status, SuccessWithMessage(data, message))
}

// ResponseList sends a paginated list response via Gin context
func ResponseList(c *gin.Context, data interface{}, pagination interface{}, message string) {
	c.JSON(http.StatusOK, List(data, pagination, message))
}

// ResponseError sends an error response via Gin context
func ResponseError(c *gin.Context, httpStatus int, code, message, traceID string, details ...ErrorIssue) {
	c.JSON(httpStatus, Error(code, message, traceID, details...))
}

// ============================================================
// Convenience Error Helpers (with Gin)
// ============================================================

// ResponseBadRequest sends a 400 error
func ResponseBadRequest(c *gin.Context, code, message, traceID string, details ...ErrorIssue) {
	ResponseError(c, http.StatusBadRequest, code, message, traceID, details...)
}

// ResponseUnauthorized sends a 401 error
func ResponseUnauthorized(c *gin.Context, code, message, traceID string, details ...ErrorIssue) {
	ResponseError(c, http.StatusUnauthorized, code, message, traceID, details...)
}

// ResponseForbidden sends a 403 error
func ResponseForbidden(c *gin.Context, code, message, traceID string, details ...ErrorIssue) {
	ResponseError(c, http.StatusForbidden, code, message, traceID, details...)
}

// ResponseNotFound sends a 404 error
func ResponseNotFound(c *gin.Context, code, message, traceID string, details ...ErrorIssue) {
	ResponseError(c, http.StatusNotFound, code, message, traceID, details...)
}

// ResponseConflict sends a 409 error
func ResponseConflict(c *gin.Context, code, message, traceID string, details ...ErrorIssue) {
	ResponseError(c, http.StatusConflict, code, message, traceID, details...)
}

// ResponseValidationFailed sends a 422 error
func ResponseValidationFailed(c *gin.Context, code, message, traceID string, details ...ErrorIssue) {
	ResponseError(c, http.StatusUnprocessableEntity, code, message, traceID, details...)
}

// ResponseInternalError sends a 500 error
func ResponseInternalError(c *gin.Context, code, message, traceID string, details ...ErrorIssue) {
	ResponseError(c, http.StatusInternalServerError, code, message, traceID, details...)
}

// ResponseServiceUnavailable sends a 503 error
func ResponseServiceUnavailable(c *gin.Context, code, message, traceID string, details ...ErrorIssue) {
	ResponseError(c, http.StatusServiceUnavailable, code, message, traceID, details...)
}

// ResponseGatewayTimeout sends a 504 error
func ResponseGatewayTimeout(c *gin.Context, code, message, traceID string, details ...ErrorIssue) {
	ResponseError(c, http.StatusGatewayTimeout, code, message, traceID, details...)
}
