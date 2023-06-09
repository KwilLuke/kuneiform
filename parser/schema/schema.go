package schema

type Schema struct {
	Owner   string   `json:"owner"`
	Name    string   `json:"name"`
	Tables  []Table  `json:"tables,omitempty"`
	Actions []Action `json:"actions,omitempty"`
}

type Table struct {
	Name    string   `json:"name"`
	Columns []Column `json:"columns,omitempty"`
	Indexes []Index  `json:"indexes,omitempty"`
}

type Column struct {
	Name       string      `json:"name"`
	Type       ColumnType  `json:"type"`
	Attributes []Attribute `json:"attributes,omitempty"`
}

type Attribute struct {
	Type  AttributeType `json:"type"`
	Value any           `json:"value,omitempty"`
}

type Index struct {
	Name    string    `json:"name"`
	Columns []string  `json:"columns"`
	Type    IndexType `json:"type,omitempty"`
}
type Action struct {
	Name       string   `json:"name"`
	Public     bool     `json:"public"`
	Inputs     []string `json:"inputs"`
	Statements []string `json:"statements,omitempty"`
}

func (s *Schema) Node()    {}
func (t *Table) Node()     {}
func (c *Column) Node()    {}
func (a *Attribute) Node() {}
func (i *Index) Node()     {}
func (a *Action) Node()    {}

//func (s *Schema) Validate() error {
//	return nil
//}
