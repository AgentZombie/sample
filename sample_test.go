package sample

import (
	"encoding/json"
	"fmt"
	"testing"
)

var (
	expInt   = fmt.Sprint(Int)
	expUint  = fmt.Sprint(Uint)
	expBool  = fmt.Sprint(Bool)
	expFloat = fmt.Sprint(Float)
)

func TestSamplePrimitives(t *testing.T) {
	t.Parallel()

	testStr := "bar"
	testInt := int(-35)
	testInt32 := int(-32)
	testInt64 := int(-64)
	testUint := uint(35)
	testUint32 := uint(32)
	testUint64 := uint(64)
	testBool := true
	testFloat32 := float32(-5432.1)
	testFloat64 := float64(-65432.1)

	for label, tCase := range map[string]struct {
		in   interface{}
		want string
	}{
		"string": {
			in:   testStr,
			want: String,
		},
		"*string": {
			in:   &testStr,
			want: String,
		},
		"int": {
			in:   testInt,
			want: expInt,
		},
		"*int": {
			in:   &testInt,
			want: expInt,
		},
		"int32": {
			in:   testInt32,
			want: expInt,
		},
		"*int32": {
			in:   &testInt32,
			want: expInt,
		},
		"int64": {
			in:   testInt64,
			want: expInt,
		},
		"*int64": {
			in:   &testInt64,
			want: expInt,
		},
		"uint": {
			in:   testUint,
			want: expUint,
		},
		"*uint": {
			in:   &testUint,
			want: expUint,
		},
		"uint32": {
			in:   testUint32,
			want: expUint,
		},
		"*uint32": {
			in:   &testUint32,
			want: expUint,
		},
		"uint64": {
			in:   testUint64,
			want: expUint,
		},
		"*uint64": {
			in:   &testUint64,
			want: expUint,
		},
		"bool": {
			in:   testBool,
			want: expBool,
		},
		"*bool": {
			in:   &testBool,
			want: expBool,
		},
		"float32": {
			in:   testFloat32,
			want: expFloat,
		},
		"*float32": {
			in:   &testFloat32,
			want: expFloat,
		},
		"float64": {
			in:   testFloat64,
			want: expFloat,
		},
		"*float64": {
			in:   &testFloat64,
			want: expFloat,
		},
		"byte slice": {
			in:   []byte{1, 2, 3, 4},
			want: Bytes,
		},
	} {
		out := Sample(tCase.in)
		t.Logf("%s: %T, %v\n", label, out, out)
		got := fmt.Sprint(out)
		if got != tCase.want {
			t.Fatalf("got %q, want %q", got, tCase.want)
		}
	}
}

// This test depends on the ordering of maps which shouldn't be relied on.
// If it becomes a problem in the future, create a normalizable structure.
func TestStruct(t *testing.T) {
	t.Parallel()

	f := Foo{}
	out := Sample(f)

	b, err := json.Marshal(out)
	if err != nil {
		t.Fatal("error marshalling: ", err)
	}

	fmt.Println("got: ", string(b))
	expect := `{"C":true,"D":{"BadTag":[-123456]},"a":-123456}`
	if expect != string(b) {
		t.Fatalf("got %s, want %s", string(b), expect)
	}
}

type Foo struct {
	A int    `json:"a"`
	B string `json:"-"`
	C *bool
	D Bar `json:",omitempty"`
}

type Bar struct {
	BadTag []int `json"name"`
}
