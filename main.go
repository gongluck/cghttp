/*
 * @Author: gongluck
 * @Date: 2020-06-12 11:18:45
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-12 16:09:08
 */

package main

import (
	"io/ioutil"
	"net/http"
	"reflect"
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

	tmp, _ := ioutil.ReadAll(respon.Body)
	*bodylen = (C.size_t)(len(tmp))
	*body = (*C.char)(C.malloc(C.size_t(*bodylen)))
	C.memcpy(unsafe.Pointer(*body), unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&tmp)).Data), *bodylen)

	return C.int(respon.StatusCode)
}

func main() {
	// var body *C.char
	// var bodylen C.size_t
	// C.Get(C.CString("http://www.gongluck.icu/web/"), &body, &bodylen)
	// C.puts(body)
	// C.Release(&body)
}
