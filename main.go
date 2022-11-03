package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"log"
	"os"
	"text/template"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("unexpected number of args")
	}
	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	bv, err := toValidJSON(b)
	if err != nil {
		log.Fatal(err)
	}

	var reqData map[string]interface{}
	if err = json.Unmarshal(bv, &reqData); err != nil {
		log.Fatal(err)
	}

	out, err := execTemplate(reqData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

const (
	typeString = iota
	typeEnum
	typeInt
	typeFloat
	typeList
	typeMap
	typeUnknown

	typeTextString  = "schema.TypeString,"
	typeTextEnum    = "schema.TypeString, // TODO: enum, add validation"
	typeTextInt     = "schema.TypeInt,"
	typeTextFloat   = "schema.TypeFloat,"
	typeTextList    = "schema.TypeList,"
	typeTextMap     = "schema.TypeList,"
	typeTextUnknown = "schema.TypeString, // TODO: unable to identify type"
)

type schemaEntry struct {
	KeyName  string
	Type     int
	TypeText string
	Elems    map[string]schemaEntry
}

var schemaTemplate = `
{{ define "schemaItem" -}}
    "{{.KeyName}}": {
        Type: {{.TypeText}}
        {{- if eq .Type 5 }}
        MaxItems: 1,
        {{- end }}
        {{- if .Elems }}
        Elems: []map[string]*schema.Resource{
            {{- range .Elems }}
                {{ template "schemaItem" . }}
            {{- end }}
        },
        {{- end }}
    },
{{- end }}
package main

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func ResourceNewThing() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
            {{ range . -}}
            {{ template "schemaItem" . }}
            {{ end }}
        },
    }
}
`

func newSchemaEntry(keyName string, entry interface{}) schemaEntry {
	sa := schemaEntry{KeyName: keyName}
	if mapItem, ok := entry.(map[string]interface{}); ok {
		elems := make(map[string]schemaEntry)
		for k, v := range mapItem {
			elems[k] = newSchemaEntry(k, v)
		}
		sa.Type = typeMap
		sa.TypeText = typeTextMap
		sa.Elems = elems
		return sa
	}
	if listItem, ok := entry.([]interface{}); ok {
		elems := make(map[string]schemaEntry)
		// TODO: handle non-object (ex. array of strings) items with another assertion
		if mapItem, ok := listItem[0].(map[string]interface{}); ok {
			for k, v := range mapItem {
				elems[k] = newSchemaEntry(k, v)
			}
		}
		sa.Type = typeList
		sa.TypeText = typeTextList
		sa.Elems = elems
		return sa
	}
	if v, ok := entry.(string); ok {
		if v == "string" {
			sa.Type = typeString
			sa.TypeText = typeTextString
			return sa
		}
		if v == "enum" {
			sa.Type = typeEnum
			sa.TypeText = typeTextEnum
			return sa
		}
	}
	if _, ok := entry.(int); ok {
		sa.Type = typeInt
		sa.TypeText = typeTextInt
		return sa
	}
	if _, ok := entry.(float64); ok {
		sa.Type = typeFloat
		sa.TypeText = typeTextFloat
		return sa
	}
	return schemaEntry{KeyName: keyName, Type: typeUnknown, TypeText: typeTextUnknown}
}

func execTemplate(reqData map[string]interface{}) ([]byte, error) {
	entries := make(map[string]schemaEntry)
	for k, v := range reqData {
		entries[k] = newSchemaEntry(k, v)
	}

	tmpl, err := template.New("schema").Parse(schemaTemplate)
	if err != nil {
		return nil, err
	}

	b := &bytes.Buffer{}
	if err := tmpl.Execute(b, entries); err != nil {
		return nil, err
	}

	// gofmt the output so the template formatting doesn't need to be perfect
	return format.Source(b.Bytes())
}

// toValidJSON replaces placeholder documentation values with values that
// are valid JSON to allow for the request to be unmarshaled
func toValidJSON(b []byte) ([]byte, error) {
	bv := bytes.ReplaceAll(b, []byte("integer"), []byte("0"))
	bv = bytes.ReplaceAll(bv, []byte("enum"), []byte(`"enum"`))
	bv = bytes.ReplaceAll(bv, []byte("number"), []byte("0.0"))
	return bv, nil
}
