# Sample

Create a populated sample of a data type as a `map[string]interface{}`. This approach is slow and lossy but useful for generating sample data.

Useful for generating example JSON. It is currently tailored to JSON but could be adapted to other destination formats.

## Example

```
type Foo struct {
        A int    `json:"a"`
        B string `json:"-"`
        C *bool
        D Bar `json:",omitempty"`
}

type Bar struct {
        BadTag []int `json"name"`
}

f := Foo{}
out := Sample(f)

b, err := json.Marshal(out)
if err != nil {
	panic("error marshalling: " + err.Error())
}

// b == []byte(`{"C":true,"D":{"BadTag":[-123456]},"a":-123456}`)
```

## What it does

It recurses over a struct accumulating a `map[string]interface{}` applying the following rules:

* Pointers are dereferenced until reaching a concrete type
* Primitives are given sample values corresponding to their types:
  * Unhandled types are omitted
  * `string` becomes `"Foo"`
  * Signed integer types become `-123456`
  * Unsigned integer types become `123456`
  * `float32`, `float64` become `12345.6`
  * `bool` becomes `true`
  * Interface types become `nil`
  * `struct` becomes a nested `map[string]interface{}` which is recursed into
    * `json` tags are obeyed for field name and omission (`-`)
  * Arrays and slices become a single-item `[]interface{}` containing once instance of the appropriate type

## Limitations

* Complex values aren't handled
* Sample integer values don't reflect the integer size
* Ironically, maps aren't handled
