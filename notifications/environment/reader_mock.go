// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package environment

import (
	"sync"
)

var (
	lockReaderMockRead sync.RWMutex
)

// ReaderMock is a mock implementation of Reader.
//
//     func TestSomethingThatUsesReader(t *testing.T) {
//
//         // make and configure a mocked Reader
//         mockedReader := &ReaderMock{
//             ReadFunc: func(in1 string) string {
// 	               panic("TODO: mock out the Read method")
//             },
//         }
//
//         // TODO: use mockedReader in code that requires Reader
//         //       and then make assertions.
//
//     }
type ReaderMock struct {
	// ReadFunc mocks the Read method.
	ReadFunc func(in1 string) string

	// calls tracks calls to the methods.
	calls struct {
		// Read holds details about calls to the Read method.
		Read []struct {
			// In1 is the in1 argument value.
			In1 string
		}
	}
}

// Read calls ReadFunc.
func (mock *ReaderMock) Read(in1 string) string {
	if mock.ReadFunc == nil {
		panic("ReaderMock.ReadFunc: method is nil but Reader.Read was just called")
	}
	callInfo := struct {
		In1 string
	}{
		In1: in1,
	}
	lockReaderMockRead.Lock()
	mock.calls.Read = append(mock.calls.Read, callInfo)
	lockReaderMockRead.Unlock()
	return mock.ReadFunc(in1)
}

// ReadCalls gets all the calls that were made to Read.
// Check the length with:
//     len(mockedReader.ReadCalls())
func (mock *ReaderMock) ReadCalls() []struct {
	In1 string
} {
	var calls []struct {
		In1 string
	}
	lockReaderMockRead.RLock()
	calls = mock.calls.Read
	lockReaderMockRead.RUnlock()
	return calls
}
