package main


import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var products = map[int]*Product{}

var id = 0

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.Use(middleware.Recover())

	productGroup := e.Group("/products")

	productGroup.POST("", createProduct)

	productGroup.GET("", getProducts)

	productGroup.GET("/:id", getProduct)

	productGroup.PUT("/:id", updateProduct)

	productGroup.DELETE("/:id", deleteProduct)

	e.Logger.Fatal(e.Start(":8080"))
}

func createProduct(c echo.Context) error {
	product := &Product{}
	if err := c.Bind(product); err != nil {
		return err
	}
	id++
	product.ID = id
	products[id] = product
	return c.JSON(http.StatusCreated, product)
}

func getProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func getProduct(c echo.Context) error {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	product, ok := products[productID]
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, "Produkt nie istnieje")
	}
	return c.JSON(http.StatusOK, product)
}

func updateProduct(c echo.Context) error {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	product, ok := products[productID]
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, "Produkt nie istnieje")
	}
	newProduct := &Product{}
	if err := c.Bind(newProduct); err != nil {
		return err
	}
	product.Name = newProduct.Name
	product.Price = newProduct.Price
	return c.JSON(http.StatusOK, product)
}

func deleteProduct(c echo.Context) error {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	_, ok := products[productID]
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, "Produkt nie istnieje")
	}
	delete(products, productID)
	return c.NoContent(http.StatusNoContent)
}
