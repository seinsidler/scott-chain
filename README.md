# The Scott-Chain
The notorious blockchain that is described in the RTRG blog. To build and run, go to the blockchain folder and run:
```bash
go build
```
and then
```bash
./blockchain -l 10000 -secio
```

TODO:

Fixes:
- The carrot in the console needs to show up when a block is found on another terminal
- automatic node discovery or persistent node discovery
- network needs to not shut down when one node fails!

Future Plans:
- Turn the blockchain history into a state (Merkle DAG)
- Add more info to block
- give way for miners to access transactions (not them creating them themselves)
- wallets
- Proof of Stake
    - implement as many other consensus algs as possible.
- EVM compatible