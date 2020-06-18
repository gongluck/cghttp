/*
 * @Author: gongluck
 * @Date: 2020-06-12 11:18:45
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-12 16:09:08
 */

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"sync"
	"time"
	"unsafe"
)

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "cghttp.h"
*/
import "C"

//export Release
func Release(data **C.char) {
	C.free(unsafe.Pointer(*data))
	*data = nil
}

//export Get
func Get(url *C.char, body **C.char, bodylen *C.size_t) C.int {
	var str string
	var strHeader = (*reflect.StringHeader)(unsafe.Pointer(&str))
	strHeader.Data = uintptr(unsafe.Pointer(url))
	strHeader.Len = int(C.strlen(url))

	respon, err := http.Get(str)
	if err != nil {
		return -1
	}

	if body != nil || bodylen != nil {
		tmp, _ := ioutil.ReadAll(respon.Body)
		if bodylen != nil {
			*bodylen = (C.size_t)(len(tmp))
		}
		if body != nil {
			bodysize := (C.size_t)(len(tmp))
			*body = (*C.char)(C.malloc(bodysize))
			C.memcpy(unsafe.Pointer(*body), unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&tmp)).Data), bodysize)
		}
	}

	return C.int(respon.StatusCode)
}

func test_get() {
	var body *C.char
	var bodylen C.size_t
	C.Get(C.CString("https://www.baidu.com"), &body, &bodylen)
	//C.puts(body)
	C.Release(&body)
	wg.Done()
}

var (
	TESTTIMES = 1000
	wg        sync.WaitGroup
)

func main() {
	t1 := time.Now()
	for i := 0; i < TESTTIMES; i++ {
		wg.Add(1)
		go test_get()
	}
	wg.Wait()
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}
