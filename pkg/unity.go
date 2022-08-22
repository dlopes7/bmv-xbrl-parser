package pkg

type Unit struct {
	ID           string         `json:"Id"`
	Type         int            `json:"Tipo"`
	Measurements []Measurement  `json:"Medidas"`
	Numerator    *[]Measurement `json:"MedidasNumerador"`
	Denominator  *[]Measurement `json:"MedidasDenominador"`
}

type Measurement struct {
	Name string `json:"Nombre"`
	Tag  string `json:"Etiqueta"`
}
