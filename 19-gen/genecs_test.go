package gen_test

import (
	gen "genecs"
	"testing"
)

func TestAssertFunctions(t *testing.T) {
	t.Run("on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
	t.Run("on strings", func(t *testing.T) {
		AssertEqual(t, "1", "1")
		AssertNotEqual(t, "1", "2")
	})
	t.Run("on boolean", func(t *testing.T) {
		AssertTrue(t, true)
		AssertFalse(t, false)
	})
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}

func AssertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v but want %+v", got, want)
	}
}

func AssertNotEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got == want {
		t.Errorf("got %+v equal to %+v", got, want)
	}
}

func TestStackInt(t *testing.T) {
	t.Run("init", func(t *testing.T) {
		stack := new(gen.Stack[int])
		if stack.Length() != 0 {
			t.Errorf("new stack length must be zero")
		}
	})
	t.Run("push", func(t *testing.T) {
		stack := new(gen.Stack[int])
		stack.Push(1)
		if stack.Length() != 1 {
			t.Errorf("expected len 1")
		}
		stack.Push(2)
		if stack.Length() != 2 {
			t.Errorf("expected len 2")
		}
	})
	t.Run("pop", func(t *testing.T) {
		stack := new(gen.Stack[int])
		stack.Push(1)
		stack.Push(2)
		item, err := stack.Pop()
		if err == false {
			t.Error("unable to pop")
		}
		if item != 2 {
			t.Error("expected pop is 2")
		}
		item, err = stack.Pop()
		if err == false {
			t.Error("unable to pop")
		}
		if item != 1 {
			t.Error("expected pop is 1")
		}
	})
	t.Run("pop failure", func(t *testing.T) {
		stack := new(gen.Stack[int])
		_, err := stack.Pop()
		if err == true {
			t.Error("expected failure")
		}
	})
	t.Run("type test", func(t *testing.T) {
		stack := new(gen.Stack[int])
		stack.Push(1)
		stack.Push(2)
		a, _ := stack.Pop()
		b, _ := stack.Pop()
		if a+b != 3 {
			t.Error("expected 3")
		}
	})
}

func TestStackString(t *testing.T) {
	t.Run("init", func(t *testing.T) {
		stack := new(gen.Stack[string])
		if stack.Length() != 0 {
			t.Errorf("new stack length must be zero")
		}
	})
	t.Run("push", func(t *testing.T) {
		stack := new(gen.Stack[string])
		stack.Push("test")
		if stack.Length() != 1 {
			t.Errorf("expected len 1")
		}
	})
	t.Run("pop", func(t *testing.T) {
		stack := new(gen.Stack[string])
		stack.Push("test")
		item, err := stack.Pop()
		if err == false {
			t.Error("unable to pop")
		}
		if item != "test" {
			t.Error("expected pop is test")
		}
	})
	t.Run("pop failure", func(t *testing.T) {
		stack := new(gen.Stack[string])
		_, err := stack.Pop()
		if err == true {
			t.Error("expected failure")
		}
	})
	t.Run("type test", func(t *testing.T) {
		stack := new(gen.Stack[string])
		stack.Push("ab")
		stack.Push("cd")
		a, _ := stack.Pop()
		b, _ := stack.Pop()
		if a+b != "cdab" {
			t.Error("expected cdab")
		}
	})
}
