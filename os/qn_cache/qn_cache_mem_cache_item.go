// Copyright 2018 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_cache

import "github.com/qnsoft/common/os/qn_time"

// IsExpired checks whether <item> is expired.
func (item *memCacheItem) IsExpired() bool {
	// Note that it should use greater than or equal judgement here
	// imaqn_intng that the cache time is only 1 millisecond.
	if item.e >= qn_time.TimestampMilli() {
		return false
	}
	return true
}
