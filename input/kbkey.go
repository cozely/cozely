// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import (
	"github.com/cozely/cozely/internal"
)

type kbKey struct {
	keycode       keyCode
	target        Action
	pressed bool
}

func (a *kbKey) bind(c ContextID, target Action) {
	aa := *a
	aa.target = target
	devices.bindings[kbmouse][c] =
		append(devices.bindings[kbmouse][c], &aa)
}

func (a *kbKey) device() DeviceID {
	return kbmouse
}

func (a *kbKey) action() Action {
	return a.target
}

func (a *kbKey) activate(d DeviceID) {
	a.target.activate(d, a)
}

func (a *kbKey) asBool() (just bool, value bool) {
	v := internal.Key(a.keycode)
	j := (v != a.pressed) //TODO: no need to store?
	a.pressed = v
	return j, a.pressed
}
