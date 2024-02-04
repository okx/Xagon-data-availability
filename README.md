# X1 Data Availability

### Data Availability Layer for X1 Validium

The x1-data-availability project is a specialized Data Availability Node (DA Node) that is part of X1's CDK (Chain Development Kit) Validium.

## Overview of Validium
The X1 Validium solution is made up of several components; start with the [X1 Node](https://github.com/okx/x1-node). For quick reference, the complete list of components are outlined below:

| Component                                                                     | Description                                                          |
| ----------------------------------------------------------------------------- | -------------------------------------------------------------------- |
| [X1 Node](https://github.com/okx/x1-node)           | Node implementation for the X1 networks in Validium mode            |
| [X1 Contracts](https://github.com/okx/x1-contracts) | Smart contract implementation for the X1 networks in Validium mode |
| [X1 Data Availability](https://github.com/okx/x1-data-availability)   | Data availability implementation for the X1 networks          |
| [Prover / Executor](https://github.com/okx/x1-prover)          | zkEVM engine and prover implementation                               |
| [Bridge Service](https://github.com/okx/x1-bridge-service)     | Bridge service implementation for X1 networks                       |

---

## Introduction

As blockchain networks grow, the volume of data that needs to be stored and validated increases, posing challenges in scalability and efficiency. Storing all data on-chain can lead to bloated blockchains, slow transactions, and high fees.

Data Availability Nodes facilitate a separation between transaction execution and data storage. They allow transaction data to reside off-chain while remaining accessible for validation. This significantly improves scalability and reduces costs. Within the framework of X1's CDK, Data Availability Committees (DAC) members run DA nodes to ensure the security, accessibility, and reliability of off-chain data.

To learn more about how the data availability layer works in the validium, please see the CDK documentation [here](https://wiki.polygon.technology/docs/cdk/dac-overview/).

### Off-Chain Data

The off-chain data is stored in a distributed manner and managed by a data availability committee, ensuring that it is available for validation. The data availability committee is defined as a core smart contract, available [here](https://github.com/okx/x1-contracts/blob/main/contracts/DataCommittee.sol). This is crucial for the Validium model, where data computation happens off-chain but needs to be verifiable on-chain.

### Running

Instructions on how to run this software can be found [here](./docs/running.md)

## License

The x1-node project is licensed under the [GNU Affero General Public License](LICENSE) free software license.
