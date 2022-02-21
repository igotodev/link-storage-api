package service

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"link-storage-api/internal/storage"
	"link-storage-api/internal/storage/db"
	"testing"
)

func TestNewService(t *testing.T) {
	var tdb *sql.DB
	st := db.NewStorage(tdb)
	ns := NewService(st)

	testCases := []struct {
		name    string
		storage storage.StorageImpl
		want    *Service
	}{
		{
			name:    "create",
			storage: st,
			want:    ns,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			assert.Equal(t, test.want, NewService(test.storage))
		})
	}
}
