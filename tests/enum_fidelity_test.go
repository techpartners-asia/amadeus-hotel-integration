package tests

import (
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/techpartners-asia/amadeus-hotel-integration/constants"
	amadeusIntegration "github.com/techpartners-asia/amadeus-hotel-integration/integrations/amadeus"
	responseContentDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/content/dto/response"
	requestHotelListCityDTO "github.com/techpartners-asia/amadeus-hotel-integration/modules/list/dto/request/city"
)

// The response DTOs declare hundreds of enum constants as named string types
// (type Segment string; const SegmentLuxury Segment = "LUXURY"). Because any
// JSON string decodes into a named string type, a wrong or incomplete constant
// list is invisible to both encoding/json and the shape checks in
// dto_fidelity_test.go. This test parses the declared constants out of the
// source and compares them against the values Amadeus actually sends, so a
// value the SDK has no constant for is reported rather than silently accepted.

// declaredEnums parses a package directory and returns namedType -> set of the
// string constant values declared for it.
func declaredEnums(t *testing.T, dir string) map[string]map[string]bool {
	t.Helper()
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, 0)
	if err != nil {
		t.Fatalf("parse %s: %v", dir, err)
	}

	enums := map[string]map[string]bool{}
	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				gen, ok := decl.(*ast.GenDecl)
				if !ok || gen.Tok != token.CONST {
					continue
				}
				var currentType string
				for _, spec := range gen.Specs {
					vs, ok := spec.(*ast.ValueSpec)
					if !ok {
						continue
					}
					if id, ok := vs.Type.(*ast.Ident); ok {
						currentType = id.Name
					}
					if currentType == "" {
						continue
					}
					for _, v := range vs.Values {
						lit, ok := v.(*ast.BasicLit)
						if !ok || lit.Kind != token.STRING {
							continue
						}
						s, err := strconv.Unquote(lit.Value)
						if err != nil {
							continue
						}
						if enums[currentType] == nil {
							enums[currentType] = map[string]bool{}
						}
						enums[currentType][s] = true
					}
				}
			}
		}
	}
	return enums
}

// enumFieldPaths walks a Go type and records json path -> enum type name for
// every field whose type is a named string type with declared constants.
func enumFieldPaths(t reflect.Type, prefix string, seen map[reflect.Type]bool, enums map[string]map[string]bool, out map[string]string) {
	for t.Kind() == reflect.Ptr || t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return
	}
	if seen[t] && prefix != "" {
		return
	}
	seen[t] = true
	defer delete(seen, t)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get("json")
		if tag == "-" {
			continue
		}
		name := strings.Split(tag, ",")[0]
		if name == "" {
			name = f.Name
		}
		p := name
		if prefix != "" {
			p = prefix + "." + name
		}

		ft := f.Type
		for ft.Kind() == reflect.Ptr || ft.Kind() == reflect.Slice || ft.Kind() == reflect.Array {
			ft = ft.Elem()
		}
		if ft.Kind() == reflect.String && ft.Name() != "" && ft.Name() != "string" {
			if _, ok := enums[ft.Name()]; ok {
				out[p] = ft.Name()
			}
		}
		enumFieldPaths(f.Type, p, seen, enums, out)
	}
}

// collectStrings records every string value observed at each json path.
func collectStrings(v any, prefix string, out map[string]map[string]bool) {
	switch x := v.(type) {
	case map[string]any:
		for k, val := range x {
			p := k
			if prefix != "" {
				p = prefix + "." + k
			}
			if s, ok := val.(string); ok {
				if out[p] == nil {
					out[p] = map[string]bool{}
				}
				out[p][s] = true
			}
			collectStrings(val, p, out)
		}
	case []any:
		for _, item := range x {
			collectStrings(item, prefix, out)
		}
	}
}

// isOpenVocabularyValue reports values that legitimately have no constant
// because the field is not a closed enum. CategoryCode carries the IATA airport
// code itself for airport points of interest (CDG, ORY, ...), which cannot be
// enumerated ahead of time.
func isOpenVocabularyValue(enumType, v string) bool {
	if enumType != "CategoryCode" {
		return false
	}
	if len(v) != 3 {
		return false
	}
	for _, r := range v {
		if r < 'A' || r > 'Z' {
			return false
		}
	}
	return true
}

// TestContentEnumsCoverAPIValues checks that every value Amadeus returns for an
// enum-typed content field has a matching declared constant.
func TestContentEnumsCoverAPIValues(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live API test in -short mode")
	}
	s := newSDK(t)

	enums := declaredEnums(t, filepath.Join("..", "modules", "content", "dto", "response"))
	if len(enums) == 0 {
		t.Fatal("no enum constants parsed")
	}

	fields := map[string]string{}
	enumFieldPaths(reflect.TypeOf(responseContentDTO.HotelContentResponse{}), "",
		map[reflect.Type]bool{}, enums, fields)
	t.Logf("parsed %d enum types, %d enum-typed fields", len(enums), len(fields))

	hotels, err := s.List.HotelListByCityCode(requestHotelListCityDTO.HotelListByCityCodeRequest{CityCode: "PAR"})
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	client := amadeusIntegration.NewClient(constants.CONTENT_BASE_URL)

	n := 25
	if len(hotels) < n {
		n = len(hotels)
	}
	values := map[string]map[string]bool{}
	for _, h := range hotels[:n] {
		res, err := client.R().SetQueryParams(map[string]string{
			"hotelID": h.HotelId, "view": "FULL",
		}).Get("/reference-data/locations/by-hotel")
		if err != nil || res.StatusCode() != 200 {
			continue
		}
		var top map[string]any
		if err := json.Unmarshal([]byte(res.String()), &top); err != nil {
			continue
		}
		if d, ok := top["data"]; ok {
			collectStrings(d, "", values)
		}
	}

	var unknown []string
	for path, enumType := range fields {
		observed, ok := values[path]
		if !ok {
			continue
		}
		for v := range observed {
			if enums[enumType][v] || isOpenVocabularyValue(enumType, v) {
				continue
			}
			unknown = append(unknown, path+" ("+enumType+"): "+strconv.Quote(v))
		}
	}
	sort.Strings(unknown)

	for _, u := range unknown {
		t.Errorf("value returned by API has no declared constant: %s", u)
	}
	t.Logf("checked %d enum-typed fields against live values", len(fields))
}
