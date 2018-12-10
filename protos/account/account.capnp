using Go = import "../go.capnp";
@0x9a4e56aa8fe168ad;
$Go.package("account");
$Go.import("protos/account");

interface AccountFactory {
	createAccount @0 (initialBalance :Int64) -> (account: Account);
	deleteAccount @1 (accountNumber: Text) -> (success: Bool);
	transferFunds @2 (sourceAccount :Text, destinationAccount :Text, amount :UInt64) -> (record: TransactionalRecord);
}

struct Account {
	accountNumber @0 :Text;
	balance @1 :Int64;
}

struct TransactionalRecord {
	sourceAccount @0 :Text;
	sourceBalance @1 :Int64;
	destinationAccount @2 :Text;
	destinationBalance @3 :Int64;
	amount @4 :UInt64;
}
