# XGON Data Availability
### Data Availability Layer for XGON Validium

The xgon-data-availability project is a specialized Data Availability Node (DA Node) that is part of XGON's CDK (Chain Development Kit) Validium.

## Overview of Validium
The XGON Validium solution is made up of several components; start with the [XGON Node](https://github.com/okx/xgon-node). For quick reference, the complete list of components are outlined below:

| Component                                                                     | Description                                                          |
| ----------------------------------------------------------------------------- | -------------------------------------------------------------------- |
| [XGON Node](https://github.com/okx/xgon-node)           | Node implementation for the XGON networks in Validium mode            |
| [XGON Contracts](https://github.com/okx/xgon-contracts) | Smart contract implementation for the XGON networks in Validium mode |
| [XGON Data Availability](https://github.com/okx/xgon-data-availability)   | Data availability implementation for the XGON networks          |
| [Prover / Executor](https://github.com/okx/xgon-prover)          | zkEVM engine and prover implementation                               |
| [Bridge Service](https://github.com/okx/xgon-bridge-service)     | Bridge service implementation for XGON networks                       |

---

## Introduction

As blockchain networks grow, the volume of data that needs to be stored and validated increases, posing challenges in scalability and efficiency. Storing all data on-chain can lead to bloated blockchains, slow transactions, and high fees.

Data Availability Nodes facilitate a separation between transaction execution and data storage. They allow transaction data to reside off-chain while remaining accessible for validation. This significantly improves scalability and reduces costs. Within the framework of XGON's CDK, Data Availability Committees (DAC) members run DA nodes to ensure the security, accessibility, and reliability of off-chain data.

To learn more about how the data availability layer works in the validium, please see the CDK documentation [here](https://wiki.polygon.technology/docs/cdk/dac-overview/).

### Off-Chain Data

The off-chain data is stored in a distributed manner and managed by a data availability committee, ensuring that it is available for validation. The data availability committee is defined as a core smart contract, available [here](https://github.com/okx/xgon-contracts/blob/main/contracts/DataCommittee.sol). This is crucial for the Validium model, where data computation happens off-chain but needs to be verifiable on-chain.

## License

The xgon-node project is licensed under the [GNU Affero General Public License](LICENSE) free software license.
