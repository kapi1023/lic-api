package utils

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	jsonTag         = "json"
	validationError = `JSON nie został zwalidowany: Brak pola "%s"`
	tagRequired     = ",required"
)

var basicTypes = []string{"string", "int", "float64", "time.Time", "bool"}

func waliduj(doWalidacji interface{}) error {
	val1 := reflect.ValueOf(doWalidacji)
	if czyBazowyTyp(val1.Type().String()) {
		return nil
	}

	wartosc := val1.Elem()
	if wartosc.Kind() == reflect.Slice {
		for i := 0; i < wartosc.Len(); i++ {
			if err := waliduj(wartosc.Index(i).Addr().Interface()); err != nil {
				return err
			}
		}

		return nil
	}

	typ := wartosc.Type()

	for i := 0; i < wartosc.NumField(); i++ {
		poleWartosc := wartosc.Field(i)
		poleParams := typ.Field(i)

		if tag, ok := poleParams.Tag.Lookup(jsonTag); ok && strings.Contains(tag, tagRequired) && puste(poleWartosc) {
			return fmt.Errorf(validationError, poleParams.Name)
		}

		if poleWartosc.Kind() == reflect.Slice {
			for i := 0; i < poleWartosc.Len(); i++ {
				if err := waliduj(poleWartosc.Index(i).Addr().Interface()); err != nil {
					return err
				}
			}
		} else if poleWartosc.Kind() == reflect.Struct {
			if err := waliduj(poleWartosc.Addr().Interface()); err != nil {
				return err
			}
		}
	}

	return nil
}

func puste(wynik reflect.Value) bool {
	return wynik.IsZero() || ((wynik.Kind() == reflect.Slice || wynik.Kind() == reflect.Map) && wynik.Len() == 0)
}

func czyBazowyTyp(typ string) bool {
	for _, e := range basicTypes {
		if e == typ || "*"+e == typ || "[]"+e == typ {
			return true
		}
	}

	return false
}
