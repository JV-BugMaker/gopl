package tempconv

import "fmt"

type Celsius float64 //定义摄氏度类型

type Fahrenheit float64 //定义华氏温度类型


const (
	AbsoluteZeroC Celsius = -273.15 // 絶對零度
	FreezingC     Celsius = 0       // 結冰點溫度
	BoilingC      Celsius = 100     // 沸水溫度
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
