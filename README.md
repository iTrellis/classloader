# classloader
class loader in Go

## [Example](example/main.go)

```go
package main

import (
	"fmt"
	"reflect"

	"github.com/iTrellis/classloader"
)

type Test struct{}

func (p *Test) Construct() {}

func (p Test) Hello(name string) {
	fmt.Println("hello:", name)
}

func main() {
	c := classloader.Default

	c.LoadClass("classloader:Test", (*Test)(nil))

	c.LoadClass("", (*Test)(nil))

	// unsupported class' type: string
	c.LoadClass("classloader:Aha", "name")

	NameTest()

	ClassLoaderTest()

	DynamicAccess()
}

func ClassLoaderTest() {

	t, e := classloader.Default.FindClass((*Test)(nil))
	if !e {
		fmt.Println("error:", "classloader:Test not exist")
		return
	}

	test := reflect.New(t)

	if !test.IsValid() {
		fmt.Println("error:", "classloader:Test is invalid")
		return
	}

	testI, _ := test.Interface().(*Test)

	testI.Hello("class loaded")

}

func NameTest() {

	t, e := classloader.Default.FindClass("classloader:Test")
	if !e {
		fmt.Println("error:", "name class not exist")
		return
	}

	test := reflect.New(t)

	if !test.IsValid() {
		fmt.Println("error:", "name class is invalid")
		return
	}

	testI, _ := test.Interface().(*Test)

	testI.Hello("name class loaded")
}

func DynamicAccess() {
	da := classloader.NewReflectiveDynamicAccess(classloader.Default)

	fmt.Println(da.GetClassFor("classloader:Test"))

	ins, e := da.CreateInstanceByName("classloader:Test")
	if e != nil {
		fmt.Println("error:", e)
		return
	}

	testI, _ := ins.(*Test)

	testI.Hello("DynamicAccess")
}
```