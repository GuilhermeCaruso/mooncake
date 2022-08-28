package mooncake

import (
	"errors"
	"os"
	"os/exec"
	"reflect"
	"testing"
)

func TestController(t *testing.T) {
	t.Run("agent_controller", func(t *testing.T) {
		t.Run("should generate an valid controller", func(t *testing.T) {
			nAgent := NewAgent()
			cAgent := nAgent.SetCall("test", reflect.TypeOf(func() string { return "" }))

			if cAgent.key != "test" {
				t.Errorf("invalid controller key. Expected=%q Got=%q", "test", cAgent.key)
			}
			if cAgent.lifeTime != LT_ONE_CALL {
				t.Errorf("invalid controller lifetime. Expected=%q Got=%q", LT_ONE_CALL, cAgent.lifeTime)
			}
			if cAgent.lifeTimeCount != 1 {
				t.Errorf("invalid controller lifetimeCounter. Expected=%q Got=%q", 1, cAgent.lifeTimeCount)
			}
			if len(cAgent.returnValues) != 1 {
				t.Errorf("invalid controller number of returns. Expected=%q Got=%q", 1, len(cAgent.returnValues))
			}
		})
		t.Run("should set the correct lifetime", func(t *testing.T) {
			nAgent := NewAgent()
			t.Run("AnyTime", func(t *testing.T) {
				cAgent := nAgent.SetCall("anyTime", reflect.TypeOf(func() string { return "" }))
				cAgent.AnyTime()

				if cAgent.lifeTime != LT_ANY_TIME {
					t.Errorf("invalid controller lifetime. Expected=%q Got=%q", LT_ANY_TIME, cAgent.lifeTime)
				}
			})
			t.Run("Repeat", func(t *testing.T) {
				cAgent := nAgent.SetCall("repeate", reflect.TypeOf(func() string { return "" }))
				cAgent.Repeat(10)

				if cAgent.lifeTime != LT_REPEAT {
					t.Errorf("invalid controller lifetime. Expected=%q Got=%q", LT_REPEAT, cAgent.lifeTime)
				}
				if cAgent.lifeTimeCount != 10 {
					t.Errorf("invalid controller lifetimeCounter. Expected=%q Got=%q", 10, cAgent.lifeTimeCount)
				}
			})

		})
		t.Run("SetReturn", func(t *testing.T) {
			t.Run("should fatal when the type of return is incorrect", func(t *testing.T) {
				cmd := exec.Command(os.Args[0], "-test.run=TestSetReturnWithFatalForIncorrectType")
				cmd.Env = append(os.Environ(), "SHOULD_FATAL=1")
				err := cmd.Run()
				if e, ok := err.(*exec.ExitError); ok && !e.Success() {
					return
				}
				t.Errorf("expected to return error. Expected=\"exit status 1\"")

			})

			t.Run("should fatal when the number of returns is incorrect", func(t *testing.T) {
				cmd := exec.Command(os.Args[0], "-test.run=TestSetReturnWithFatalForIncorrectNumberOfArgs")
				cmd.Env = append(os.Environ(), "SHOULD_FATAL=1")
				err := cmd.Run()
				if e, ok := err.(*exec.ExitError); ok && !e.Success() {
					return
				}
				t.Errorf("expected to return error. Expected=\"exit status 1\"")

			})

			t.Run("should define the correct type and number of return", func(t *testing.T) {
				nAgent := NewAgent()
				cAgent := nAgent.SetCall("repeate", reflect.TypeOf(func() string { return "" }))
				cAgent.SetReturn("return")

				if len(cAgent.returnValues) != 1 {
					t.Errorf("invalid number of returns. Expected=%q Got=%q", 1, len(cAgent.returnValues))

				}

				if cAgent.returnValues[0].Value != "return" {
					t.Errorf("invalid return value. Expected=%q Got=%q", "return", cAgent.returnValues[0].Value)

				}
			})

			t.Run("should define the correct number of repetitions", func(t *testing.T) {
				nAgent := NewAgent()
				cAgent := nAgent.SetCall("repeate", reflect.TypeOf(func() string { return "" }))
				cAgent.SetReturn("return")

				cAgent = nAgent.SetCall("repeate", reflect.TypeOf(func() string { return "" }))
				cAgent.SetReturn("return")

				if len(cAgent.returnValues) != 1 {
					t.Errorf("invalid number of returns. Expected=%q Got=%q", 1, len(cAgent.returnValues))

				}

				if cAgent.returnValues[0].Value != "return" {
					t.Errorf("invalid return value. Expected=%q Got=%q", "return", cAgent.returnValues[0].Value)

				}
			})
		})
	})
}

func TestSetReturnWithFatalForIncorrectType(t *testing.T) {
	if os.Getenv("SHOULD_FATAL") == "1" {
		t.Run("should fatal when the type of return is incorrect", func(t *testing.T) {
			nAgent := NewAgent()
			cAgent := nAgent.SetCall("repeate", reflect.TypeOf(func() string { return "" }))
			cAgent.SetReturn(1)
		})
	}
}

func TestSetReturnWithFatalForIncorrectNumberOfArgs(t *testing.T) {
	if os.Getenv("SHOULD_FATAL") == "1" {
		t.Run("should fatal when the number of returns is incorrect", func(t *testing.T) {
			nAgent := NewAgent()
			cAgent := nAgent.SetCall("repeate", reflect.TypeOf(func() string { return "" }))
			cAgent.SetReturn("1", "2")
		})
	}
}

func TestErrInvalidNumberOfReturns(t *testing.T) {
	testCase := []struct {
		key           string
		expected      int
		got           int
		expectedError error
	}{
		{
			key:           "test",
			expected:      1,
			got:           2,
			expectedError: errors.New("invalid number of returns for test. expected=1 got=2"),
		},
		{
			key:           "new",
			expected:      10,
			got:           2,
			expectedError: errors.New("invalid number of returns for new. expected=10 got=2"),
		},
	}

	for _, tc := range testCase {
		result := ErrInvalidNumberOfReturns(tc.key, tc.expected, tc.got)
		if result.Error() != tc.expectedError.Error() {
			t.Errorf("invalid function return. expected=%s got=%s", tc.expectedError.Error(), result.Error())
		}
	}
}

func TestErrInvalidTypeOfReturn(t *testing.T) {
	testCase := []struct {
		key           string
		expected      interface{}
		got           interface{}
		expectedError error
	}{
		{
			key:           "test",
			expected:      1,
			got:           2,
			expectedError: errors.New("invalid type of return for test. expected=1 got=2"),
		},
		{
			key:           "new",
			expected:      10,
			got:           2,
			expectedError: errors.New("invalid type of return for new. expected=10 got=2"),
		},
	}

	for _, tc := range testCase {
		result := ErrInvalidTypeOfReturn(tc.key, tc.expected, tc.got)
		if result.Error() != tc.expectedError.Error() {
			t.Errorf("invalid function return. expected=%v got=%v", tc.expectedError.Error(), result.Error())
		}
	}
}
