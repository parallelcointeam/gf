// Copyright 2017 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

// 其他工具包
package gutil

import (
    "fmt"
    "bytes"
    "encoding/json"
    "reflect"
    "gitee.com/johng/gf/g/util/gconv"
)

// 格式化打印变量(类似于PHP-vardump)
func Dump(i...interface{}) {
    for _, v := range i {
        if b, ok := v.([]byte); ok {
            fmt.Print(string(b))
        } else {
            // 主要针对 map[interface{}]* 进行处理，json无法进行encode，
            // 这里强制对所有map进行反射处理转换
            refValue := reflect.ValueOf(v)
            if refValue.Kind() == reflect.Map {
                m := make(map[string]interface{})
                keys := refValue.MapKeys()
                for _, k := range keys {
                    key   := gconv.String(k.Interface())
                    m[key] = refValue.MapIndex(k).Interface()
                }
                v = m
            }
            // json encode并打印到终端
            buffer  := &bytes.Buffer{}
            encoder := json.NewEncoder(buffer)
            encoder.SetEscapeHTML(false)
            encoder.SetIndent("", "\t")
            if err := encoder.Encode(v); err == nil {
                fmt.Print(buffer.String())
            } else {
                fmt.Errorf("%s", err.Error())
            }
        }
        //fmt.Println()
    }
}

