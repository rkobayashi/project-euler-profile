package profsvg

import (
	"fmt"
	"io"
	"net/url"
)

// Write writes profile SVG to writer.
func Write(writer io.Writer, query url.Values) error {
	name, err := parseUserNameFromQuery(query)
	if err != nil {
		fmt.Fprint(writer, errSVG)
		return err
	}

	profile, err := getProfileFromServer(name)
	if err != nil {
		fmt.Fprint(writer, errSVG)
		return err
	}

	t := newSVGTemplate()
	return t.Execute(writer, profile)
}

func parseUserNameFromQuery(q url.Values) (string, error) {
	if len(q["username"]) == 1 {
		return q["username"][0], nil
	}

	return "", fmt.Errorf("cannot find username. query=%v", q)
}
