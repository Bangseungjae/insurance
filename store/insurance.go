package store

type UserInsurance struct {
	ID        int       `json:"id"`
	Name      string    `json:"user_name"`
	Insurance Insurance `json:"insurance"`
}

type Insurance struct {
	InsuranceName string  `json:"insurance_name"`
	Prices        []Price `json:"prices"`
}

type Price struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func NewInsurance(insuranceName string) *Insurance {
	return &Insurance{
		InsuranceName: insuranceName,
		Prices:        make([]Price, 0, 10),
	}
}

var (
	UserInsurances map[int]UserInsurance
)

func init() {
	UserInsurances = make(map[int]UserInsurance)
}
