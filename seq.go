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
	"github.com/Dviih/bin"
	"net"
)

type SeqListener struct {
	ctx context.Context

	l    net.Listener
	addr net.Addr
}

func (listener *SeqListener) Close() error {
	if listener.l == nil {
		return nil
	}

	return listener.l.Close()
}

func (listener *SeqListener) Listen() error {
	var lc net.ListenConfig

	l, err := lc.Listen(listener.ctx, listener.addr.Network(), listener.addr.String())
	if err != nil {
		return err
	}

	listener.l = l
	return nil
}

