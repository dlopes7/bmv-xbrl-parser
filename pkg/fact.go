package pkg

type Fact struct {
	ID           string   `json:"Id"`
	ConceptID    string   `json:"IdConcepto"`
	ConceptName  string   `json:"NombreConcepto"`
	ContextID    string   `json:"IdContexto"`
	UnityID      string   `json:"IdUnidad"`
	Decimals     string   `json:"Decimales"`
	Type         int      `json:"Tipo"`
	Value        string   `json:"Valor"`
	NumericValue int      `json:"ValorNumerico"`
	Denominator  *float64 `json:"ValorDenominador"`
	Numerator    *float64 `json:"ValorNumerador"`
}
