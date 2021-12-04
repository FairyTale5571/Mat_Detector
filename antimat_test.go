package swears_test

import (
	"swears"
	"testing"
)


func main() {

}

func TestContainsMat(t *testing.T) {
	type args struct {
		s string
	}
	var tests []struct {
		name string
		args args
		want bool
	}
	toCheck := []string{
		"да опять импульсы по КД просто , скучно же, еще и курицы еще эти ебанутые, смысл не у всех интерес играть чтобы рашить, кто АВП покупает изначально готов в кустах сидеть и снайперить, а тут курицы по всюду",
		"в любом гараже достанешь",

	}
	tests = append(tests)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swears.ContainsMat(tt.args.s); got != tt.want {
				t.Errorf("ContainsMat() = %v, want %v", got, tt.want)
			}
		})
	}
}