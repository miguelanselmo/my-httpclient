package demo

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/miguelanselmo/my-httpclient/mime"
)

type AuthResponse struct {
	session_id               string "json:\"session_id\""
	access_token             string "json:\"access_token\""
	access_token_expires_at  string "json:\"access_token_expires_at\""
	refresh_token            string "json:\"refresh_token\""
	refresh_token_expires_at string "json:\"refresh_token_expires_at\""
	user                     User   "json:\"user\""
}

func Demo() {
	fmt.Println("About to start 'Demo'")
	url := "http://localhost:8082/"
	var authResp AuthResponse
	//Auth
	auth := Authentication{UserName: "TestUser3", Password: "123456"}
	response, err := httpClient.Post(url+"login", auth)
	if err != nil {
		fmt.Errorf("Error: %s\n", err)
		return
	} else {
		fmt.Printf("Status Code: %d\n", response.StatusCode)
		fmt.Printf("Status: %s\n", response.Status)
		//fmt.Printf("Body: %s\n", response.String())
		if err := response.UnmarshalJson(&authResp); err != nil {
			return
		}

	}
	//fmt.Printf("Acess token: %s\n", authResp)
	headers := make(http.Header)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjQ3ZGMzOTljLWMxNDEtNDVjNi1iNGJkLThjYWZmODIzMWVlYiIsInVzZXJuYW1lIjoiVGVzdFVzZXIzIiwiaXNzdWVkX2F0IjoiMjAyMi0wNy0yNVQwOToyMDowMy4xODczOTYzKzAxOjAwIiwiZXhwaXJlZF9hdCI6IjIwMjItMDctMjVUMDk6MzU6MDMuMTg3Mzk2MyswMTowMCJ9.pXzqPwYa0BVMPwUgLFMzeRVvwO1GvnYytVwsMemrLfo"
	headers.Set(mime.HeaderContentType, mime.ContentTypeJson)
	headers.Set(mime.HeaderAuthorization, "Bearer "+token)
	for i := 0; i < 10; i++ {
		go func() {
			//GetUserAll
			response, err = httpClient.Get(url+"users?page_id=1&page_size=5", headers)
			if err != nil {
				fmt.Errorf("Error: %s\n", err)
				return
			} else {
				fmt.Printf("Status Code: %d\n", response.StatusCode)
				fmt.Printf("Status: %s\n", response.Status)
				//fmt.Printf("Body: %s\n", response.String())
			}
			//GetUserById
			response, err = httpClient.Get(url+"users/3", headers)
			if err != nil {
				fmt.Errorf("Error: %s\n", err)
				return
			} else {
				fmt.Printf("Status Code: %d\n", response.StatusCode)
				fmt.Printf("Status: %s\n", response.Status)
				//fmt.Printf("Body: %s\n", response.String())
			}
			//UpdateUser
			user := User{Name: "testeuser" + strconv.Itoa(i), Email: "testeuser1@mail.com"}
			response, err = httpClient.Put(url+"users/3", user, headers)
			if err != nil {
				fmt.Errorf("Error: %s\n", err)
				return
			} else {
				fmt.Printf("Status Code: %d\n", response.StatusCode)
				fmt.Printf("Status: %s\n", response.Status)
				//fmt.Printf("Body: %s\n", response.String())
			}
		}()
	}
}

func GetDemoEndpoints() error {
	response, err := httpClient.Get("http://localhost:8082")
	if err != nil {
		return err
	}

	fmt.Printf(fmt.Sprintf("Status Code: %d\n", response.StatusCode))
	fmt.Printf(fmt.Sprintf("Status: %s\n", response.Status))
	fmt.Printf(fmt.Sprintf("Body: %s\n", response.String()))
	return nil
}

func Auth(auth Authentication) error {
	fmt.Println("About to start 'Auth'")
	response, err := httpClient.Post("http://localhost:8082/login", auth)
	if err != nil {
		fmt.Errorf("Error: %s\n", err)
		return err
	} else {
		fmt.Printf("Status Code: %d\n", response.StatusCode)
		fmt.Printf("Status: %s\n", response.Status)
		fmt.Printf("Body: %s\n", response.String())
		return nil
	}
}
