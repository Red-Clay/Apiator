package main

import "github.com/tidwall/gjson"
import (
	"flag"
	"strings"
	"fmt"
	"os"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Init() (inputMaxMachines int, inputOS string, inputPlatform string, inputDifficulty string, inputCertification string,blacklistCert string) {
	var help bool

	flag.IntVar(&inputMaxMachines, "max", 10, "Define the maximum number of machines to display.")
	flag.StringVar(&inputDifficulty, "dif", "all", "Search for machines with a specific level of difficulty.")
	flag.StringVar(&inputOS, "os", "all", "Search for machines with a specific operating system.")
	flag.StringVar(&inputCertification, "cert", "all", "Search for machines that have a specific certification.")
	flag.StringVar(&inputPlatform, "plat", "all", "Add a certification to the whitelist.")
	flag.StringVar(&blacklistCert, "bcert", "all", "Add a certification to the blacklist.")
	flag.BoolVar(&help, "help", false, "Display information about usage.")

	flag.Parse()

	if help {
		color.Cyan("Uso:")
		flag.PrintDefaults()
		os.Exit(0)
	}
  return
	// return inputMaxMachines, inputOS, inputPlatform, inputDifficulty, inputCertification,blac
}

func main() {
	inputMaxMachines, inputOS, inputPlatform, inputDifficulty, inputCertification,blacklistCert := Init()
	fmt.Printf("All flags:\n")
	fmt.Printf("-max: %v\n",inputMaxMachines)
	fmt.Printf("-dif: %v\n",inputDifficulty)
	fmt.Printf("-os: %v\n",inputOS)
	fmt.Printf("-plat: %v\n",inputPlatform)
	fmt.Printf("-cert: %v\n",inputCertification)
	fmt.Printf("-bcert: %v\n",blacklistCert)

	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Video", "Nombre", "Plataforma", "Dificultad", "OS", "Nº Certificaciones","Certificaciones"})

	dat, err := os.ReadFile("./machines.json")
	check(err)
	data := string(dat)

	parsePlatform := fmt.Sprintf("|#(platform%%*%s*)#",inputPlatform);
	parseCertification := fmt.Sprintf("|#(certification%%*%s*)#",inputCertification);
	parseDifficulty := fmt.Sprintf("|#(state%%*%s*)#",inputDifficulty);
  parseOS := fmt.Sprintf("|#(os%%*%s*)#",inputOS);
  parseBlackCertification := fmt.Sprintf("|#(certification!%%*%s*)#",blacklistCert);
     
	parser := fmt.Sprintf("newData%s%s%s%s%s", parsePlatform, parseCertification, parseDifficulty, parseOS,parseBlackCertification)
    
	fmt.Printf("GJSON: %v\n",parser)

// https://github.com/tidwall/gjson/issues/64
	// result := gjson.Get(data, "newData.#(platform==HackTheBox)#|#(certification%*eWPT*)#|#(certification%*eWPTXv2*)#|#(certification!%*OSWE*)#")
	result := gjson.Get(data, parser)


	result.ForEach(func(key, value gjson.Result) bool {
		name := gjson.Get(value.Raw, "name")
		platform := gjson.Get(value.Raw, "platform")
		youtube_link := gjson.Get(value.Raw, "video")
		operative_system := gjson.Get(value.Raw, "os")
		difficulty := gjson.Get(value.Raw, "state")
		certs := gjson.Get(value.Raw, "certification")
		n_certs := strings.Split(certs.Raw, "\\n")
		fmt.Printf("%T",certs)
		// strings.Replace(certs, "fg", "FG", -1)

		t.AppendRows([]table.Row{
			{youtube_link, name, platform, difficulty, operative_system, len(n_certs), certs},
		})

		t.AppendSeparator()
		return true // keep iterating
	})
	t.SetStyle(table.StyleLight)
	t.Render()
}
