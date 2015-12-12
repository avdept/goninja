package goninja

import "testing"
import "reflect"

func TestNewRouter(t *testing.T) {
	if reflect.ValueOf(NewRouter()) != new(Router) {
		t.Error("New Router doesnt match empty Router struct")
	}

}

