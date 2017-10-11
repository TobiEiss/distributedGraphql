package distributedGraphql

type Schema struct {
	Directives   []Directive                   `json:"directives"`
	MutationType []map[string]interface{} `json:"mutationType"`
	QueryType    struct {
		Name string `name`
	} `json:"queryType"`
	SubscriptionType []map[string]interface{} `json:"subscriptionType"`
	Types            []Type                   `json:"types"`
}

type Type struct {
	Description   string    `json:"description"`
	EnumValues    []Field   `json:"enumValues"`
	Fields        []Field   `json:"fields"`
	InputFields   []TypeRef `json:"inputFields"`
	Interfaces    []TypeRef `json:"interfaces"`
	Kind          string    `json:"kind"`
	Name          string    `json:"name"`
	PossibleTypes []TypeRef `json:"possibleTypes"`
}

type Field struct {
	Args              []Arg  `json:"args"`
	DeprecationReason string `json:"deprecationReason"`
	Description       string `json:"description"`
	IsDeprecated      bool   `json:"isDeprecated"`
	Name              string `json:"name"`
	Type              Type   `json:"type"`
}

type Arg struct {
	DefaultValue interface{} `json:"defaultValue"`
	Description  string      `json:"description"`
	Name         string      `json:"name"`
	Type         TypeRef     `json:"type"`
}

type TypeRef struct {
	Kind   string `json:"kind"`
	Name   string `json:"name"`
	OfType struct {
		Kind   string `json:"kind"`
		Name   string `json:"name"`
		OfType struct {
			Kind   string `json:"kind"`
			Name   string `json:"name"`
			OfType struct {
				Kind   string      `json:"kind"`
				Name   string      `json:"name"`
				OfType interface{} `json:"ofType"`
			} `json:"ofType"`
		} `json:"ofType"`
	} `json:"ofType"`
}

type Directive struct {
	Args        []Arg    `json:"args"`
	Description string   `json:"description"`
	Locations   []string `json:"locations"`
	Name        string   `json:"name"`
	OnField     bool     `json:"onField"`
	OnFragment  bool     `json:"onFragment"`
	OnOperation bool     `json:"onOperation"`
}
