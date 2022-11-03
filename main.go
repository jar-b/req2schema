package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	fmt.Println("parsing request schema")

	b, err := os.ReadFile("req.json")
	if err != nil {
		log.Fatal(err)
	}

	bv, err := toValidJSON(b)
	if err != nil {
		log.Fatal(err)
	}

	var in map[string]interface{}
	if err = json.Unmarshal(bv, &in); err != nil {
		log.Fatal(err)
	}

	if err := populateTemplate(in); err != nil {
		log.Fatal(err)
	}
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
    Schema: map[string]*schema.Schema{
    	{{ range $key, $attr := . -}}
		"{{$key}}": {
    		Type: {{$attr.TypeText}}
    		{{if eq $attr.Type 5}}MaxItems: 1,{{end}}
    	},
    	{{ end }}
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

func populateTemplate(in map[string]interface{}) error {
	entries := make(map[string]schemaEntry)
	for k, v := range in {
		entries[k] = newSchemaEntry(k, v)
	}
	tmp, _ := json.MarshalIndent(entries, "", "  ")
	fmt.Printf("writing processed object to: %s\n", "processed.json")
	os.WriteFile("processed.json", tmp, 0644)

	tmpl, err := template.New("schema").Parse(schemaTemplate)
	if err != nil {
		return err
	}

	return tmpl.Execute(os.Stdout, entries)
}

// toValidJSON replaces placeholder documentation values with values that
// are valid JSON to allow for the request to be unmarshaled
func toValidJSON(b []byte) ([]byte, error) {
	bv := bytes.ReplaceAll(b, []byte("integer"), []byte("0"))
	bv = bytes.ReplaceAll(bv, []byte("enum"), []byte(`"enum"`))
	bv = bytes.ReplaceAll(bv, []byte("number"), []byte("0.0"))
	return bv, nil
}
