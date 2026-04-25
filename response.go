package response

import (
	"time"
)

// ============================================================
// Response Types
// ============================================================

// SuccessResponse represents a successful API response
type SuccessResponse struct {
	Success bool        `json:"success"`
	Status  bool        `json:"status"` // backward compat
	Data    interface{} `json:"data"`
	Message string      `json:"message,omitempty"` // backward compat
}

// ErrorResponse represents an error API response
type ErrorResponse struct {
	Success bool         `json:"success"`
	Status  bool         `json:"status"`  // backward compat
	Error   *ErrorDetail `json:"error"`
	Message string       `json:"message"` // backward compat flat
	TraceID string       `json:"traceId,omitempty"` // backward compat flat
}

// ErrorDetail contains detailed error information
type ErrorDetail struct {
	Code      string       `json:"code"`
	Message   string       `json:"message"`
	Details   []ErrorIssue `json:"details,omitempty"`
	TraceID   string       `json:"traceId"`
	Timestamp string       `json:"timestamp"`
}

// ErrorIssue represents a specific error issue
type ErrorIssue struct {
	Service string `json:"service"`
	Issue   string `json:"issue"`
}

// Pagination represents pagination information
type Pagination struct {
	Total      int `json:"total"`
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalPages int `json:"total_pages"`
}

// ListResponse represents a paginated list response (backward compat)
type ListResponse struct {
	Success    bool        `json:"success"`
	Status     bool        `json:"status"`  // backward compat
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination,omitempty"` // backward compat flat
}

// ============================================================
// Constructor Functions
// ============================================================

// Success creates a successful response
func Success(data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Success: true,
		Status:  true,
		Data:    data,
	}
}

// SuccessWithMessage creates a successful response with message
func SuccessWithMessage(data interface{}, message string) *SuccessResponse {
	return &SuccessResponse{
		Success: true,
		Status:  true,
		Data:    data,
		Message: message,
	}
}

// List creates a paginated list response (backward compat — pagination at top level)
func List(data interface{}, pagination interface{}, message string) *ListResponse {
	return &ListResponse{
		Success:    true,
		Status:     true,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	}
}

// Error creates an error response
func Error(code, message, traceID string, details ...ErrorIssue) *ErrorResponse {
	return &ErrorResponse{
		Success: false,
		Status:  false,
		Error: &ErrorDetail{
			Code:      code,
			Message:   message,
			Details:   details,
			TraceID:   traceID,
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		},
		Message: message,
		TraceID: traceID,
	}
}

// NewPagination creates a pagination object with calculated total pages
func NewPagination(total, page, pageSize int) Pagination {
	totalPages := total / pageSize
	if total%pageSize != 0 {
		totalPages++
	}
	if totalPages == 0 && total > 0 {
		totalPages = 1
	}

	return Pagination{
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}
}

// ============================================================
// Error Code Constants — Per Service
// ============================================================

// Zone B — Web Application (BFF)
const (
	BNYBBadRequest     = "B-NYB-400"
	BNYBUnauthorized   = "B-NYB-401"
	BNYBForbidden      = "B-NYB-403"
	BNYBNotFound       = "B-NYB-404"
	BNYBConflict       = "B-NYB-409"
	BNYBValidation     = "B-NYB-422"
	BNYBTooMany        = "B-NYB-429"
	BNYBInternal       = "B-NYB-500"
	BNYBUnavailable    = "B-NYB-503"
	BNYBGatewayTimeout = "B-NYB-504"

	BCOBBadRequest     = "B-COB-400"
	BCOBUnauthorized   = "B-COB-401"
	BCOBForbidden      = "B-COB-403"
	BCOBNotFound       = "B-COB-404"
	BCOBInternal       = "B-COB-500"
	BCOBUnavailable    = "B-COB-503"
	BCOBGatewayTimeout = "B-COB-504"
)

// Zone C — Core Services
const (
	// C-AUT — Auth Service
	CAUTBadRequest   = "C-AUT-400"
	CAUTUnauthorized = "C-AUT-401"
	CAUTForbidden    = "C-AUT-403"
	CAUTNotFound     = "C-AUT-404"
	CAUTConflict     = "C-AUT-409"
	CAUTValidation   = "C-AUT-422"
	CAUTInternal     = "C-AUT-500"
	CAUTUnavailable  = "C-AUT-503"

	// C-LST — Listing Service
	CLSTBadRequest   = "C-LST-400"
	CLSTUnauthorized = "C-LST-401"
	CLSTForbidden    = "C-LST-403"
	CLSTNotFound     = "C-LST-404"
	CLSTConflict     = "C-LST-409"
	CLSTValidation   = "C-LST-422"
	CLSTInternal     = "C-LST-500"
	CLSTUnavailable  = "C-LST-503"

	// C-MED — Media Service
	CMEDBadRequest     = "C-MED-400"
	CMEDNotFound       = "C-MED-404"
	CMEDPayloadTooLarge = "C-MED-413"
	CMEDUnsupportedMedia = "C-MED-415"
	CMEDInternal       = "C-MED-500"
	CMEDBadGateway     = "C-MED-502"

	// C-SRC — Search Service
	CSRCBadRequest  = "C-SRC-400"
	CSRCInternal    = "C-SRC-500"
	CSRCUnavailable = "C-SRC-503"

	// C-NOT — Notification Service
	CNOTInternal = "C-NOT-500"

	// C-RPT — Report Service
	CRPTBadRequest = "C-RPT-400"
	CRPTNotFound   = "C-RPT-404"
	CRPTValidation = "C-RPT-422"
	CRPTInternal   = "C-RPT-500"

	// C-INT — Integration Service
	CINTBadRequest = "C-INT-400"
	CINTInternal   = "C-INT-500"
)

// Zone X — External Integration
const (
	XINTBadGateway     = "X-INT-502"
	XINTUnavailable    = "X-INT-503"
	XINTGatewayTimeout = "X-INT-504"
)

// Zone D — Worker / Job
const (
	DWRKBadRequest = "D-WRK-400"
	DWRKInternal   = "D-WRK-500"
)
