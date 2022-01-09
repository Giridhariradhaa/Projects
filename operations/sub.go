package operations
import "fmt"


var Radha uint8 = 222 //global variable //this variable isn't working in main.go file. Why?
var Shannu uint8 = 11 //global variable

//used two variables from int8
func Subint8(){

	var a int8 = 22
	var b int8 = 122

	fmt.Println(a - b)
}

//used two variables from int16

func Subint16(){

	var a int16 = 226
	var b int16 = -12269

	fmt.Println(a - b)
}

//used two variables from int32

func Subint32(){

	var trine int32 = -6598485
	var university int32 = -6589423 //result printed is (-9062) but that is incorrect when you actually do -6598485 -6589423
	fmt.Println(trine - university)
}

//used two varibales from int64

func Subint64(){

var shannu = 2698745
 var radha = 526489
	fmt.Println(shannu - radha)
}

func SubUint8(){

	fmt.Println(Radha - Shannu)
}

func SubUint16(){

	var Trine = 126
	var University = 22

	fmt.Println( Trine - University)
}

func SubUint32(){

	fmt.Println(Radha + Shannu)
}

func SubUint64(){

	fmt.Println (Radha + c) //used global varible "Radha" from sub.go and package variable "c" from add.go file
}

func SubFloat32(){

	fmt.Println(riya - srija)
}

func SubFloat64(){

	fmt.Println(surya - rishi)
}