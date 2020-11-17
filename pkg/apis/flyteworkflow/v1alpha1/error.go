package v1alpha1

import (
	"bytes"
	"encoding/json"

	"github.com/golang/protobuf/jsonpb"
	"github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
)

// Wrapper around core.Execution error. Execution Error has a protobuf enum and hence needs to be wrapped by custom marshaller
type ExecutionError struct {
	*core.ExecutionError
}

func (in *ExecutionError) UnmarshalJSON(b []byte) error {
	in.ExecutionError = &core.ExecutionError{}
	return jsonpb.Unmarshal(bytes.NewReader(b), in.ExecutionError)
}

func (in *ExecutionError) MarshalJSON() ([]byte, error) {
	if in == nil {
		return json.Marshal(nil)
	}
	var buf bytes.Buffer
	if err := marshaler.Marshal(&buf, in.ExecutionError); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (in *ExecutionError) DeepCopyInto(out *ExecutionError) {
	*out = *in
}
