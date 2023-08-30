## Take-Home Assignment: Transaction Broadcasting and Monitoring Client Module

### Objective:
Your task is to develop a client module in either Go or Rust that interacts with an HTTP server. This client module will enable the broadcasting of a transaction and subsequently monitor its status until finalization.

### Requirements:
Create a client module in Go or a client crate in Rust that is designed to be integrated into another application, with the following capabilities:

1. **Broadcast Transaction**: Construct a JSON payload and send a POST request to `https://mock-node-wgqbnxruha-as.a.run.app/broadcast`. The payload structure is as follows:

    ```json
    {
        "symbol": "string",          // Transaction symbol, e.g., BTC
        "price": uint64,             // Symbol price, e.g., 100000
        "timestamp": uint64          // Timestamp of price retrieval
    }
    ```

    Example payload and server response:
    ```json
    // Payload
    {
        "symbol": "ETH",
        "price": 4500,
        "timestamp": 1678912345
    }
    
    // Server response
    {
        "tx_hash": "e97ae379d666eed7576bf285c451baf9f0edba93ce718bf15b06c8a85d07b8d1"
    }
    ```

2. **Transaction Status Monitoring**: Utilize the transaction hash obtained from the response to periodically issue GET requests to `https://mock-node-wgqbnxruha-as.a.run.app/check/<tx_hash>`. The response will be plaintext indicating the transaction status, which can be one of the following:
   - `CONFIRMED`: Transaction has been processed and confirmed
   - `FAILED`: Transaction failed to process
   - `PENDING`: Transaction is awaiting processing
   - `DNE`: Transaction does not exist

   An example response is shown below:
   ```json
   {
       "tx_status": "CONFIRMED"
   }
   ```
   
3. **Documentation**: Provide clear documentation and a example script that implements your client module. The documentation should explain how to use and integrate it, along with your strategies for handling each transaction status.

### Submission:
Choose one of the following submission methods:

1. **Zip Package**: Create a zip file containing your client module's code and the README.
2. **GitHub Repository**: Establish a GitHub repository for your client module, including both the code and README.

Please feel free to reach out if you need any clarifications or have inquiries. We look forward to reviewing your submission!
