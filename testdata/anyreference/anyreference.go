package anyreference

import (
	"github.com/openagent-md/guts/testdata/anyreference/reference"
	"github.com/openagent-md/guts/testdata/prefix"
)

type ExampleStruct struct {
	String string
	Int    int
}

type Example reference.Struct[map[string]string]

type String string

type UsesPrefixPackage struct {
	Field       prefix.Struct
	FieldString prefix.String
	FieldSlice  prefix.StructSlice
}

type ExternalString prefix.String
