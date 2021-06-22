package adexp

import (
	"encoding/json"
	"testing"
)


func TestNewGenerator(t *testing.T) {
	gen := NewGenerator()
	if gen == nil {
		t.Fatalf("%v is nil and should be a valid reference to a generator", gen)
	}
}

func TestGenerator_Next(t *testing.T) {
	gen := NewGenerator()
		jsonFPL := gen.Next()
		generated := &Fpl{}
		err := json.Unmarshal(jsonFPL, generated)
		if err != nil {
			t.Fatalf("%s is not parseable as FPL error is %s", jsonFPL, err)
		}

	if len(generated.Route) == 0 {
		t.Fatalf("%s has no route", jsonFPL)
	}

	if len(generated.Addr) == 0 {
		t.Fatalf("%s has no addr", jsonFPL)
	}
	if generated.Title != "IFPL" {
		t.Fatalf("%s title is not IFPL", jsonFPL)
	}
}

func TestGenerator_Next10(t *testing.T) {
	gen := NewGenerator()

	generated := make([]*Fpl, 10)
	for i := 0; i < 10; i++ {
		jsonFPL := gen.Next()
		generated[i] = &Fpl{}
		err := json.Unmarshal(jsonFPL, generated[i])
		if err != nil {
			t.Fatalf("%v is not parseable as FPL error is %s", jsonFPL, err)
		}
		if generated[i] == nil {
			t.Fatalf("%v is nil", generated[i])
		}

		if i > 0 {
			if !isDifferent(generated[i-1], generated[i]) {
				t.Fatalf("%+v is the same as next generated fpl %+v",generated[i-1], generated[i])
			}
		}
	}
}


func isDifferent(fpl1 *Fpl, fpl2 *Fpl) bool {
	if fpl1.IfplId == fpl2.IfplId {
		return false
	}

	if fpl1.Arcid == fpl2.Arcid {
		return false
	}
	return true
}