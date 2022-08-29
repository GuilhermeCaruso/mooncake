package example

import (
	"testing"

	"github.com/GuilhermeCaruso/mooncake"
)

func checkValue(t *testing.T, es GenericInterface[string, string], expectedResult, secondeExpectedResult string) {
	v, z := es.Other(expectedResult)
	if v != expectedResult {
		t.Errorf("unexpected result. expected=%v got=%v", expectedResult, v)
	}
	if z != secondeExpectedResult {
		t.Errorf("unexpected result. expected=%v got=%v", secondeExpectedResult, v)
	}

}
func TestWithoutMock(t *testing.T) {
	keyValue := "fixed"
	internalKeyValue := "test"
	es := GenericExample[string, string]{
		Value: internalKeyValue,
	}
	checkValue(t, es, keyValue, internalKeyValue)
}

func TestWithMock(t *testing.T) {
	a := mooncake.NewAgent()
	keyValue := "fixed_mock"
	internalKeyValue := "test_mock"

	ac := NewMockGenericInterface[string, string](a)
	ac.Prepare().Other(string(mooncake.ANY)).SetReturn(keyValue, internalKeyValue)

	checkValue(t, ac, keyValue, internalKeyValue)

}
