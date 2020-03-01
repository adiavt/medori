package main

import (
	"encoding/json"
	"testing"
	"time"
)

type MedoriTimeTest struct {
	PointerTime *Time `json:"pointer_time"`
	NormalTime  Time  `json:"normal_time"`
}

func TestMedoriTimeType(t *testing.T) {
	t.Run("PointersDefaultShouldReturnNil", func(t *testing.T) {
		t.Helper()

		v := new(MedoriTimeTest)
		input := `{"normal_time":null}`

		if err := json.Unmarshal([]byte(input), &v); err != nil {
			t.Error("Error unmarshalling input ", err)
		}

		if v.PointerTime != nil {
			t.Error("Pointer default value should be nil")
		}
	})

	t.Run("EmptyStringShouldReturnNow", func(t *testing.T) {
		t.Helper()

		v := new(MedoriTimeTest)
		input := `{"normal_time":""}`

		if err := json.Unmarshal([]byte(input), &v); err != nil {
			t.Error("Error unmarshalling input ", err)
		}

		got := v.NormalTime.String()
		want := time.Now().Format("15:04:05")

		if got != want {
			t.Errorf("Expected '%s' but got '%s'", want, got)
		}
	})

	t.Run("ShouldMatchGivenTime", func(t *testing.T) {
		t.Helper()

		v := new(MedoriTimeTest)
		input := `{"pointer_time":"16:57:01", "normal_time":"16:57:01"}`

		if err := json.Unmarshal([]byte(input), &v); err != nil {
			t.Error("Error unmarshalling input ", err)
		}

		gotPointer := v.PointerTime.String()
		gotNormal := v.NormalTime.String()

		want := "16:57:01"

		if gotPointer != want {
			t.Errorf("Expected pointer value '%s' but got '%s'", want, gotPointer)
		}

		if gotNormal != want {
			t.Errorf("Expected normal value '%s' but got '%s'", want, gotNormal)
		}
	})

	t.Run("MarshallingJSON", func(t *testing.T) {
		t.Helper()

		now := Time{time.Now()}

		v := MedoriTimeTest{
			PointerTime: &now,
			NormalTime:  now,
		}

		_, err := json.Marshal(v)

		if err != nil {
			t.Error("Error marshalling input ", err)
		}
	})
}
