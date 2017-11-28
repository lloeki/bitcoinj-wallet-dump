# bitcoinj-wallet-dump

Dump wallets serialized by bitcoinj (e.g from Android Bitcoin app backups)

## Decrypt the wallet file

```
openssl enc -d -aes-256-cbc -md md5 -a -in my_encrypted_wallet > my_decrypted_wallet
```

If unsuccessful, you'll get a `bad decrypt` error message.

## Dump the protobuf-serialized wallet

```
# set your GOPATH and PATH if necessary ie:
export GOPATH="$HOME/where/you/store/your/go/projects"
export PATH="$GOPATH/bin:$PATH"

# tell Go to get the code to bitcoinj-wallet-dump
go get github.com/lloeki/bitcoinj-wallet-dump

# run bitcoinj-wallet-dump on your wallet
bitcoinj-wallet-dump my_decrypted_wallet
```
