package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

// Get the request body `r` and parses it to JSON
// The response `w` is a PDF with the contents of the request body if no errors happen
func HandlePdfGen(w http.ResponseWriter, r *http.Request) {

	// Returns an error if the request is not of Method POST
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		fmt.Fprint(w, `{"error":"Only POST method allowed"}`)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"error":"Error reading request body"}`)
		fmt.Println(err)
		return
	}

	var jsonData map[string]interface{}

	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error":"Send valid JSON"}`)
		return
	}

	err = generatePdf(jsonData, w)
	if err != nil {
		fmt.Print(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Could not send the PDF"}`))
		fmt.Println(err)
	}
}

// Create a PDF through iterating over `content` and outputs it to `out`
// Returns errors if any
func generatePdf(content map[string]interface{}, out io.Writer) error {
	pdf := gofpdf.New("L", "mm", "A4", "")
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.AddPage()
	pdf.SetFont("Times", "", 20)

	formattedString := getFormattedString(content, 2, 0)

	pdf.Write(10, tr(formattedString))

	if pdf.Err() == true {
		return errors.New(fmt.Sprintf("Error during PDF creation: %s", pdf.Error()))
	}

	err := pdf.Output(out)

	if err != nil {
		return errors.New(fmt.Sprintf("Couldn't ouput the PDF to the writer: %s", err))
	}

	return nil
}

/*
Returns a formatted string with new lines and indentation (in the case of nested values).
Iterates over `data` recursively and for each possible nesting adds the `indent`.
`initialIndent` is the amount of characters all strings will start,
for example if `initialIndent` is 4 all the strings will start with 4 empty spaces

NOTE: If a SLICE contains another slice or a map, this map or slice will NOT be returned.
*/
func getFormattedString(data map[string]interface{}, indent, initialIndent int) string {
	var formattedString string
	indentSpaces := strings.Repeat(" ", initialIndent)

	for k, v := range data {
		_, isArray := v.([]interface{})
		_, isMap := v.(map[string]interface{})

		if isArray == true {
			filteredArr := getSliceWithoutNesting(v.([]interface{}))
			formattedString += fmt.Sprintf("%v%v = %v\n", indentSpaces, k, filteredArr)
		} else if isMap == true {
			formattedString += fmt.Sprintf("%v%v:\n", indentSpaces, k)
			str := getFormattedString(v.(map[string]interface{}), indent, initialIndent+indent)
			formattedString += str
		} else {
			formattedString += fmt.Sprintf("%v%v = %v\n", indentSpaces, k, v)
		}
	}
	return formattedString
}

// Returns a new slice that does not contain nesting of another slice or map
// For example [1, [1, 2]] returns [1]
func getSliceWithoutNesting(arr []interface{}) []interface{} {

	var filteredArr []interface{}

	for _, v := range arr {
		if fmt.Sprint((reflect.TypeOf(v).Kind())) == "map" || fmt.Sprint((reflect.TypeOf(v).Kind())) == "slice" {
			continue
		}
		filteredArr = append(filteredArr, v)
	}
	return filteredArr
}
