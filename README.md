# Venus Fly Trap

_A network anomaly detection engine_

## Summary

VFT straddles the line between a network intrusion detection system (NIDS), host-based intrustion detection system (HIDS), and a honeypot. It is simultaneously all and none of these. Designed to work best across a fleet of clients, VFT detects connections and potential scanning activity that may evade your NIDS or HIDS, but without requiring setup of heavier systems such ELK or relying on event correlation from Datadog (although that is totally an option). 

### How it works

VFT is a single binary that runs in three modes: server, client, and standalone. 

#### Server mode

VFT server mode listens for incoming client connections, correlates activity, and reports any potentially interesting anomalies. Events such as frequent hits to the same "trap" port across all clients, frequent hits from the same address across all clients, and frequent hits to the same client are all examples of activity that will be alerted on.

The server binds to `127.0.0.1:9999` by default.

Example of starting the server:
```
$ vft --bind 0.0.0.0:9999 server
```

#### Client mode

VFT client mode starts a TCP listener on 5 random "common" ports and listens for connections. Upon receiving a connection, VFT closes it and sends a report of the connection to the server which will then attempt to correlate the activity. Client mode requires a VFT server to operate properly.

Example of starting the client:
```
$ vft --server vft.yourcompany.net:9999 client
```

#### Standalone

Many organizations already have event correlation in place. With those users in mind, VFT also supports standalone mode, which behaves like a client but does not try to perform any interaction with a VFT server. Instead, it simply sends a JSON message to the destination IP and port so that the user's pre-installed event correlation can use the data. This is ideal for people who already have an existing ELK or Splunk installation and want to add VFT to their threat detection toolkit.

Example of starting a standalone instance:
```
vft --dest elk.yourcompany.net:9300 standalone
```

## Installation

### Linux
#### Bulding
Clone the repo + Build the binary
```
git clone git@gitlab.com:bbriggs1/vft.git
cd vft
./build.sh
```

### Windows
Check back later.