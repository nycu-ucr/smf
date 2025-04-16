package util_oauth_test

import (
	"testing"

	"github.com/nycu-ucr/gonet/http/httptest"

	"github.com/nycu-ucr/gonet/http"

	"github.com/nycu-ucr/gin"
	"github.com/pkg/errors"

	"github.com/nycu-ucr/openapi/models"
	util_oauth "github.com/nycu-ucr/smf/internal/util/oauth"
)

const (
	Valid   = "valid"
	Invalid = "invalid"
)

type mockSMFContext struct{}

func newMockSMFContext() *mockSMFContext {
	return &mockSMFContext{}
}

func (m *mockSMFContext) AuthorizationCheck(token string, serviceName models.ServiceName) error {
	if token == Valid {
		return nil
	}

	return errors.New("invalid token")
}

func TestRouterAuthorizationCheck_Check(t *testing.T) {
	// Mock gin.Context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	var err error
	c.Request, err = http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("error on http request: %+v", err)
	}

	type Args struct {
		token string
	}
	type Want struct {
		statusCode int
	}

	tests := []struct {
		name string
		args Args
		want Want
	}{
		{
			name: "Valid Token",
			args: Args{
				token: Valid,
			},
			want: Want{
				statusCode: http.StatusOK,
			},
		},
		{
			name: "Invalid Token",
			args: Args{
				token: Invalid,
			},
			want: Want{
				statusCode: http.StatusUnauthorized,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w = httptest.NewRecorder()
			c, _ = gin.CreateTestContext(w)
			c.Request, err = http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Errorf("error on http request: %+v", err)
			}
			c.Request.Header.Set("Authorization", tt.args.token)

			rac := util_oauth.NewRouterAuthorizationCheck(models.ServiceName("testService"))
			rac.Check(c, newMockSMFContext())
			if w.Code != tt.want.statusCode {
				t.Errorf("StatusCode should be %d, but got %d", tt.want.statusCode, w.Code)
			}
		})
	}
}
