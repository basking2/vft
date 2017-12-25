# Venus Fly Trap

_A network anomaly detection engine_

## Summary

VFT straddles the line between a network intrusion detection system (NIDS), host-based intrustion detection system (HIDS), and a honeypot. It is simultaneously all and none of these. Designed to work best across a fleet of clients, VFT detects connections and potential scanning activity that may evade your NIDS or HIDS, but without requiring setup of heavier systems such ELK or relying on event correlation from Datadog (although that is totally an option). 

### How it works

VFT is a pair of binaries that work as a client and a server.

#### Server

`vft-server` listens for incoming client connections, correlates activity, and reports any potentially interesting anomalies. Events such as frequent hits to the same "trap" port across all clients, frequent hits from the same address across all clients, and frequent hits to the same client are all examples of activity that will be alerted on.

The server binds to `0.0.0.0:9999` by default.

Example of starting the server:
```
$ vft-server --bind 0.0.0.0:9999
```

Or with SSL:
```
$ vft-server --ssl --cert ./cert.pem --key ./key.pem --bind 0.0.0.0:9999
```

#### Client mode

`vft-client` starts a TCP listener on 5 random "common" ports and listens for connections. Upon receiving a connection, VFT closes it and sends a report of the connection to the server which will then attempt to correlate the activity. Client mode requires a VFT server to operate properly.

Example of starting the client:
```
$ vft-client --server vft.yourcompany.net:9999
```

Passing the `--cert` option autmatically enables SSL:
```
$ vft-client --cert ./cert.pem --server vft.yourcompany.net:9999
```

### Installation

Instructions forthcoming for the following:
  - Downloads
  - Install as package
  - Docker container
  - Build from source

