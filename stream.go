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
	"net"
)

type Listener interface {
	io.Closer

	Listen() error
	Addr() net.Addr

	Accept() (Stream, error)
}

type Stream interface {
	io.Closer

	Addr() net.Addr
	Encode(interface{}) error
	Decode(interface{}) error
}

type Family string

const (
	TCP        Family = "tcp"
	TCP4       Family = "tcp4"
	TCP6       Family = "tcp6"
	Unix       Family = "unix"
	UnixPacket Family = "unixpacket"

	UDP      Family = "udp"
	UDP4     Family = "udp4"
	UDP6     Family = "UDP6"
	UnixGram Family = "unixgram"
	IP       Family = "ip"
	IP4      Family = "ip4"
	IP6      Family = "ip6"
)

