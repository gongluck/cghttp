/*
 * @Author: gongluck
 * @Date: 2020-06-12 11:18:45
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-18 17:27:18
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

static char* parse(char** datas, size_t index)
{
    return datas[index];
}
*/
import "C"

func ParseString(cstring *C.char) string {
	var str string
	var strHeader = (*reflect.StringHeader)(unsafe.Pointer(&str))
	strHeader.Data = uintptr(unsafe.Pointer(cstring))
	strHeader.Len = int(C.strlen(cstring))
	return str
}

//export Release
func Release(data **C.char) {
	C.free(unsafe.Pointer(*data))
	*data = nil
}

//export Get
func Get(url *C.char, body **C.char, bodylen *C.size_t) C.int {
	str := ParseString(url)

	respon, err := http.Get(str)
	if err != nil {
		return -1
	}

	if body != nil || bodylen != nil {
		tmp, _ := ioutil.ReadAll(respon.Body)
		defer respon.Body.Close()
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

//export Post
func Post(posturl *C.char, keys **C.char, values **C.char, keynum C.size_t, response **C.char, responselen *C.size_t) C.int {
	str := ParseString(posturl)

	forms := make(map[string][]string)
	for i := 0; (C.size_t)(i) < keynum; i++ {
		value := make([]string, 1)
		value[0] = ParseString(C.parse(values, (C.size_t)(i)))
		forms[ParseString(C.parse(keys, (C.size_t)(i)))] = value
	}

	respon, err := http.PostForm(str, forms)
	if err != nil {
		return -1
	}

	if response != nil || responselen != nil {
		tmp, _ := ioutil.ReadAll(respon.Body)
		defer respon.Body.Close()
		if responselen != nil {
			*responselen = (C.size_t)(len(tmp))
		}
		if response != nil {
			bodysize := (C.size_t)(len(tmp))
			*response = (*C.char)(C.malloc(bodysize))
			C.memcpy(unsafe.Pointer(*response), unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&tmp)).Data), bodysize)
		}
	}

	return C.int(respon.StatusCode)
}

var (
	TESTTIMES = 1000
	wg        sync.WaitGroup
)

func test_get() {
	var body *C.char
	var bodylen C.size_t
	C.Get(C.CString("https://www.baidu.com"), &body, &bodylen)
	//C.puts(body)
	C.Release(&body)
	wg.Done()
}

func test_post() {
	key := C.CString("title")
	value := C.CString("testpost")
	var body *C.char
	var bodylen C.size_t
	C.Post(C.CString("http://47.115.57.81/api/postvideo"), &key, &value, 1, &body, &bodylen)
	//C.puts(body)
	C.Release(&body)
	wg.Done()
}

func main() {
	t1 := time.Now()
	for i := 0; i < TESTTIMES; i++ {
		wg.Add(1)
		go test_get()
	}
	wg.Wait()
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))

	t1 = time.Now()
	for i := 0; i < TESTTIMES; i++ {
		wg.Add(1)
		go test_post()
	}
	wg.Wait()
	t2 = time.Now()
	fmt.Println(t2.Sub(t1))
}
