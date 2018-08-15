package input

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/libretro/go-playthemall/libretro"
	"github.com/libretro/go-playthemall/notifications"
	"github.com/libretro/go-playthemall/settings"
	"github.com/libretro/go-playthemall/state"
)

const numPlayers = 5

type joybinds map[bind]uint32

const (
	menuActionMenuToggle       uint32 = libretro.DeviceIDJoypadR3 + 1
	menuActionFullscreenToggle uint32 = libretro.DeviceIDJoypadR3 + 2
	menuActionShouldClose      uint32 = libretro.DeviceIDJoypadR3 + 3
	menuActionLast             uint32 = libretro.DeviceIDJoypadR3 + 4
)

var keyBinds = map[glfw.Key]uint32{
	glfw.KeyX:         libretro.DeviceIDJoypadA,
	glfw.KeyZ:         libretro.DeviceIDJoypadB,
	glfw.KeyA:         libretro.DeviceIDJoypadY,
	glfw.KeyS:         libretro.DeviceIDJoypadX,
	glfw.KeyUp:        libretro.DeviceIDJoypadUp,
	glfw.KeyDown:      libretro.DeviceIDJoypadDown,
	glfw.KeyLeft:      libretro.DeviceIDJoypadLeft,
	glfw.KeyRight:     libretro.DeviceIDJoypadRight,
	glfw.KeyEnter:     libretro.DeviceIDJoypadStart,
	glfw.KeyBackspace: libretro.DeviceIDJoypadSelect,
	glfw.KeyP:         menuActionMenuToggle,
	glfw.KeyF:         menuActionFullscreenToggle,
	glfw.KeyEscape:    menuActionShouldClose,
}

const btn = 0
const axis = 1

type bind struct {
	kind      uint32
	index     uint32
	direction float32
	threshold float32
}

type inputstate [numPlayers][menuActionLast]bool

// Input state for all the players
var (
	NewState inputstate // input state for the current frame
	OldState inputstate // input state for the previous frame
	Released inputstate // keys just released during this frame
	Pressed  inputstate // keys just pressed during this frame
)

func joystickCallback(joy int, event int) {
	switch event {
	case 262145:
		notifications.DisplayAndLog("Input", "Joystick #%d plugged: %s.", joy, glfw.GetJoystickName(glfw.Joystick(joy)))
	case 262146:
		notifications.DisplayAndLog("Input", "Joystick #%d unplugged.", joy)
	default:
		notifications.DisplayAndLog("Input", "Joystick #%d unhandled event: %d.", joy, event)
	}
}

type keygetter interface {
	GetKey(glfw.Key) glfw.Action
}

type displayLogger interface {
	DisplayAndLog(prefix, message string, vars ...interface{})
}

var window keygetter

func Init(w keygetter) {
	window = w
	glfw.SetJoystickCallback(joystickCallback)
}

// Reset all retropad buttons to false
func inputPollReset(state inputstate) inputstate {
	for p := range state {
		for k := range state[p] {
			state[p][k] = false
		}
	}
	return state
}

// Process joypads of all players
func inputPollJoypads(state inputstate) inputstate {
	for p := range state {
		buttonState := glfw.GetJoystickButtons(glfw.Joystick(p))
		axisState := glfw.GetJoystickAxes(glfw.Joystick(p))
		name := glfw.GetJoystickName(glfw.Joystick(p))
		jb := joyBinds[name]
		if len(buttonState) > 0 {
			for k, v := range jb {
				switch k.kind {
				case btn:
					if glfw.Action(buttonState[k.index]) == glfw.Press {
						state[p][v] = true
					}
				case axis:
					if k.direction*axisState[k.index] > k.threshold*k.direction {
						state[p][v] = true
					}
				}
			}
		}
	}
	return state
}

// Process keyboard keys
func inputPollKeyboard(state inputstate) inputstate {
	for k, v := range keyBinds {
		if window.GetKey(k) == glfw.Press {
			state[0][v] = true
		}
	}
	return state
}

// Compute the keys pressed or released during this frame
func inputGetPressedReleased(new inputstate, old inputstate) (inputstate, inputstate) {
	for p := range new {
		for k := range new[p] {
			Pressed[p][k] = new[p][k] && !old[p][k]
			Released[p][k] = !new[p][k] && old[p][k]
		}
	}
	return Pressed, Released
}

func Poll() {
	NewState = inputPollReset(NewState)
	NewState = inputPollJoypads(NewState)
	NewState = inputPollKeyboard(NewState)
	Pressed, Released = inputGetPressedReleased(NewState, OldState)

	// Toggle the menu if menuActionMenuToggle is pressed
	if Released[0][menuActionMenuToggle] && state.Global.CoreRunning {
		state.Global.MenuActive = !state.Global.MenuActive
	}

	// Toggle fullscreen if menuActionFullscreenToggle is pressed
	if Released[0][menuActionFullscreenToggle] {
		settings.Settings.VideoFullscreen = !settings.Settings.VideoFullscreen
		//videoConfigure(settings.Settings.VideoFullscreen)
		settings.Save()
	}

	// Close on escape
	if Pressed[0][menuActionShouldClose] {
		//window.SetShouldClose(true)
	}

	// Store the old input state for comparisions
	OldState = NewState
}

// State is a callback passed to core.SetInputState
// It returns 1 if the button corresponding to the parameters is pressed
func State(port uint, device uint32, index uint, id uint) int16 {
	if id >= 255 || index > 0 || device != libretro.DeviceJoypad {
		return 0
	}

	if NewState[port][id] {
		return 1
	}
	return 0
}