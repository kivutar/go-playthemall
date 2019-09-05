package menu

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/libretro/ludo/video"
)

// Used to easily compose different hint bars based on the context.
func stackHint(stack *float32, icon uint32, label string, h int) {
	c := video.Color{R: 0.25, G: 0.25, B: 0.25, A: 1}
	vid.Font.SetColor(0.25, 0.25, 0.25, 1.0)
	*stack += 30 * menu.ratio
	vid.DrawImage(icon, *stack, float32(h)-70*menu.ratio, 70*menu.ratio, 70*menu.ratio, 1.0, c)
	*stack += 70 * menu.ratio
	vid.Font.Printf(*stack, float32(h)-23*menu.ratio, 0.5*menu.ratio, label)
	*stack += vid.Font.Width(0.5*menu.ratio, label)
}

func hintIcons() (arrows, upDown, leftRight, a, b, x, y, start, slct, guide uint32) {
	if glfw.JoystickPresent(glfw.Joystick(0)) {
		return menu.icons["pad-arrows"],
			menu.icons["pad-up-down"],
			menu.icons["pad-left-right"],
			menu.icons["pad-a"],
			menu.icons["pad-b"],
			menu.icons["pad-x"],
			menu.icons["pad-y"],
			menu.icons["pad-start"],
			menu.icons["pad-select"],
			menu.icons["pad-guide"]
	}
	return menu.icons["key-arrows"],
		menu.icons["key-up-down"],
		menu.icons["key-left-right"],
		menu.icons["key-x"],
		menu.icons["key-z"],
		menu.icons["key-s"],
		menu.icons["key-a"],
		menu.icons["key-return"],
		menu.icons["key-shift"],
		menu.icons["key-p"]
}
