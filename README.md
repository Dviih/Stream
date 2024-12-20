# Stream
### The first implementation of the [Bin](https://github.com/Dviih/bin) Protocol on top of network protocols.
##### This project is licensed under [AGPL](https://github.com/Dviih/Stream/blob/main/LICENSE).

---

## Listener
### A generic implementation for both `net.Listener` and `net.PacketConn` to work as the same way.

## Stream
### Handling packets by encoding and decoding it, dialing or accepting will result return it.

## Addr
### `net.Addr` implementation by user.

## Family
### An enum including each supported family.

---

# Usage

## Listener
- `Accept` - accepts a new incoming stream.
- `Addr` - returns set address or listening address after `Listen` call.
- `Close` - closes the listener.
- `Listen` - starts the listener.

## Stream
- `Addr` - returns the address of stream.
- `Close` - closes the stream.
- `Encode` - encodes data and sends it.
- `Decode` - decodes data after receiving it.

---

## Utilities

- `Conn` - returns a Stream from `net.Conn`.
- `Packet` - returns a Stream from `net.PacketListener` and `net.Addr`.
- `NewAddr` - returns a `net.Addr` by user input.
- `NewSeqListener` - returns a Listener.
- `NewPacketListener` - returns a Listener.

---

## Listen & Dial

- `Listen` - listens to specified family and address, returns Listener.
- `Dial` - dials to specified family and address, returns Stream.

---

### Made for Gophers by @Dviih
