/*
 *     Unified data streaming across protocols
 *     Copyright (C) 2024  Dviih
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Affero General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Affero General Public License for more details.
 *
 *     You should have received a copy of the GNU Affero General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 *
 */

package stream

import (
	"context"
	"errors"
	"github.com/Dviih/bin"
	"github.com/Dviih/bin/buffer"
	"net"
)

type ba struct {
	buffer *buffer.Buffer
	addr   net.Addr
}

type PacketListener struct {
	ctx context.Context

	pc   net.PacketConn
	addr net.Addr

	m map[string]*ba
	c chan string
}

func (listener *PacketListener) Close() error {
	if listener.pc == nil {
		return nil
	}

	return listener.pc.Close()
}

func (listener *PacketListener) Listen() error {
	var lc net.ListenConfig

	pc, err := lc.ListenPacket(listener.ctx, listener.addr.Network(), listener.addr.String())
	if err != nil {
		return err
	}

	listener.pc = pc
	go listener.handler()
	return nil
}

func (listener *PacketListener) Addr() net.Addr {
	if listener.pc == nil {
		return listener.addr
	}

	return listener.pc.LocalAddr()
}

func (listener *PacketListener) Accept() (Stream, error) {
	name := <-listener.c

	m, ok := listener.m[name]
	if !ok {
		return nil, errors.New("invalid")
	}

	return Packet(listener.pc, m.addr, m.buffer), nil
}

func (listener *PacketListener) handler() {
	for {
		data := make([]byte, 512)

		n, addr, err := listener.pc.ReadFrom(data)
		if err != nil {
			return
		}

		name := addr.Network() + addr.String()

		m, ok := listener.m[name]
		if !ok {
			b := buffer.New()
			_, _ = b.Write(data[:n])

			listener.m[name] = &ba{
				buffer: b,
				addr:   addr,
			}

			listener.c <- name
			continue
		}

		_, _ = m.buffer.Write(data[:n])
	}
}

func NewPacketListener(ctx context.Context, addr net.Addr) Listener {
	return &PacketListener{
		ctx:  ctx,
		addr: addr,
		m:    make(map[string]*ba),
		c:    make(chan string, 64),
	}
}

type pcStream struct {
	pc   net.PacketConn
	addr net.Addr

	encoder *bin.Encoder
	decoder *bin.Decoder
}

func (stream *pcStream) Close() error {
	if _, err := stream.pc.WriteTo(nil, stream.Addr()); err != nil {
		return err
	}

	return nil
}

func (stream *pcStream) Addr() net.Addr {
	return stream.addr
}

func (stream *pcStream) Write(data []byte) (int, error) {
	return stream.pc.WriteTo(data, stream.Addr())
}

func (stream *pcStream) Encode(v interface{}) error {
	return stream.encoder.Encode(v)
}

func (stream *pcStream) Decode(v interface{}) error {
	return stream.decoder.Decode(v)
}

func Packet(pc net.PacketConn, addr net.Addr, b *buffer.Buffer) Stream {
	stream := &pcStream{
		pc:      pc,
		addr:    addr,
		decoder: bin.NewDecoder(b),
	}
	stream.encoder = bin.NewEncoder(stream)

	return stream
}
