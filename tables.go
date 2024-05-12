package main


import (
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/jedib0t/go-pretty/v6/table"
)

var PersonalizedStyle = table.Style{
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
	}
