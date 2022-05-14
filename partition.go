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
	"sync"
)

type partitionGroup []*partition

type partition struct {
	ctx  map[any]any
	lock sync.RWMutex
}

// ----------------------------------------------------------------

func createPartition() *partition {
	return &partition{
		ctx: make(map[any]any),
	}
}

func (pt *partition) get(key PartitionKey) (any, bool) {
	keyValue := key.Value()
	pt.lock.RLock() // read lock
	defer pt.lock.RUnlock()
	v, ok := pt.ctx[keyValue]

	return v, ok
}

func (pt *partition) put(key PartitionKey, v any) {
	keyValue := key.Value()
	pt.lock.Lock()
	defer pt.lock.Unlock()
	pt.ctx[keyValue] = v
}

func (pt *partition) remove(key PartitionKey) {
	keyValue := key.Value()
	pt.lock.Lock()
	defer pt.lock.Unlock()
	delete(pt.ctx, keyValue)
}
