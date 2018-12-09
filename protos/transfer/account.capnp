using Go = import "../go.capnp";
@0x8e7192767ff40215;
$Go.package("transfer");
$Go.import("protos/transfer");

interface TransferFactory {
	transferFunds @0 (sourceAccount :Text, destinationAccount :Text, amount :UInt64) -> (account: Account);
}

struct Account {
	accountNumber @0 :Text;
	balance @1 :Int64;
}
