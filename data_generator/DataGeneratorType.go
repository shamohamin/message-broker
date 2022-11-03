package data_generator

func X2(val float32) float32 {
	return val * val
}

func X3(val float32) float32 {
	return val * val * val
}

func X4(val float32) float32 {
	return val * val * val * val
}

func X5(val float32) float32 {
	return val * val * val * val * val
}

type DataGenerator struct {
	InitialValueX 	float32
	FuncExc			func(float32) float32
	Step 			float32
}

func (generator *DataGenerator) Generate() float32 {
	if generator.FuncExc == nil {
		generator.FuncExc = X2
	}

	if generator.Step == float32(0) {
		generator.Step = float32(1)
	}

	retVal := generator.FuncExc(generator.InitialValueX)
	generator.InitialValueX += generator.Step

	return retVal
}

