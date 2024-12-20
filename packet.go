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

