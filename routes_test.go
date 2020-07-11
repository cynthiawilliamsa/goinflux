package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_handleVersion(t *testing.T) {
	srv := server{router: http.NewServeMux()}
	srv.Routes()

	type test struct {
		name       string
		r          *http.Request
		wantStatus int
		wantError  string
	}

	tests := []test{
		{
			name: "successful request",
			r: func() *http.Request {
				req, _ := http.NewRequest("GET", "/version", nil)
				return req
			}(),
			wantStatus: http.StatusOK,
			wantError:  "{\"version\":\"v",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			srv.router.ServeHTTP(rr, tt.r)

			if rr.Code != tt.wantStatus {
				t.Errorf("handler return wrong status code: got %v, want %v", rr.Code, tt.wantStatus)
			}
			if !strings.Contains(rr.Body.String(), tt.wantError) {
				t.Errorf("handler returned unexpected body: got %v, should contain %v", rr.Body.String(), tt.wantError)
			}
		})
	}
}

// func Test_server_ListContainers(t *testing.T) {
// 	type fields struct {
// 		router *http.ServeMux
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   http.HandlerFunc
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			s := &server{
// 				router: tt.fields.router,
// 			}
// 			if got := s.ListContainers(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("server.ListContainers() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
