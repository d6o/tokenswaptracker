# TokenSwapTracker

TokenSwapTracker is a tool that connects to any Ethereum-like blockchain, subscribes to new blocks, and scans all
transactions for ERC20 token swaps. When a swap is found, it prints detailed information about the swap.

## Features

- Connects to any Ethereum-like blockchain via WebSocket.
- Subscribes to new blocks in real-time.
- Scans all transactions in each block for ERC20 token swaps.
- Logs detailed information about each swap found.

## Requirements

- Node.js
- An environment variable `ETH_WS` set with a valid WebSocket connection to an Ethereum-like node.

## Installation

1. Clone the repository:

```bash
git clone https://github.com/d6o/tokenswaptracker.git
```

2. Change into the project directory:

```bash
cd tokenswaptracker
```

## Configuration

Ensure that you have the `ETH_WS` environment variable set with a valid WebSocket connection URL to your Ethereum-like
node. You can set this environment variable in your terminal session as follows:

```bash
export ETH_WS=wss://your-node-url
```

## Usage

To start the TokenSwapTracker, run the following command:

```bash
go run tokenswapwatcher.go
```

## Output

The output logs will look like this:

```plaintext
time=2024-06-30T17:06:51.121-03:00 level=INFO msg="Found new block" block_number=20206784
time=2024-06-30T17:06:52.116-03:00 level=INFO msg="Swap found" block_number=20206784 txn=0x28a2dea5ba9bf817810e9586023f233a40182558f9f4ba7ab04108bb5690ecee swapIndex=4 sender=0x3328F7f4A1D1C57c35df56bBf0c9dCAFCA309C49 to=0xD90DC302A451E10C68BBBD0e2Ddc5ddB7BF210C2 token0_address=0x07E6d1A3839c5D037d66dabe4fFc59a1DAb77631 token1_address=0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2 token0_symbol=ANDEI token1_symbol=WETH amount0_in=0 amount1_in=995024875621890547 amount0_out=2706282388548402081 amount1_out=0
time=2024-06-30T17:06:53.338-03:00 level=INFO msg="Swap found" block_number=20206784 txn=0x01946f505753922425ecfa85492cc97d2bbb4c0c7b383e5f0b1249613634ac2e swapIndex=4 sender=0x3328F7f4A1D1C57c35df56bBf0c9dCAFCA309C49 to=0x92D23D04D97bfF74eE5278c03482AF85Bb79F039 token0_address=0x07E6d1A3839c5D037d66dabe4fFc59a1DAb77631 token1_address=0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2 token0_symbol=ANDEI token1_symbol=WETH amount0_in=0 amount1_in=995024875621890547 amount0_out=2171210719502646335 amount1_out=0
...
```

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/YourFeature`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Open a Pull Request.

## License

This project is licensed under the MIT License.

---

Feel free to reach out if you have any questions or need further assistance. Happy coding!
