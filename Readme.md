# Transaction Broadcasting and Monitoring Client Module

A boardcasting and monitoring transaction go client module.

## Install

clone this repository.
``` bash
git clone github.com/devbyP/bt
```

cd to your project, then add replace directive to your go.mod.
```
module your-project-name

go 1.20

replace github.com/devbyP/bt => path/to/bt/module
```
or you can use this go command.
``` bash
go mod edit -replace github.com/devbyP/bt=../path/to/bt/module
```

import the module in your go file in your project.
``` go
package main

import "github.com/devbyP/bt"
```
then run
``` bash
go mod tidy
```

## Getting start

you can boardcast a transaction or monitoring transaction by calling function
directly from a package module.

### Boardcast a transaction

use function `BoardcastTransaction(BoardcasePayload)` to boardcast a transaction.
`BoardcastTransaction` receive `BoardcastPayload` as an argument.

example
``` go
package main

import (
	"log"
	"time"

	"github.com/devbyP/bt"
)

func main() {
	hash, err := bt.BoardcastTransaction(bt.BoardcastPayload{
		Symbol: "BTC",
		Price:  100000,
		// use unix time and convert to uint64
		Timestamp: uint64(time.Now().Unix()),
	})
	if err != nil {
		log.Fatal(err)
	}
	// do something with the returned hash.
}
```

It return a hash of the transaction as `string` type.

### Monitoring a transaction (get status)

you can check a statas of the transaction by using hashes string.

example
``` go
package main

import (
	"log"

	"github.com/devbyP/bt"
)

func main() {
    // ...get a hash

	status, err := bt.GetTransactionStatus(hash)
	if err != nil {
		log.Fatal(err)
	}
    // do something.
}
```

`bt.GetTransactionStatus(hash)` receive "hash" of `string` type
and return "status" of `string` type.

the module also provided Status constant for you to check status and handle each scenario.

``` go
package main

import (
	"log"

	"github.com/devbyP/bt"
)

func main() {
    // ...get a hash.
	status, err := bt.GetTransactionStatus(hash)
	if err != nil {
		log.Fatal(err)
	}
	// handle code by status.
	switch status {
	case bt.StatusConfirmed:
		// handle confirmed transaction.
	case bt.StatusPending:
		// handle pending transaction.
	case bt.StatusFailed:
		// handle failed transaction.
	case bt.StatusDoesNotExist:
		// handle non exist transaction.
	default:
		// handle unknown status.
	}
    // do something.
}
```
