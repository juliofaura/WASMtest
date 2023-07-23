package main

import (
	"fmt"
	"strconv"
	"syscall/js"
	"time"

	"github.com/juliofaura/WASMtest/src/dom"
)

func performCalculation() {
	firstNumber, firstNumberErr := strconv.Atoi(dom.GetString("first-number", "value"))
	secondNumber, secondNumberErr := strconv.Atoi(dom.GetString("second-number", "value"))

	if firstNumberErr == nil && secondNumberErr == nil {
		dom.SetValue("result", "value", strconv.Itoa(firstNumber+secondNumber))
		dom.RemoveClass("result", "error")
	} else {
		dom.SetValue("result", "value", "ERR")
		dom.AddClass("result", "error")
	}
}

func main() {

	// Start interacting with the DOM
	dom.Hide("loading")
	dom.Show("calc")

	dom.AddEventListener("myCanvas", "click", performCalculation)

	//dom.AddEventListener("first-number", "input", performCalculation)
	//dom.AddEventListener("second-number", "input", performCalculation)

	document := js.Global().Get("document")
	myCanvas := document.Call("getElementById", "myCanvas")
	ctx := myCanvas.Call("getContext", "2d")
	fmt.Println(ctx)
	ctx.Set("fillStyle", "red")
	ctx.Set("font", "30px Arial")
	ctx.Call("fillRect", 0, 0, 200, 50)
	ctx.Call("strokeText", "Calcular", 10, 35)

	for {
		dom.Show("mrBean")
		time.Sleep(time.Second)
		dom.Hide("mrBean")
		time.Sleep(time.Second)
	}
}
