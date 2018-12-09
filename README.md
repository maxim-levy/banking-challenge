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

## Client

## Server

## TODO's
* Add redundancy for the server-side.
* Add auth layer to controll who can modify accounts and balances.
* Vendor client and server applications
* ~~Add data persitance using bolt db~~
* ~~Add create account method~~
* ~~Add delete account method~~
* Add transfer funds method
* Add list accounts method
