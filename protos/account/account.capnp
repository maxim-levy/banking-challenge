using Go = import "../go.capnp";
@0x9a4e56aa8fe168ad;
$Go.package("account");
$Go.import("protos/account");

interface AccountFactory {
	createAccount @0 (initialBalance :Int64) -> (account: Account);
	deleteAccount @1 (accountNumber: Text) -> (success: Bool);
}

struct Account {
	accountNumber @0 :Text;
	balance @1 :Int64;
}
