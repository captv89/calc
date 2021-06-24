package actions

import (
	"net/http"
	"strconv"

	"github.com/gobuffalo/buffalo"
	"github.com/gorilla/mux"
)

// CalcAdd default implementation.
func CalcAdd(c buffalo.Context) error {
	params := mux.Vars(c.Request())
	num1, _ := strconv.Atoi(params["num1"])
	num2, _ := strconv.Atoi(params["num2"])
	println(num1, num2)
	result := int(num1) + int(num2)
	return c.Render(http.StatusOK, r.JSON(result))
}

// CalcSub default implementation.
func CalcSub(c buffalo.Context) error {
	params := mux.Vars(c.Request())
	num1, _ := strconv.Atoi(params["num1"])
	num2, _ := strconv.Atoi(params["num2"])
	println(num1, num2)
	result := int(num1) - int(num2)
	return c.Render(http.StatusOK, r.JSON(result))
}

// CalcMul default implementation.
func CalcMul(c buffalo.Context) error {
	params := mux.Vars(c.Request())
	num1, _ := strconv.Atoi(params["num1"])
	num2, _ := strconv.Atoi(params["num2"])
	println(num1, num2)
	result := int(num1) * int(num2)
	return c.Render(http.StatusOK, r.JSON(result))
}

// CalcDiv default implementation.
func CalcDiv(c buffalo.Context) error {
	params := mux.Vars(c.Request())
	num1, _ := strconv.Atoi(params["num1"])
	num2, _ := strconv.Atoi(params["num2"])
	println(num1, num2)
	result := int(num1) / int(num2)
	return c.Render(http.StatusOK, r.JSON(result))
}
