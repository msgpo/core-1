// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: internal

package internal

import (
	"fmt"
	"io"
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/vdl"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

type Struct struct {
	X int32
	Y int32
}

func (Struct) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/cmd/vrpc/internal.Struct"`
}) {
}

func (x Struct) VDLIsZero() bool {
	return x == Struct{}
}

func (x Struct) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_1); err != nil {
		return err
	}
	if x.X != 0 {
		if err := enc.NextFieldValueInt(0, vdl.Int32Type, int64(x.X)); err != nil {
			return err
		}
	}
	if x.Y != 0 {
		if err := enc.NextFieldValueInt(1, vdl.Int32Type, int64(x.Y)); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Struct) VDLRead(dec vdl.Decoder) error {
	*x = Struct{}
	if err := dec.StartValue(__VDLType_struct_1); err != nil {
		return err
	}
	decType := dec.Type()
	for {
		index, err := dec.NextField()
		switch {
		case err != nil:
			return err
		case index == -1:
			return dec.FinishValue()
		}
		if decType != __VDLType_struct_1 {
			index = __VDLType_struct_1.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			switch value, err := dec.ReadValueInt(32); {
			case err != nil:
				return err
			default:
				x.X = int32(value)
			}
		case 1:
			switch value, err := dec.ReadValueInt(32); {
			case err != nil:
				return err
			default:
				x.Y = int32(value)
			}
		}
	}
}

type Array2Int [2]int32

func (Array2Int) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/cmd/vrpc/internal.Array2Int"`
}) {
}

func (x Array2Int) VDLIsZero() bool {
	return x == Array2Int{}
}

func (x Array2Int) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_array_2); err != nil {
		return err
	}
	for _, elem := range x {
		if err := enc.NextEntryValueInt(vdl.Int32Type, int64(elem)); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Array2Int) VDLRead(dec vdl.Decoder) error {
	if err := dec.StartValue(__VDLType_array_2); err != nil {
		return err
	}
	for index := 0; index < 2; index++ {
		switch done, elem, err := dec.NextEntryValueInt(32); {
		case err != nil:
			return err
		case done:
			return fmt.Errorf("short array, got len %d < 2 %T)", index, *x)
		default:
			x[index] = int32(elem)
		}
	}
	switch done, err := dec.NextEntry(); {
	case err != nil:
		return err
	case !done:
		return fmt.Errorf("long array, got len > 2 %T", *x)
	}
	return dec.FinishValue()
}

//////////////////////////////////////////////////
// Interface definitions

// TypeTesterClientMethods is the client interface
// containing TypeTester methods.
//
// TypeTester methods are listed in alphabetical order, to make it easier to
// test Signature output, which sorts methods alphabetically.
type TypeTesterClientMethods interface {
	// Methods to test support for primitive types.
	EchoBool(_ *context.T, I1 bool, _ ...rpc.CallOpt) (O1 bool, _ error)
	EchoFloat32(_ *context.T, I1 float32, _ ...rpc.CallOpt) (O1 float32, _ error)
	EchoFloat64(_ *context.T, I1 float64, _ ...rpc.CallOpt) (O1 float64, _ error)
	EchoInt32(_ *context.T, I1 int32, _ ...rpc.CallOpt) (O1 int32, _ error)
	EchoInt64(_ *context.T, I1 int64, _ ...rpc.CallOpt) (O1 int64, _ error)
	EchoString(_ *context.T, I1 string, _ ...rpc.CallOpt) (O1 string, _ error)
	EchoByte(_ *context.T, I1 byte, _ ...rpc.CallOpt) (O1 byte, _ error)
	EchoUint32(_ *context.T, I1 uint32, _ ...rpc.CallOpt) (O1 uint32, _ error)
	EchoUint64(_ *context.T, I1 uint64, _ ...rpc.CallOpt) (O1 uint64, _ error)
	// Methods to test support for composite types.
	XEchoArray(_ *context.T, I1 Array2Int, _ ...rpc.CallOpt) (O1 Array2Int, _ error)
	XEchoMap(_ *context.T, I1 map[int32]string, _ ...rpc.CallOpt) (O1 map[int32]string, _ error)
	XEchoSet(_ *context.T, I1 map[int32]struct{}, _ ...rpc.CallOpt) (O1 map[int32]struct{}, _ error)
	XEchoSlice(_ *context.T, I1 []int32, _ ...rpc.CallOpt) (O1 []int32, _ error)
	XEchoStruct(_ *context.T, I1 Struct, _ ...rpc.CallOpt) (O1 Struct, _ error)
	// Methods to test support for different number of arguments.
	YMultiArg(_ *context.T, I1 int32, I2 int32, _ ...rpc.CallOpt) (O1 int32, O2 int32, _ error)
	YNoArgs(*context.T, ...rpc.CallOpt) error
	// Methods to test support for streaming.
	ZStream(_ *context.T, NumStreamItems int32, StreamItem bool, _ ...rpc.CallOpt) (TypeTesterZStreamClientCall, error)
}

// TypeTesterClientStub adds universal methods to TypeTesterClientMethods.
type TypeTesterClientStub interface {
	TypeTesterClientMethods
	rpc.UniversalServiceMethods
}

// TypeTesterClient returns a client stub for TypeTester.
func TypeTesterClient(name string) TypeTesterClientStub {
	return implTypeTesterClientStub{name}
}

type implTypeTesterClientStub struct {
	name string
}

func (c implTypeTesterClientStub) EchoBool(ctx *context.T, i0 bool, opts ...rpc.CallOpt) (o0 bool, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "EchoBool", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implTypeTesterClientStub) EchoFloat32(ctx *context.T, i0 float32, opts ...rpc.CallOpt) (o0 float32, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "EchoFloat32", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implTypeTesterClientStub) EchoFloat64(ctx *context.T, i0 float64, opts ...rpc.CallOpt) (o0 float64, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "EchoFloat64", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implTypeTesterClientStub) EchoInt32(ctx *context.T, i0 int32, opts ...rpc.CallOpt) (o0 int32, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "EchoInt32", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implTypeTesterClientStub) EchoInt64(ctx *context.T, i0 int64, opts ...rpc.CallOpt) (o0 int64, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "EchoInt64", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implTypeTesterClientStub) EchoString(ctx *context.T, i0 string, opts ...rpc.CallOpt) (o0 string, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "EchoString", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implTypeTesterClientStub) EchoByte(ctx *context.T, i0 byte, opts ...rpc.CallOpt) (o0 byte, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "EchoByte", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implTypeTesterClientStub) EchoUint32(ctx *context.T, i0 uint32, opts ...rpc.CallOpt) (o0 uint32, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "EchoUint32", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implTypeTesterClientStub) EchoUint64(ctx *context.T, i0 uint64, opts ...rpc.CallOpt) (o0 uint64, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "EchoUint64", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implTypeTesterClientStub) XEchoArray(ctx *context.T, i0 Array2Int, opts ...rpc.CallOpt) (o0 Array2Int, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "XEchoArray", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implTypeTesterClientStub) XEchoMap(ctx *context.T, i0 map[int32]string, opts ...rpc.CallOpt) (o0 map[int32]string, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "XEchoMap", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implTypeTesterClientStub) XEchoSet(ctx *context.T, i0 map[int32]struct{}, opts ...rpc.CallOpt) (o0 map[int32]struct{}, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "XEchoSet", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implTypeTesterClientStub) XEchoSlice(ctx *context.T, i0 []int32, opts ...rpc.CallOpt) (o0 []int32, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "XEchoSlice", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implTypeTesterClientStub) XEchoStruct(ctx *context.T, i0 Struct, opts ...rpc.CallOpt) (o0 Struct, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "XEchoStruct", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implTypeTesterClientStub) YMultiArg(ctx *context.T, i0 int32, i1 int32, opts ...rpc.CallOpt) (o0 int32, o1 int32, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "YMultiArg", []interface{}{i0, i1}, []interface{}{&o0, &o1}, opts...)
	return
}

func (c implTypeTesterClientStub) YNoArgs(ctx *context.T, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "YNoArgs", nil, nil, opts...)
	return
}

func (c implTypeTesterClientStub) ZStream(ctx *context.T, i0 int32, i1 bool, opts ...rpc.CallOpt) (ocall TypeTesterZStreamClientCall, err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "ZStream", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	ocall = &implTypeTesterZStreamClientCall{ClientCall: call}
	return
}

// TypeTesterZStreamClientStream is the client stream for TypeTester.ZStream.
type TypeTesterZStreamClientStream interface {
	// RecvStream returns the receiver side of the TypeTester.ZStream client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() bool
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
}

// TypeTesterZStreamClientCall represents the call returned from TypeTester.ZStream.
type TypeTesterZStreamClientCall interface {
	TypeTesterZStreamClientStream
	// Finish blocks until the server is done, and returns the positional return
	// values for call.
	//
	// Finish returns immediately if the call has been canceled; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless the call
	// has been canceled or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() error
}

type implTypeTesterZStreamClientCall struct {
	rpc.ClientCall
	valRecv bool
	errRecv error
}

func (c *implTypeTesterZStreamClientCall) RecvStream() interface {
	Advance() bool
	Value() bool
	Err() error
} {
	return implTypeTesterZStreamClientCallRecv{c}
}

type implTypeTesterZStreamClientCallRecv struct {
	c *implTypeTesterZStreamClientCall
}

func (c implTypeTesterZStreamClientCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implTypeTesterZStreamClientCallRecv) Value() bool {
	return c.c.valRecv
}
func (c implTypeTesterZStreamClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implTypeTesterZStreamClientCall) Finish() (err error) {
	err = c.ClientCall.Finish()
	return
}

// TypeTesterServerMethods is the interface a server writer
// implements for TypeTester.
//
// TypeTester methods are listed in alphabetical order, to make it easier to
// test Signature output, which sorts methods alphabetically.
type TypeTesterServerMethods interface {
	// Methods to test support for primitive types.
	EchoBool(_ *context.T, _ rpc.ServerCall, I1 bool) (O1 bool, _ error)
	EchoFloat32(_ *context.T, _ rpc.ServerCall, I1 float32) (O1 float32, _ error)
	EchoFloat64(_ *context.T, _ rpc.ServerCall, I1 float64) (O1 float64, _ error)
	EchoInt32(_ *context.T, _ rpc.ServerCall, I1 int32) (O1 int32, _ error)
	EchoInt64(_ *context.T, _ rpc.ServerCall, I1 int64) (O1 int64, _ error)
	EchoString(_ *context.T, _ rpc.ServerCall, I1 string) (O1 string, _ error)
	EchoByte(_ *context.T, _ rpc.ServerCall, I1 byte) (O1 byte, _ error)
	EchoUint32(_ *context.T, _ rpc.ServerCall, I1 uint32) (O1 uint32, _ error)
	EchoUint64(_ *context.T, _ rpc.ServerCall, I1 uint64) (O1 uint64, _ error)
	// Methods to test support for composite types.
	XEchoArray(_ *context.T, _ rpc.ServerCall, I1 Array2Int) (O1 Array2Int, _ error)
	XEchoMap(_ *context.T, _ rpc.ServerCall, I1 map[int32]string) (O1 map[int32]string, _ error)
	XEchoSet(_ *context.T, _ rpc.ServerCall, I1 map[int32]struct{}) (O1 map[int32]struct{}, _ error)
	XEchoSlice(_ *context.T, _ rpc.ServerCall, I1 []int32) (O1 []int32, _ error)
	XEchoStruct(_ *context.T, _ rpc.ServerCall, I1 Struct) (O1 Struct, _ error)
	// Methods to test support for different number of arguments.
	YMultiArg(_ *context.T, _ rpc.ServerCall, I1 int32, I2 int32) (O1 int32, O2 int32, _ error)
	YNoArgs(*context.T, rpc.ServerCall) error
	// Methods to test support for streaming.
	ZStream(_ *context.T, _ TypeTesterZStreamServerCall, NumStreamItems int32, StreamItem bool) error
}

// TypeTesterServerStubMethods is the server interface containing
// TypeTester methods, as expected by rpc.Server.
// The only difference between this interface and TypeTesterServerMethods
// is the streaming methods.
type TypeTesterServerStubMethods interface {
	// Methods to test support for primitive types.
	EchoBool(_ *context.T, _ rpc.ServerCall, I1 bool) (O1 bool, _ error)
	EchoFloat32(_ *context.T, _ rpc.ServerCall, I1 float32) (O1 float32, _ error)
	EchoFloat64(_ *context.T, _ rpc.ServerCall, I1 float64) (O1 float64, _ error)
	EchoInt32(_ *context.T, _ rpc.ServerCall, I1 int32) (O1 int32, _ error)
	EchoInt64(_ *context.T, _ rpc.ServerCall, I1 int64) (O1 int64, _ error)
	EchoString(_ *context.T, _ rpc.ServerCall, I1 string) (O1 string, _ error)
	EchoByte(_ *context.T, _ rpc.ServerCall, I1 byte) (O1 byte, _ error)
	EchoUint32(_ *context.T, _ rpc.ServerCall, I1 uint32) (O1 uint32, _ error)
	EchoUint64(_ *context.T, _ rpc.ServerCall, I1 uint64) (O1 uint64, _ error)
	// Methods to test support for composite types.
	XEchoArray(_ *context.T, _ rpc.ServerCall, I1 Array2Int) (O1 Array2Int, _ error)
	XEchoMap(_ *context.T, _ rpc.ServerCall, I1 map[int32]string) (O1 map[int32]string, _ error)
	XEchoSet(_ *context.T, _ rpc.ServerCall, I1 map[int32]struct{}) (O1 map[int32]struct{}, _ error)
	XEchoSlice(_ *context.T, _ rpc.ServerCall, I1 []int32) (O1 []int32, _ error)
	XEchoStruct(_ *context.T, _ rpc.ServerCall, I1 Struct) (O1 Struct, _ error)
	// Methods to test support for different number of arguments.
	YMultiArg(_ *context.T, _ rpc.ServerCall, I1 int32, I2 int32) (O1 int32, O2 int32, _ error)
	YNoArgs(*context.T, rpc.ServerCall) error
	// Methods to test support for streaming.
	ZStream(_ *context.T, _ *TypeTesterZStreamServerCallStub, NumStreamItems int32, StreamItem bool) error
}

// TypeTesterServerStub adds universal methods to TypeTesterServerStubMethods.
type TypeTesterServerStub interface {
	TypeTesterServerStubMethods
	// Describe the TypeTester interfaces.
	Describe__() []rpc.InterfaceDesc
}

// TypeTesterServer returns a server stub for TypeTester.
// It converts an implementation of TypeTesterServerMethods into
// an object that may be used by rpc.Server.
func TypeTesterServer(impl TypeTesterServerMethods) TypeTesterServerStub {
	stub := implTypeTesterServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implTypeTesterServerStub struct {
	impl TypeTesterServerMethods
	gs   *rpc.GlobState
}

func (s implTypeTesterServerStub) EchoBool(ctx *context.T, call rpc.ServerCall, i0 bool) (bool, error) {
	return s.impl.EchoBool(ctx, call, i0)
}

func (s implTypeTesterServerStub) EchoFloat32(ctx *context.T, call rpc.ServerCall, i0 float32) (float32, error) {
	return s.impl.EchoFloat32(ctx, call, i0)
}

func (s implTypeTesterServerStub) EchoFloat64(ctx *context.T, call rpc.ServerCall, i0 float64) (float64, error) {
	return s.impl.EchoFloat64(ctx, call, i0)
}

func (s implTypeTesterServerStub) EchoInt32(ctx *context.T, call rpc.ServerCall, i0 int32) (int32, error) {
	return s.impl.EchoInt32(ctx, call, i0)
}

func (s implTypeTesterServerStub) EchoInt64(ctx *context.T, call rpc.ServerCall, i0 int64) (int64, error) {
	return s.impl.EchoInt64(ctx, call, i0)
}

func (s implTypeTesterServerStub) EchoString(ctx *context.T, call rpc.ServerCall, i0 string) (string, error) {
	return s.impl.EchoString(ctx, call, i0)
}

func (s implTypeTesterServerStub) EchoByte(ctx *context.T, call rpc.ServerCall, i0 byte) (byte, error) {
	return s.impl.EchoByte(ctx, call, i0)
}

func (s implTypeTesterServerStub) EchoUint32(ctx *context.T, call rpc.ServerCall, i0 uint32) (uint32, error) {
	return s.impl.EchoUint32(ctx, call, i0)
}

func (s implTypeTesterServerStub) EchoUint64(ctx *context.T, call rpc.ServerCall, i0 uint64) (uint64, error) {
	return s.impl.EchoUint64(ctx, call, i0)
}

func (s implTypeTesterServerStub) XEchoArray(ctx *context.T, call rpc.ServerCall, i0 Array2Int) (Array2Int, error) {
	return s.impl.XEchoArray(ctx, call, i0)
}

func (s implTypeTesterServerStub) XEchoMap(ctx *context.T, call rpc.ServerCall, i0 map[int32]string) (map[int32]string, error) {
	return s.impl.XEchoMap(ctx, call, i0)
}

func (s implTypeTesterServerStub) XEchoSet(ctx *context.T, call rpc.ServerCall, i0 map[int32]struct{}) (map[int32]struct{}, error) {
	return s.impl.XEchoSet(ctx, call, i0)
}

func (s implTypeTesterServerStub) XEchoSlice(ctx *context.T, call rpc.ServerCall, i0 []int32) ([]int32, error) {
	return s.impl.XEchoSlice(ctx, call, i0)
}

func (s implTypeTesterServerStub) XEchoStruct(ctx *context.T, call rpc.ServerCall, i0 Struct) (Struct, error) {
	return s.impl.XEchoStruct(ctx, call, i0)
}

func (s implTypeTesterServerStub) YMultiArg(ctx *context.T, call rpc.ServerCall, i0 int32, i1 int32) (int32, int32, error) {
	return s.impl.YMultiArg(ctx, call, i0, i1)
}

func (s implTypeTesterServerStub) YNoArgs(ctx *context.T, call rpc.ServerCall) error {
	return s.impl.YNoArgs(ctx, call)
}

func (s implTypeTesterServerStub) ZStream(ctx *context.T, call *TypeTesterZStreamServerCallStub, i0 int32, i1 bool) error {
	return s.impl.ZStream(ctx, call, i0, i1)
}

func (s implTypeTesterServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implTypeTesterServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{TypeTesterDesc}
}

// TypeTesterDesc describes the TypeTester interface.
var TypeTesterDesc rpc.InterfaceDesc = descTypeTester

// descTypeTester hides the desc to keep godoc clean.
var descTypeTester = rpc.InterfaceDesc{
	Name:    "TypeTester",
	PkgPath: "v.io/x/ref/cmd/vrpc/internal",
	Doc:     "// TypeTester methods are listed in alphabetical order, to make it easier to\n// test Signature output, which sorts methods alphabetically.",
	Methods: []rpc.MethodDesc{
		{
			Name: "EchoBool",
			Doc:  "// Methods to test support for primitive types.",
			InArgs: []rpc.ArgDesc{
				{"I1", ``}, // bool
			},
			OutArgs: []rpc.ArgDesc{
				{"O1", ``}, // bool
			},
		},
		{
			Name: "EchoFloat32",
			InArgs: []rpc.ArgDesc{
				{"I1", ``}, // float32
			},
			OutArgs: []rpc.ArgDesc{
				{"O1", ``}, // float32
			},
		},
		{
			Name: "EchoFloat64",
			InArgs: []rpc.ArgDesc{
				{"I1", ``}, // float64
			},
			OutArgs: []rpc.ArgDesc{
				{"O1", ``}, // float64
			},
		},
		{
			Name: "EchoInt32",
			InArgs: []rpc.ArgDesc{
				{"I1", ``}, // int32
			},
			OutArgs: []rpc.ArgDesc{
				{"O1", ``}, // int32
			},
		},
		{
			Name: "EchoInt64",
			InArgs: []rpc.ArgDesc{
				{"I1", ``}, // int64
			},
			OutArgs: []rpc.ArgDesc{
				{"O1", ``}, // int64
			},
		},
		{
			Name: "EchoString",
			InArgs: []rpc.ArgDesc{
				{"I1", ``}, // string
			},
			OutArgs: []rpc.ArgDesc{
				{"O1", ``}, // string
			},
		},
		{
			Name: "EchoByte",
			InArgs: []rpc.ArgDesc{
				{"I1", ``}, // byte
			},
			OutArgs: []rpc.ArgDesc{
				{"O1", ``}, // byte
			},
		},
		{
			Name: "EchoUint32",
			InArgs: []rpc.ArgDesc{
				{"I1", ``}, // uint32
			},
			OutArgs: []rpc.ArgDesc{
				{"O1", ``}, // uint32
			},
		},
		{
			Name: "EchoUint64",
			InArgs: []rpc.ArgDesc{
				{"I1", ``}, // uint64
			},
			OutArgs: []rpc.ArgDesc{
				{"O1", ``}, // uint64
			},
		},
		{
			Name: "XEchoArray",
			Doc:  "// Methods to test support for composite types.",
			InArgs: []rpc.ArgDesc{
				{"I1", ``}, // Array2Int
			},
			OutArgs: []rpc.ArgDesc{
				{"O1", ``}, // Array2Int
			},
		},
		{
			Name: "XEchoMap",
			InArgs: []rpc.ArgDesc{
				{"I1", ``}, // map[int32]string
			},
			OutArgs: []rpc.ArgDesc{
				{"O1", ``}, // map[int32]string
			},
		},
		{
			Name: "XEchoSet",
			InArgs: []rpc.ArgDesc{
				{"I1", ``}, // map[int32]struct{}
			},
			OutArgs: []rpc.ArgDesc{
				{"O1", ``}, // map[int32]struct{}
			},
		},
		{
			Name: "XEchoSlice",
			InArgs: []rpc.ArgDesc{
				{"I1", ``}, // []int32
			},
			OutArgs: []rpc.ArgDesc{
				{"O1", ``}, // []int32
			},
		},
		{
			Name: "XEchoStruct",
			InArgs: []rpc.ArgDesc{
				{"I1", ``}, // Struct
			},
			OutArgs: []rpc.ArgDesc{
				{"O1", ``}, // Struct
			},
		},
		{
			Name: "YMultiArg",
			Doc:  "// Methods to test support for different number of arguments.",
			InArgs: []rpc.ArgDesc{
				{"I1", ``}, // int32
				{"I2", ``}, // int32
			},
			OutArgs: []rpc.ArgDesc{
				{"O1", ``}, // int32
				{"O2", ``}, // int32
			},
		},
		{
			Name: "YNoArgs",
		},
		{
			Name: "ZStream",
			Doc:  "// Methods to test support for streaming.",
			InArgs: []rpc.ArgDesc{
				{"NumStreamItems", ``}, // int32
				{"StreamItem", ``},     // bool
			},
		},
	},
}

// TypeTesterZStreamServerStream is the server stream for TypeTester.ZStream.
type TypeTesterZStreamServerStream interface {
	// SendStream returns the send side of the TypeTester.ZStream server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item bool) error
	}
}

// TypeTesterZStreamServerCall represents the context passed to TypeTester.ZStream.
type TypeTesterZStreamServerCall interface {
	rpc.ServerCall
	TypeTesterZStreamServerStream
}

// TypeTesterZStreamServerCallStub is a wrapper that converts rpc.StreamServerCall into
// a typesafe stub that implements TypeTesterZStreamServerCall.
type TypeTesterZStreamServerCallStub struct {
	rpc.StreamServerCall
}

// Init initializes TypeTesterZStreamServerCallStub from rpc.StreamServerCall.
func (s *TypeTesterZStreamServerCallStub) Init(call rpc.StreamServerCall) {
	s.StreamServerCall = call
}

// SendStream returns the send side of the TypeTester.ZStream server stream.
func (s *TypeTesterZStreamServerCallStub) SendStream() interface {
	Send(item bool) error
} {
	return implTypeTesterZStreamServerCallSend{s}
}

type implTypeTesterZStreamServerCallSend struct {
	s *TypeTesterZStreamServerCallStub
}

func (s implTypeTesterZStreamServerCallSend) Send(item bool) error {
	return s.s.Send(item)
}

// Hold type definitions in package-level variables, for better performance.
var (
	__VDLType_struct_1 *vdl.Type
	__VDLType_array_2  *vdl.Type
)

var __VDLInitCalled bool

// __VDLInit performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
//    var _ = __VDLInit()
//
// The purpose of this function is to ensure that vdl initialization occurs in
// the right order, and very early in the init sequence.  In particular, vdl
// registration and package variable initialization needs to occur before
// functions like vdl.TypeOf will work properly.
//
// This function returns a dummy value, so that it can be used to initialize the
// first var in the file, to take advantage of Go's defined init order.
func __VDLInit() struct{} {
	if __VDLInitCalled {
		return struct{}{}
	}
	__VDLInitCalled = true

	// Register types.
	vdl.Register((*Struct)(nil))
	vdl.Register((*Array2Int)(nil))

	// Initialize type definitions.
	__VDLType_struct_1 = vdl.TypeOf((*Struct)(nil)).Elem()
	__VDLType_array_2 = vdl.TypeOf((*Array2Int)(nil))

	return struct{}{}
}