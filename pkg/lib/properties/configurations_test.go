package properties

import (
	"reflect"
	"testing"
)

func TestConfigurations(t *testing.T) {
	p := reflect.ValueOf(&configurations{})
	v := reflect.Indirect(p)

	for i := 0; i < v.Type().NumMethod(); i++ {
		methodSpec := v.Type().Method(i)

		t.Run(methodSpec.Name, func(t *testing.T) {
			results := v.MethodByName(methodSpec.Name).Call([]reflect.Value{})

			if results[0].IsNil() || !results[1].IsNil() {
				t.Errorf("Configurations.%s() returns (%v, %v)", methodSpec.Name, results[0].Interface(), results[1].Interface())
				return
			}
		})
	}
}
