# The Scott-Chain
The notorious blockchain that is described in the RTRG blog. To build and run, go to the blockchain folder and run:
```bash
go build
```
and then
```bash
./blockchain -listen /ip4/127.0.0.1/tcp/6666
```
On windows, this will create an exe file to execute. For other nodes, run the same command but with a different port or number at the end of the address. For example, 6668 is a pretty good one (you got a lot of options).

When you run the build command, you might be prompted that certain modules are missing. If that's the case, simply run the ```go mod``` command that the output gives until you can build. I'll figure out a good way to make that work better in the future. 

Added a DHT! Yay! 

Future Plans:
- Turn the blockchain history into a state (Merkle DAG)
- Add more info to block
- give way for miners to access transactions (not them creating them themselves)
- wallets
- Proof of Stake
    - implement as many other consensus algs as possible.
- EVM compatible