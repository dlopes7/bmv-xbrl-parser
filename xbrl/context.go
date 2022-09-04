package xbrl

type Context struct {
	ID string `json:"Id"`

	Period Period `json:"Periodo"`
	Entity Entity `json:"Entidad"`
	Facts  []Fact
}
