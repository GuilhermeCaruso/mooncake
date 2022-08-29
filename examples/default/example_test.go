package example

import (
	"testing"

	"github.com/GuilhermeCaruso/mooncake"
)

func checkValue(t *testing.T, es SimpleInterface, expectedResult string) {
	v, err := es.Get()
	if v != expectedResult {
		t.Errorf("unexpected result. expected=%v got=%v", expectedResult, v)
	}

	if err != nil {
		t.Errorf("unexpected error. expected=<nil> got=%v", err.Error())
	}
}
func TestWithoutMock(t *testing.T) {
	keyValue := "fixed"
	es := ExampleStructure{}
	checkValue(t, es, keyValue)
}

func TestWithMock(t *testing.T) {
	a := mooncake.NewAgent()

	ac := NewMockSimpleInterface(a)
	ac.Prepare().Get().SetReturn("mocked_value", nil)

	checkValue(t, ac, "mocked_value")

}
