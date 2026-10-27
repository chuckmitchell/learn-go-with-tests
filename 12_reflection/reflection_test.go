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

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{Name: "Chris"},
			[]string{"Chris"},
		},
		{
			"struct with 2 string fields",
			struct {
				Name string
				City string
			}{Name: "Chris", City: "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{Name: "Chris", Age: 19},
			[]string{"Chris"},
		},
		{
			"nested struct",
			Person{
				Name: "Chris",
				Profile: Profile{
					Age:  33,
					City: "London",
				},
			},
			[]string{"Chris", "London"},
		},
		{
			"struct witha pointer",
			&Person{
				Name: "Chris",
				Profile: Profile{
					Age:  33,
					City: "London",
				},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"},
				{34, "Halifax"},
			},
			[]string{"London", "Halifax"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{34, "Halifax"},
			},
			[]string{"London", "Halifax"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, wanted %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baaa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baaa")
	})
}
