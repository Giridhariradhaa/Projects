package main
import ( 
"Calculatormodule/operations"
"fmt"

)
func main (){

	//used int8, int16, int32 and int64 ---- ADDITION

	operations.Add()	
	operations.Addint8()
	operations.Addint16()
	operations.Addint32()
	operations.Addint64()
	//
	//used Uint8, Uint16, Uint32 and Uint64 ----------ADDITION
	operations.AddUint8()
	operations.AddUint8()
	operations.AddUint16()
	operations.AddUint32()
	operations.AddUint64()

	//Used float32 and float64 -------- ADDITION
	operations.AddFloat32()
	operations.AddFloat64()
	//

	//used int8, int16, int32 and int64 -----------SUBTRACTION
	operations.Subint8()
	operations.SubUint8()
	operations.Subint16()
	operations.Subint32()
	operations.Subint64()

	//
	//used Uint8, Uint16, Uint32 and Uint64 -------SUBTRACTION
	operations.SubUint8()
	operations.SubUint16()
	operations.SubUint32()
	operations.SubUint64()
	//
	//used Float32 and Float64 ----------- SUBTRACTION
	operations.SubFloat32()
	operations.SubFloat64()
	//

	fmt.Println(operations.Radha) //used global pack of type Uint8
	//
	//used int8, int16, int32 and int64 --------------- MULTIPLICATION
	operations.Multint8()
	operations.Multint16()
	operations.Multint32()
	operations.Multint64()
	//
	//used Uint8, Uint16, Uint32 and Uint64 ------------- MULTIPLICATION
	operations.MultUint8()
	operations.MultUint16()
	operations.MultUint32() 
	operations.MultUint64() 
	//
	//used Float32 and Float64 ----------MULTIPLICATION

	operations.MultFloat32()
	operations.MultFloat64()

	//
	//used int8, int16, int32 and int64 ----------DIVISION
	operations.Divideint8()
		operations.Divideint16()
	operations.Divideint32()
	operations.Divideint64()

	//
	//used Uint8, Uint16, Uint32 and Uint64 -------------DIVISION
	operations.DivideUint8()
	operations.DivideUint16()
	operations.DivideUint32()
	operations.DivideUint64()
	operations.DivideUint64()
	//
	//used Float32 and Float64 -----------------DIVISION
	operations.DivideFloat32()
	operations.DivideFloat64()

	

}