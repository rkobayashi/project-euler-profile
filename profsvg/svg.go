package profsvg

import (
	"html/template"
	"io"
	"math"
)

type profile struct {
	UserName string `xml:"username"`
	Country  string `xml:"country"`
	Language string `xml:"language"`
	Solved   int    `xml:"solved"`
	Level    int    `xml:"level"`
}

// Write writes profile SVG to writer.
func Write(writer io.Writer) error {
	profile := profile{
		UserName: "rkobayashi",
		Country:  "Japan",
		Language: "Haskell",
		Solved:   2,
		Level:    0,
	}

	t := newSVGTemplate()
	return t.Execute(writer, profile)
}

func newSVGTemplate() *template.Template {
	funcs := template.FuncMap{
		"circumference": circumference,
	}

	return template.Must(template.New("svg").Funcs(funcs).Parse(profileTemplate))
}

func circumference(radius int) float32 {
	return 2.0 * math.Pi * float32(radius)
}

const profileTemplate = `{{$rankRadius := 40}}
<svg xmlns="http://www.w3.org/2000/svg"
  width="495" height="195" viewBox="0,0,495,195">

  <style>
    .level-circle {
      animation: levelCircle ease-in-out 1s forwards;
    }

    @keyframes levelCircle {
      to {
        stroke-dashoffset: 243;
      }
      from {
        stroke-dashoffset: {{circumference $rankRadius}};
      }
    }
  </style>

  <text x="10" y="20" font-weight="bold" fill="#6b4e3d">
    {{.UserName}}'s Project Euler Profile
  </text>

  <text x="10" y="50" font-weight="bold">
    Solved
  </text>
  <text x="100" y="50">
    {{.Solved}}
  </text>

  <text x="10" y="80" font-weight="bold">
    Languege
  </text>
  <text x="100" y="80">
    {{.Language}}
  </text>

  <text x="10" y="110" font-weight="bold">
    Country
  </text>
  <text x="100" y="110">
    {{.Country}}
  </text>

  <text x="230" y="80" font-weight="bold" font-size="30" fill="#ff9933">
    {{.Level}}
  </text>
  <circle cx="240" cy="70" r="{{$rankRadius}}" stroke="#ff9933" stroke-width="8" fill="none"
    stroke-dasharray="{{circumference $rankRadius}}" class="level-circle" transform="rotate(-90,240,70)"
  />
  <circle cx="240" cy="70" r="{{$rankRadius}}" stroke="#ff9933" stroke-width="8" fill="none" opacity="0.2"/>
</svg>`
