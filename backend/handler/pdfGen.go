package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

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

	for k, v := range content {
		pdf.Write(10, tr(
			fmt.Sprintf("%v: %+v\n", k, v),
		))
	}

	if pdf.Err() == true {
		return errors.New(fmt.Sprintf("Error during PDF creation: %s", pdf.Error()))
	}

	err := pdf.Output(out)

	if err != nil {
		return errors.New(fmt.Sprintf("Couldn't ouput the PDF to the writer: %s", err))
	}

	return nil
}
