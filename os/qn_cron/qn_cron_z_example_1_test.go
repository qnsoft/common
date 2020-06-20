// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

package qn_cron_test

import (
	"time"

	"github.com/qnsoft/common/os/gcron"
	"github.com/qnsoft/common/os/qn_log"
)

func Example_cronAddSingleton() {
	gcron.AddSingleton("* * * * * *", func() {
		qn_log.Println("doing")
		time.Sleep(2 * time.Second)
	})
	select {}
}
