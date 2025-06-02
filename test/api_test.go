package test

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const address = "http://localhost:8080"

var client = http.Client{
	Timeout: 30 * time.Second,
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
	Age      int    `json:"age"`
}

type Post struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	ImageURL  string    `json:"image_url"`
	Caption   string    `json:"caption"`
	CreatedAt time.Time `json:"created_at"`
}

func TestCreateUser(t *testing.T) {
	user := User{Username: "testuser1", Email: "jaba@jaba.com", Bio: "Very jaba from jaba town"}
	payload, err := json.Marshal(user)
	require.NoError(t, err)

	resp, err := client.Post(address+"/api/action/create", "application/json", bytes.NewBuffer(payload))
	require.NoError(t, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	require.Equal(t, http.StatusCreated, resp.StatusCode, string(body))
}

func TestCreatePost(t *testing.T) {
	post := Post{
		ID:       "post_test_1",
		UserID:   "testuser1",
		ImageURL: "https://xkcd.com/info.0.json",
		Caption:  "CAPA",
	}
	payload, err := json.Marshal(post)
	require.NoError(t, err)

	resp, err := client.Post(address+"/api/action/create/post", "application/json", bytes.NewBuffer(payload))
	require.NoError(t, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, resp.StatusCode, string(body))
}

func createUser(t *testing.T, username string) {
	user := User{Username: username, Email: "jaba@jaba.com", Bio: "Very jaba from jaba town"}
	payload, err := json.Marshal(user)
	require.NoError(t, err)

	resp, err := client.Post(address+"/api/action/create", "application/json", bytes.NewBuffer(payload))
	require.NoError(t, err)
	defer resp.Body.Close()
}

func createPost(t *testing.T, postId string) {
	post := Post{
		ID:       postId,
		UserID:   "testuser1",
		ImageURL: "https://xkcd.com/info.0.json",
		Caption:  "CAPA",
	}
	payload, err := json.Marshal(post)
	require.NoError(t, err)

	resp, err := client.Post(address+"/api/action/create/post", "application/json", bytes.NewBuffer(payload))
	require.NoError(t, err)
	defer resp.Body.Close()
}

// TestFollowUser отправляет POST-запрос для подписки одного пользователя на другого.
func TestFollowUser(t *testing.T) {
	createUser(t, "testuser1")
	createUser(t, "testuser2")
	time.Sleep(1 * time.Second)
	params := url.Values{}
	params.Add("username", "testuser1")
	params.Add("usernameToFollow", "testuser2")
	u := address + "/api/action/follow?" + params.Encode()

	req, err := http.NewRequest(http.MethodPost, u, nil)
	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode, string(body))
}

// TestLikePost отправляет POST-запрос для регистрации лайка поста.
func TestLikePost(t *testing.T) {
	createUser(t, "testuser1")
	createPost(t, "post_test_1")
	time.Sleep(1 * time.Second)

	params := url.Values{}
	params.Add("user_id", "testuser1")
	params.Add("post_id", "post_test_1")
	u := address + "/api/action/like?" + params.Encode()

	req, err := http.NewRequest(http.MethodPost, u, nil)
	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode, string(body))
}

func TestGetFeed(t *testing.T) {
	params := url.Values{}
	params.Add("username", "testuser1")
	u := address + "/api/action/feed?" + params.Encode()

	resp, err := client.Get(u)
	require.NoError(t, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode, string(body))
}

type userData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
	Age      int    `json:"age"`
}

func TestManyUsers(t *testing.T) {
	sema := make(chan struct{}, 30)
	var wg sync.WaitGroup
	cnt := 1000
	wg.Add(cnt)
	for range cnt {
		go func() {
			sema <- struct{}{}
			defer func() {
				<-sema
				wg.Done()
			}()
			user := User{Username: uuid.NewString(), Email: uuid.NewString() + "@jaba.com", Bio: uuid.NewString(), Age: rand.Int()}
			payload, err := json.Marshal(user)
			require.NoError(t, err)

			resp, err := client.Post(address+"/api/action/create", "application/json", bytes.NewBuffer(payload))
			require.NoError(t, err)

			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)

			require.Equal(t, http.StatusCreated, resp.StatusCode, string(body))
		}()
	}
	wg.Wait()
}
