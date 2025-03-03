package pkg

type Calculator struct {
	Data string
}

func (receiver Calculator) NewCalc() string {
	return receiver.Data
}
