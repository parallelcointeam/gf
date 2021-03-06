// Copyright 2018 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

package gdb

import (
    "time"
    "gitee.com/johng/gf/g/util/gconv"
)

func (v Value) IsNil()          bool            { return v == nil }
func (v Value) Bytes()          []byte          { return []byte(v) }
func (v Value) String()         string          { return string(v.Bytes()) }
func (v Value) Bool()           bool            { return gconv.Bool(v.Bytes()) }

func (v Value) Int()            int             { return gconv.Int(v.Bytes()) }
func (v Value) Int8()           int8            { return gconv.Int8(v.Bytes()) }
func (v Value) Int16()          int16           { return gconv.Int16(v.Bytes()) }
func (v Value) Int32()          int32           { return gconv.Int32(v.Bytes()) }
func (v Value) Int64()          int64           { return gconv.Int64(v.Bytes()) }

func (v Value) Uint()           uint            { return gconv.Uint(v.Bytes()) }
func (v Value) Uint8()          uint8           { return gconv.Uint8(v.Bytes()) }
func (v Value) Uint16()         uint16          { return gconv.Uint16(v.Bytes()) }
func (v Value) Uint32()         uint32          { return gconv.Uint32(v.Bytes()) }
func (v Value) Uint64()         uint64          { return gconv.Uint64(v.Bytes()) }

func (v Value) Float32()        float32         { return gconv.Float32(v.Bytes()) }
func (v Value) Float64()        float64         { return gconv.Float64(v.Bytes()) }

func (v Value) Time(format...string) time.Time       { return gconv.Time(v.Bytes(), format...) }
func (v Value) TimeDuration()        time.Duration   { return gconv.TimeDuration(v.Bytes()) }