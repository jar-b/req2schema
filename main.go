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
		log.Fatal(fmt.Errorf("reading input: %w", err))
	}

	out, err := requestToSchema(b)
	if err != nil {
		log.Fatal(fmt.Errorf("schema generation: %w", err))
	}
	fmt.Println(string(out))
}

const (
	typeEnum = iota
	typeFloat
	typeInt
	typeListOfMaps
	typeListOfType
	typeMap
	typeString
	typeUnknown

	typeTextEnum       = "schema.TypeString, // TODO: enum, add validation"
	typeTextFloat      = "schema.TypeFloat,"
	typeTextInt        = "schema.TypeInt,"
	typeTextListOfMaps = "schema.TypeList,"
	typeTextListOfType = "schema.TypeList,"
	typeTextMap        = "schema.TypeList,"
	typeTextString     = "schema.TypeString,"
	typeTextUnknown    = "schema.TypeString, // TODO: unable to identify type"
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
        Elems: &schema.Resource{
			Schema: map[string]*schema.Schema{
            {{- range .Elems }}
                {{ template "schemaItem" . }}
            {{- end }}
			},
        },
        {{- end }}
        {{- if eq .Type 4 }}
        Elems: &schema.Schema{Type: schema.TypeString},
        {{- end }}
    },
{{- end }}
package someservice

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

func ResourceSomeThing() *schema.Resource {
    return &schema.Resource{
        Schema: map[string]*schema.Schema{
            {{ range . -}}
            {{ template "schemaItem" . }}
            {{ end -}}
			"tags":     tftags.TagsSchema(),
			"tags_all": tftags.TagsSchemaComputed(),
        },
    }
}
`

func execTemplate(reqData map[string]interface{}) ([]byte, error) {
	entries := make(map[string]schemaEntry)
	for k, v := range reqData {
		if skipKey(k) {
			continue
		}
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
		if mapItem, ok := listItem[0].(map[string]interface{}); ok {
			elems := make(map[string]schemaEntry)
			for k, v := range mapItem {
				elems[k] = newSchemaEntry(k, v)
			}
			sa.Type = typeListOfMaps
			sa.TypeText = typeTextListOfMaps
			sa.Elems = elems
		} else {
			// TODO: more granular handling of assertions on lists of concrete types (string, int, more?)
			sa.Type = typeListOfType
			sa.TypeText = typeTextListOfType
		}
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

func requestToSchema(b []byte) ([]byte, error) {
	bv, err := toValidJSON(b)
	if err != nil {
		return nil, fmt.Errorf("cleaning request: %w", err)
	}

	var reqData map[string]interface{}
	if err = json.Unmarshal(bv, &reqData); err != nil {
		return nil, fmt.Errorf("unmarshaling request: %w", err)
	}

	out, err := execTemplate(reqData)
	if err != nil {
		return nil, fmt.Errorf("executing template: %w", err)
	}
	return out, nil
}

func skipKey(k string) bool {
	// keysToSkip is a list of keys to skip adding to the terraform schema
	//
	// requestId - this is specific to AWS HTTP requests and has no context in terraform
	// tags - the "tags" and "tags_all" attibutes are accounted for in the base template
	keysToSkip := []string{"requestId", "tags"}
	for _, sk := range keysToSkip {
		if k == sk {
			return true
		}
	}
	return false
}

// toValidJSON replaces placeholder documentation values with values that
// are valid JSON to allow for the request to be unmarshaled
func toValidJSON(b []byte) ([]byte, error) {
	bv := bytes.ReplaceAll(b, []byte("integer"), []byte("0"))
	bv = bytes.ReplaceAll(bv, []byte("enum"), []byte(`"enum"`))
	bv = bytes.ReplaceAll(bv, []byte("number"), []byte("0.0"))
	return bv, nil
}
