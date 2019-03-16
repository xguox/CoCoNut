package model

import (
	"github.com/xguox/coconut/db"
	"fmt"
	"os"
	"testing"
	"time"

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

func TestCategories(t *testing.T) {
	asserts := assert.New(t)
	var category Category
	var categories []Category
	var oldCategoriesCount, currentCategoriesCount int
	categoriesMocker(6)

	category, err := GetCategoryByID("1")
	asserts.NoError(err, "category should exist")
	asserts.Equal("name-1", category.Name, "GetCategoryByID() should return category with the right name")

	_, err = GetCategoryByID("10000")
	asserts.Error(err, "category not found should return err")

	testDB.Model(&Category{}).Count(&oldCategoriesCount)
	categoriesMocker(3)
	categories = GetCategories()
	currentCategoriesCount = oldCategoriesCount + 3
	asserts.Equal(currentCategoriesCount, len(categories), "GetCategories() should return all categories where deleted_at is null without pagination")

	category.SetDeletedAt(time.Now())
	categories = GetCategories()
	asserts.Equal(currentCategoriesCount-1, len(categories), "category SetDeletedAt() should remove from GetCategories()")
}

func newTestProduct(title, slug string) Product {
	return Product{
		Title: title,
		Slug:  slug,
	}
}

func productsMocker(n int) []Product {
	var offset int
	testDB.Model(&Product{}).Count(&offset)
	var pArr []Product
	for i := offset + 1; i <= offset+n; i++ {
		product := newTestProduct(fmt.Sprintf("title-%v", i), fmt.Sprintf("slug-%v", i))
		testDB.Create(&product)
		pArr = append(pArr, product)
	}
	return pArr
}

func TestProducts(t *testing.T) {
	asserts := assert.New(t)
	var product Product
	var products []Product
	var oldProductsCount, currentProductsCount int
	productsMocker(1)
	productsMocker(6)

	product, err := GetProductByID("1")
	asserts.NoError(err, "product should exist")
	asserts.Equal("title-1", product.Title, "GetProductByID() should return product with the right title")

	_, err = GetProductByID("5000")
	asserts.Error(err, "product not found should return err")

	testDB.Model(&Product{}).Count(&oldProductsCount)
	productsMocker(3)
	products = GetProducts()
	currentProductsCount = oldProductsCount + 3
	asserts.Equal(currentProductsCount, len(products), "GetProducts() should return all Products where deleted_at is null without pagination")

	// product.SetDeletedAt(time.Now())
	// products = GetProducts()
	// asserts.Equal(currentProductsCount-1, len(products), "product SetDeletedAt() should remove from GetProducts()")
}
