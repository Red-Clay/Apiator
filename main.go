package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
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

func longSentenceFormat(sentence string) string {
	if len(sentence) > 6 {
		return sentence[:6] + "..."
	}
	return sentence
}

func Init() (inputMaxMachines uint64, inputOS string, inputPlatform string, inputDifficulty string, inputCertification string, inputName string, inputTechs string) {
	var help bool

	// Yet not pointers
	flag.Uint64Var(&inputMaxMachines, "max", 10, "Maximum number of machines to display.")
	flag.StringVar(&inputName, "n", "-1", "Search machine by name.")
	flag.StringVar(&inputTechs, "t", "-1", "Search machine by techniques.")
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
	inputMaxMachines, inputOS, inputPlatform, inputDifficulty, inputCertification, inputName, inputTechs := Init()

	tableHeaderMaxMachines := strconv.FormatUint(uint64(inputMaxMachines), 10)

	mappedFields := map[string]string{
		"platform":      inputPlatform,
		"name":          capitalizeFirstLetter(inputName),
		"techniques":    inputTechs,
		"certification": inputCertification,
		"video":         "-1",
		"ip":            "-1",
		"os":            capitalizeFirstLetter(inputOS),
		"state":         capitalizeFirstLetter(inputDifficulty),
	}

	parser := "newData"

	for k, v := range mappedFields {
		if v == "-1" {
			continue
		}
		if string(v[0]) == "!" {
			parser = fmt.Sprintf("%s|#(%s!%%*%s*)#", parser, k, string(v[1:]))
			continue
		}

		parser = fmt.Sprintf("%s|#(%s%%*%s*)#", parser, k, v)

	}

	dat, err := os.ReadFile("./machines.json")
	check(err)
	data := string(dat)

	// TABLE
	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)

	n_results := gjson.Get(data, parser+"|#")

	t.AppendHeader(
		table.Row{
			"Nombre (" + tableHeaderMaxMachines + "/" + n_results.String() + ")",
			"Plataforma",
			"Dificultad",
			"OS",
			"Nº / Certs",
			"Nº / Tecnicas",
		},
	)

	// https://github.com/tidwall/gjson/issues/64
	// result := gjson.Get(data, "newData.#(platform==HackTheBox)#|#(certification%*eWPT*)#|#(certification%*eWPTXv2*)#|#(certification!%*OSWE*)#")
	// fmt.Printf("\nGJSON: %v\n",gjson.Get(data, parser+"|#"))

	result := gjson.Get(data, parser)
	// fmt.Printf("%v\n",result)

	result.ForEach(func(key, value gjson.Result) bool {
		name := gjson.Get(value.Raw, "name")
		platform := gjson.Get(value.Raw, "platform")
		// youtube_link := gjson.Get(value.Raw, "video")
		operative_system := gjson.Get(value.Raw, "os")
		filterDifficultyArg := gjson.Get(value.Raw, "state").String()
		techniques := gjson.Get(value.Raw, "techniques")
		certs := gjson.Get(value.Raw, "certification")
		n_certs := strings.Split(certs.Raw, "\\n")
		techniques_sentences := strings.Split(techniques.Raw, "\\n")
		n_techs := strings.Split(certs.Raw, "\\n")

		if key.Uint() == inputMaxMachines {
			return false // stop iterating
		}

		short_techniques := ""
		// fmt.Printf("%v",techniques_sentences)
		for i, sentence := range techniques_sentences {
			short_techniques = short_techniques + longSentenceFormat(sentence)
			// fmt.Printf("Resultado:%s\n",short_techniques)
			if i%4 == 0 && i >= 3 {
				short_techniques = short_techniques + "\n"
			}
		}

		easy_color := color.New(color.FgGreen).SprintFunc()
		medium_color := color.New(color.FgYellow).SprintFunc()
		hard_color := color.New(color.FgRed).SprintFunc()
		insane_color := color.New(color.FgMagenta).SprintFunc()

		switch diff := filterDifficultyArg; diff {
		case "Easy":
			filterDifficultyArg = easy_color(filterDifficultyArg)
		case "Medium":
			filterDifficultyArg = medium_color(filterDifficultyArg)
		case "Hard":
			filterDifficultyArg = hard_color(filterDifficultyArg)
		case "Insane":
			filterDifficultyArg = insane_color(filterDifficultyArg)
		default:
		}
		formatedCerts := longSentenceFormat(strings.Join(n_certs, ","))
		strCountCerts := strconv.FormatUint(uint64(len(n_certs)), 10)
		certsValue := strCountCerts + " / " + formatedCerts

		formatedTechs := short_techniques
		strCountTechs := strconv.FormatUint(uint64(len(n_techs)), 10)
		techsValue := strCountTechs + " / " + formatedTechs

		t.AppendRows([]table.Row{
			{
				name.String(), /* + " /\n" + youtube_link.String()  */
				platform,
				filterDifficultyArg,
				operative_system,
				certsValue,
				techsValue,
			},
		})

		t.AppendSeparator()
		return true // keep iterating
	})
	// t.SetStyle(table.StyleLight)

	// t.SortBy([]table.SortBy{
	//   {Name: "Dificultad", Mode: table.Asc},
	//   // {Name: "OS", Mode: table.Asc},
	//  })

	// t.SetAllowedRowLength(91)

	t.SetStyle(table.Style{
		Name: "myNewStyle",
		Box: table.BoxStyle{
			BottomLeft:       "┗",
			BottomRight:      "┛",
			BottomSeparator:  "┻",
			EmptySeparator:   text.RepeatAndTrim(" ", text.RuneWidthWithoutEscSequences("╋")),
			Left:             "┃",
			LeftSeparator:    "┣",
			MiddleHorizontal: "━",
			MiddleSeparator:  "╋",
			MiddleVertical:   "┃",
			PaddingLeft:      " ",
			PaddingRight:     " ",
			PageSeparator:    "\n",
			Right:            "┃",
			RightSeparator:   "┫",
			TopLeft:          "┏",
			TopRight:         "┓",
			TopSeparator:     "┳",
			UnfinishedRow:    " ≈",
		},
		Color: table.ColorOptions{
			Header:       text.Colors{text.BgHiGreen, text.FgBlack, text.Bold},
			Row:          text.Colors{text.BgHiBlack},
			RowAlternate: text.Colors{text.BgHiBlack},
		},
		Format: table.FormatOptions{
			Footer: text.FormatUpper,
			Header: text.FormatUpper,
			Row:    text.FormatDefault,
		},
		Options: table.Options{
			DrawBorder:      true,
			SeparateColumns: true,
			SeparateFooter:  true,
			SeparateHeader:  true,
			SeparateRows:    false,
		},
	})

	t.Render()
}
