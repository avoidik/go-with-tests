package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name     string
		Input    interface{}
		Expected []string
	}{
		{
			Name: "single string",
			Input: struct {
				Name string
			}{"Chris"},
			Expected: []string{"Chris"},
		},
		{
			Name: "two string elements",
			Input: struct {
				Name string
				City string
			}{"Chris", "London"},
			Expected: []string{"Chris", "London"},
		},
		{
			Name: "non-string elements",
			Input: struct {
				Name string
				Age  int
			}{"Chris", 10},
			Expected: []string{"Chris"},
		},
		{
			Name:     "nested struct",
			Input:    Person{"Chris", Profile{10, "London"}},
			Expected: []string{"Chris", "London"},
		},
		{
			Name:     "pointer nested struct",
			Input:    &Person{"Chris", Profile{10, "London"}},
			Expected: []string{"Chris", "London"},
		},
		{
			Name: "slices struct",
			Input: []Profile{
				{11, "Moscow"},
				{46, "Chicago"},
			},
			Expected: []string{"Moscow", "Chicago"},
		},
		{
			Name: "arrays struct",
			Input: [2]Profile{
				{11, "Moscow"},
				{46, "Chicago"},
			},
			Expected: []string{"Moscow", "Chicago"},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.Name, func(t *testing.T) {
			var got []string
			walk(testCase.Input, func(s string) {
				got = append(got, s)
			})
			if !reflect.DeepEqual(got, testCase.Expected) {
				t.Errorf("expected %q, but got %q", testCase.Expected, got)
			}
		})
	}

	t.Run("maps", func(t *testing.T) {
		testData := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}
		var got []string
		walk(testData, func(s string) {
			got = append(got, s)
		})
		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("chans", func(t *testing.T) {
		testData := make(chan Profile)
		go func() {
			testData <- Profile{33, "Berlin"}
			testData <- Profile{34, "Katowice"}
			close(testData)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(testData, func(s string) {
			got = append(got, s)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v got %v", want, got)
		}
	})

	t.Run("functions", func(t *testing.T) {
		testData := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(testData, func(s string) {
			got = append(got, s)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v got %v", want, got)
		}
	})
}

func assertContains(t testing.TB, got []string, item string) {
	t.Helper()
	contains := false
	for _, v := range got {
		if v == item {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q", got, item)
	}
}
