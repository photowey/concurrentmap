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

type IntKey struct {
	value int
}

func (key IntKey) Value() any {
	return key.value
}

func (key IntKey) HashCode() int64 {
	return int64(key.value) // value as hash code
}

func NewIntKey(src int) IntKey {
	return IntKey{src}
}
