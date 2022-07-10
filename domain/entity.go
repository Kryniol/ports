package domain

type Port struct {
	ID      PortID
	Name    string
	Address Address
	Alias   []string
	Regions []string
	Unlocs  []PortID
	Code    string
}

func NewPort(
	id PortID,
	name string,
	address Address,
	alias []string,
	regions []string,
	unlocs []PortID,
	code string,
) *Port {
	return &Port{
		ID:      id,
		Name:    name,
		Address: address,
		Alias:   alias,
		Regions: regions,
		Unlocs:  unlocs,
		Code:    code,
	}
}
