/*
 * Copyright © 2022 photowey (photowey@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package concurrentmap

import (
	"errors"
	"fmt"
)

var _ ConcurrentMap = (*concurrentMap)(nil)

type concurrentMap struct {
	partitions partitionGroup // partition slice
	buckets    int            // the size of partitions
}

// ----------------------------------------------------------------

func (cmap *concurrentMap) Put(key PartitionKey, value any) ConcurrentMap {
	pt := cmap.determinePartition(key)
	pt.put(key, value)

	return cmap
}

func (cmap *concurrentMap) PutKeyInt(key int, value any) ConcurrentMap {
	pk := NewIntKey(key)
	pt := cmap.determinePartition(pk)
	pt.put(pk, value)

	return cmap
}

func (cmap *concurrentMap) PutKeyInt64(key int64, value any) ConcurrentMap {
	pk := NewInt64Key(key)
	pt := cmap.determinePartition(pk)
	pt.put(pk, value)

	return cmap
}

func (cmap *concurrentMap) PutKeyString(key string, value any) ConcurrentMap {
	pk := NewStringKey(key)
	pt := cmap.determinePartition(pk)
	pt.put(pk, value)

	return cmap
}

func (cmap *concurrentMap) Get(key PartitionKey, standBy any) (any, bool) {
	pk := key
	if v, ok := cmap.determinePartition(pk).get(pk); ok {
		return v, true
	}

	return standBy, false // FIXME: return nil maybe Inaccurate
}

func (cmap *concurrentMap) GetKeyInt(key int, standBy any) (any, bool) {
	pk := NewIntKey(key)
	if v, ok := cmap.determinePartition(pk).get(pk); ok {
		return v, true
	}

	return standBy, false
}

func (cmap *concurrentMap) GetKeyInt64(key int64, standBy any) (any, bool) {
	pk := NewInt64Key(key)
	if v, ok := cmap.determinePartition(pk).get(pk); ok {
		return v, true
	}

	return standBy, false
}

func (cmap *concurrentMap) GetKeyString(key string, standBy any) (any, bool) {
	pk := NewStringKey(key)
	if v, ok := cmap.determinePartition(pk).get(pk); ok {
		return v, true
	}

	return standBy, false
}

// GetString - get the value from cmap as string,
// but not recommended unless you can determine the type
func (cmap *concurrentMap) GetString(key PartitionKey) (string, bool) {
	pk := key
	if v, ok := cmap.determinePartition(pk).get(pk); ok {
		return String(v), true // String() -> fmt.Sprintf("%v", src)
	}

	return emptyString, false
}

func (cmap *concurrentMap) Remove(key PartitionKey) ConcurrentMap {
	pk := key
	km := cmap.determinePartition(pk)
	km.remove(pk)

	return cmap
}

func (cmap *concurrentMap) RemoveKeyInt(key int) ConcurrentMap {
	pk := NewIntKey(key)
	km := cmap.determinePartition(pk)
	km.remove(pk)

	return cmap
}

func (cmap *concurrentMap) RemoveKeyInt64(key int64) ConcurrentMap {
	pk := NewInt64Key(key)
	km := cmap.determinePartition(pk)
	km.remove(pk)

	return cmap
}

func (cmap *concurrentMap) RemoveKeyString(key string) ConcurrentMap {
	pk := NewStringKey(key)
	km := cmap.determinePartition(pk)
	km.remove(pk)

	return cmap
}

func (cmap *concurrentMap) Has(key PartitionKey) bool {
	pk := key
	km := cmap.determinePartition(pk)
	_, ok := km.get(pk)

	return ok
}

func (cmap *concurrentMap) HasKeyInt(key int) bool {
	pk := NewIntKey(key)
	km := cmap.determinePartition(pk)
	_, ok := km.get(pk)

	return ok
}

func (cmap *concurrentMap) HasKeyInt64(key int64) bool {
	pk := NewInt64Key(key)
	km := cmap.determinePartition(pk)
	_, ok := km.get(pk)

	return ok
}

func (cmap *concurrentMap) HasKeyString(key string) bool {
	pk := NewStringKey(key)
	km := cmap.determinePartition(pk)
	_, ok := km.get(pk)

	return ok
}

func (cmap *concurrentMap) UnsafeLength() int64 {
	length := cmap.length()

	return length
}

func (cmap *concurrentMap) Length() int64 {
	panic("not supported now")
}

// ----------------------------------------------------------------

func (cmap *concurrentMap) GetBool(key PartitionKey) (bool, bool) {
	panic("not supported now")
}

func (cmap *concurrentMap) GetInt(key PartitionKey) (int, bool) {
	panic("not supported now")
}

func (cmap *concurrentMap) GetInt64(key PartitionKey) (int64, bool) {
	panic("not supported now")
}

func (cmap *concurrentMap) GetFloat64(key PartitionKey) (float64, bool) {
	panic("not supported now")
}

// ----------------------------------------------------------------

func (cmap *concurrentMap) determinePartition(key PartitionKey) *partition {
	partitionId := key.HashCode() & (Int64(cmap.buckets) - 1) // %

	return cmap.partitions[partitionId]
}

func (cmap *concurrentMap) length() int64 {
	var length int64
	for _, pt := range cmap.partitions {
		pt.lock.RLock()
		length += Int64(len(pt.ctx))
		pt.lock.RUnlock()
	}

	return length
}

// ----------------------------------------------------------------

func NewConcurrentMap(capacity int) (*concurrentMap, error) {
	if capacity < 2 {
		return nil, errors.New("capacity must be an integer multiple of 2")
	}
	zero := capacity & (capacity - 1)
	if 0 != zero {
		return nil, errors.New("capacity must be an integer multiple of 2")
	}

	var partitions partitionGroup
	for i := 0; i < capacity; i++ {
		partitions = append(partitions, createPartition())
	}

	return &concurrentMap{
		partitions: partitions,
		buckets:    capacity,
	}, nil
}

// ----------------------------------------------------------------

func String(src any) string {
	return fmt.Sprintf("%v", src)
}

func Int64(src int) int64 {
	return int64(src)
}
