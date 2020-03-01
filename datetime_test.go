package main

import (
	"encoding/json"
	"testing"
	"time"
)

type MedoriDatetimeTest struct {
	PointerDate *Datetime `json:"pointer_date"`
	NormalDate  Datetime  `json:"normal_date"`
}

func TestMedoriDatetimeType(t *testing.T) {
	t.Run("PointersDefaultShouldReturnNil", func(t *testing.T) {
		t.Helper()

		v := new(MedoriDatetimeTest)
		input := `{"normal_date":null}`

		if err := json.Unmarshal([]byte(input), &v); err != nil {
			t.Error("Error unmarshalling input ", err)
		}

		if v.PointerDate != nil {
			t.Error("Pointer default value should be nil")
		}
	})

	t.Run("EmptyStringShouldReturnNow", func(t *testing.T) {
		t.Helper()

		v := new(MedoriDatetimeTest)
		input := `{"normal_date":""}`

		if err := json.Unmarshal([]byte(input), &v); err != nil {
			t.Error("Error unmarshalling input ", err)
		}

		got := v.NormalDate.String()
		want := time.Now().Format("2006-01-02 15:04:05")

		if got != want {
			t.Errorf("Expected '%s' but got '%s'", want, got)
		}
	})

	t.Run("ShouldMatchGivenDatetime", func(t *testing.T) {
		t.Helper()

		v := new(MedoriDatetimeTest)
		input := `{"pointer_date":"2020-03-01 16:57:01", "normal_date":"2020-03-01 16:57:01"}`

		if err := json.Unmarshal([]byte(input), &v); err != nil {
			t.Error("Error unmarshalling input ", err)
		}

		gotPointer := v.PointerDate.String()
		gotNormal := v.NormalDate.String()

		want := "2020-03-01 16:57:01"

		if gotPointer != want {
			t.Errorf("Expected pointer value '%s' but got '%s'", want, gotPointer)
		}

		if gotNormal != want {
			t.Errorf("Expected normal value '%s' but got '%s'", want, gotNormal)
		}
	})

	t.Run("UnmarshalDashedDMYHis", func(t *testing.T) {
		t.Helper()

		v := new(MedoriDatetimeTest)
		input := `{"pointer_date":"01-03-2020 16:57:01", "normal_date":"01-03-2020 16:57:01"}`

		if err := json.Unmarshal([]byte(input), &v); err != nil {
			t.Error("Error unmarshalling input slashed DMY ", err)
		}
	})

	t.Run("UnmarshalSlashedYMDHis", func(t *testing.T) {
		t.Helper()

		v := new(MedoriDatetimeTest)
		input := `{"pointer_date":"2020/03/01 16:57:01", "normal_date":"2020/03/01 16:57:01"}`

		if err := json.Unmarshal([]byte(input), &v); err != nil {
			t.Error("Error unmarshalling input slashed YMD ", err)
		}
	})

	t.Run("UnmarshalSlashedDMYHis", func(t *testing.T) {
		t.Helper()

		v := new(MedoriDatetimeTest)
		input := `{"pointer_date":"01/03/2020 16:57:01", "normal_date":"01/03/2020 16:57:01"}`

		if err := json.Unmarshal([]byte(input), &v); err != nil {
			t.Error("Error unmarshalling input slashed DMY", err)
		}
	})

	t.Run("MarshallingJSON", func(t *testing.T) {
		t.Helper()

		now := Datetime{time.Now()}

		v := MedoriDatetimeTest{
			PointerDate: &now,
			NormalDate:  now,
		}

		_, err := json.Marshal(v)

		if err != nil {
			t.Error("Error marshalling input ", err)
		}
	})
}
