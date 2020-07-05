package main


func Test_handleVersion(t *testing.T) {
	src := server{router: http.NewServeMus()}
	srv.routes()

	type test struct {
		name string
		r *http.Request
		wantStatus int
		wantError string
	}

	tests := []test{
		{
			name: "successful request"
			r: func() *http.Request {
				req, _ := http.Request {
					("GET", "/version", nil)
					return req
				}(),
				wantStatus: http.StatusOK,
				wantError: "{\"version\":\"v",
			},
		}
		t.Run(tt.name, func(t*testing.T) {
			rr := httptest.NewRecorder()
			srv.router.ServeHTTP(rr, tt.r)

			if rr.Code != tt.wantStatus {
				t.Errorf("handler return wrong status cod: got %v, want %v", rr.Code, tt.wantStatus)
			}
			if !strings.Contains(rr.Body.String(), tt.wantError) {
				t.Errorf("handler returned unexpected body: got %v, should container %v", rr.Body.String(), tt.wantError)
				)
			}
		})
	}
}