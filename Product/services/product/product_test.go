package product_test

import (
	m "Golang-Rest-API/product/models"
	productService "Golang-Rest-API/product/services/product"
	"Golang-Rest-API/product/stores"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

var products = []m.Product{
	{
		Id:    1,
		Name:  "Milk",
		Price: 26.5,
	},
	{
		Id:    1,
		Name:  "Grapes",
		Price: 70,
	},
}

func TestGetProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := stores.NewMockProduct(ctrl)
	mockStore.EXPECT().GetProduct().Return(products, nil)

	productServiceClient := productService.NewProduct(mockStore)

	expectedProducts, err := productServiceClient.GetProduct()
	fmt.Println(expectedProducts)
	if !reflect.DeepEqual(err, nil) {
		t.Errorf("Expected: %v, Got: %v", nil, err)
	}
	if !reflect.DeepEqual(products, expectedProducts) {
		t.Errorf("Expected: %v, Got: %v", products, expectedProducts)
	}

}
