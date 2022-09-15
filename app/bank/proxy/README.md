# Bank Proxy

The "proxy" contract model uses two contracts: a dedicated Proxy and an
external API. The Proxy calls out to the API for business logic, creating
a system that can be upgraded without losing the Proxy contract's stored
data by updating the address of the external API within the Proxy.

The Proxy contract will be both the access point and the data store for the
overall "Bank" system.

The API contracts are where any complex data operations should take place. They
can be replaced at minimal cost, and without needing to re-deploy the original
Proxy contract, therefore losing access to the data it stores or performing a
costly migration.

Because the contracts are fully separated, rollback operations are also
supported. If API v3 introduces a new bug, the Proxy contract can "upgrade" back
to API v2.

```
                                                       ┌────────────┐
                                                       │            │
                                                 ┌────►│   API_v1   │
                                                 │     │            │
                                                 │     └────────────┘
                                                 │
 ┌────────────┐       ┌───────────┐              │     ┌────────────┐
 │            │       │           │              │     │            │
 │  Customer  ├──TX──►│   Proxy   ├─delegatecall─┼────►│   API_v2   │
 │            │       │           │              │     │            │
 └────────────┘       └───────────┘              │     └────────────┘
                                                 │
                                                 │     ┌────────────┐
                                                 │     │            │
                                                 └────►│   API_v3   │
                                                       │            │
                                                       └────────────┘
```

A caveat is that the storage models for both the Proxy and the API must match
almost exactly. This is because the methods in the API expect the variables
it interacts with to be at certain points in the contract's storage.

Smart contract storage is a largely "slots" based system. Each slot can
potentially contain multiple variables, if they are small enough to fit, but
they will always be allocated according to the order in which they are declared
and not via any automated optimization strategy. This makes the allocations
predictable.

Because they are predictable, as long as the same variables are declared in the
same order for both the Proxy and the API, they will match up for the purposes
of the delegatecall.

```
+--------------+------------+------------+
| Storage Slot | Proxy      | API        |
+--------------+------------+------------+
| Slot 1       | Variable 1 | Variable 1 |
+--------------+------------+------------+
| Slot 2       | Variable 2 | Variable 2 |
+--------------+------------+------------+
| Slot 3       | Variable 3 | Variable 3 |
+--------------+------------+------------+
```

A mismatch will cause the delegatecall to reference the wrong location in
storage, possibly corrupting the data.

```
+--------------+------------+------------+
| Storage Slot | Proxy      | API        |
+--------------+------------+------------+
| Slot 1       | Variable 1 | Variable 1 |
+--------------+------------+------------+
| Slot 2       | Variable 2 | Variable 3 | <- Mismatch, delegatecall trying to reach
+--------------+------------+------------+    Variable 3 will instead be working with
| Slot 3       | Variable 3 |            |    Variable 2.
+--------------+------------+------------+
```

## Initial Setup

The make commands below are to be run from the repository root. They will
assume you have already followed the [repository instructions](README.md) to
install the appropriate tooling and run a local geth instance.

## Building the Bank Proxy and APIs

The Proxy smart contract can be found in `app/bank/proxy/contract/src/proxy/proxy.sol`.

The API smart contracts can be found in `app/bank/proxy/contract/src/api/{v1,v2,v3}/api.sol`.

To build all of the contracts, run the following make command.

```
$ make bank-proxy-build
```

## Deploying the Bank Proxy and APIs

The initial deployment must be performed in order. First the version 1 API
contract must be deployed and its Ethereum address retrieved, then that address
must be provided to the Proxy contract when it is deployed.

```
$ make bank-proxy-api-v1-deploy
...
Contract Details
----------------------------------------------------
contract id     : 0xA5B76e49bD18E952502f1eB4c4B281B91C727CBD
export BANK_API_V1_CONTRACT_ID=0xA5B76e49bD18E952502f1eB4c4B281B91C727CBD
...
```

The output of a successful deploy will include an `export` line which you can
copy and paste into your terminal to store the address of the newly-deployed
API contract in your local environment. Other make commands will depend upon
this environment variable and may error if it is empty.

```
$ export BANK_API_V1_CONTRACT_ID=0xA5B76e49bD18E952502f1eB4c4B281B91C727CBD
```

With that environment variable in place, you can then deploy the Proxy contract.

```
$ make bank-proxy-proxy-deploy
...
Contract Details
----------------------------------------------------
contract id     : 0x840bB1050D4d79c998667C35EEE4223Ef97127B2
export PROXY_CONTRACT_ID=0x840bB1050D4d79c998667C35EEE4223Ef97127B2
...
```

Again, note that an export line is provided and must be run in your shell before
proceeding.

```
$ export PROXY_CONTRACT_ID=0x840bB1050D4d79c998667C35EEE4223Ef97127B2
```

From here, you may deploy the remaining API contracts.

```
$ make bank-proxy-api-v2-deploy
$ make bank-proxy-api-v3-deploy
```

Be sure to run the appropriate `export` command from the output of each.

## Upgrading the Proxy's API

The Proxy stores the address of the API it uses to perform delegatecalls. An API
update simply exchanges the old address for a new one.

To upgrade the Proxy to API v2, run the following command.

```
$ make bank-proxy-upgrade-v2
...
Logs
----------------------------------------------------
upgrade[eb380d740ec33adf803abe0d6b14ee29ae6194a9]
...
```

A successful upgrade will log the address of the new API contract, as shown
above.
