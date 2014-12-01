# CenturyLink SDK

This project is a Go package that provides interfaces to CenturyLink Cloud API v2.0. 

## How does it help?

The main goal of CenturyLink SDK is to create wrappers for all methods in CenturyLink REST API, provide developers with access to strongly typed models, instead of raw json messages, take care of all details, such as setting up all necessary headers and handling authorization, as well as provide developers with some very usefully additional features, such as logging and link resolution.

## Who can benefit from it?

This project is mainly created for Go developers, who want to use CenturyLinks Cloud platform from there application.

## Using CenturyLink SDK: Prerequisites

In order to use this project [Go](https://golang.org/doc/install) should be installed. Also, please, note that in order to use go tools, GOPATH environment variable should be set (look [here](https://golang.org/doc/code.html) for more details how to set GOPATH) 

## Getting CenturyLink SDK 

The easiest way to get CenturyLink SDK is to use `go get` command.

	go get github.com/s-matyukevich/centurylink_sdk

This command will checkout the source of `centurylink_sdk` and inspect it for any unmet Go package dependencies, downloading those as well. `go get` will also build and install `centurylink_sdk` and its dependencies.

## Using CenturyLink SDK

The following program demonstrate the basic usage of `centurylink_sdk`

	package main

	import (
		sdk "github.com/s-matyukevich/centurylink_sdk"
	)

	func main() {
		client := sdk.NewClient()
		err := client.Connect("username", "password")
		if err != nil {
			return
		}
		client.GetGroup("someGroup")
	}
Basically, the process of working with `centurylink_sdk` is the following

**Import `centurylink_sdk` package**  
	
	import (
		sdk "github.com/s-matyukevich/centurylink_sdk"
	)

**Create new client.** 

	client := sdk.NewClient()

**Connect the client**

	err := client.Connect("username", "password")

After this line is executed `centurylink_sdk` sends login request and if it succeeded sdk stores internally account alias and bearer token. After this, it can use them for sending all other requests. 
Alternatively, if you already have you bearer token, you can create client, that is already connected

	client = sdk.NewClientInitialized("accountAlias", "bearerToken")

**Start sending requests to API**

For each documented method in CenturyLink API v2.0 we have corresponding method in `client` struct. This methods returns strongly typed objects, that are created by unmarshalling json responses. 
Also, if response contains links collection, for each link in response class separate method is created, that, when executed, sends request to link href and parse response. 
Another important feature, you should be aware of, is Logging. by default `centurylink_sdk` skips all logs, but you can set logger to client at any time. The following example shows how to configure client to set logs to standard output.

	client.SetLogger(log.New(os.Stdout, "", log.LstdFlags))

## Testing

Tests are written using [gockeck](https://labix.org/gocheck) library. You can install it with the following command

	go get gopkg.in/check.v1

After this package is installed, you can use `go test github.com/s-matyukevich/centurylink_sdk/...` command to execute all tests.

## Current Limitations

- Currently, not all API is fully covered with integration tests. You can see all existing integration tests in api_test.go
- Currently, not for all types of links methods that resolves them have been implemented. This is mostly because there is no documentation for the responses, that are returned, when executing some types of links.

## Collaborate

You are welcome to contribute via
[pull request](https://help.github.com/articles/using-pull-requests).



