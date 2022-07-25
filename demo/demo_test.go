package demo

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("About to start test cases for package 'examples'")

	// Tell the HTTP library to mock any further requests from here.
	//mock.MockupServer.Stop()

	os.Exit(m.Run())
}

func TestGetDemoEndpoints(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "GetDemoEndpoints", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetDemoEndpoints(); (err != nil) != tt.wantErr {
				t.Errorf("GetDemoEndpoints() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuth(t *testing.T) {
	type args struct {
		auth Authentication
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Auth",
			args:    args{Authentication{UserName: "teste1", Password: "123456"}},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Auth(tt.args.auth); (err != nil) != tt.wantErr {
				t.Errorf("Auth() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDemo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "Demo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Demo()
		})
	}
}
