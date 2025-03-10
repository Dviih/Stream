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

import "net"

type Addr struct {
	network string
	address string
}

func (addr *Addr) Network() string {
	return addr.network
}

func (addr *Addr) String() string {
	return addr.address
}

func NewAddr(network, address string) net.Addr {
	return &Addr{
		network: network,
		address: address,
	}
}
