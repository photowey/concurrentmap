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

type Int64Key struct {
	value int64
}

func (key Int64Key) Value() any {
	return key.value
}

func (key Int64Key) HashCode() int64 {
	return key.value // value as hash code
}

func NewInt64Key(src int64) Int64Key {
	return Int64Key{src}
}
