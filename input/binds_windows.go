package input

import "github.com/libretro/ludo/libretro"

var joyBinds = map[string]joybinds{
	"Microsoft X-Box 360 pad":    xbox360JoyBinds,
	"Xbox 360 Controller":        xboxOneJoyBinds,
	"Xbox Controller":            xboxOneJoyBinds,
	"Wireless Controller":        ds4JoyBinds,
	"PLAYSTATION(R)3 Controller": ds3JoyBinds,
}

var xbox360JoyBinds = joybinds{
	bind{btn, 0, 0, 0}:  libretro.DeviceIDJoypadB,
	bind{btn, 1, 0, 0}:  libretro.DeviceIDJoypadA,
	bind{btn, 2, 0, 0}:  libretro.DeviceIDJoypadY,
	bind{btn, 3, 0, 0}:  libretro.DeviceIDJoypadX,
	bind{btn, 4, 0, 0}:  libretro.DeviceIDJoypadL,
	bind{btn, 5, 0, 0}:  libretro.DeviceIDJoypadR,
	bind{btn, 6, 0, 0}:  libretro.DeviceIDJoypadSelect,
	bind{btn, 7, 0, 0}:  libretro.DeviceIDJoypadStart,
	bind{btn, 8, 0, 0}:  libretro.DeviceIDJoypadL3,
	bind{btn, 9, 0, 0}:  libretro.DeviceIDJoypadR3,
	bind{btn, 10, 0, 0}: libretro.DeviceIDJoypadUp,
	bind{btn, 11, 0, 0}: libretro.DeviceIDJoypadRight,
	bind{btn, 12, 0, 0}: libretro.DeviceIDJoypadDown,
	bind{btn, 13, 0, 0}: libretro.DeviceIDJoypadLeft,
	bind{axis, 4, 1, 0}: libretro.DeviceIDJoypadL2,
	bind{axis, 5, 1, 0}: libretro.DeviceIDJoypadL3,
}

var xboxOneJoyBinds = joybinds{
	bind{btn, 0, 0, 0}:  libretro.DeviceIDJoypadB,
	bind{btn, 1, 0, 0}:  libretro.DeviceIDJoypadA,
	bind{btn, 2, 0, 0}:  libretro.DeviceIDJoypadY,
	bind{btn, 3, 0, 0}:  libretro.DeviceIDJoypadX,
	bind{btn, 4, 0, 0}:  libretro.DeviceIDJoypadL,
	bind{btn, 5, 0, 0}:  libretro.DeviceIDJoypadR,
	bind{btn, 6, 0, 0}:  libretro.DeviceIDJoypadSelect,
	bind{btn, 7, 0, 0}:  libretro.DeviceIDJoypadStart,
	bind{btn, 8, 0, 0}:  libretro.DeviceIDJoypadL3,
	bind{btn, 9, 0, 0}:  libretro.DeviceIDJoypadR3,
	bind{btn, 10, 0, 0}: libretro.DeviceIDJoypadUp,
	bind{btn, 11, 0, 0}: libretro.DeviceIDJoypadRight,
	bind{btn, 12, 0, 0}: libretro.DeviceIDJoypadDown,
	bind{btn, 13, 0, 0}: libretro.DeviceIDJoypadLeft,
	bind{axis, 4, 1, 0}: libretro.DeviceIDJoypadL2,
	bind{axis, 5, 1, 0}: libretro.DeviceIDJoypadL3,
}

var ds4JoyBinds = joybinds{
	bind{btn, 0, 0, 0}:  libretro.DeviceIDJoypadX,
	bind{btn, 1, 0, 0}:  libretro.DeviceIDJoypadB,
	bind{btn, 2, 0, 0}:  libretro.DeviceIDJoypadA,
	bind{btn, 3, 0, 0}:  libretro.DeviceIDJoypadY,
	bind{btn, 4, 0, 0}:  libretro.DeviceIDJoypadL,
	bind{btn, 5, 0, 0}:  libretro.DeviceIDJoypadR,
	bind{btn, 4, 0, 0}:  libretro.DeviceIDJoypadL2,
	bind{btn, 5, 0, 0}:  libretro.DeviceIDJoypadR2,
	bind{btn, 6, 0, 0}:  libretro.DeviceIDJoypadSelect,
	bind{btn, 7, 0, 0}:  libretro.DeviceIDJoypadStart,
	bind{btn, 8, 0, 0}:  libretro.DeviceIDJoypadL3,
	bind{btn, 9, 0, 0}:  libretro.DeviceIDJoypadR3,
	bind{btn, 10, 0, 0}: ActionMenuToggle,
	bind{btn, 14, 0, 0}: libretro.DeviceIDJoypadUp,
	bind{btn, 15, 0, 0}: libretro.DeviceIDJoypadRight,
	bind{btn, 16, 0, 0}: libretro.DeviceIDJoypadDown,
	bind{btn, 17, 0, 0}: libretro.DeviceIDJoypadLeft,
	bind{axis, 3, 1, 0}: libretro.DeviceIDJoypadL2,
	bind{axis, 4, 1, 0}: libretro.DeviceIDJoypadL3,
}

// Detected but doesn't send inputs
var ds3JoyBinds = joybinds{}
