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

type Timestamp struct {
	Timestamp int64 `json:"timestamp"`

	Millis int32 `json:"millis"`
}

const TimestampAvroCRC64Fingerprint = "\x8e\xdd@\fBJ[~"

func NewTimestamp() Timestamp {
	r := Timestamp{}
	return r
}

func DeserializeTimestamp(r io.Reader) (Timestamp, error) {
	t := NewTimestamp()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeTimestampFromSchema(r io.Reader, schema string) (Timestamp, error) {
	t := NewTimestamp()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeTimestamp(r Timestamp, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.Timestamp, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Millis, w)
	if err != nil {
		return err
	}
	return err
}

func (r Timestamp) Serialize(w io.Writer) error {
	return writeTimestamp(r, w)
}

func (r Timestamp) Schema() string {
	return "{\"fields\":[{\"name\":\"timestamp\",\"type\":\"long\"},{\"name\":\"millis\",\"type\":\"int\"}],\"name\":\"com.alibaba.dts.formats.avro.Timestamp\",\"type\":\"record\"}"
}

func (r Timestamp) SchemaName() string {
	return "com.alibaba.dts.formats.avro.Timestamp"
}

func (_ Timestamp) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Timestamp) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Timestamp) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Timestamp) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Timestamp) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Timestamp) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Timestamp) SetString(v string)   { panic("Unsupported operation") }
func (_ Timestamp) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Timestamp) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.Timestamp}

		return w

	case 1:
		w := types.Int{Target: &r.Millis}

		return w

	}
	panic("Unknown field index")
}

func (r *Timestamp) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Timestamp) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Timestamp) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Timestamp) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Timestamp) HintSize(int)                     { panic("Unsupported operation") }
func (_ Timestamp) Finalize()                        {}

func (_ Timestamp) AvroCRC64Fingerprint() []byte {
	return []byte(TimestampAvroCRC64Fingerprint)
}

func (r Timestamp) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["timestamp"], err = json.Marshal(r.Timestamp)
	if err != nil {
		return nil, err
	}
	output["millis"], err = json.Marshal(r.Millis)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Timestamp) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["timestamp"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Timestamp); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for timestamp")
	}
	val = func() json.RawMessage {
		if v, ok := fields["millis"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Millis); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for millis")
	}
	return nil
}