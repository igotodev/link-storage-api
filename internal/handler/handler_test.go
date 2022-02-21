package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"link-storage-api/internal/service"
	"link-storage-api/internal/storage"
	"link-storage-api/internal/storage/db"
	"link-storage-api/internal/storage/model"
	"link-storage-api/pkg/config"
	"link-storage-api/pkg/logger"
	"link-storage-api/pkg/psql"
	"link-storage-api/pkg/router"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

var (
	once sync.Once

	// WARNING! use ONLY test-db
	cfg = &config.Config{
		PostgresDB: config.PostgresDB{
			Host:     "127.0.0.1",
			Port:     "5432",
			User:     "postgres",
			Password: "postgres",
			DBName:   "my_db",
		},
		Listen: config.Listen{
			Port:   "8080",
			BindIP: "0.0.0.0",
		},
	}
)

func TestNewHandler(t *testing.T) {
	var sImpl storage.StorageImpl
	s := service.NewService(sImpl)
	l := logger.NewLogger()
	r := router.Router(l)

	h := Handler{
		router:    r,
		service:   s,
		appLogger: l,
	}

	testCases := []struct {
		name      string
		router    *chi.Mux
		service   service.ServiceImpl
		appLogger *logger.Logger
		want      *Handler
	}{
		{
			name:      "create",
			router:    r,
			service:   s,
			appLogger: l,
			want:      &h,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			assert.Equal(t, test.want, NewHandler(test.router, test.service, test.appLogger))
		})
	}
}

func TestLink(t *testing.T) {
	testCases := []struct {
		name    string
		address string
		item    int
		badItem string
		want    int
	}{
		{
			name:    "get",
			address: "/api/v1/link/",
			item:    1,
			want:    200,
		},
		{
			name:    "bad-item",
			address: "/api/v1/link/",
			badItem: "fsdf3sd",
			want:    400,
		},
	}

	var sImpl storage.StorageImpl = db.NewStorage(psql.NewPSLQ(cfg))
	s := service.NewService(sImpl)
	l := logger.NewLogger()
	r := router.Router(l)

	h := NewHandler(r, s, l)
	h.RegisterRouting()

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			addr := ""

			if strings.Contains(test.name, "bad") {
				addr = fmt.Sprintf(test.address + test.badItem)
			} else {
				addr = fmt.Sprintf(test.address + strconv.Itoa(test.item))
			}

			req := httptest.NewRequest(http.MethodGet, addr, nil)

			h.router.ServeHTTP(rec, req)

			assert.Equal(t, test.want, rec.Code)
		})
	}
}

func TestAllLink(t *testing.T) {
	testCases := []struct {
		name    string
		address string
		want    int
	}{
		{
			name:    "get",
			address: "/api/v1/link/",
			want:    200,
		},
		{
			name:    "bad-address",
			address: "/api/v000/link/",
			want:    404,
		},
	}

	var sImpl storage.StorageImpl = db.NewStorage(psql.NewPSLQ(cfg))
	s := service.NewService(sImpl)
	l := logger.NewLogger()
	r := router.Router(l)

	h := NewHandler(r, s, l)
	h.RegisterRouting()

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, test.address, nil)

			h.router.ServeHTTP(rec, req)

			assert.Equal(t, test.want, rec.Code)
		})
	}
}

func TestAddLink(t *testing.T) {
	m := model.Link{
		Name:     "Test",
		Category: "Test",
		URL:      "test@test.test",
		Date:     time.Now(),
	}

	marshal, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
	}

	testCases := []struct {
		name    string
		address string
		body    []byte
		want    int
	}{
		{
			name:    "post",
			address: "/api/v1/link/",
			body:    marshal,
			want:    200,
		},
		{
			name:    "bad-post",
			address: "/api/v1/link/",
			body:    []byte("fwsdfsd"),
			want:    400,
		},
	}

	var sImpl storage.StorageImpl = db.NewStorage(psql.NewPSLQ(cfg))
	s := service.NewService(sImpl)
	l := logger.NewLogger()
	r := router.Router(l)

	h := NewHandler(r, s, l)
	h.RegisterRouting()

	var answ answer

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, test.address, bytes.NewReader(test.body))

			h.router.ServeHTTP(rec, req)

			err := json.Unmarshal(rec.Body.Bytes(), &answ)
			if err != nil {
				log.Println(err)
			}

			once.Do(func() {
				writeIDToFile(strconv.Itoa(answ.ID))
			})

			assert.Equal(t, test.want, rec.Code)
		})
	}

}

func TestUpdateLink(t *testing.T) {
	idStr := readIDFromFile()

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
	}

	m := model.Link{
		ID:       id,
		Name:     "Test22",
		Category: "Test22",
		URL:      "test22@test22.test",
		Date:     time.Now(),
	}

	marshal, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
	}

	testCases := []struct {
		name    string
		address string
		item    int
		badItem string
		body    []byte
		want    int
	}{
		{
			name:    "put",
			address: "/api/v1/link/",
			item:    id,
			body:    marshal,
			want:    200,
		},
		{
			name:    "bad-item",
			address: "/api/v1/link/",
			item:    id,
			badItem: "sfdf43",
			body:    marshal,
			want:    400,
		},
	}

	var sImpl storage.StorageImpl = db.NewStorage(psql.NewPSLQ(cfg))
	s := service.NewService(sImpl)
	l := logger.NewLogger()
	r := router.Router(l)

	h := NewHandler(r, s, l)
	h.RegisterRouting()

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			addr := ""

			if strings.Contains(test.name, "bad") {
				addr = fmt.Sprintf(test.address + test.badItem)
			} else {
				addr = fmt.Sprintf(test.address + strconv.Itoa(test.item))
			}

			req := httptest.NewRequest(http.MethodPut, addr, bytes.NewReader(test.body))

			h.router.ServeHTTP(rec, req)

			assert.Equal(t, test.want, rec.Code)
		})
	}
}

func TestDeleteLink(t *testing.T) {
	idStr := readIDFromFile()

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
	}

	testCases := []struct {
		name    string
		address string
		item    int
		badItem string
		body    []byte
		want    int
	}{
		{
			name:    "delete",
			address: "/api/v1/link/",
			item:    id,
			want:    204,
		},
		{
			name:    "bad-delete",
			address: "/api/v1/link/",
			badItem: "dfsd3f",
			want:    400,
		},
	}

	var sImpl storage.StorageImpl = db.NewStorage(psql.NewPSLQ(cfg))
	s := service.NewService(sImpl)
	l := logger.NewLogger()
	r := router.Router(l)

	h := NewHandler(r, s, l)
	h.RegisterRouting()

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			addr := ""
			if strings.Contains(test.name, "bad") {
				addr = fmt.Sprintf(test.address + test.badItem)
			} else {
				addr = fmt.Sprintf(test.address + strconv.Itoa(test.item))
			}

			req := httptest.NewRequest(http.MethodDelete, addr, nil)

			h.router.ServeHTTP(rec, req)

			assert.Equal(t, test.want, rec.Code)
		})
	}
}

func readIDFromFile() string {
	all, err := ioutil.ReadFile("test")
	if err != nil {
		log.Fatal(err)
	}

	strID := strings.TrimSpace(string(all))

	return strID
}

func writeIDToFile(idStr string) {
	idBytes := []byte(idStr)
	err := ioutil.WriteFile("test", idBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
