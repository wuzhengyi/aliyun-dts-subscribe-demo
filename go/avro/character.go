// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Character struct {
	Charset string `json:"charset"`

	Value Bytes `json:"value"`
}

const CharacterAvroCRC64Fingerprint = "\xc1a\xeb\xfb\x9e\x97\\\xe7"

func NewCharacter() Character {
	r := Character{}
	return r
}

func DeserializeCharacter(r io.Reader) (Character, error) {
	t := NewCharacter()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCharacterFromSchema(r io.Reader, schema string) (Character, error) {
	t := NewCharacter()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCharacter(r Character, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Charset, w)
	if err != nil {
		return err
	}
	err = vm.WriteBytes(r.Value, w)
	if err != nil {
		return err
	}
	return err
}

func (r Character) Serialize(w io.Writer) error {
	return writeCharacter(r, w)
}

func (r Character) Schema() string {
	return "{\"fields\":[{\"name\":\"charset\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"com.alibaba.dts.formats.avro.Character\",\"type\":\"record\"}"
}

func (r Character) SchemaName() string {
	return "com.alibaba.dts.formats.avro.Character"
}

func (_ Character) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Character) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Character) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Character) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Character) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Character) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Character) SetString(v string)   { panic("Unsupported operation") }
func (_ Character) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Character) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Charset}

		return w

	case 1:
		w := BytesWrapper{Target: &r.Value}

		return w

	}
	panic("Unknown field index")
}

func (r *Character) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Character) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Character) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Character) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Character) HintSize(int)                     { panic("Unsupported operation") }
func (_ Character) Finalize()                        {}

func (_ Character) AvroCRC64Fingerprint() []byte {
	return []byte(CharacterAvroCRC64Fingerprint)
}

func (r Character) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["charset"], err = json.Marshal(r.Charset)
	if err != nil {
		return nil, err
	}
	output["value"], err = json.Marshal(r.Value)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Character) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["charset"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Charset); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for charset")
	}
	val = func() json.RawMessage {
		if v, ok := fields["value"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Value); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for value")
	}
	return nil
}