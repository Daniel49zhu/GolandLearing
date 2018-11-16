package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/8 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
