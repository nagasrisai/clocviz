package utils

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"math/rand"
)

func GetLangColor(lang string) string {
	if color, ok := COLORS[lang]; ok {
		return color
	}
	// If color is not in COLORS, we save a random color for future usage.
	rand := getRandomColorInHex()
	COLORS[lang] = rand
	return rand
}

type RGB struct {
	red   int
	green int
	blue  int
}

func GradateHex(hex string, percentage float32) string {
	rgb := hexToRGB(hex)
	red := strconv.FormatInt(int64(float32(rgb.red)*percentage), 16)
	green := strconv.FormatInt(int64(float32(rgb.green)*percentage), 16)
	blue := strconv.FormatInt(int64(float32(rgb.blue)*percentage), 16)
	return fmt.Sprintf("#%v%v%v", red, green, blue)
}

func getRandomColorInHex() string {
	rand.Seed(time.Now().UnixNano())
	red := rand.Intn(255)
	green := rand.Intn(255)
	blue := rand.Intn(255)
	hex := "#" + getHex(red) + getHex(green) + getHex(blue)
	return hex
}

func getHex(num int) string {
	hex := fmt.Sprintf("%x", num)
	if len(hex) == 1 {
		hex = "0" + hex
	}
	return hex
}

func hexToRGB(hex string) *RGB {
	rgb := &RGB{}
	_, err := fmt.Sscanf(hex, "#%02x%02x%02x", &rgb.red, &rgb.green, &rgb.blue)
	if err != nil {
		log.Fatal(err)
	}
	return rgb
}

// Associations between language name and official color (in hex)
// Source: https://github.com/doda/github-language-colors/blob/master/colors.json
var COLORS = map[string]string{
	"ABAP":                  "#E8274B",
	"AGS Script":            "#B9D9FF",
	"AMPL":                  "#E6EFBB",
	"ANTLR":                 "#9DC3FF",
	"API Blueprint":         "#2ACCA8",
	"APL":                   "#5A8164",
	"ASP":                   "#6a40fd",
	"ATS":                   "#1ac620",
	"ActionScript":          "#882B0F",
	"Ada":                   "#02f88c",
	"Agda":                  "#315665",
	"Alloy":                 "#64C800",
	"Arc":                   "#aa2afe",
	"Arduino":               "#bd79d1",
	"AspectJ":               "#a957b0",
	"Assembly":              "#6E4C13",
	"AutoHotkey":            "#6594b9",
	"AutoIt":                "#1C3552",
	"BlitzMax":              "#cd6400",
	"Boo":                   "#d4bec1",
	"Brainfuck":             "#2F2530",
	"C Sharp":               "#178600",
	"C":                     "#555555",
	"CSS":                   "#563d7c",
	"Chapel":                "#8dc63f",
	"Cirru":                 "#ccccff",
	"Clarion":               "#db901e",
	"Clean":                 "#3F85AF",
	"Click":                 "#E4E6F3",
	"Clojure":               "#db5855",
	"CoffeeScript":          "#244776",
	"ColdFusion CFC":        "#ed2cd6",
	"ColdFusion":            "#ed2cd6",
	"Common Lisp":           "#3fb68b",
	"Component Pascal":      "#b0ce4e",
	"Crystal":               "#776791",
	"D":                     "#ba595e",
	"DM":                    "#447265",
	"Dart":                  "#00B4AB",
	"Diff":                  "#88dddd",
	"Dogescript":            "#cca760",
	"Dylan":                 "#6c616e",
	"E":                     "#ccce35",
	"ECL":                   "#8a1267",
	"Eagle":                 "#814C05",
	"Eiffel":                "#946d57",
	"Elixir":                "#6e4a7e",
	"Elm":                   "#60B5CC",
	"Emacs Lisp":            "#c065db",
	"EmberScript":           "#FFF4F3",
	"Erlang":                "#B83998",
	"F#":                    "#b845fc",
	"FLUX":                  "#88ccff",
	"FORTRAN":               "#4d41b1",
	"Factor":                "#636746",
	"Fancy":                 "#7b9db4",
	"Fantom":                "#dbded5",
	"Forth":                 "#341708",
	"FreeMarker":            "#0050b2",
	"Frege":                 "#00cafe",
	"Game Maker Language":   "#8fb200",
	"Glyph":                 "#e4cc98",
	"Gnuplot":               "#f0a9f0",
	"Go":                    "#375eab",
	"Golo":                  "#88562A",
	"Gosu":                  "#82937f",
	"Grammatical Framework": "#79aa7a",
	"Groovy":                "#e69f56",
	"HTML":                  "#e44b23",
	"Handlebars":            "#01a9d6",
	"Harbour":               "#0e60e3",
	"Haskell":               "#29b544",
	"Haxe":                  "#df7900",
	"Hy":                    "#7790B2",
	"IDL":                   "#a3522f",
	"Io":                    "#a9188d",
	"Ioke":                  "#078193",
	"Isabelle":              "#FEFE00",
	"J":                     "#9EEDFF",
	"JFlex":                 "#DBCA00",
	"JSONiq":                "#40d47e",
	"Java":                  "#b07219",
	"JavaScript":            "#f1e05a",
	"Julia":                 "#a270ba",
	"Jupyter Notebook":      "#DA5B0B",
	"KRL":                   "#28431f",
	"Kotlin":                "#F18E33",
	"LFE":                   "#004200",
	"LOLCODE":               "#cc9900",
	"LSL":                   "#3d9970",
	"Lasso":                 "#999999",
	"Latte":                 "#A8FF97",
	"Lex":                   "#DBCA00",
	"LiveScript":            "#499886",
	"LookML":                "#652B81",
	"Lua":                   "#000080",
	"MAXScript":             "#00a6a6",
	"MTML":                  "#b7e1f4",
	"Makefile":              "#427819",
	"Mask":                  "#f97732",
	"Matlab":                "#bb92ac",
	"Max":                   "#c4a79c",
	"Mercury":               "#ff2b2b",
	"Metal":                 "#8f14e9",
	"Mirah":                 "#c7a938",
	"NCL":                   "#28431f",
	"Nemerle":               "#3d3c6e",
	"NetLinx":               "#0aa0ff",
	"NetLinx+ERB":           "#747faa",
	"NetLogo":               "#ff6375",
	"NewLisp":               "#87AED7",
	"Nimrod":                "#37775b",
	"Nit":                   "#009917",
	"Nix":                   "#7e7eff",
	"Nu":                    "#c9df40",
	"OCaml":                 "#3be133",
	"Objective-C":           "#438eff",
	"Objective-C++":         "#6866fb",
	"Objective-J":           "#ff0c5a",
	"Omgrofl":               "#cabbff",
	"Opal":                  "#f7ede0",
	"Oxygene":               "#cdd0e3",
	"Oz":                    "#fab738",
	"PAWN":                  "#dbb284",
	"PHP":                   "#4F5D95",
	"PLSQL":                 "#dad8d8",
	"Pan":                   "#cc0000",
	"Papyrus":               "#6600cc",
	"Parrot":                "#f3ca0a",
	"Pascal":                "#b0ce4e",
	"Perl":                  "#0298c3",
	"Perl6":                 "#0000fb",
	"PigLatin":              "#fcd7de",
	"Pike":                  "#005390",
	"PogoScript":            "#d80074",
	"Processing":            "#0096D8",
	"Prolog":                "#74283c",
	"Propeller Spin":        "#7fa2a7",
	"Puppet":                "#302B6D",
	"Pure Data":             "#91de79",
	"PureBasic":             "#5a6986",
	"PureScript":            "#1D222D",
	"Python":                "#3572A5",
	"QML":                   "#44a51c",
	"R":                     "#198ce7",
	"RAML":                  "#77d9fb",
	"Racket":                "#22228f",
	"Ragel in Ruby Host":    "#9d5200",
	"Rebol":                 "#358a5b",
	"Red":                   "#ee0000",
	"Ren'Py":                "#ff7f7f",
	"Rouge":                 "#cc0088",
	"Ruby":                  "#701516",
	"Rust":                  "#dea584",
	"SAS":                   "#B34936",
	"SQF":                   "#3F3F3F",
	"SaltStack":             "#646464",
	"Scala":                 "#DC322F",
	"Scheme":                "#1e4aec",
	"Self":                  "#0579aa",
	"Shell":                 "#89e051",
	"Shen":                  "#120F14",
	"Slash":                 "#007eff",
	"Slim":                  "#ff8f77",
	"Smalltalk":             "#596706",
	"SourcePawn":            "#5c7611",
	"Squirrel":              "#800000",
	"Stan":                  "#b2011d",
	"Standard ML":           "#dc566d",
	"SuperCollider":         "#46390b",
	"Swift":                 "#ffac45",
	"SystemVerilog":         "#DAE1C2",
	"Tcl":                   "#e4cc98",
	"TeX":                   "#3D6117",
	"Turing":                "#45f715",
	"TypeScript":            "#2b7489",
	"Unified Parallel C":    "#4e3617",
	"Unity3D Asset":         "#ab69a1",
	"UnrealScript":          "#a54c4d",
	"VHDL":                  "#adb2cb",
	"Vala":                  "#fbe5cd",
	"Verilog":               "#b2b7f8",
	"VimL":                  "#199f4b",
	"Visual Basic":          "#945db7",
	"Volt":                  "#1F1F1F",
	"Vue":                   "#2c3e50",
	"Web Ontology Language": "#9cc9dd",
	"X10":                   "#4B6BEF",
	"XC":                    "#99DA07",
	"XQuery":                "#5232e7",
	"Zephir":                "#118f9e",
	"cpp":                   "#f34b7d",
	"eC":                    "#913960",
	"edn":                   "#db5855",
	"nesC":                  "#94B0C7",
	"ooc":                   "#b0b77e",
	"wisp":                  "#7582D1",
	"xBase":                 "#403a40",
}
