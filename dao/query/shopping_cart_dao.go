package query

import (
	"encoding/json"
	"fmt"
	dbPackage "tmf663/dao"
	model "tmf663/dao/model"
	models "tmf663/models"
)

var Q Query

func CreateShoppingCart(shoppingCart models.ShoppingCartCreate) error {

	db := dbPackage.GetDB()
	//var cart model.SHOPPINGCART
	var cart = convertShoppingCartToDbModel(shoppingCart)
	db.Create(&cart)

	/*
		u := Q.SHOPPINGCART

		fmt.Println("calling Create start")

		ctx := context.Background()
		err := u.WithContext(ctx).Create(&cart)

		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println("calling Create start")

		/*
			var cart_items model.CARTITEM

			ci := Q.CARTITEM

			err1 := ci.WithContext(ctx).Create(&cart_items)

			if err1 != nil {
				fmt.Println(err1)
				return err1
			}

			var product model.PRODUCT
			p := Q.PRODUCT

			err2 := p.WithContext(ctx).Create(&product)

			if err2 != nil {
				fmt.Println(err2)
				return err2
			}
	*/
	return nil

}

func convertShoppingCartToDbModel(shoppingCart models.ShoppingCartCreate) (result model.SHOPPINGCART) {
	var cart model.SHOPPINGCART

	cart.STARTDATETIME = shoppingCart.ValidFor.StartDateTime
	cart.ENDDATETIME = shoppingCart.ValidFor.EndDateTime

	fmt.Println(cart)
	return cart
}

func convertCartItemToDbModel(shoppingCart models.CartItem) (result model.CARTITEM) {
	var cartitem model.CARTITEM

	return cartitem
}

func convertProductToDbModel(shoppingCart models.ProductRefOrValue) (result model.PRODUCT) {
	var product model.PRODUCT

	return product
}

// Converting JSON data to a Go struct
func UnmarshalJSON(jsonStr string) (models.CartPrice, error) {
	var object models.CartPrice
	err := json.Unmarshal([]byte(jsonStr), &object)
	if err != nil {
		return models.CartPrice{}, err
	}
	return object, nil
}

// Converting a Go struct to JSON data
func MarshalJSON(object models.CartPrice) (string, error) {
	jsonStr, err := json.Marshal(object)
	if err != nil {
		return "", err
	}
	return string(jsonStr), nil
}
