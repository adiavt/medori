package main

import (
	"encoding/json"
	"testing"
	"time"
)

type MedoriDateTest struct {
	PointerDate *Date `json:"pointer_date"`
	NormalDate  Date  `json:"normal_date"`
}

func TestMedoriDateType(t *testing.T) {
	t.Run("PointersDefaultShouldReturnNil", func(t *testing.T) {
		t.Helper()

		v := new(MedoriDateTest)
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

		v := new(MedoriDateTest)
		input := `{"normal_date":""}`

		if err := json.Unmarshal([]byte(input), &v); err != nil {
			t.Error("Error unmarshalling input ", err)
		}

		got := v.NormalDate.String()
		want := time.Now().Format("2006-01-02")

		if got != want {
			t.Errorf("Expected '%s' but got '%s'", want, got)
		}
	})

	t.Run("ShouldMatchGivenDate", func(t *testing.T) {
		t.Helper()

		v := new(MedoriDateTest)
		input := `{"pointer_date":"2020-03-01", "normal_date":"2020-03-01"}`

		if err := json.Unmarshal([]byte(input), &v); err != nil {
			t.Error("Error unmarshalling input ", err)
		}

		gotPointer := v.PointerDate.String()
		gotNormal := v.NormalDate.String()

		want := "2020-03-01"

		if gotPointer != want {
			t.Errorf("Expected pointer value '%s' but got '%s'", want, gotPointer)
		}

		if gotNormal != want {
			t.Errorf("Expected normal value '%s' but got '%s'", want, gotNormal)
		}
	})

	t.Run("UnmarshalDashedDMY", func(t *testing.T) {
		t.Helper()

		v := new(MedoriDateTest)
		input := `{"pointer_date":"01-03-2020", "normal_date":"01-03-2020"}`

		if err := json.Unmarshal([]byte(input), &v); err != nil {
			t.Error("Error unmarshalling input slashed DMY ", err)
		}
	})

	t.Run("UnmarshalSlashedYMD", func(t *testing.T) {
		t.Helper()

		v := new(MedoriDateTest)
		input := `{"pointer_date":"2020/03/01", "normal_date":"2020/03/01"}`

		if err := json.Unmarshal([]byte(input), &v); err != nil {
			t.Error("Error unmarshalling input slashed YMD ", err)
		}
	})

	t.Run("UnmarshalSlashedDMY", func(t *testing.T) {
		t.Helper()

		v := new(MedoriDateTest)
		input := `{"pointer_date":"01/03/2020", "normal_date":"01/03/2020"}`

		if err := json.Unmarshal([]byte(input), &v); err != nil {
			t.Error("Error unmarshalling input slashed DMY", err)
		}
	})

	t.Run("MarshallingJSON", func(t *testing.T) {
		t.Helper()

		now := Date{time.Now()}

		v := MedoriDateTest{
			PointerDate: &now,
			NormalDate:  now,
		}

		_, err := json.Marshal(v)

		if err != nil {
			t.Error("Error marshalling input ", err)
		}
	})
}
