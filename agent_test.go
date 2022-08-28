package mooncake

import (
	"errors"
	"os"
	"os/exec"
	"reflect"
	"testing"
)

func TestAgent(t *testing.T) {
	t.Run("queue", func(t *testing.T) {
		t.Run("should return an empty agent", func(t *testing.T) {
			nAgent := NewAgent()
			expectedQueueSize := 0
			queueSize := len(nAgent.queue)
			if queueSize != expectedQueueSize {
				t.Errorf("invalid queue size. Expected=%q Got=%q", expectedQueueSize, queueSize)
			}
		})

		t.Run("should clear a not empty queue", func(t *testing.T) {
			nAgent := NewAgent()
			_ = nAgent.SetCall("test", reflect.TypeOf(func() {}))

			preTestQueueSize := len(nAgent.queue)
			preTestExpectedQueueSize := 1

			if preTestQueueSize != preTestExpectedQueueSize {
				t.Errorf("invalid queue size. Expected=%q Got=%q", preTestExpectedQueueSize, preTestQueueSize)
			}

			nAgent.CleanQueue()

			expectedQueueSize := 0
			queueSize := len(nAgent.queue)
			if queueSize != expectedQueueSize {
				t.Errorf("invalid queue size. Expected=%q Got=%q", expectedQueueSize, queueSize)
			}

		})
	})
	t.Run("call", func(t *testing.T) {
		t.Run("should register call and return valid agent controller", func(t *testing.T) {
			nAgent := NewAgent()
			cAgent := nAgent.SetCall("test", reflect.TypeOf(func() string { return "" }))

			expectedLifeTime := 1
			expectedLifeTimeCount := 1
			expectedNumberOfReturns := 1
			expectedReturnType := reflect.TypeOf("")

			if cAgent.lifeTime != LT_ONE_CALL {
				t.Errorf("invalid lifetime. Expected=%q Got=%q", expectedLifeTime, cAgent.lifeTime)

			}

			if cAgent.lifeTimeCount != expectedLifeTimeCount {
				t.Errorf("invalid lifetime counter. Expected=%q Got=%q", expectedLifeTimeCount, cAgent.lifeTimeCount)

			}

			if len(cAgent.returnValues) != expectedNumberOfReturns {
				t.Errorf("invalid number of returns. Expected=%q Got=%q", expectedNumberOfReturns, len(cAgent.returnValues))
			}

			if cAgent.returnValues[0].DType != expectedReturnType {
				t.Errorf("invalid type of return. Expected=%q Got=%q", expectedReturnType, cAgent.returnValues[0].DType)
			}
		})
		t.Run("should consume a call", func(t *testing.T) {
			nAgent := NewAgent()

			callName := "test"

			cAgent := nAgent.SetCall(callName, reflect.TypeOf(func() string { return "" }))
			cAgent.Repeat(2)

			if nAgent.queue[callName].lifeTimeCount != 2 {
				t.Errorf("invalid lifetime counter. Expected=%v Got=%v", 2, nAgent.queue[callName].lifeTimeCount)

			}

			_ = nAgent.GetCall(callName)

			if nAgent.queue[callName].lifeTimeCount != 1 {
				t.Errorf("invalid lifetime counter. Expected=%v Got=%v", 1, nAgent.queue[callName].lifeTimeCount)

			}

			_ = nAgent.GetCall(callName)

			if len(nAgent.queue) != 0 {
				t.Errorf("invalid number of itens on queue. Expected=%v Got=%v", 0, len(nAgent.queue))

			}

		})
		t.Run("should fatal when try to get invalid call", func(t *testing.T) {
			cmd := exec.Command(os.Args[0], "-test.run=TestGetInvalidCall")
			cmd.Env = append(os.Environ(), "SHOULD_FATAL=1")
			err := cmd.Run()
			if e, ok := err.(*exec.ExitError); ok && !e.Success() {
				return
			}
			t.Errorf("expected to return error. Expected=\"exit status 1\"")

		})
	})
}

func TestGetInvalidCall(t *testing.T) {
	if os.Getenv("SHOULD_FATAL") == "1" {
		t.Run("should fatal when try to get invalid call", func(t *testing.T) {
			nAgent := NewAgent()
			nAgent.GetCall("repeate")
		})
	}
}
func TestErrNoCallsRegistered(t *testing.T) {
	testCase := []struct {
		key           string
		expectedError error
	}{
		{
			key:           "test",
			expectedError: errors.New("no calls registered for test"),
		},
		{
			key:           "new",
			expectedError: errors.New("no calls registered for new"),
		},
	}

	for _, tc := range testCase {
		result := ErrNoCallsRegistered(tc.key)
		if result.Error() != tc.expectedError.Error() {
			t.Errorf("invalid function return. expected=%v got=%v", tc.expectedError.Error(), result.Error())
		}
	}
}
