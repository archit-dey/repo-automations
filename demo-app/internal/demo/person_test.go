package demo

import "testing"

func TestPersonIntroduce(t *testing.T) {
	tests := []struct {
		name   string
		person Person
		want   string
	}{
		{"adult", Person{"Alice", 30}, "Hi, I'm Alice and I'm 30 years old."},
		{"child", Person{"Bob", 10}, "Hi, I'm Bob and I'm 10 years old."},
		{"zero age", Person{"Charlie", 0}, "Hi, I'm Charlie and I'm 0 years old."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.person.Introduce()
			if got != tt.want {
				t.Errorf("Introduce() = %q; want %q", got, tt.want)
			}
		})
	}
}

func TestPersonIsAdult(t *testing.T) {
	tests := []struct {
		name   string
		person Person
		want   bool
	}{
		{"adult", Person{"Alice", 30}, true},
		{"exactly 18", Person{"Bob", 18}, true},
		{"minor", Person{"Charlie", 17}, false},
		{"child", Person{"David", 5}, false},
		{"zero age", Person{"Eve", 0}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.person.IsAdult()
			if got != tt.want {
				t.Errorf("IsAdult() = %v; want %v", got, tt.want)
			}
		})
	}
}

func TestPersonHaveBirthday(t *testing.T) {
	person := Person{"Alice", 25}
	person.HaveBirthday()
	if person.Age != 26 {
		t.Errorf("After HaveBirthday(), Age = %d; want 26", person.Age)
	}

	person.HaveBirthday()
	person.HaveBirthday()
	if person.Age != 28 {
		t.Errorf("After 3 birthdays, Age = %d; want 28", person.Age)
	}
}
