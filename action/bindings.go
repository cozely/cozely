// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package action

var bindings = map[string]binding{
	// Gamepad
	"Left Stick":    gamepadStick{},
	"Left Stick X":  gamepadStick{},
	"Left Stick Y":  gamepadStick{},
	"Right Stick":   gamepadStick{},
	"Right Stick X": gamepadStick{},
	"Right Stick Y": gamepadStick{},
	"Left Trigger":  gamepadTrigger{},
	"Right Trigger": gamepadTrigger{},
	"Left Bumper":   gamepadButton{},
	"Right Bumper":  gamepadButton{},
	"Dpad Up":       gamepadButton{},
	"Dpad Left":     gamepadButton{},
	"Dpad Down":     gamepadButton{},
	"Dpad Right":    gamepadButton{},
	"Button Y":      gamepadButton{},
	"Button X":      gamepadButton{},
	"Button A":      gamepadButton{},
	"Button B":      gamepadButton{},
	"Button Back":   gamepadButton{},
	"Button Start":  gamepadButton{},
	// Mouse
	"Mouse":             mouse{},
	"Mouse X":           mouse{},
	"Mouse Y":           mouse{},
	"Mouse Left":        mouseButton{},
	"Mouse Middle":      mouseButton{},
	"Mouse Right":       mouseButton{},
	"Mouse Back":        mouseButton{},
	"Mouse Forward":     mouseButton{},
	"Mouse Button 6":    mouseButton{},
	"Mouse Button 7":    mouseButton{},
	"Mouse Button 8":    mouseButton{},
	"Mouse Button 9":    mouseButton{},
	"Mouse Button 10":   mouseButton{},
	"Mouse Button 11":   mouseButton{},
	"Mouse Button 12":   mouseButton{},
	"Mouse Button 13":   mouseButton{},
	"Mouse Button 14":   mouseButton{},
	"Mouse Button 15":   mouseButton{},
	"Mouse Button 16":   mouseButton{},
	"Mouse Button 17":   mouseButton{},
	"Mouse Button 18":   mouseButton{},
	"Mouse Button 19":   mouseButton{},
	"Mouse Button 20":   mouseButton{},
	"Mouse Scroll Up":   mouseButton{},
	"Mouse Scroll down": mouseButton{},
	// Keyboard
	"A": keyboard{},
	"B": keyboard{},
	"C": keyboard{},
	"D": keyboard{},
}
