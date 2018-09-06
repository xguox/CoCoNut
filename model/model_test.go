package model

import (
	"coconut/db"
	"fmt"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {
	testDB = db.TestDBInit()
	exitVal := m.Run()
	db.ResetTestDB(testDB)
	os.Exit(exitVal)
}

func resetDBWithMock() {
	db.ResetTestDB(testDB)
	testDB = db.TestDBInit()
	categoriesMocker(5)
}

func newTestCategory(name, slug string) Category {
	return Category{
		Name: name,
		Slug: slug,
	}
}

func categoriesMocker(n int) []Category {
	var offset int
	testDB.Model(&Category{}).Count(&offset)
	var cArr []Category
	for i := offset + 1; i <= offset+n; i++ {
		category := newTestCategory(fmt.Sprintf("name-%v", i), fmt.Sprintf("slug-%v", i))
		testDB.Create(&category)
		cArr = append(cArr, category)
	}
	return cArr
}

func TestGetCategoryByID(t *testing.T) {
	categoriesMocker(2)
	asserts := assert.New(t)
	category, err := GetCategoryByID("1")
	asserts.NoError(err, "category should exist")
	asserts.Equal("name-1", category.Name, "GetCategoryByID() should return category with the right name")

}