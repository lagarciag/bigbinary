//bigbinary does this blah blah
package bigbinary

import (

	//"encoding/binary"
	"reflect"
	"fmt"
	"encoding/binary"
	"bytes"
)

func Read(buf []byte, offset int, length int) (interface{}, error) {
	//Check size:
	byteOffSet := offset / 8
	bitsOffset := uint((offset % 8)-1)

	bytesLength := length / 8
	if ((length % 8) > 0) {
		bytesLength++
	}

	fmt.Println(byteOffSet)
	fmt.Println(bitsOffset)

	sType := size(length)

	fmt.Println(sType)
	fmt.Println("Buffer:",buf)
	fmt.Println("Bytes Length:",bytesLength)
	fmt.Println("Bits offset:",bitsOffset)
	newBuf := buf[byteOffSet:byteOffSet + bytesLength]
	fmt.Println(newBuf)

	if (bytesLength == 1) {
		ret, err :=  Read8(newBuf)
		fmt.Println("Ret;",ret)
		ret = ret << bitsOffset
		fmt.Println("Ret;",ret)
		return ret, err
	}else if (bytesLength == 2) {
		return Read16(buf)
	}else if (bytesLength <= 4) {
		return Read32(buf)
	}else if (bytesLength <= 8) {
		return Read64(buf)
	}

	return 0,nil
}

func Read8(buf []byte) (uint8, error) {
	fmt.Println("Read 8")
	var res uint8
	reader := bytes.NewReader(buf)
	err := binary.Read(reader,binary.LittleEndian,&res)
	if (err != nil) {
		fmt.Println(err)
		return 0,err
	}
	return res, nil
}

func Read16(buf []byte) (uint16, error) {
	fmt.Println("Read 16")
	var res uint16
	reader := bytes.NewReader(buf)
	err := binary.Read(reader,binary.LittleEndian,&res)
	if (err != nil) {
		fmt.Println(err)
		return 0,err
	}
	return res, nil
}


func Read32(buf []byte) (interface{}, error) {
	fmt.Println("Read 32")
	var res uint32
	reader := bytes.NewReader(buf)
	err := binary.Read(reader,binary.LittleEndian,&res)
	if (err != nil) {
		fmt.Println(err)
		return 0,err
	}
	return res, nil
}


func Read64(buf []byte) (interface{}, error) {
	fmt.Println("Read 64")
	var res uint64
	reader := bytes.NewReader(buf)
	err := binary.Read(reader,binary.LittleEndian,&res)
	if (err != nil) {
		fmt.Println(err)
		return 0,err
	}
	return res, nil
}





func size(s int) int {

	if (s <= 8) {
		return 8
	}else if (s <= 16) {
		return 16
	}else if (s <= 32) {
		return 32
	}else if (s <= 64) {
		return 64
	}else {
		nbytes := s / 8
		mbytes := s % 8

		returnSize := nbytes * 8

		if (mbytes > 0) {
			returnSize = returnSize + 8
		}

		return returnSize

	}

}


// dataSize returns the number of bytes the actual data represented by v occupies in memory.
// For compound structures, it sums the sizes of the elements. Thus, for instance, for a slice
// it returns the length of the slice times the element size and does not count the memory
// occupied by the header. If the type of v is not acceptable, dataSize returns -1.
func dataSize(v reflect.Value) int {
	if v.Kind() == reflect.Slice {
		if s := sizeof(v.Type().Elem()); s >= 0 {
			return s * v.Len()
		}
		return -1
	}
	return sizeof(v.Type())
}

// sizeof returns the size >= 0 of variables for the given type or -1 if the type is not acceptable.
func sizeof(t reflect.Type) int {
	switch t.Kind() {
	case reflect.Array:
		if s := sizeof(t.Elem()); s >= 0 {
			return s * t.Len()
		}

	case reflect.Struct:
		sum := 0
		for i, n := 0, t.NumField(); i < n; i++ {
			s := sizeof(t.Field(i).Type)
			if s < 0 {
				return -1
			}
			sum += s
		}
		return sum

	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return int(t.Size())
	}

	return -1
}