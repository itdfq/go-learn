package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age  int
	Sex  int
}

func (this *Person) String() string {
	buffer := bytes.NewBufferString("This is ")
	buffer.WriteString(this.Name + ", ")
	if this.Sex == 0 {
		buffer.WriteString("He ")
	} else {
		buffer.WriteString("She ")
	}

	buffer.WriteString("is ")
	buffer.WriteString(strconv.Itoa(this.Age))
	buffer.WriteString(" years old.")
	return buffer.String()
}
func (this *Person) GoString() string {
	return "&Person{Name is " + this.Name + ", Age is " + strconv.Itoa(this.Age) + ", Sex is " + strconv.Itoa(this.Sex) + "}"
}
func main() {
	p := &Person{"polaris", 28, 0}
	fmt.Printf("%#v", p)
}
