package main

import (
	"Bangseungjae/insurance/danger"
	_ "Bangseungjae/insurance/danger"
	"Bangseungjae/insurance/store"
	"Bangseungjae/insurance/user"
	_ "Bangseungjae/insurance/user"
	"Bangseungjae/insurance/util"
	"context"
	"fmt"
	"github.com/go-yaml/yaml"
	"log"
	"net"
	"os"
	"path/filepath"
)

type Config struct {
	Products   []ProductsConfig `yaml:"products"`
	CoverRates CoverRatesConfig `yaml:"cover_rates"`
}
type ProductsConfig struct {
	Name   string         `yaml:"name"`
	Covers []CoversConfig `yaml:"covers"`
}

type CoversConfig struct {
	Name        string  `yaml:"name"`
	BasePremium float64 `yaml:"base_premium"`
}

type CoverRatesConfig struct {
	Sanghe  float64 `yaml:"sanghe"`
	Jilbung float64 `yaml:"jilbung"`
	Am      float64 `yaml:"am"`
}

var (
	config Config
)

func init() {
	filename, err := filepath.Abs("config/insurance.yml")
	util.Check(err)
	yamlFile, err := os.ReadFile(filename)
	util.Check(err)
	err = yaml.Unmarshal(yamlFile, &config)
	util.Check(err)
	//s, _ := json.MarshalIndent(config, "", "  ")
	log.Println("main init")
}

const (
	Sanghe  = "상해사망"
	Jilbung = "질병사망"
	Am      = "암진단"
)

func main() {
	setting()
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminated server: %v\n", err)
		os.Exit(1)
	}

}
func run(ctx context.Context) error {
	port := 8080
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", port, err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)

	mux, err := NewMux()
	s := NewServer(l, mux)
	return s.Run(ctx)
}

func setting() {
	Users := user.Users
	for _, u := range Users {
		for _, d := range danger.Dangers {
			if d.ToAge >= u.Age && d.FromAge <= u.Age && d.Gender == u.Gender {
				fmt.Printf("user name: %s\n", u.Name)
				var userInsurance store.UserInsurance
				userInsurance.Name = u.Name
				userInsurance.ID = u.Id

				for _, cp := range config.Products {
					fmt.Printf("%s : ", cp.Name)
					userInsurance.Insurance = *store.NewInsurance(cp.Name)
					for _, cover := range cp.Covers {
						riskRate := 0.0
						if cover.Name == Sanghe {
							riskRate = config.CoverRates.Sanghe
						} else if cover.Name == Jilbung {
							riskRate = config.CoverRates.Jilbung
						} else if cover.Name == Am {
							riskRate = config.CoverRates.Am
						}
						p := store.Price{
							Name:  cover.Name,
							Price: int(riskRate * cover.BasePremium * d.RiskRage),
						}
						userInsurance.Insurance.Prices = append(userInsurance.Insurance.Prices, p)
						//fmt.Printf("name: %s price: %f  ", cover.Name, riskRate*cover.BasePremium*d.RiskRage)
						fmt.Printf("name: %s price: %d  ", cover.Name, int(riskRate*cover.BasePremium*d.RiskRage))
					}
					fmt.Println()
				}
				store.UserInsurances[u.Id] = userInsurance
			}
		}
	}
}
