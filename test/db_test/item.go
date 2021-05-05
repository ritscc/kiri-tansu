package db

import (
	"gorm.io/gorm"
	"reflect"
	model "ritscc/kiri-tansu/db/model"
	"testing"
)

func test_db_item(db *gorm.DB, t *testing.T) {
	result := db.Create(&model.Item{
		ID:        1,
		Name:      "Vim",
		Overview:  "Vimの本です",
		Type:      1,
		CreatedBy: 1,
		UpdatedBy: 1,
	}).Error

	if result != nil {
		t.Error(result)
		return
	}

	testCases := []struct {
		expected model.Item
	}{
		{
			expected: model.Item{
				ID:        1,
				Name:      "Vim",
				Overview:  "Vimの本です",
				Type:      1,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
		},
	}

	for i, testCase := range testCases {
		var actual model.Item
		db.First(&actual)
		expected := testCase.expected
		if !reflect.DeepEqual(actual.ID, expected.ID) {
			t.Errorf("errorNumber:%d actual: %v, expect: %v", i, actual, expected)
		}

		if !reflect.DeepEqual(actual.Name, expected.Name) {
			t.Errorf("errorNumber:%d actual: %v, expect: %v", i, actual, expected)
		}

		if !reflect.DeepEqual(actual.Overview, expected.Overview) {
			t.Errorf("errorNumber:%d actual: %v, expect: %v", i, actual, expected)
		}

		if !reflect.DeepEqual(actual.Type, expected.Type) {
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
