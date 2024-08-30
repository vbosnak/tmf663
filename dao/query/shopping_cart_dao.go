package query

import (
	"encoding/json"
	"fmt"
	"strconv"
	dbPackage "tmf663/dao"
	model "tmf663/dao/model"
	models "tmf663/models"
)

var Q Query

func CreateShoppingCart(shoppingCart models.ShoppingCartCreate) error {

	db := dbPackage.GetDB()
	var cart = convertShoppingCartToDbModel(shoppingCart)
	Q.SHOPPINGCART.sHOPPINGCARTDo.UseDB(db, nil)
	Q.SHOPPINGCART.sHOPPINGCARTDo.Create(&cart)

	fmt.Println("Added Cart ID is : ", cart.ID)
	cartitems := shoppingCart.CartItem
	var product model.PRODUCT
	for _, opt := range cartitems {
		product = convertProductToDbModel(opt.Product)
	}

	Q.PRODUCT.pRODUCTDo.UseDB(db, nil)
	Q.PRODUCT.pRODUCTDo.Create(&product)
	fmt.Println("Added Product ID is : ", product.ID)

	var cart_items model.CARTITEM

	for _, opt := range cartitems {
		cart_items = convertCartItemToDbModel(opt)
	}
	cart_items.SHOPPINGCARTID = cart.ID
	cart_items.PRODUCTID = product.ID

	Q.CARTITEM.cARTITEMDo.UseDB(db, nil)
	Q.CARTITEM.cARTITEMDo.Create(&cart_items)

	fmt.Println("Added Cart Item ID is : ", cart_items.ID)

	return nil

}

func GetShoppingCartById(id string) models.ShoppingCart {
	//var cartAPI models.ShoppingCart
	var cart model.SHOPPINGCART
	db := dbPackage.GetDB()

	n, err := strconv.Atoi(id)
	if err != nil {
		return models.ShoppingCart{}
	}
	cart.ID = int32(n)

	fmt.Println("Get Cart ID is : ", cart.ID)

	result := db.First(&cart, cart.ID)
	if result.Error != nil {
		return models.ShoppingCart{}
	}

	var cartItem model.CARTITEM
	cartItem.SHOPPINGCARTID = cart.ID
	fmt.Println("Get Cart Item ID is : ", cartItem.SHOPPINGCARTID)

	result1 := db.Where(&model.CARTITEM{SHOPPINGCARTID: cart.ID}).First(&cartItem)

	if result1.Error != nil {
		return models.ShoppingCart{}
	}

	//fmt.Println("Cart Item ID Before is : ", cartItem)

	cartItemAPI := mapDbModelToCartItem(cartItem)
	var product model.PRODUCT
	result2 := db.Where(&model.PRODUCT{ID: cartItem.PRODUCTID}).First(&product)

	if result2.Error != nil {
		return models.ShoppingCart{}
	}

	productAPI := mapDbModelToProduct(product)
	cartItemAPI.Product = productAPI
	cartAPI := mapDbModelToShoppingCart(cart)
	cartAPI.CartItem = []models.CartItem{cartItemAPI}
	// Convert to models.ShoppingCart and return
	//cartAPI.CartItem = append(cartAPI.CartItem, cartItemAPI)

	return cartAPI
}

func convertShoppingCartToDbModel(shoppingCart models.ShoppingCartCreate) (result model.SHOPPINGCART) {
	var cart model.SHOPPINGCART

	cart.STARTDATETIME = shoppingCart.ValidFor.StartDateTime
	cart.ENDDATETIME = shoppingCart.ValidFor.EndDateTime

	return cart
}

func mapDbModelToShoppingCart(shoppingCart model.SHOPPINGCART) (result models.ShoppingCart) {
	cart := models.ShoppingCart{}

	cart.ValidFor.StartDateTime = shoppingCart.STARTDATETIME
	cart.ValidFor.EndDateTime = shoppingCart.ENDDATETIME

	return cart
}

func mapDbModelToCartItem(shoppingCart model.CARTITEM) (result models.CartItem) {
	cart := models.CartItem{}

	cart.Quantity = shoppingCart.QUANTITY
	cart.Action = models.CartItemActionType(shoppingCart.ACTION)
	cart.Status = models.CartItemStatusType(shoppingCart.STATUS)
	cartTerm := models.CartTerm{}
	cartTerm.Duration.Amount = float32(shoppingCart.ITEMTERM)
	cart.ItemTerm = []models.CartTerm{cartTerm}

	itemPrice := models.Price{TaxIncludedAmount: models.Money{Value: shoppingCart.ITEMPRICE}}
	cartItemPrice := models.CartPrice{Price: itemPrice}
	cart.ItemPrice = []models.CartPrice{cartItemPrice}

	itemTotalPrice := models.Price{TaxIncludedAmount: models.Money{Value: shoppingCart.ITEMTTOTALPRICE}}
	cartTotalItemPrice := models.CartPrice{Price: itemTotalPrice}
	cart.ItemTotalPrice = []models.CartPrice{cartTotalItemPrice}

	return cart
}

func mapDbModelToProduct(shoppingCart model.PRODUCT) (result models.ProductRefOrValue) {
	cart := models.ProductRefOrValue{}

	cart.Description = shoppingCart.DESCRIPTION
	cart.Name = shoppingCart.NAME
	cart.IsBundle = shoppingCart.ISBUNDLE
	cart.IsCustomerVisible = shoppingCart.ISCUSTOMERVISIBLE
	cart.OrderDate = shoppingCart.ORDERDATE
	cart.StartDate = shoppingCart.STARTDATE
	cart.TerminationDate = shoppingCart.TERMINATIONDATE
	cart.ProductSerialNumber = shoppingCart.PRODUCTSERIALNUMBER
	productSpecification := models.ProductSpecificationRef{Name: shoppingCart.PRODUCTSPECIFICATION}
	cart.ProductSpecification = productSpecification

	productTerm := models.ProductTerm{}
	productTerm.Duration.Amount = float32(shoppingCart.PRODUCTTERM)
	cart.ProductTerm = []models.ProductTerm{productTerm}

	itemPrice := models.Price{TaxIncludedAmount: models.Money{Value: shoppingCart.PRODUCTPRICE}}
	productPrice := models.ProductPrice{Price: itemPrice}
	cart.ProductPrice = []models.ProductPrice{productPrice}

	//fmt.Println("Cart Product is : ", cart)

	return cart
}

func convertCartItemToDbModel(shoppingCart models.CartItem) (result model.CARTITEM) {
	var cartitem model.CARTITEM

	cartitem.QUANTITY = shoppingCart.Quantity
	cartitem.ACTION = string(shoppingCart.Action)
	cartitem.STATUS = string(shoppingCart.Status)
	term := shoppingCart.ItemTerm
	for _, opt := range term {
		cartitem.ITEMTERM = int32(opt.Duration.Amount)
	}

	itemPrice := shoppingCart.ItemPrice
	for _, opt := range itemPrice {
		cartitem.ITEMPRICE = opt.Price.TaxIncludedAmount.Value
	}

	itemTotalPrice := shoppingCart.ItemTotalPrice
	for _, opt := range itemTotalPrice {
		cartitem.ITEMTTOTALPRICE = opt.Price.TaxIncludedAmount.Value
	}

	return cartitem
}

func convertProductToDbModel(shoppingCart models.ProductRefOrValue) (result model.PRODUCT) {
	var product model.PRODUCT

	product.DESCRIPTION = shoppingCart.Description
	product.NAME = shoppingCart.Name
	product.ISBUNDLE = shoppingCart.IsBundle
	product.ISCUSTOMERVISIBLE = shoppingCart.IsCustomerVisible
	product.ORDERDATE = shoppingCart.OrderDate
	product.STARTDATE = shoppingCart.StartDate
	product.TERMINATIONDATE = shoppingCart.TerminationDate
	product.PRODUCTSERIALNUMBER = shoppingCart.ProductSerialNumber
	product.PRODUCTSPECIFICATION = shoppingCart.ProductSpecification.Name
	term := shoppingCart.ProductTerm
	for _, opt := range term {
		product.PRODUCTTERM = int32(opt.Duration.Amount)
	}

	price := shoppingCart.ProductPrice
	for _, opt := range price {
		product.PRODUCTPRICE = opt.Price.TaxIncludedAmount.Value
	}

	relatedParty := shoppingCart.RelatedParty
	for _, opt := range relatedParty {
		product.RELATEDPARTY = opt.Name
	}

	return product
}

func DeleteShoppingCartById(id string) (bool, error) {

	var cart model.SHOPPINGCART
	db := dbPackage.GetDB()

	n, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	cart.ID = int32(n)

	fmt.Println("Get Cart ID is : ", cart.ID)

	result := db.First(&cart, cart.ID)
	if result.Error != nil {
		fmt.Println(err)
		return false, err
	}

	var cartItem model.CARTITEM
	cartItem.SHOPPINGCARTID = cart.ID

	fmt.Println("Get Cart Item ID is : ", cartItem.SHOPPINGCARTID)

	result3 := db.Where(&model.CARTITEM{SHOPPINGCARTID: cart.ID}).Find(&cartItem)

	if result3.Error != nil {
		return false, err
	}

	producId := cartItem.PRODUCTID

	fmt.Println("Deleting Cart Item ID is : ", cartItem.SHOPPINGCARTID)

	if cartItem.ID != 0 {
		result1 := db.Where(&model.CARTITEM{SHOPPINGCARTID: cart.ID}).Delete(&cartItem)

		if result1.Error != nil {
			return false, err
		}
	}
	var product model.PRODUCT
	fmt.Println("Deleting Product ID is : ", producId)

	if producId != 0 {
		result2 := db.Where(&model.PRODUCT{ID: producId}).Delete(&product)

		if result2.Error != nil {
			return false, err
		}
	}
	result = db.Delete(&cart)

	if result.Error != nil {
		panic("failed to delete Cart: " + result.Error.Error())
	} else if result.RowsAffected == 0 {
		panic("no Cart record was deleted")
	} else {
		fmt.Println("Cart record deleted successfully")
	}
	return true, err
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

func RetrieveShoppingCart(cartID int32) (models.ShoppingCart, error) {
	db := dbPackage.GetDB()

	Q.SHOPPINGCART.sHOPPINGCARTDo.UseDB(db, nil)

	cart, err := Q.SHOPPINGCART.sHOPPINGCARTDo.Where(Q.SHOPPINGCART.ID.Eq(cartID)).First()
	if err != nil {
		fmt.Println("Error retrieving shopping cart:", err)
		return models.ShoppingCart{}, err
	}
	fmt.Println("Retrieved Cart ID is:", cart.ID)

	Q.CARTITEM.cARTITEMDo.UseDB(db, nil)

	cartItems, err := Q.CARTITEM.cARTITEMDo.Where(Q.CARTITEM.SHOPPINGCARTID.Eq(cart.ID)).Find()
	if err != nil {
		fmt.Println("Error retrieving cart items:", err)
		return models.ShoppingCart{}, err
	}

	var cartt model.SHOPPINGCART = *cart
	cartAPI := mapDbModelToShoppingCart(cartt)

	var products []model.PRODUCT
	productMap := make(map[int32]model.PRODUCT)
	for _, cartItem := range cartItems {
		var item model.CARTITEM = *cartItem
		cartItemAPI := mapDbModelToCartItem(item)
		var product model.PRODUCT
		if prod, exists := productMap[cartItem.PRODUCTID]; exists {
			product = prod
		} else {
			productPtr, err := Q.PRODUCT.pRODUCTDo.Where(Q.PRODUCT.ID.Eq(cartItem.PRODUCTID)).First()
			if err != nil {
				fmt.Println("Error retrieving product:", err)
				return models.ShoppingCart{}, err
			}
			product = *productPtr
		}
		productAPI := mapDbModelToProduct(product)

		products = append(products, product)
		cartItemAPI.Product = productAPI

		cartAPI.CartItem = []models.CartItem{cartItemAPI}
	}

	return cartAPI, nil
}

func GetShoppingCartList() ([]models.ShoppingCart, error) {

	db := dbPackage.GetDB()

	Q.SHOPPINGCART.sHOPPINGCARTDo.UseDB(db, nil)

	var carts []model.SHOPPINGCART

	var returnCart []models.ShoppingCart

	cart := db.Find(&carts)

	if cart.Error != nil {
		fmt.Println("Error retrieving shopping cart: ", cart.Error)
		return []models.ShoppingCart{}, cart.Error
	}

	for _, cart := range carts {
		var cartItems []model.CARTITEM
		fmt.Println("Retrieved Cart ID is:", cart.ID)
		//Q.CARTITEM.cARTITEMDo.UseDB(db, nil)

		//cartItems, err := Q.CARTITEM.cARTITEMDo.Where(Q.CARTITEM.SHOPPINGCARTID.Eq(cart.ID)).Find()
		result := db.Where(&model.CARTITEM{SHOPPINGCARTID: cart.ID}).Find(&cartItems)

		if result.Error != nil {
			fmt.Println("Error retrieving cart items:", result.Error)
			return []models.ShoppingCart{}, result.Error
		}

		var items []models.CartItem
		for _, cartItem := range cartItems {
			cartItemAPI := mapDbModelToCartItem(cartItem)
			/*
				var item model.CARTITEM = *cartItem
				cartItemAPI := mapDbModelToCartItem(item)
				var product model.PRODUCT
				productPtr, err := Q.PRODUCT.pRODUCTDo.Where(Q.PRODUCT.ID.Eq(cartItem.PRODUCTID)).First()
				if err != nil {
					fmt.Println("Error retrieving product:", err)
					return nil, err
				}
			*/
			//var productPtr *model.PRODUCT
			var product model.PRODUCT
			result1 := db.Where(&model.PRODUCT{ID: cartItem.PRODUCTID}).First(&product)
			if result1.Error != nil {
				fmt.Println("Error retrieving Product:", result1.Error)
				return []models.ShoppingCart{}, result1.Error
			}
			//product = *productPtr
			productAPI := mapDbModelToProduct(product)
			cartItemAPI.Product = productAPI
			items = append(items, cartItemAPI)
		}
		cartAPI := mapDbModelToShoppingCart(cart)
		cartAPI.CartItem = items
		returnCart = append(returnCart, cartAPI)
	}

	return returnCart, nil
}
