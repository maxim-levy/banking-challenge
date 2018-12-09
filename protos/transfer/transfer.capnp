using Go = import "../go.capnp";
@0x8e7192767ff40215;
$Go.package("transfer");
$Go.import("protos/transfer");

interface TransferFactory {
	transferFunds @0 (sourceAccount :Text, destinationAccount :Text, amount :UInt64) -> (record: TransactionalRecord);
}

struct TransactionalRecord {
	sourceAccount @0 :Text;
	sourceBalance @1 :Int64;
	destinationAccount @2 :Text;
	destinationBalance @3 :Int64;
	amount @4 :UInt64;
}
