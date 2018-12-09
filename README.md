# quoine-banking
Interview test for quoine

## Install
You will need to install the cap'n libs inorder to edit the protocols.
First install capnp for your platform.
https://capnproto.org/install.html

After that you need to install the go code generator plugin for capnp,
you can find this code here.
https://godoc.org/zombiezen.com/go/capnproto2

You will need to use go get, then after generate the bin file using this command (ajust to your linking).
`go build zombiezen.com/go/capnproto2/capnpc-go`
this will place a file named `capnpc-go` in `$GOPATH/bin`
you will now need to move this file in to your `/usr/local/go/bin`
dir so the `capnpc` command can reach it (or you can add it directly to your $PATH).

After this you should be able to run `./build.sh` with no warnings.

## ENV
You can change the default settings using ENV variables.

Variable | Default value | Description
--- | --- | ---
SERVER_ADDR | 127.0.0.1 | The address the server should bind to.
SERVER_PORT | 8080 | The port the server should run on.
CLIENT_SERVER_ADDR | 127.0.0.1 | Address of the server the client should connect to.
CLIENT_SERVER_PORT | 8080 | Port of the server the client should connect to.

## Client
The client is a CLI.
You can run this client directly by cloning the repo and running `go run main.go` from within the client folder.

COMMANDS
```
create-account, ca  Register an new account
delete-account, da  Delete an existing account
transfer-funds, tf  Transfer funds from one account to another
help, h             Shows a list of commands or help for one command
```

For and extended list of options, please type the command you want to execute and apped `--help`,
This will bring up more info about this command.

## Server

## Tests
Both the server and client folder has their own separate tests.
You can run them by using the following command `go test ./...` in the client or server folder.

## TODO's
* ~~Add support for ENV runtime configuration.~~
* Add redundancy for the server-side.
* Add auth layer to controll who can modify accounts and balances.
* Vendor client and server applications
* ~~Write unit-tests.~~
* ~~Add data persitance using bolt db~~
* ~~Add create account method~~
* ~~Add delete account method~~
* ~~Add transfer funds method~~
* Add list accounts method
