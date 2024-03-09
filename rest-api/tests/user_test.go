package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

const (
	nameUser = "John"
	sexUser = "male"
	birthdate = "2002-02-02"
	age = 34
	weight = 88
)

func TestGetUser_HappyPath(t *testing.T) {
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

	httpURL = fmt.Sprintf("http://localhost:1234/user/get/?email=%s&password=%s", email, password)

	resp, err = http.Get(httpURL)
	resp.Body.Close()
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200)
}

func TestUpdateUser_HappyPath(t *testing.T) {
	t.Parallel()

	email, password := randomEmailandPassword()
	obj := fmt.Sprintf(`{"email": "%s"", "password": "%s"}`, email, password)

	httpURL := fmt.Sprintf("http://localhost:1234/auth/reg/?email=%s&password=%s", email, password)
	jsonBody := []byte(obj)
	bodyReader := bytes.NewReader(jsonBody)

	resp, err := http.Post(httpURL, "", bodyReader)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200)
	resp.Body.Close()
	
	httpURL = fmt.Sprintf("http://localhost:1234/user/up/?email=%s&password=%s&name=%s&sex=%s&birthdate=%s&age=%d&weight=%d",  email, password, nameUser, sexUser, birthdate, age, weight)

	obj = fmt.Sprintf(`{"email": "%s"", "password": "%s", "name": "%s"", "sex": "%s", "birthdate": "%s"", "age": "%d", "weight"="%d"}`, email, password, nameUser, sexUser, birthdate, age, weight)
	jsonBody = []byte(obj)
	bodyReader = bytes.NewReader(jsonBody)

	res, err := http.Post(httpURL, "", bodyReader)
	res.Body.Close()
	assert.NoError(t, err)
	assert.Equal(t, res.StatusCode, 200)
}

func TestGetUser_FailCases(t *testing.T) {
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
			name: "Get User with empty email",
            email: "",
            password: randomPassword(),
            expectedStatusCode: 400,
		},
		{
			name: "Get User with empty password",
            email: gofakeit.Email(),
            password: "",
            expectedStatusCode: 400,
		},
		{
			name: "Get User with empty both",
            email: "",
            password: "",
            expectedStatusCode: 400,
		},
		{
			name: "Get User with not valid email",
            email: gofakeit.Email(),
            password: "",
            expectedStatusCode: 400,
		},
		{
			name: "Get User with not valid password",
            email: email,
            password: randomPassword(),
            expectedStatusCode: 400,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			httpURL := fmt.Sprintf(
                "http://localhost:1234/user/get/?email=%s&password=%s", 
                tt.email, tt.password)
            obj = fmt.Sprintf(`{"email": "%s"", "password": "%s"}`, email, password)

            resp, err := http.Get(httpURL)
			resp.Body.Close()
            assert.NoError(t, err)
            assert.Equal(t, resp.StatusCode, tt.expectedStatusCode)
		})
	}
}

func TestUpdateUser_FailCases(t *testing.T) {
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
		userName string
		sex string
		birthdate string
		age int
		weight int
		expectedStatusCode int
	}{
		{
			name: "Update User with empty email",
            email: "",
            password: randomPassword(),
			userName: nameUser,
			sex: sexUser,
            birthdate: birthdate,
            age: age,
            weight: weight,				
            expectedStatusCode: 400,
		},
		{
			name: "Update User with empty password",
            email: gofakeit.Email(),
            password: "",
			userName: nameUser,
			sex: sexUser,
            birthdate: birthdate,
            age: age,
            weight: weight,				
            expectedStatusCode: 400,
		},
		{
			name: "Update User with empty name",
            email: gofakeit.Email(),
            password: randomPassword(),
			userName: "",
			sex: sexUser,
            birthdate: birthdate,
            age: age,
            weight: weight,				
            expectedStatusCode: 400,
		},
		{
			name: "Update User with empty sex",
            email: gofakeit.Email(),
            password: randomPassword(),
			userName: nameUser,
			sex: "",
            birthdate: birthdate,
            age: age,
            weight: weight,				
            expectedStatusCode: 400,
		},
		{
			name: "Update User with empty birthdate",
            email: gofakeit.Email(),
            password: randomPassword(),
			userName: nameUser,
			sex: sexUser,
            birthdate: "",
            age: age,
            weight: weight,				
            expectedStatusCode: 400,
		},
		{
			name: "Update User with empty age",
            email: gofakeit.Email(),
            password: randomPassword(),
			userName: nameUser,
			sex: sexUser,
            birthdate: birthdate,
            age: 0,
            weight: weight,				
            expectedStatusCode: 400,
		},
		{
			name: "Update User with empty weight",
            email: gofakeit.Email(),
            password: randomPassword(),
			userName: nameUser,
			sex: sexUser,
            birthdate: birthdate,
            age: age,
            weight: 0,				
            expectedStatusCode: 400,
		},
		{
			name: "Update User with empty both",
            email: "",
            password: "",
			userName: "",
			sex: "",
            birthdate: "",
            age: 0,
            weight: 0,				
            expectedStatusCode: 400,
		},
		{
			name: "Update User with not valid email",
            email: gofakeit.Email(),
            password: "",
			userName: nameUser,
			sex: sexUser,
            birthdate: birthdate,
            age: age,
            weight: weight,
            expectedStatusCode: 400,
		},
		{
			name: "Update User with not valid password",
            email: email,
            password: randomPassword(),
			userName: nameUser,
			sex: sexUser,
            birthdate: birthdate,
            age: age,
            weight: weight,
            expectedStatusCode: 400,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			httpURL := fmt.Sprintf("http://localhost:1234/user/up/?email=%s&password=%s&name=%s&sex=%s&birthdate=%s&age=%d&weight=%d",  tt.email, tt.password, tt.userName, tt.sex, tt.birthdate, tt.age, tt.weight)

			obj := fmt.Sprintf(`{"email": "%s"", "password": "%s", "name": "%s"", "sex": "%s", "birthdate": "%s"", "age": "%d", "weight"="%d"}`, tt.email, tt.password, tt.userName, tt.sex, tt.birthdate, tt.age, tt.weight)
			jsonBody := []byte(obj)
			bodyReader := bytes.NewReader(jsonBody)
			resp, _ := http.Post(httpURL, "", bodyReader)
			resp.Body.Close()
            // assert.NoError(t, err)
            assert.Equal(t, resp.StatusCode, tt.expectedStatusCode)

		})
	}
}