package db

import (
	"reflect"
	model "ritscc/kiri-tansu/db/model"
	"testing"
	"gorm.io/gorm"
)

func test_db_tag(db *gorm.DB, t *testing.T) {
	testCases := []struct {
		expected model.Tag
	}{
		{
			expected: model.Tag{
				ID:       1,
				Name: "book",
				CreatedAt: stringToTime("Jan 2, 2006 at 3:04pm (JST)"),
				CreatedBy: 1,
				UpdatedAt: stringToTime("Feb 3, 2013 at 7:54pm (JST)"),
				UpdatedBy: 1,
			},
		},
	}

	for i, testCase := range testCases {
		var actual model.Tag
		db.First(&actual)
		expected := testCase.expected
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("errorNumber:%d actual: %v, expect: %v", i, actual, expected)
		}
	}
}
