# Bank Proxy

The "proxy" contract model uses two contracts: a dedicated Proxy and an
external API. The Proxy calls out to the API for business logic, creating
a system that can be upgraded without losing the Proxy contract's stored
data by updating the address of the external API within the Proxy.

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

A mismatch will cause the delegatecall to reference the wrong location in storage,
possibly corrupting the data.

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

The API contract should be deployed first, so that it is available to the Proxy when
it is deployed. Other API versions can be switched to via the Proxy's UpgradeAPI
function.

