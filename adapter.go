package main

import "fmt"

type OldInterface interface {
	Operation1()
	Operation2()
}

type OldExampleStruct struct {
}

func getOldExampleStruct() (oldExampleStruct *OldExampleStruct) {
	oldExampleStruct = &OldExampleStruct{}
	return
}

func (OldExampleStruct *OldExampleStruct) Operation1() {
	fmt.Println("Operation 1 of OldExampleStruct")
}

func (OldExampleStruct *OldExampleStruct) Operation2() {
	fmt.Println("Operation 2 of OldExampleStruct")
}

type NewInterface interface {
	NewOperation1()
	NewOperation2()
	NewOperation3()
}

type NewExampleStruct struct {
}

func getNewExampleStruct() (newExampleStruct *NewExampleStruct) {
	newExampleStruct = &NewExampleStruct{}
	return
}

func (newExampleStruct *NewExampleStruct) NewOperation1() {
	fmt.Println("NewOperation 1 of NewExampleStruct")
}

func (newExampleStruct *NewExampleStruct) NewOperation2() {
	fmt.Println("NewOperation 2 of NewExampleStruct")
}

func (newExampleStruct *NewExampleStruct) NewOperation3() {
	fmt.Println("NewOperation 3 of NewExampleStruct")
}

type Adapter struct {
	oldInterface OldInterface
}

func getAdapter(oldInterface OldInterface) (adapter *Adapter) {
	adapter = &Adapter{}
	adapter.oldInterface = oldInterface
	return
}

func (adapter *Adapter) NewOperation1() {

	fmt.Println("Inside NewOperation1")

	adapter.oldInterface.Operation1()
}

func (adapter *Adapter) NewOperation2() {
	fmt.Println("Inside NewOperation2")

	adapter.oldInterface.Operation2()
}

func (adapter *Adapter) NewOperation3() {
	panic("Operation not avaiable")
}

func main() {
	oldStruct := getOldExampleStruct()
	adapter := getAdapter(oldStruct)

	adapter.NewOperation1()
	adapter.NewOperation2()
}
