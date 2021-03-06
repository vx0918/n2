/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.12
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: n2gomodule.i

package n2

/*
#cgo CXXFLAGS: -std=c++11 -lgomp -I./third_party/spdlog/include
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef _gostring_ swig_type_1;
typedef _gostring_ swig_type_2;
typedef _gostring_ swig_type_3;
typedef _gostring_ swig_type_4;
typedef _gostring_ swig_type_5;
typedef _gostring_ swig_type_6;
typedef _goslice_ swig_type_7;
typedef _goslice_ swig_type_8;
typedef _goslice_ swig_type_9;
extern void _wrap_Swig_free_n2_ab9ec64a84e85a64(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_n2_ab9ec64a84e85a64(swig_intgo arg1);
extern uintptr_t _wrap_new_HnswIndex__SWIG_0_n2_ab9ec64a84e85a64(swig_intgo arg1);
extern uintptr_t _wrap_new_HnswIndex__SWIG_1_n2_ab9ec64a84e85a64(swig_intgo arg1, swig_type_1 arg2);
extern void _wrap_delete_HnswIndex_n2_ab9ec64a84e85a64(uintptr_t arg1);
extern void _wrap_HnswIndex_Build_n2_ab9ec64a84e85a64(uintptr_t arg1, swig_intgo arg2, swig_intgo arg3, swig_intgo arg4, swig_intgo arg5, float arg6, swig_type_2 arg7, swig_type_3 arg8);
extern _Bool _wrap_HnswIndex_SaveModel_n2_ab9ec64a84e85a64(uintptr_t arg1, swig_type_4 arg2);
extern _Bool _wrap_HnswIndex_LoadModel__SWIG_0_n2_ab9ec64a84e85a64(uintptr_t arg1, swig_type_5 arg2);
extern _Bool _wrap_HnswIndex_LoadModel__SWIG_1_n2_ab9ec64a84e85a64(uintptr_t arg1, swig_type_6 arg2, _Bool arg3);
extern void _wrap_HnswIndex_UnloadModel_n2_ab9ec64a84e85a64(uintptr_t arg1);
extern void _wrap_HnswIndex_AddData_n2_ab9ec64a84e85a64(uintptr_t arg1, swig_type_7 arg2);
extern void _wrap_HnswIndex_SearchByVector__SWIG_0_n2_ab9ec64a84e85a64(uintptr_t arg1, swig_type_8 arg2, swig_intgo arg3, swig_intgo arg4, swig_voidp arg5);
extern void _wrap_HnswIndex_SearchByVector__SWIG_1_n2_ab9ec64a84e85a64(uintptr_t arg1, swig_type_9 arg2, swig_intgo arg3, swig_intgo arg4, swig_voidp arg5, swig_voidp arg6);
extern void _wrap_HnswIndex_SearchById__SWIG_0_n2_ab9ec64a84e85a64(uintptr_t arg1, swig_intgo arg2, swig_intgo arg3, swig_intgo arg4, swig_voidp arg5);
extern void _wrap_HnswIndex_SearchById__SWIG_1_n2_ab9ec64a84e85a64(uintptr_t arg1, swig_intgo arg2, swig_intgo arg3, swig_intgo arg4, swig_voidp arg5, swig_voidp arg6);
extern void _wrap_HnswIndex_PrintDegreeDist_n2_ab9ec64a84e85a64(uintptr_t arg1);
extern void _wrap_HnswIndex_PrintConfigs_n2_ab9ec64a84e85a64(uintptr_t arg1);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_n2_ab9ec64a84e85a64(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_n2_ab9ec64a84e85a64(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrHnswIndex uintptr

func (p SwigcptrHnswIndex) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrHnswIndex) SwigIsHnswIndex() {
}

func NewHnswIndex__SWIG_0(arg1 int) (_swig_ret HnswIndex) {
	var swig_r HnswIndex
	_swig_i_0 := arg1
	swig_r = (HnswIndex)(SwigcptrHnswIndex(C._wrap_new_HnswIndex__SWIG_0_n2_ab9ec64a84e85a64(C.swig_intgo(_swig_i_0))))
	return swig_r
}

func NewHnswIndex__SWIG_1(arg1 int, arg2 string) (_swig_ret HnswIndex) {
	var swig_r HnswIndex
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (HnswIndex)(SwigcptrHnswIndex(C._wrap_new_HnswIndex__SWIG_1_n2_ab9ec64a84e85a64(C.swig_intgo(_swig_i_0), *(*C.swig_type_1)(unsafe.Pointer(&_swig_i_1)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func NewHnswIndex(a ...interface{}) HnswIndex {
	argc := len(a)
	if argc == 1 {
		return NewHnswIndex__SWIG_0(a[0].(int))
	}
	if argc == 2 {
		return NewHnswIndex__SWIG_1(a[0].(int), a[1].(string))
	}
	panic("No match for overloaded function call")
}

func DeleteHnswIndex(arg1 HnswIndex) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_HnswIndex_n2_ab9ec64a84e85a64(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrHnswIndex) Build(arg2 int, arg3 int, arg4 int, arg5 int, arg6 float32, arg7 string, arg8 string) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	_swig_i_4 := arg5
	_swig_i_5 := arg6
	_swig_i_6 := arg7
	_swig_i_7 := arg8
	C._wrap_HnswIndex_Build_n2_ab9ec64a84e85a64(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2), C.swig_intgo(_swig_i_3), C.swig_intgo(_swig_i_4), C.float(_swig_i_5), *(*C.swig_type_2)(unsafe.Pointer(&_swig_i_6)), *(*C.swig_type_3)(unsafe.Pointer(&_swig_i_7)))
	if Swig_escape_always_false {
		Swig_escape_val = arg7
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg8
	}
}

func (arg1 SwigcptrHnswIndex) SaveModel(arg2 string) (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (bool)(C._wrap_HnswIndex_SaveModel_n2_ab9ec64a84e85a64(C.uintptr_t(_swig_i_0), *(*C.swig_type_4)(unsafe.Pointer(&_swig_i_1))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrHnswIndex) LoadModel__SWIG_0(arg2 string) (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (bool)(C._wrap_HnswIndex_LoadModel__SWIG_0_n2_ab9ec64a84e85a64(C.uintptr_t(_swig_i_0), *(*C.swig_type_5)(unsafe.Pointer(&_swig_i_1))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrHnswIndex) LoadModel__SWIG_1(arg2 string, arg3 bool) (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (bool)(C._wrap_HnswIndex_LoadModel__SWIG_1_n2_ab9ec64a84e85a64(C.uintptr_t(_swig_i_0), *(*C.swig_type_6)(unsafe.Pointer(&_swig_i_1)), C._Bool(_swig_i_2)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (p SwigcptrHnswIndex) LoadModel(a ...interface{}) bool {
	argc := len(a)
	if argc == 1 {
		return p.LoadModel__SWIG_0(a[0].(string))
	}
	if argc == 2 {
		return p.LoadModel__SWIG_1(a[0].(string), a[1].(bool))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrHnswIndex) UnloadModel() {
	_swig_i_0 := arg1
	C._wrap_HnswIndex_UnloadModel_n2_ab9ec64a84e85a64(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrHnswIndex) AddData(arg2 []float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_HnswIndex_AddData_n2_ab9ec64a84e85a64(C.uintptr_t(_swig_i_0), *(*C.swig_type_7)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrHnswIndex) SearchByVector__SWIG_0(arg2 []float32, arg3 int, arg4 int, arg5 *[]int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	_swig_i_4 := arg5
	C._wrap_HnswIndex_SearchByVector__SWIG_0_n2_ab9ec64a84e85a64(C.uintptr_t(_swig_i_0), *(*C.swig_type_8)(unsafe.Pointer(&_swig_i_1)), C.swig_intgo(_swig_i_2), C.swig_intgo(_swig_i_3), C.swig_voidp(_swig_i_4))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrHnswIndex) SearchByVector__SWIG_1(arg2 []float32, arg3 int, arg4 int, arg5 *[]int, arg6 *[]float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	_swig_i_4 := arg5
	_swig_i_5 := arg6
	C._wrap_HnswIndex_SearchByVector__SWIG_1_n2_ab9ec64a84e85a64(C.uintptr_t(_swig_i_0), *(*C.swig_type_9)(unsafe.Pointer(&_swig_i_1)), C.swig_intgo(_swig_i_2), C.swig_intgo(_swig_i_3), C.swig_voidp(_swig_i_4), C.swig_voidp(_swig_i_5))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (p SwigcptrHnswIndex) SearchByVector(a ...interface{}) {
	argc := len(a)
	if argc == 4 {
		p.SearchByVector__SWIG_0(a[0].([]float32), a[1].(int), a[2].(int), a[3].(*[]int))
		return
	}
	if argc == 5 {
		p.SearchByVector__SWIG_1(a[0].([]float32), a[1].(int), a[2].(int), a[3].(*[]int), a[4].(*[]float32))
		return
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrHnswIndex) SearchById__SWIG_0(arg2 int, arg3 int, arg4 int, arg5 *[]int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	_swig_i_4 := arg5
	C._wrap_HnswIndex_SearchById__SWIG_0_n2_ab9ec64a84e85a64(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2), C.swig_intgo(_swig_i_3), C.swig_voidp(_swig_i_4))
}

func (arg1 SwigcptrHnswIndex) SearchById__SWIG_1(arg2 int, arg3 int, arg4 int, arg5 *[]int, arg6 *[]float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	_swig_i_4 := arg5
	_swig_i_5 := arg6
	C._wrap_HnswIndex_SearchById__SWIG_1_n2_ab9ec64a84e85a64(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2), C.swig_intgo(_swig_i_3), C.swig_voidp(_swig_i_4), C.swig_voidp(_swig_i_5))
}

func (p SwigcptrHnswIndex) SearchById(a ...interface{}) {
	argc := len(a)
	if argc == 4 {
		p.SearchById__SWIG_0(a[0].(int), a[1].(int), a[2].(int), a[3].(*[]int))
		return
	}
	if argc == 5 {
		p.SearchById__SWIG_1(a[0].(int), a[1].(int), a[2].(int), a[3].(*[]int), a[4].(*[]float32))
		return
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrHnswIndex) PrintDegreeDist() {
	_swig_i_0 := arg1
	C._wrap_HnswIndex_PrintDegreeDist_n2_ab9ec64a84e85a64(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrHnswIndex) PrintConfigs() {
	_swig_i_0 := arg1
	C._wrap_HnswIndex_PrintConfigs_n2_ab9ec64a84e85a64(C.uintptr_t(_swig_i_0))
}

type HnswIndex interface {
	Swigcptr() uintptr
	SwigIsHnswIndex()
	Build(arg2 int, arg3 int, arg4 int, arg5 int, arg6 float32, arg7 string, arg8 string)
	SaveModel(arg2 string) (_swig_ret bool)
	LoadModel(a ...interface{}) bool
	UnloadModel()
	AddData(arg2 []float32)
	SearchByVector(a ...interface{})
	SearchById(a ...interface{})
	PrintDegreeDist()
	PrintConfigs()
}
