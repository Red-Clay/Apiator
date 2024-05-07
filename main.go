package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/tidwall/gjson"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// func stringInSlice(a string, list []string) bool {
//     for _, b := range list {
//         if b == a {
//             return true
//         }
//     }
//     return false
// }

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}


func Init() (inputMaxMachines uint64, inputOS string, inputPlatform string, inputDifficulty string, inputCertification string, inputName string, inputTechs string) {
	var help bool

	// Yet not pointers
	flag.Uint64Var(&inputMaxMachines, "max", 10, "Maximum number of machines to display.")
	flag.StringVar(&inputName, "n", "-1" , "Search machine by name.")
	flag.StringVar(&inputTechs, "t", "-1" , "Search machine by techniques.")
	flag.StringVar(&inputDifficulty, "d", "-1", "Search machines by difficulty.")
	flag.StringVar(&inputOS, "o", "-1", "Search machines by operating system.")
	flag.StringVar(&inputCertification, "c", "-1", "Search machines by certifications.")
	flag.StringVar(&inputPlatform, "p", "-1", "Search by platform.")
	flag.BoolVar(&help, "help", false, "Display information about usage.")

	flag.Parse()

	if help {
		color.Cyan("\nUso:")
		flag.PrintDefaults()
		os.Exit(0)
	}


  // if strings.EqualFold("htb",inputPlatform) || strings.EqualFold("hackthebox",inputPlatform){
		// inputPlatform = "HackTheBox"
  // }

  return
}

func main() {


  // caser := cases.Title(language.English)

	// ARGUMENTS
	inputMaxMachines, inputOS, inputPlatform, inputDifficulty, inputCertification,inputName,inputTechs := Init()

	str_MaxMachines := strconv.FormatUint(uint64(inputMaxMachines), 10)

	mappedFields := map[string]string{
		"platform" : inputPlatform,
		"name" : capitalizeFirstLetter(inputName),
		"techniques" : inputTechs,
		"certification" :inputCertification,
		"video" : "-1",
		"ip" : "-1",
		"os" : capitalizeFirstLetter(inputOS),
		"state" : capitalizeFirstLetter(inputDifficulty),
	}


	parser:= "newData"

	for k, v := range mappedFields {
		if v == "-1" {
			continue
		}
		if string(v[0]) == "!" {
			parser = fmt.Sprintf("%s|#(%s!%%*%s*)#",parser,k,string(v[1:])) 
			continue
		}

		 parser = fmt.Sprintf("%s|#(%s%%*%s*)#",parser,k,v) 

  }

	dat, err := os.ReadFile("./machines.json")
	check(err)
	data := string(dat)

	// TABLE 
	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)

	n_results := gjson.Get(data, parser+"|#")

	t.AppendHeader(table.Row{"Nombre / Video (" + str_MaxMachines + "/" + n_results.String() + ")", "Plataforma", "Dificultad", "OS", "Certificaciones/ Nº","Tecnicas"})



// https://github.com/tidwall/gjson/issues/64
	// result := gjson.Get(data, "newData.#(platform==HackTheBox)#|#(certification%*eWPT*)#|#(certification%*eWPTXv2*)#|#(certification!%*OSWE*)#")
	// fmt.Printf("\nGJSON: %v\n",gjson.Get(data, parser+"|#"))

	result := gjson.Get(data, parser)
  // fmt.Printf("%v\n",result)

	result.ForEach(func(key, value gjson.Result) bool {
		name := gjson.Get(value.Raw, "name")
		platform := gjson.Get(value.Raw, "platform")
		youtube_link := gjson.Get(value.Raw, "video")
		operative_system := gjson.Get(value.Raw, "os")
		difficulty := gjson.Get(value.Raw, "state")
		techniques := gjson.Get(value.Raw, "techniques")
		certs := gjson.Get(value.Raw, "certification")
		n_certs := strings.Split(certs.Raw, "\\n")
		if key.Uint() == inputMaxMachines {
			return false // stop iterating
		}


		
		t.AppendRows([]table.Row{
			{ name.String() + " /\n" + youtube_link.String()  , platform, difficulty, operative_system,  strings.Join(n_certs, ",") + " / " + strconv.FormatUint(uint64(len(n_certs)), 10), techniques },
		})

		t.AppendSeparator()
		return true // keep iterating
	})
	t.SetStyle(table.StyleLight)
	t.Render()
}
