package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
	
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

const(
	notValidEmail = "notvalid@email@gmail.com"
)

func TestRegister_HappyPath(t *testing.T) {
	t.Parallel()
	email, password := randomEmailandPassword()
	obj := fmt.Sprintf(`{"email": "%s"", "password": "%s"}`, email, password)

	httpURL := fmt.Sprintf("http://localhost:1234/auth/reg/?email=%s&password=%s", email, password)
	jsonBody := []byte(obj)
	bodyReader := bytes.NewReader(jsonBody)

	resp, err := http.Post(httpURL, "", bodyReader)
	resp.Body.Close()
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200)
}

func TestLogin_HappyPath(t *testing.T) {
	t.Parallel()
	email, password := randomEmailandPassword()
	obj := fmt.Sprintf(`{"email": "%s"", "password": "%s"}`, email, password)

	httpURL := fmt.Sprintf("http://localhost:1234/auth/reg/?email=%s&password=%s", email, password)
	jsonBody := []byte(obj)
	bodyReader := bytes.NewReader(jsonBody)

	resp, err := http.Post(httpURL, "", bodyReader)
	resp.Body.Close()
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200)

	resp, err = http.Get("http://localhost:1234/auth/get/?email=serpan2002@mail.ru&password=54321")
	resp.Body.Close()
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200)
}

func TestUpdatePassword_HappyPath(t *testing.T) {
	t.Parallel()
	
	email, password := randomEmailandPassword()
	obj := fmt.Sprintf(`{"email": "%s"", "password": "%s"}`, email, password)

	httpURL := fmt.Sprintf("http://localhost:1234/auth/reg/?email=%s&password=%s", email, password)
	jsonBody := []byte(obj)
	bodyReader := bytes.NewReader(jsonBody)

	resp, err := http.Post(httpURL, "", bodyReader)
	resp.Body.Close()
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200)
	
	passwordnew := randomPassword()
	obj = fmt.Sprintf(`{"email": "%s"", "password": "%s", "passwordnew": "%s"}`, email, password, passwordnew)

	httpURL = fmt.Sprintf("http://localhost:1234/auth/up/?email=%s&password=%s&passwordnew=%s", email, password, passwordnew)
	jsonBody = []byte(obj)
	bodyReader = bytes.NewReader(jsonBody)

	resp, err = http.Post(httpURL, "", bodyReader)
	resp.Body.Close()
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200)
}

func TestDelete_HappyPath(t *testing.T) {
	t.Parallel()
	
	email, password := randomEmailandPassword()
	obj := fmt.Sprintf(`{"email": "%s"", "password": "%s"}`, email, password)

	httpURL := fmt.Sprintf("http://localhost:1234/auth/reg/?email=%s&password=%s", email, password)
	jsonBody := []byte(obj)
	bodyReader := bytes.NewReader(jsonBody)

	resp, err := http.Post(httpURL, "", bodyReader)
	resp.Body.Close()
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200)
	
	obj = fmt.Sprintf(`{"email": "%s"", "password": "%s"}`, email, password)

	httpURL = fmt.Sprintf("http://localhost:1234/auth/del/?email=%s&password=%s", email, password)
	jsonBody = []byte(obj)
	bodyReader = bytes.NewReader(jsonBody)

	resp, err = http.Post(httpURL, "", bodyReader)
	resp.Body.Close()
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200)
}

func TestRecover_HappyPath(t *testing.T) {
	t.Parallel()
	
	email, password := randomEmailandPassword()
	obj := fmt.Sprintf(`{"email": "%s"", "password": "%s"}`, email, password)

	httpURL := fmt.Sprintf("http://localhost:1234/auth/reg/?email=%s&password=%s", email, password)
	jsonBody := []byte(obj)
	bodyReader := bytes.NewReader(jsonBody)

	resp, err := http.Post(httpURL, "", bodyReader)
	resp.Body.Close()
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200)
	
	obj = fmt.Sprintf(`{"email": "%s"", "password": "%s"}`, email, password)

	httpURL = fmt.Sprintf("http://localhost:1234/auth/rec/?email=%s&password=%s", email, password)
	jsonBody = []byte(obj)
	bodyReader = bytes.NewReader(jsonBody)

	resp, err = http.Post(httpURL, "", bodyReader)
	resp.Body.Close()
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200)
}

func TestRegister_FailCases(t *testing.T){
	t.Parallel()
	
	email, password := randomEmailandPassword()
	obj := fmt.Sprintf(`{"email": "%s"", "password": "%s"}`, email, password)

	httpURL := fmt.Sprintf("http://localhost:1234/auth/reg/?email=%s&password=%s", email, password)
	jsonBody := []byte(obj)
	bodyReader := bytes.NewReader(jsonBody)

	resp, err := http.Post(httpURL, "", bodyReader)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200)

	cases := []struct {
		name string
        email string
        password string
		expectedStatusCode int
	}{
		{
			name: "Register with empty email",
            email: "",
            password: randomPassword(),
            expectedStatusCode: 400,
		},
		{
			name: "Register with empty password",
            email: gofakeit.Email(),
            password: "",
            expectedStatusCode: 400,
		},
		{
			name: "Register with empty both",
            email: "",
            password: "",
            expectedStatusCode: 400,
		},
		{
			name: "Register with existing email",
            email: email,
            password: password,
            expectedStatusCode: 400,
		},
		{
			name: "Register with not valid email",
            email: notValidEmail,
            password: randomPassword(),
            expectedStatusCode: 400,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
            obj := fmt.Sprintf(`{"email": "%s"", "password": "%s"}`, tt.email, tt.password)

            httpURL := fmt.Sprintf("http://localhost:1234/auth/reg/?email=%s&password=%s", tt.email, tt.password)
            jsonBody := []byte(obj)
            bodyReader := bytes.NewReader(jsonBody)

            resp, err := http.Post(httpURL, "", bodyReader)
			resp.Body.Close()
			if err != nil {
				t.Error(err)
			}
            assert.Equal(t, resp.StatusCode, tt.expectedStatusCode)
        })
	}
}

func TestLogin_FailCases(t *testing.T){
	t.Parallel()
	
	email, password := randomEmailandPassword()
	obj := fmt.Sprintf(`{"email": "%s"", "password": "%s"}`, email, password)

	httpURL := fmt.Sprintf("http://localhost:1234/auth/reg/?email=%s&password=%s", email, password)
	jsonBody := []byte(obj)
	bodyReader := bytes.NewReader(jsonBody)

	resp, err := http.Post(httpURL, "", bodyReader)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200)

	cases := []struct {
		name string
        email string
        password string
        expectedStatusCode int
    }{
		{
			name: "Login with empty email",
            email: "",
            password: randomPassword(),
            expectedStatusCode: 400,
		},
		{
			name: "Login with empty password",
            email: email,
            password: "",
            expectedStatusCode: 400,
		},
		{
			name: "Login with empty both",
            email: "",
            password: "",
            expectedStatusCode: 400,
		},
		{
			name: "Login with not valid email",
            email: "notvalid@email@gmail.com",
            password: randomPassword(),
            expectedStatusCode: 400,
		},
		{
			name: "Login with not valid password",
            email: email,
            password: randomPassword(),
            expectedStatusCode: 400,
		},
		{
			name: "Login with not exist",
            email: gofakeit.Email(),
            password: randomPassword(),
            expectedStatusCode: 400,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
            httpURL := fmt.Sprintf("http://localhost:1234/auth/get/?email=%s&password=%s", tt.email, tt.password)

            resp, err = http.Get(httpURL)
			resp.Body.Close()
			if err != nil {
				t.Error(err)
			}
            assert.Equal(t, resp.StatusCode, tt.expectedStatusCode)
        })
	}
}

func TestUpdatePassword_FailCases(t *testing.T) {
	t.Parallel()
	email, password := randomEmailandPassword()
	obj := fmt.Sprintf(`{"email": "%s"", "password": "%s"}`, email, password)

	httpURL := fmt.Sprintf("http://localhost:1234/auth/reg/?email=%s&password=%s", email, password)
	jsonBody := []byte(obj)
	bodyReader := bytes.NewReader(jsonBody)

	resp, err := http.Post(httpURL, "", bodyReader)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200)

	cases := []struct {
		name string
        email string
        password string
		passwordnew string
        expectedStatusCode int
    }{
		{
			name: "Update with empty email",
            email: "",
            password: randomPassword(),
			passwordnew: randomPassword(),
            expectedStatusCode: 400,
		},
		{
			name: "Update with empty password",
            email: email,
            password: "",
			passwordnew: randomPassword(),
            expectedStatusCode: 400,
		},
		{
			name: "Update with empty passwordnew",
            email: email,
            password: randomPassword(),
			passwordnew: "",
            expectedStatusCode: 400,
		},
		{
			name: "Update with empty all",
            email: "",
            password: "",
			passwordnew: "",
            expectedStatusCode: 400,
		},
		{
			name: "Update with not valid email",
            email: "notvalid@email@gmail.com",
            password: randomPassword(),
			passwordnew: randomPassword(),
            expectedStatusCode: 400,
		},
		{
			name: "Update with not valid password",
            email: email,
            password: randomPassword(),
			passwordnew: randomPassword(),
            expectedStatusCode: 400,
		},
		{
			name: "Update with not exist",
            email: gofakeit.Email(),
            password: randomPassword(),
			passwordnew: randomPassword(),
            expectedStatusCode: 400,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
            httpURL := fmt.Sprintf(
				"http://localhost:1234/auth/up/?email=%s&password=%s&passwordnew=%s", 
				tt.email, tt.password, tt.passwordnew)
			obj = fmt.Sprintf(`{"email": "%s"", "password": "%s", "passwordnew": "%s"}`, email, password, tt.passwordnew)
			jsonBody := []byte(obj)
			bodyReader := bytes.NewReader(jsonBody)
			
            resp, err = http.Post(httpURL, "", bodyReader)
			resp.Body.Close()
			if err != nil {
				t.Error(err)
			}
            assert.Equal(t, resp.StatusCode, tt.expectedStatusCode)
        })
	}
}	

func TestDelete_FailCases(t *testing.T) {
	t.Parallel()
	email, password := randomEmailandPassword()
	obj := fmt.Sprintf(`{"email": "%s"", "password": "%s"}`, email, password)

	httpURL := fmt.Sprintf("http://localhost:1234/auth/reg/?email=%s&password=%s", email, password)
	jsonBody := []byte(obj)
	bodyReader := bytes.NewReader(jsonBody)

	resp, err := http.Post(httpURL, "", bodyReader)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200)

	cases := []struct {
		name string
        email string
        password string
        expectedStatusCode int
    }{
		{
			name: "Delete with empty email",
            email: "",
            password: randomPassword(),
            expectedStatusCode: 400,
		},
		{
			name: "Delete with empty password",
            email: email,
            password: "",
            expectedStatusCode: 400,
		},
		{
			name: "Delete with empty all",
            email: "",
            password: "",
            expectedStatusCode: 400,
		},
		{
			name: "Delete with not valid email",
            email: "notvalid@email@gmail.com",
            password: randomPassword(),
            expectedStatusCode: 400,
		},
		{
			name: "Delete with not valid password",
            email: email,
            password: randomPassword(),
            expectedStatusCode: 400,
		},
		{
			name: "Delete with not exist",
            email: gofakeit.Email(),
            password: randomPassword(),
            expectedStatusCode: 400,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
            httpURL := fmt.Sprintf(
				"http://localhost:1234/auth/del/?email=%s&password=%s", 
				tt.email, tt.password)
			obj = fmt.Sprintf(`{"email": "%s"", "password": "%s"}`, email, password)
			jsonBody := []byte(obj)
			bodyReader := bytes.NewReader(jsonBody)
			
            resp, err = http.Post(httpURL, "", bodyReader)
			resp.Body.Close()
			if err != nil {
				t.Error(err)
			}
            assert.Equal(t, resp.StatusCode, tt.expectedStatusCode)
        })
	}
}

func TestRecover_FailCases(t *testing.T) {
	t.Parallel()

	email, password := randomEmailandPassword()
	obj := fmt.Sprintf(`{"email": "%s"", "password": "%s"}`, email, password)

	httpURL := fmt.Sprintf("http://localhost:1234/auth/reg/?email=%s&password=%s", email, password)
	jsonBody := []byte(obj)
	bodyReader := bytes.NewReader(jsonBody)

	resp, err := http.Post(httpURL, "", bodyReader)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200)

	cases := []struct {
		name string
        email string
        password string
        expectedStatusCode int
    }{
		{
			name: "Recover with empty email",
            email: "",
            password: randomPassword(),
            expectedStatusCode: 400,
		},
		{
			name: "Recover with empty password",
            email: email,
            password: "",
            expectedStatusCode: 400,
		},
		{
			name: "Recover with empty all",
            email: "",
            password: "",
            expectedStatusCode: 400,
		},
		{
			name: "Recover with not valid email",
            email: "notvalid@email@gmail.com",
            password: randomPassword(),
            expectedStatusCode: 400,
		},
		{
			name: "Recover with not valid password",
            email: email,
            password: randomPassword(),
            expectedStatusCode: 400,
		},
		{
			name: "Recover with not exist",
            email: gofakeit.Email(),
            password: randomPassword(),
            expectedStatusCode: 400,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
            httpURL := fmt.Sprintf(
				"http://localhost:1234/auth/rec/?email=%s&password=%s", 
				tt.email, tt.password)
			obj = fmt.Sprintf(`{"email": "%s"", "password": "%s"}`, email, password)
			jsonBody := []byte(obj)
			bodyReader := bytes.NewReader(jsonBody)
			
            resp, err = http.Post(httpURL, "", bodyReader)
			resp.Body.Close()
			if err != nil {
				t.Error(err)
			}
            assert.Equal(t, resp.StatusCode, tt.expectedStatusCode)
        })
	}
}
func randomPassword() string {
	return gofakeit.Password(true, true, true, true, false, 10)
}
func randomEmailandPassword() (string, string) {
	return gofakeit.Email(), randomPassword()
}