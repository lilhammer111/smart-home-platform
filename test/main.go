package main

import (
	"fmt"
	"reflect"
)

type OriginalResponse struct {
	Meta struct {
		Success bool
		Code    int
		Message string
	}
	Data struct {
		Name string
		Age  int
	}
}

type FlattenedResponse struct {
	Success bool
	Code    int
	Message string
	Data    struct {
		Name string
		Age  int
	}
}

// CopyFields copies fields from src to dst struct by matching field names.
func CopyFields(src, dst interface{}) {
	srcVal := reflect.ValueOf(src).Elem()
	dstVal := reflect.ValueOf(dst).Elem()

	fieldSuccess := srcVal.FieldByName("Meta").FieldByName("Success")
	fieldCode := srcVal.FieldByName("Meta").FieldByName("Code")
	fieldMessage := srcVal.FieldByName("Meta").FieldByName("Message")
	srcDataField := srcVal.FieldByName("Data")

	dstVal.FieldByName("Success").Set(fieldSuccess)
	dstVal.FieldByName("Code").Set(fieldCode)
	dstVal.FieldByName("Message").Set(fieldMessage)
	dstVal.FieldByName("Data").Set(srcDataField)
}

func getCopy() {
	original := OriginalResponse{
		Meta: struct {
			Success bool
			Code    int
			Message string
		}{Success: true, Code: 200, Message: "yes"},
		Data: struct {
			Name string
			Age  int
		}{Name: "John Doe", Age: 30},
	}

	var flattened FlattenedResponse
	CopyFields(&original, &flattened)

	fmt.Printf("Flattened: %+v\n", flattened)
}

const (
	Normal int8 = -64 >> iota
	PowerOn
	Connected
	Idle
	Charging
)

func PrintPowerOn() {
	fmt.Println(PowerOn)
}

func main() {
	PrintPowerOn()
}
