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

//
// Use custom concurrent map instead of sync.map
//

const (
	emptyString = ""
)

// ConcurrentMap - custom concurrent map as cmap instead of sync.map
type ConcurrentMap interface {
	// Put - put the value(any) into cmap by partitionKey wrapped
	Put(key PartitionKey, value any) ConcurrentMap
	PutKeyInt(key int, value any) ConcurrentMap
	PutKeyInt64(key int64, value any) ConcurrentMap
	PutKeyString(key string, value any) ConcurrentMap

	// Get - get a value with default (standBy) from cmap  by partitionKey
	Get(key PartitionKey, standBy any) (any, bool)
	GetKeyInt(key int, standBy any) (any, bool)
	GetKeyInt64(key int64, standBy any) (any, bool)
	GetKeyString(key string, standBy any) (any, bool)

	// GetString - get a value as string from cmap by partitionKey
	GetString(key PartitionKey) (string, bool)

	// Remove - remove the value from cmap by partitionKey
	Remove(key PartitionKey) ConcurrentMap
	RemoveKeyInt(key int) ConcurrentMap
	RemoveKeyInt64(key int64) ConcurrentMap
	RemoveKeyString(key string) ConcurrentMap

	// Has - check if a key exists in cmap
	Has(key PartitionKey) bool
	HasKeyInt(key int) bool
	HasKeyInt64(key int64) bool
	HasKeyString(key string) bool
	// UnsafeLength - calculate the length of cmap, but maybe unsafe
	UnsafeLength() int64
	// Length - calculate the length of cmap, but not supported now
	Length() int64

	// TODO: not implement

	GetBool(key PartitionKey) (bool, bool) // real, ok
	GetInt(key PartitionKey) (int, bool)
	GetInt64(key PartitionKey) (int64, bool)
	GetFloat64(key PartitionKey) (float64, bool)
}
