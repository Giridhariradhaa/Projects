package operations

import 

"fmt"


var c uint8  = 22 //package variable


	var riya float32 = 216.9 //package variable
	
var srija float32 = 22.22 //package variable

var surya float64 = 221.8 //package variable

	var rishi float64 = 12.364 //package variable


func Add(){

//	var a int = 125   This a+b give the value of 135 which is outside the range. Hence this is not a correct answer. We have to declare the value of a variable in the way where the output of it, will be within the range of that particular data type. For example from -128 to 127 for int8.
//var b int = 10

var a int = 110
var b int = 15

	fmt.Println(a + b)
}

//Used two variables from int8

func Addint8(){

var a int8 = 29 //int8

var b int8 = 122

// var b int8 = -129// " -129 overflows, hence it doesn't work for int8 = error occurrs"

fmt.Println (a + b)}

//used two variables from int16

func Addint16(){

	var a int16 = -32762

	var b int16 = 32568

fmt.Println (a + b)
}

//used two variables from int32

func Addint32(){
var radha int32 = 214748
var shannu int32 = -214748

fmt.Println(radha + shannu)

}

//used two variables from int64

func Addint64(){

	var trine int64 = 123456987
	var university int64 = -987654321

	fmt.Println(trine + university)
}

//used two variables from Uint8

func AddUint8(){
  
	var a uint8
	a = 129
	var b uint8
	b = 26

	fmt.Println(a + b + c)

}

func AddUint16(){

	var a uint16
	a = 1981
	var b uint16 = 2481

	fmt.Println (a + b)

}

func AddUint32(){

	var a uint32 = 1236
	var b uint32 = 2

	fmt.Println(a + b)
}

func AddUint64(){

	var radha int64 = 138
	var shannu int64 = -16 

	fmt.Println(radha + shannu)
}

func AddFloat32(){

	//var riya float32 
	//riya = 216.9

//var srija float32
//srija = 22.22

	fmt.Println(riya + srija)
}

func AddFloat64(){

	//var surya float64 = 221.8
	//	var rishi float64 = 12.364

	fmt.Println(surya + rishi)
}


		

