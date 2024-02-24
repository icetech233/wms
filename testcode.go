package main

import (
	"fmt"
	"time"
	"wms/util"
)

func testCopy() {

	_age := 666
	src := StructB{DateTime: time.Now(), Age: _age}
	dst := StructA{}
	// go copy(dst Writer, src Reader)
	for i := 0; i < 10; i++ {
		_age = _age + 1
		src.Age = _age
		src.DateTime = time.Now()
		t1 := time.Now()
		er2 := util.CopyStruct(&dst, &src)
		if er2 != nil {
			panic(er2)
		}
		fmt.Println("复制耗时:", time.Since(t1))
		fmt.Println("结果", src, dst)
	}
}
