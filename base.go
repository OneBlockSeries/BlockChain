package main

import (
    "encoding/binary"
    "unsafe"
    "reflect"
)

func IntToBytes(i int) []byte {
    var buf = make([]byte, 4)
    binary.BigEndian.PutUint32(buf, uint32(i))
    return buf
}

func BytesToInt(buf []byte) int32 {
    return int32(binary.BigEndian.Uint32(buf))
}
func BytesToString(b []byte) string {
    return *(*string)(unsafe.Pointer(&b))
}

func StringToBytes(s string) []byte {
    sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
    bh := reflect.SliceHeader{sh.Data, sh.Len, 0}
    return *(*[]byte)(unsafe.Pointer(&bh))
}