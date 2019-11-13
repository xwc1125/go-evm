// description: XChain 
// 
// @author: xwc1125
// @date: 2019/9/30
package hashmap

import (
	"reflect"
	"sort"
	"sync"
)

type HashMapInterface interface {
	Ini(args ...string);
	Put(k string, v interface{});
	Get(k string) (interface{}, string, bool);
	ContainsKey(k string) (bool, string);
	Remove(k string) (interface{}, bool);
	ForEach() map[string]interface{};
	Size() int;
	Sort() map[string]interface{};
}

//hashmap 只能单协程中运行 比起HashMapConcurrent执行效率更高。
type HashMap struct {
	m map[string]interface{}
}

func New() *HashMap {
	hashMap := new(HashMap)
	hashMap.Ini()
	return hashMap
}

//必须初始化
func (this *HashMap) Ini(args ...string) {
	this.m = map[string]interface{}{}
}

//增加或者修改一个元素
func (this *HashMap) Put(k string, v interface{}) {
	this.m[k] = v;
}

//返回值，返回值类型，是否有返回
func (this *HashMap) Get(k string) (interface{}, string, bool) {
	v, cb := this.m[k];
	var rv interface{} = nil;
	var rt string = "";
	var rs bool = false;
	if (cb) {
		rv = v;
		rs = true;
		rt = reflect.TypeOf(v).String();
	}
	return rv, rt, rs;
}

//判断是否包括key，如果包含key返回value的类型
func (this *HashMap) ContainsKey(k string) (bool, string) {
	v, cb := this.m[k];
	var rs bool = false;
	var rt string = "";
	if (cb) {
		rs = true;
		rt = reflect.TypeOf(v).String();
	}
	return rs, rt;
}

//移除一个元素
func (this *HashMap) Remove(k string) (interface{}, bool) {
	v, cb := this.m[k];
	var rs bool = false;
	var rv interface{} = nil;
	if (cb) {
		rv = v;
		rs = true;
		delete(this.m, k);
	}
	return rv, rs;
}

//复制map用于外部遍历
func (this *HashMap) ForEach() map[string]interface{} {
	mb := map[string]interface{}{};
	for k, v := range this.m {
		mb[k] = v;
	}
	return mb;
}

//放回现在的个数
func (this *HashMap) Size() int {
	return len(this.m)
}

//排序
func (this *HashMap) Sort() map[string]interface{} {
	newm := map[string]interface{}{};
	var keyArray []string;
	for k, _ := range this.m {
		keyArray = append(keyArray, k);
	}
	sort.Strings(keyArray);
	for _, v := range keyArray {
		newm[v] = this.m[v];
	}
	return newm;
}

//并发hashmapConcurrent 多协程，使用安全
type HashMapConcurrent struct {
	m    map[string]interface{}
	lock *sync.Mutex
}

//初始化
func (this *HashMapConcurrent) Ini(args ...string) {
	this.m = map[string]interface{}{}
	this.lock = new(sync.Mutex);
}

//加入或修改
func (this *HashMapConcurrent) Put(k string, v interface{}) {
	this.lock.Lock();
	this.m[k] = v;
	this.lock.Unlock();
}

//返回值，返回值类型，是否有返回
func (this *HashMapConcurrent) Get(k string) (interface{}, string, bool) {
	this.lock.Lock();
	v, cb := this.m[k];
	var rv interface{} = nil;
	var rt string = "";
	var rs bool = false;
	if (cb) {
		rv = v;
		rs = true;
		rt = reflect.TypeOf(v).String();
	}
	this.lock.Unlock();
	return rv, rt, rs;
}

//判断是否包括key，如果包含key返回value的类型
func (this *HashMapConcurrent) ContainsKey(k string) (bool, string) {
	this.lock.Lock();
	v, cb := this.m[k];
	var rs bool = false;
	var rt string = "";
	if (cb) {
		rs = true;
		rt = reflect.TypeOf(v).String();
	}
	this.lock.Unlock();
	return rs, rt;
}

//移除一个对象
func (this *HashMapConcurrent) Remove(k string) (interface{}, bool) {
	this.lock.Lock();
	v, cb := this.m[k];
	var rs bool = false;
	var rv interface{} = nil;
	if (cb) {
		rv = v;
		rs = true;
		delete(this.m, k);
	}
	this.lock.Unlock();
	return rv, rs;
}

//复制map用于外部遍历
func (this *HashMapConcurrent) ForEach() map[string]interface{} {
	this.lock.Lock();
	mb := map[string]interface{}{};
	for k, v := range this.m {
		mb[k] = v;
	}
	this.lock.Unlock();
	return mb;
}

//返回个数
func (this *HashMapConcurrent) Size() int {
	this.lock.Lock();
	s := len(this.m);
	this.lock.Unlock();
	return s;
}

//排序
func (this *HashMapConcurrent) Sort() map[string]interface{} {
	newm := map[string]interface{}{};
	this.lock.Lock();
	var keyArray []string;
	for k, _ := range this.m {
		keyArray = append(keyArray, k);
	}
	sort.Strings(keyArray);
	for _, v := range keyArray {
		newm[v] = this.m[v];
	}
	this.lock.Unlock();
	return newm;
}
