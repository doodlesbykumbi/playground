package pkg

import (
	"github.com/golang/mock/gomock"
	"testing"
)

//go:generate mockgen -package pkg -source=pkg.go -destination=mock.go
func TestIsYourNameKumbi(t *testing.T) {
	tables := []struct {
		name string
		result bool
	}{
		{"Kumbi", true},
		{"NotKumbi", false},
	}

	for _, table := range tables {
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		m := NewMockNameable(ctrl)

		m.EXPECT().
			Name().
			Return(table.name)

		result := IsYourNameKumbi(m)
		if result != table.result {
			t.Errorf("IsYourNameKumbi() was incorrect, given: %v, got: %v, expect: %v", table.name, result, table.result)
		}
	}
}
