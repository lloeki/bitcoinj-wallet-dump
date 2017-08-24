# bitcoinj-wallet-dump

Dump wallets serialized by bitcoinj (e.g from Android Bitcoin app backups)

## Decrypt the wallet file

```
openssl enc -d -aes-256-cbc -a -in my_encrypted_wallet > my_decrypted_wallet
```

If unsuccessful, you'll get a `bad decrypt` error message.

## Dump the protobuf-serialized wallet

```
# set your GOPATH and PATH if necessary
go install github.com/lloeki/bitcoinj-wallet-dump
bitcoinj-wallet-dump my_decrypted_wallet
```
