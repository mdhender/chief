package docconv

// docconv is licensed per https://github.com/sajari/docconv/blob/master/LICENSE

import (
	"encoding/xml"
	"io"
)

// XMLToText converts XML to plain text given how to treat elements.
func XMLToText(r io.Reader, breaks []string, skip []string, strict bool) (string, error) {
	var result string

	dec := xml.NewDecoder(io.LimitReader(r, maxBytes))
	dec.Strict = strict
	for {
		t, err := dec.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}

		switch v := t.(type) {
		case xml.CharData:
			result += string(v)
		case xml.StartElement:
			for _, breakElement := range breaks {
				if v.Name.Local == breakElement {
					result += "\n"
				}
			}
			for _, skipElement := range skip {
				if v.Name.Local == skipElement {
					depth := 1
					for {
						t, err := dec.Token()
						if err != nil {
							// An io.EOF here is actually an error.
							return "", err
						}

						switch t.(type) {
						case xml.StartElement:
							depth++
						case xml.EndElement:
							depth--
						}

						if depth == 0 {
							break
						}
					}
				}
			}
		}
	}
	return result, nil
}

// XMLToMap converts XML to a nested string map.
func XMLToMap(r io.Reader) (map[string]string, error) {
	m := make(map[string]string)
	dec := xml.NewDecoder(io.LimitReader(r, maxBytes))
	var tagName string
	for {
		t, err := dec.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		switch v := t.(type) {
		case xml.StartElement:
			tagName = string(v.Name.Local)
		case xml.CharData:
			m[tagName] = string(v)
		}
	}
	return m, nil
}
