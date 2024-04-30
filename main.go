package main

import "github.com/tidwall/gjson"
import (
	_ "fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Video", "Nombre", "Plataforma", "Dificultad", "OS", "Certificaciones"})

	dat, err := os.ReadFile("./machines.json")
	check(err)
	data := string(dat)

	result := gjson.Get(data, "newData.#(platform==HackTheBox)#")

	result.ForEach(func(key, value gjson.Result) bool {
		name := gjson.Get(value.Raw, "name")
		platform := gjson.Get(value.Raw, "platform")
		youtube_link := gjson.Get(value.Raw, "video")
		operative_system := gjson.Get(value.Raw, "os")
		difficulty := gjson.Get(value.Raw, "state")
		certs := gjson.Get(value.Raw, "certification")
		n_certs := strings.Split(certs.Raw, "\\n")

		t.AppendRows([]table.Row{
			{youtube_link, name, platform, difficulty, operative_system, len(n_certs)},
		})

		t.AppendSeparator()
		return true // keep iterating
	})

	t.SetStyle(table.StyleLight)
	t.Render()

}
