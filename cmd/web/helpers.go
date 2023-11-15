package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-playground/form/v4"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

func formatVietnamesePrice(price int32) string {
	priceStr := fmt.Sprintf("%d", price)
	formattedPrice := ""

	length := len(priceStr)
	remainder := length % 3

	if remainder > 0 {
		formattedPrice += priceStr[:remainder] + "."
	}

	for i := remainder; i < length; i += 3 {
		formattedPrice += priceStr[i:i+3] + "."
	}

	// Remove the trailing dot and add the currency symbol if needed
	formattedPrice = strings.TrimRight(formattedPrice, ".")

	return formattedPrice
}

func formatVietnameseDate(t time.Time) string {
	return t.Format("15:04 02/01/2006")
}

func plusOne(x, y int) int {
	return x + y
}

func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template set %s does not exist", page)
		app.serverError(w, err)
		return
	}

	// TODO: Write other headers.

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, page, data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}

func (app *application) decodePostForm(r *http.Request, dst any) error {
	err := r.ParseForm()
	if err != nil {
		app.errorLog.Print(err)
		return err
	}

	err = app.formDecoder.Decode(dst, r.PostForm)
	if err != nil {
		var invalidDecoderError *form.InvalidDecoderError

		if errors.As(err, &invalidDecoderError) {
			panic(err)
		}

		return err
	}

	return nil
}

func (app *application) decodeMultipartForm(r *http.Request, dst any) error {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		app.errorLog.Print(err)
		return err
	}

	err = app.formDecoder.Decode(dst, r.PostForm)
	if err != nil {
		var invalidDecoderError *form.InvalidDecoderError

		if errors.As(err, &invalidDecoderError) {
			panic(err)
		}

		return err
	}

	return nil
}

func (app *application) saveFileToDisk(file multipart.File, header *multipart.FileHeader) error {
	dst, err := os.Create(fmt.Sprintf("./ui/static/img/%s", header.Filename))
	if err != nil {
		app.errorLog.Println(err)
		return err
	}

	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}

	return nil
}

func (app *application) isAuthenticated(r *http.Request) bool {
	isAuthenticated, ok := r.Context().Value(isAuthenticatedContextKey).(bool)
	if !ok {
		return false
	}

	return isAuthenticated
}

func (app *application) isProvider(r *http.Request) bool {
	isProvider, ok := r.Context().Value(isProviderContextKey).(bool)
	if !ok {
		return false
	}

	return isProvider
}

func (app *application) isAdmin(r *http.Request) bool {
	isAdmin, ok := r.Context().Value(isAdminContextKey).(bool)
	if !ok {
		return false
	}

	return isAdmin
}
