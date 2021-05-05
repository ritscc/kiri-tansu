
package db

import (
	"reflect"
	model "ritscc/kiri-tansu/db/model"
	"testing"
	"gorm.io/gorm"
)

func test_db_user(db *gorm.DB, t *testing.T) {
	testCases := []struct {
		expected model.User
	}{
		{
			expected: model.User{
				ID:       1,
				Nickname: "Hoge",
				Role: 1,
			},
		},
	}

	for i, testCase := range testCases {
		var actual model.User
		db.First(&actual)
		expected := testCase.expected
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("errorNumber:%d actual: %v, expect: %v", i, actual, expected)
		}
	}
}
