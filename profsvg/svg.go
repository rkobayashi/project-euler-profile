package profsvg

import (
	"html/template"
	"math"
)

const projectEulerLevelMax = 28

func newSVGTemplate() *template.Template {
	funcs := template.FuncMap{
		"circumference":  circumference,
		"arcByLevel":     arcByLevel,
		"levelPositionX": levelPositionX,
	}

	return template.Must(template.New("svg").Funcs(funcs).Parse(profileTemplate))
}

func circumference(radius int) float32 {
	return 2.0 * math.Pi * float32(radius)
}

func arcByLevel(radius, level int) float32 {
	if level < 0 {
		return circumference(radius)
	}
	if level >= projectEulerLevelMax {
		return 0.0
	}

	return (1 - float32(level+1)/(projectEulerLevelMax+1)) * circumference(radius)
}

func levelPositionX(level int) int {
	if level < 10 {
		return 230
	}

	return 220
}

const profileTemplate = `{{$rankRadius := 40}}
<svg xmlns="http://www.w3.org/2000/svg"
  width="350" height="150" viewBox="0,0,350,150">

  <style>
    .level-circle {
      animation: levelCircle ease-in-out 1s forwards;
    }

    @keyframes levelCircle {
      to {
        stroke-dashoffset: {{arcByLevel $rankRadius .Level}};
      }
      from {
        stroke-dashoffset: {{circumference $rankRadius}};
      }
    }
  </style>

  <rect width="100%" height="100%" fill="white" />

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

  <text x="{{levelPositionX .Level}}" y="80" font-weight="bold" font-size="30" fill="#ff9933">
    {{.Level}}
  </text>
  <circle cx="240" cy="70" r="{{$rankRadius}}" stroke="#ff9933" stroke-width="8" fill="none"
    stroke-dasharray="{{circumference $rankRadius}}" class="level-circle" transform="rotate(-90,240,70)"
  />
  <circle cx="240" cy="70" r="{{$rankRadius}}" stroke="#ff9933" stroke-width="8" fill="none" opacity="0.2"/>
</svg>`

const errSVG = `<svg xmlns="http://www.w3.org/2000/svg"
width="495" height="195" viewBox="0,0,495,195">
<text x="10" y="20" font-weight="bold" fill="#6b4e3d">
  Internal Server Error
</text>
</svg>`
