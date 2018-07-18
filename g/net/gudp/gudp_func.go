// Copyright 2017-2018 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

package gudp

import "net"

// 创建标准库UDP链接操作对象
func NewNetConn(raddr string, laddr...string) (*net.UDPConn, error) {
    var err error
    var rudpaddr, ludpaddr *net.UDPAddr
    rudpaddr, err = net.ResolveUDPAddr("udp", raddr)
    if err != nil {
        return nil, err
    }
    if len(laddr) > 0 {
        ludpaddr, err = net.ResolveUDPAddr("udp", laddr[0])
        if err != nil {
            return nil, err
        }
    }
    conn, err := net.DialUDP("udp", ludpaddr, rudpaddr)
    if err != nil {
        return nil, err
    }
    return conn, nil
}

// (面向短链接)发送数据
func Send(addr string, data []byte, retry...Retry) error {
    conn, err := NewConn(addr)
    if err != nil {
        return err
    }
    defer conn.Close()
    return conn.Send(data, retry...)
}

// (面向短链接)发送数据并等待接收返回数据
func SendReceive(addr string, data []byte, receive int, retry...Retry) ([]byte, error) {
    conn, err := NewConn(addr)
    if err != nil {
        return nil, err
    }
    defer conn.Close()
    return conn.SendReceive(data, receive, retry...)
}