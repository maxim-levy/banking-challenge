using Go = import "../go.capnp";
@0x9a4e56aa8fe168ad;
$Go.package("account");
$Go.import("protos/account");

interface AccountFactory {
	createAccount @0 (initialBalance :Int64) -> (account: Account);
	deleteAccount @1 (sourceAccount: Text) -> (success: Bool);
}

struct Account {
	sourceAccount @0 :Text;
	initialBalance @1 :Int64;
}
