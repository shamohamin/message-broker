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
	if generating.FuncExc == nil {
		generating.FuncExc = X2
	}

	if generating.Step == float32(0) {
		generating.Step = float32(1)
	}

	retVal := generating.FuncExc(generating.InitialValueX)
	generating.InitialValueX += generating.Step

	return retVal
}

