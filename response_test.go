package response

import (
	"encoding/json"
	"testing"
)

func TestSuccess(t *testing.T) {
	resp := Success(map[string]string{"name": "test"})
	if !resp.Success || !resp.Status {
		t.Error("Success and Status should be true")
	}
	if resp.Data == nil {
		t.Error("Data should not be nil")
	}
}

func TestSuccessWithMessage(t *testing.T) {
	resp := SuccessWithMessage("data", "get success")
	if resp.Message != "get success" {
		t.Errorf("expected 'get success', got '%s'", resp.Message)
	}
	if !resp.Success || !resp.Status {
		t.Error("Success and Status should be true")
	}
}

func TestList(t *testing.T) {
	items := []string{"a", "b"}
	pagination := NewPagination(100, 1, 10)
	resp := List(items, pagination, "get success")
	if !resp.Success || !resp.Status {
		t.Error("Success should be true")
	}
	if resp.Message != "get success" {
		t.Errorf("expected 'get success', got '%s'", resp.Message)
	}
}

func TestError(t *testing.T) {
	resp := Error(CLSTBadRequest, "invalid input", "trace-123")
	if resp.Success || resp.Status {
		t.Error("Success and Status should be false")
	}
	if resp.Error.Code != "C-LST-400" {
		t.Errorf("expected C-LST-400, got %s", resp.Error.Code)
	}
	if resp.Error.TraceID != "trace-123" {
		t.Errorf("expected trace-123, got %s", resp.Error.TraceID)
	}
	if resp.Error.Timestamp == "" {
		t.Error("Timestamp should be set")
	}
	// backward compat flat fields
	if resp.Message != "invalid input" {
		t.Errorf("flat message should be 'invalid input', got '%s'", resp.Message)
	}
	if resp.TraceID != "trace-123" {
		t.Errorf("flat traceId should be 'trace-123', got '%s'", resp.TraceID)
	}
}

func TestErrorWithDetails(t *testing.T) {
	resp := Error(CLSTInternal, "server error", "trace-456",
		ErrorIssue{Service: "Listing Service", Issue: "database timeout"},
	)
	if len(resp.Error.Details) != 1 {
		t.Errorf("expected 1 detail, got %d", len(resp.Error.Details))
	}
	if resp.Error.Details[0].Service != "Listing Service" {
		t.Errorf("expected 'Listing Service', got '%s'", resp.Error.Details[0].Service)
	}
}

func TestNewPagination(t *testing.T) {
	p := NewPagination(100, 1, 10)
	if p.TotalPages != 10 {
		t.Errorf("expected 10 pages, got %d", p.TotalPages)
	}

	p2 := NewPagination(101, 1, 10)
	if p2.TotalPages != 11 {
		t.Errorf("expected 11 pages, got %d", p2.TotalPages)
	}

	p3 := NewPagination(0, 1, 10)
	if p3.TotalPages != 0 {
		t.Errorf("expected 0 pages, got %d", p3.TotalPages)
	}
}

func TestErrorJSON_BackwardCompat(t *testing.T) {
	resp := Error(CLSTNotFound, "not found", "t-1")
	b, err := json.Marshal(resp)
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]interface{}
	json.Unmarshal(b, &m)

	// New format
	if m["success"] != false {
		t.Error("success should be false")
	}
	errObj := m["error"].(map[string]interface{})
	if errObj["code"] != "C-LST-404" {
		t.Errorf("error.code should be C-LST-404, got %v", errObj["code"])
	}

	// Backward compat flat fields
	if m["status"] != false {
		t.Error("status should be false (backward compat)")
	}
	if m["message"] != "not found" {
		t.Errorf("flat message should be 'not found', got '%v'", m["message"])
	}
}

func TestSuccessJSON_BackwardCompat(t *testing.T) {
	resp := SuccessWithMessage(map[string]int{"count": 5}, "ok")
	b, err := json.Marshal(resp)
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]interface{}
	json.Unmarshal(b, &m)

	if m["success"] != true {
		t.Error("success should be true")
	}
	if m["status"] != true {
		t.Error("status should be true (backward compat)")
	}
	if m["message"] != "ok" {
		t.Errorf("message should be 'ok', got '%v'", m["message"])
	}
}

func TestAllErrorCodes(t *testing.T) {
	codes := []string{
		BNYBBadRequest, BNYBUnauthorized, BNYBForbidden, BNYBNotFound, BNYBInternal, BNYBGatewayTimeout,
		CAUTBadRequest, CAUTUnauthorized, CAUTForbidden, CAUTNotFound, CAUTInternal,
		CLSTBadRequest, CLSTUnauthorized, CLSTNotFound, CLSTConflict, CLSTValidation, CLSTInternal,
		CMEDBadRequest, CMEDNotFound, CMEDPayloadTooLarge, CMEDInternal,
		CSRCBadRequest, CSRCInternal,
		CNOTInternal,
		CRPTBadRequest, CRPTNotFound, CRPTInternal,
		XINTBadGateway, XINTUnavailable, XINTGatewayTimeout,
		DWRKBadRequest, DWRKInternal,
	}

	for _, code := range codes {
		if code == "" {
			t.Error("Empty error code found")
		}
	}
	t.Logf("Verified %d error codes", len(codes))
}
