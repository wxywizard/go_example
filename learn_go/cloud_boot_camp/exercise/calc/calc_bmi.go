package calc

func CalcBMI(tall, weight float64) (bmi float64) {

	if tall <= 0 {
		panic("身高不能是0或者负数")
	}
	if weight <= 0 {
		panic("体重不能是0或者负数")
	}
	return weight / (tall * tall)
}
