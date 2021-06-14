package db

import (
	"gorm.io/gorm"
	"reflect"
	model "ritscc/kiri-tansu/db/model"
	"testing"
)

func test_db_tag(db *gorm.DB, t *testing.T) {
	result := db.Create(&model.Tag{
		ID:   1,
		Name: "book",
		CreatedBy: 1,
		UpdatedBy: 1,
	}).Error

	if result != nil {
		t.Error(result)
		return
	}

	testCases := []struct {
		expected model.Tag
	}{
		{
			expected: model.Tag{
				ID:        1,
				Name:      "book",
				CreatedBy: 1,
				UpdatedBy: 1,
			},
		},
	}

	for i, testCase := range testCases {
		var actual model.Tag
		db.First(&actual)
		expected := testCase.expected
		if !reflect.DeepEqual(actual.ID, expected.ID) {
			t.Errorf("errorNumber:%d actual: %v, expect: %v", i, actual, expected)
		}

		if !reflect.DeepEqual(actual.Name, expected.Name) {
			t.Errorf("errorNumber:%d actual: %v, expect: %v", i, actual, expected)
		}

		if !reflect.DeepEqual(actual.CreatedBy, expected.CreatedBy) {
			t.Errorf("errorNumber:%d actual: %v, expect: %v", i, actual, expected)
		}

		if !reflect.DeepEqual(actual.UpdatedBy, expected.UpdatedBy) {
			t.Errorf("errorNumber:%d actual: %v, expect: %v", i, actual, expected)
		}
	}
}
