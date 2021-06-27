package templatefunction

import (
	"testing"
)

func TestRepeatInt(t *testing.T) {
	a := Repeat(3, "a")
	if a != "aaa" {
		t.Error("aaa/" + a)
	}
}

func TestRepeatFloat(t *testing.T) {
	a := Repeat(3.8, "a")
	if a != "aaa" {
		t.Error("aaa/" + a)
	}
}

func TestRepeatArray(t *testing.T) {
	var r = [5]int{1, 2, 3, 4, 5}
	a := Repeat(r, "a")
	if a != "aaaaa" {
		t.Error("aaaaa/" + a)
	}
}

func TestRepeatSlice(t *testing.T) {
	var r = []int{1, 2, 3, 4, 5}
	a := Repeat(r, "a")
	if a != "aaaaa" {
		t.Error("aaaaa/" + a)
	}
}

func TestRepeatMap(t *testing.T) {
	var r = map[string]interface{}{
		"hello":  "world",
		"string": "value",
		"repeat": "test",
	}
	a := Repeat(r, "a")
	if a != "aaa" {
		t.Error("aaa/" + a)
	}
}

func TestJoinSlice(t *testing.T) {
	a := Join(".", []int{1, 2, 3, 4, 5})
	if a != "1.2.3.4.5" {
		t.Error("1.2.3.4.5/" + a)
	}
}

func TestJoinMap(t *testing.T) {
	a := Join(".", map[string]interface{}{
		"hello": "world",
	})
	if a != "hello.world" {
		t.Error("hello.world/" + a)
	}
}

func TestRemoveLast(t *testing.T) {
	a := RemoveLast("12345")
	if a != "1234" {
		t.Error("1234/" + a)
	}
}

func TestRemoveLastSlice(t *testing.T) {
	a := RemoveLast([]int{100, 200, 300, 400, 500})
	if a != "100200300400" {
		t.Error("100200300400/" + a)
	}
}

func TestRemoveFirst(t *testing.T) {
	a := RemoveFirst("12345")
	if a != "2345" {
		t.Error("2345/" + a)
	}
}

func TestRemoveFirstSlice(t *testing.T) {
	a := RemoveFirst([]int{100, 200, 300, 400, 500})
	if a != "200300400500" {
		t.Error("200300400500/" + a)
	}
}

func Test_to_snake_1(t *testing.T) {
	c := ToSnake("HelloWorld")
	if c != "hello_world" {
		t.Error("hello_world" + c)
	}
}

func Test_to_snake_2(t *testing.T) {
	c := ToSnake("GOLANG")
	if c != "g_o_l_a_n_g" {
		t.Error("g_o_l_a_n_g" + c)
	}
}

func Test_to_snake_3(t *testing.T) {
	c := ToSnake("HE_llo")
	if c != "h_e_llo" {
		t.Error("h_e_llo/" + c)
	}
}

func Test_to_snake_4(t *testing.T) {
	c := ToSnake("h_e_l_l_o_w_o_r_l_d")
	if c != "h_e_l_l_o_w_o_r_l_d" {
		t.Error("h_e_l_l_o_w_o_r_l_d/" + c)
	}
}

func TestToPascal1(t *testing.T) {
	c := ToPascal("HelloWorld")
	if c != "HelloWorld" {
		t.Error("HelloWorld" + c)
	}
}

func TestToPascal2(t *testing.T) {
	c := ToPascal("g_o_l_a_n_g")
	if c != "GOLANG" {
		t.Error("GOLANG" + c)
	}
}

func TestToPascal3(t *testing.T) {
	c := ToPascal("HE_llo")
	if c != "HELlo" {
		t.Error("HELlo/" + c)
	}
}

func TestToPascal4(t *testing.T) {
	c := ToPascal("h_e_l_l_o_w_o_r_l_d")
	if c != "HELLOWORLD" {
		t.Error("HELLOWORLD/" + c)
	}
}

func TestToCamel1(t *testing.T) {
	c := ToCamel("HelloWorld")
	if c != "helloWorld" {
		t.Error("helloWorld" + c)
	}
}

func TestToCamel2(t *testing.T) {
	c := ToCamel("g_o_l_a_n_g")
	if c != "gOLANG" {
		t.Error("gOLANG" + c)
	}
}

func TestToCamel3(t *testing.T) {
	c := ToCamel("HE_llo")
	if c != "hELlo" {
		t.Error("hELlo/" + c)
	}
}

func TestToCamel4(t *testing.T) {
	c := ToCamel("h_e_l_l_o_w_o_r_l_d")
	if c != "hELLOWORLD" {
		t.Error("hELLOWORLD/" + c)
	}
}
