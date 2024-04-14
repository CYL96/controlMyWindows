/**************************************************
*文件名：define.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/4/11
**************************************************/

package config

// Defining a bunch of constants.

type (
	KKey string
)

func (i KKey) Str() string {
	return string(i)
}

const (
	Spec0  KKey = "`"
	Spec1  KKey = "-"
	Spec2  KKey = "="
	Spec3  KKey = "["
	Spec4  KKey = "]"
	Spec5  KKey = "\\"
	Spec6  KKey = ";"
	Spec7  KKey = "'"
	Spec8  KKey = ","
	Spec9  KKey = "."
	Spec10 KKey = "/"

	// KeyA define key "a"
	KeyA KKey = "a"
	KeyB KKey = "b"
	KeyC KKey = "c"
	KeyD KKey = "d"
	KeyE KKey = "e"
	KeyF KKey = "f"
	KeyG KKey = "g"
	KeyH KKey = "h"
	KeyI KKey = "i"
	KeyJ KKey = "j"
	KeyK KKey = "k"
	KeyL KKey = "l"
	KeyM KKey = "m"
	KeyN KKey = "n"
	KeyO KKey = "o"
	KeyP KKey = "p"
	KeyQ KKey = "q"
	KeyR KKey = "r"
	KeyS KKey = "s"
	KeyT KKey = "t"
	KeyU KKey = "u"
	KeyV KKey = "v"
	KeyW KKey = "w"
	KeyX KKey = "x"
	KeyY KKey = "y"
	KeyZ KKey = "z"
	//
	CapA KKey = "A"
	CapB KKey = "B"
	CapC KKey = "C"
	CapD KKey = "D"
	CapE KKey = "E"
	CapF KKey = "F"
	CapG KKey = "G"
	CapH KKey = "H"
	CapI KKey = "I"
	CapJ KKey = "J"
	CapK KKey = "K"
	CapL KKey = "L"
	CapM KKey = "M"
	CapN KKey = "N"
	CapO KKey = "O"
	CapP KKey = "P"
	CapQ KKey = "Q"
	CapR KKey = "R"
	CapS KKey = "S"
	CapT      = "T"
	CapU KKey = "U"
	CapV KKey = "V"
	CapW KKey = "W"
	CapX KKey = "X"
	CapY KKey = "Y"
	CapZ KKey = "Z"
	//
	Key0 KKey = "0"
	Key1 KKey = "1"
	Key2 KKey = "2"
	Key3 KKey = "3"
	Key4 KKey = "4"
	Key5 KKey = "5"
	Key6 KKey = "6"
	Key7 KKey = "7"
	Key8 KKey = "8"
	Key9 KKey = "9"

	// Backspace backspace key string
	Backspace KKey = "backspace"
	Delete    KKey = "delete"
	Enter     KKey = "enter"
	Tab       KKey = "tab"
	Esc       KKey = "esc"
	Escape    KKey = "escape"
	Up        KKey = "up"    // Up arrow key
	Down      KKey = "down"  // Down arrow key
	Right     KKey = "right" // Right arrow key
	Left      KKey = "left"  // Left arrow key
	Home      KKey = "home"
	End       KKey = "end"
	Pageup    KKey = "pageup"
	Pagedown  KKey = "pagedown"

	F1  KKey = "f1"
	F2  KKey = "f2"
	F3  KKey = "f3"
	F4  KKey = "f4"
	F5  KKey = "f5"
	F6  KKey = "f6"
	F7  KKey = "f7"
	F8  KKey = "f8"
	F9  KKey = "f9"
	F10 KKey = "f10"
	F11 KKey = "f11"
	F12 KKey = "f12"
	F13 KKey = "f13"
	F14 KKey = "f14"
	F15 KKey = "f15"
	F16 KKey = "f16"
	F17 KKey = "f17"
	F18 KKey = "f18"
	F19 KKey = "f19"
	F20 KKey = "f20"
	F21 KKey = "f21"
	F22 KKey = "f22"
	F23 KKey = "f23"
	F24 KKey = "f24"

	Cmd  KKey = "cmd"  // is the "win" key for windows
	Lcmd KKey = "lcmd" // left command
	Rcmd KKey = "rcmd" // right command
	// "command"
	Alt     KKey = "alt"
	Lalt    KKey = "lalt" // left alt
	Ralt    KKey = "ralt" // right alt
	Ctrl    KKey = "ctrl"
	Lctrl   KKey = "lctrl" // left ctrl
	Rctrl   KKey = "rctrl" // right ctrl
	Control KKey = "control"
	Shift   KKey = "shift"
	Lshift  KKey = "lshift" // left shift
	Rshift  KKey = "rshift" // right shift
	// "right_shift"
	Capslock    KKey = "capslock"
	Space       KKey = "space"
	Print       KKey = "print"
	Printscreen KKey = "printscreen" // No Mac support
	Insert      KKey = "insert"
	Menu        KKey = "menu" // Windows only

	AudioMute    KKey = "audio_mute"     // Mute the volume
	AudioVolDown KKey = "audio_vol_down" // Lower the volume
	AudioVolUp   KKey = "audio_vol_up"   // Increase the volume
	AudioPlay    KKey = "audio_play"
	AudioStop    KKey = "audio_stop"
	AudioPause   KKey = "audio_pause"
	AudioPrev    KKey = "audio_prev"    // Previous Track
	AudioNext    KKey = "audio_next"    // Next Track
	AudioRewind  KKey = "audio_rewind"  // Linux only
	AudioForward KKey = "audio_forward" // Linux only
	AudioRepeat  KKey = "audio_repeat"  //  Linux only
	AudioRandom  KKey = "audio_random"  //  Linux only

	Num0    KKey = "num0" // numpad 0
	Num1    KKey = "num1"
	Num2    KKey = "num2"
	Num3    KKey = "num3"
	Num4    KKey = "num4"
	Num5    KKey = "num5"
	Num6    KKey = "num6"
	Num7    KKey = "num7"
	Num8    KKey = "num8"
	Num9    KKey = "num9"
	NumLock KKey = "num_lock"

	NumDecimal KKey = "num."
	NumPlus    KKey = "num+"
	NumMinus   KKey = "num-"
	NumMul     KKey = "num*"
	NumDiv     KKey = "num/"
	NumClear   KKey = "num_clear"
	NumEnter   KKey = "num_enter"
	NumEqual   KKey = "num_equal"

	LightsMonUp     KKey = "lights_mon_up"     // Turn up monitor brightness			No Windows support
	LightsMonDown   KKey = "lights_mon_down"   // Turn down monitor brightness		No Windows support
	LightsKbdToggle KKey = "lights_kbd_toggle" // Toggle keyboard backlight on/off		No Windows support
	LightsKbdUp     KKey = "lights_kbd_up"     // Turn up keyboard backlight brightness	No Windows support
	LightsKbdDown   KKey = "lights_kbd_down"
)

type (
	KMouseScrollDir int
)

const (
	MouseLeft  KKey = "left"
	MouseRight KKey = "right"
	MouseWheel KKey = "wheel"

	MouseScrollUp    KMouseScrollDir = 1
	MouseScrollDown  KMouseScrollDir = 2
	MouseScrollLeft  KMouseScrollDir = 3
	MouseScrollRight KMouseScrollDir = 4
)
