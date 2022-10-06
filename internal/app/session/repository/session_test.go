package sessionRepository

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepo(t *testing.T) {
	repo := New()

	test := []struct {
		id  uint
		cid string
	}{
		{
			id: 1,
		},
		{
			id: 2,
		},
		{
			id: 3,
		},
		{
			id: 3,
		},
	}

	for idx := range test {
		test[idx].cid = repo.Create(test[idx].id).Value
		require.Equal(t, repo.Storage[test[idx].cid], test[idx].id, "Create session")
	}

	for _, val := range test {
		repo.Remove(val.cid)
		if _, exists := repo.Storage[val.cid]; exists {
			t.Error("Session doesn't remove")
		}
	}
}
