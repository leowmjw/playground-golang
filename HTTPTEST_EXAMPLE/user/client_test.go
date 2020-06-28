package user

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewUserServiceClient(t *testing.T) {
	type args struct {
		baseURL string
	}
	tests := []struct {
		name string
		args args
		want UserServiceClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserServiceClient(tt.args.baseURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserServiceClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartUserService(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestUserServiceClient_HealthUserService(t *testing.T) {
	type fields struct {
		baseURL    string
		httpClient *http.Client
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usc := UserServiceClient{
				baseURL:    tt.fields.baseURL,
				httpClient: tt.fields.httpClient,
			}
			if err := usc.HealthUserService(); (err != nil) != tt.wantErr {
				t.Errorf("HealthUserService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserServiceClient_QueryExternal(t *testing.T) {
	type fields struct {
		baseURL    string
		httpClient *http.Client
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usc := UserServiceClient{
				baseURL:    tt.fields.baseURL,
				httpClient: tt.fields.httpClient,
			}
			if err := usc.QueryExternal(); (err != nil) != tt.wantErr {
				t.Errorf("QueryExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserServiceClient_QueryUserService(t *testing.T) {
	type fields struct {
		baseURL    string
		httpClient *http.Client
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usc := UserServiceClient{
				baseURL:    tt.fields.baseURL,
				httpClient: tt.fields.httpClient,
			}
			if err := usc.QueryUserService(); (err != nil) != tt.wantErr {
				t.Errorf("QueryUserService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserServiceClient_SlowUserService(t *testing.T) {
	type fields struct {
		baseURL    string
		httpClient *http.Client
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usc := UserServiceClient{
				baseURL:    tt.fields.baseURL,
				httpClient: tt.fields.httpClient,
			}
			if err := usc.SlowUserService(); (err != nil) != tt.wantErr {
				t.Errorf("SlowUserService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserServiceClient_callUserService(t *testing.T) {
	type fields struct {
		baseURL    string
		httpClient *http.Client
	}
	type args struct {
		urlpath        string
		failureHandler func(error) error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"happy #1", fields{
			baseURL:    "",
			httpClient: nil,
		}, args{
			urlpath:        "/query",
			failureHandler: nil,
		}, false},
		{"sad #1", fields{
			baseURL:    "",
			httpClient: nil,
		}, args{
			urlpath:        "/slow",
			failureHandler: nil,
		}, false},
	}
	// Setup httptest ..
	// Use niave way first ..
	usc := NewUserServiceClient("http://localhost:8080")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//usc := UserServiceClient{
			//	baseURL:    tt.fields.baseURL,
			//	httpClient: tt.fields.httpClient,
			//}
			if err := usc.callUserService(tt.args.urlpath, tt.args.failureHandler); (err != nil) != tt.wantErr {
				t.Errorf("callUserService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_isTimeoutError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTimeoutError(tt.args.err); got != tt.want {
				t.Errorf("isTimeoutError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_doGetREST(t *testing.T) {
	type args struct {
		fullURL        string
		client         *http.Client
		failureHandler func(error) error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := doGetREST(tt.args.fullURL, tt.args.client, tt.args.failureHandler); (err != nil) != tt.wantErr {
				t.Errorf("doGetREST() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
