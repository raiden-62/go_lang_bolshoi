package storage

import "testing"

type testCase struct {
	test_name  string
	key   string
	value string
	//err bool
}

func TestSetGet(t *testing.T) {
	cases := []testCase{
		{"string value", "hello", "world"},
		{"empty string as value", "key2", ""},
		{"int value", "intkey", "123456"}
	}

	s, err := NewStorage()
	if err != nil {
		t.Errorf("new storage: %v", err)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s.Set(c.key, c.value)

			sValue := s.get(c.key)

			if *sValue != c.value {
				t.Errorf("values not equal")
			}
		})

	}
}

type testCaseWithType { //add true/false functionality for tests
	test_name string
	key string
	value string
	kind Kind
	//err bool
}

func TestSetGetithType(t *testing.T) {
	cases := []testCaseWithKind{
		{"hello world", "hello", "world", KindStr}
		{"int test", "key", "123456", KindInt}
	}

	s, err := NewStorage()
	if err != nil {
		t.Errorf("new storage: %v", err)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T){
			s.Set(c.key, c.value)

			sValue := s.Get(c.key)

			if *sValue != c.value{
				t.Errorf("values not equal")
			}

			if getType(*sValue) != getType(c.value){ //getType is returning wrong type for the same values
				t.Errorf("value kinds not equal")
			}

			if getType(*sValue) != c.kind { //getType is just returning wrong
				t.Errorf("expected value kind: %v", c.kind)
			}
		})
	}
}