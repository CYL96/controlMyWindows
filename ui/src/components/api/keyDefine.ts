import {PressType} from "@/components/api/detail.js";

export const KeyDefine = {
    KeyA: "a",
    KeyB: "b",
    KeyC: "c",
    KeyD: "d",
    KeyE: "e",
    KeyF: "f",
    KeyG: "g",
    KeyH: "h",
    KeyI: "i",
    KeyJ: "j",
    KeyK: "k",
    KeyL: "l",
    KeyM: "m",
    KeyN: "n",
    KeyO: "o",
    KeyP: "p",
    KeyQ: "q",
    KeyR: "r",
    KeyS: "s",
    KeyT: "t",
    KeyU: "u",
    KeyV: "v",
    KeyW: "w",
    KeyX: "x",
    KeyY: "y",
    KeyZ: "z",

    CapA: "A",
    CapB: "B",
    CapC: "C",
    CapD: "D",
    CapE: "E",
    CapF: "F",
    CapG: "G",
    CapH: "H",
    CapI: "I",
    CapJ: "J",
    CapK: "K",
    CapL: "L",
    CapM: "M",
    CapN: "N",
    CapO: "O",
    CapP: "P",
    CapQ: "Q",
    CapR: "R",
    CapS: "S",
    CapT: "T",
    CapU: "U",
    CapV: "V",
    CapW: "W",
    CapX: "X",
    CapY: "Y",
    CapZ: "Z",
    Key0: "0",
    Key1: "1",
    Key2: "2",
    Key3: "3",
    Key4: "4",
    Key5: "5",
    Key6: "6",
    Key7: "7",
    Key8: "8",
    Key9: "9",

    Backspace: "backspace",
    Delete: "delete",
    Enter: "enter",
    Tab: "tab",
    Esc: "esc",
    Escape: "escape",
    Up: "up",
    Down: "down",
    Right: "right",
    Left: "left",
    Home: "home",
    End: "end",
    Pageup: "pageup",
    Pagedown: "pagedown",

    Spec0: "`",
    Spec1: "-",
    Spec2: "=",
    Spec3: "[",
    Spec4: "]",
    Spec5: "\\",
    Spec6: ";",
    Spec7: "'",
    Spec8: ",",
    Spec9: ".",
    Spec10: "/",

    F1: "f1",
    F2: "f2",
    F3: "f3",
    F4: "f4",
    F5: "f5",
    F6: "f6",
    F7: "f7",
    F8: "f8",
    F9: "f9",
    F10: "f10",
    F11: "f11",
    F12: "f12",
    F13: "f13",
    F14: "f14",
    F15: "f15",
    F16: "f16",
    F17: "f17",
    F18: "f18",
    F19: "f19",
    F20: "f20",
    F21: "f21",
    F22: "f22",
    F23: "f23",
    F24: "f24",

    Cmd: "cmd",
    Lcmd: "lcmd",
    Rcmd: "rcmd",

    Alt: "alt",
    Lalt: "lalt",
    Ralt: "ralt",
    Ctrl: "ctrl",
    Lctrl: "lctrl",
    Rctrl: "rctrl",
    Control: "control",
    Shift: "shift",
    Lshift: "lshift",
    Rshift: "rshift",
    Capslock: "capslock",
    Space: "space",
    Print: "print",
    Printscreen: "printscreen",
    Insert: "insert",
    Menu: "menu",
    AudioMute: "audio_mute",
    AudioVolDown: "audio_vol_down",
    AudioVolUp: "audio_vol_up",
    AudioPlay: "audio_play",
    AudioStop: "audio_stop",
    AudioPause: "audio_pause",
    AudioPrev: "audio_prev",
    AudioNext: "audio_next",
    AudioRewind: "audio_rewind",
    AudioForward: "audio_forward",
    AudioRepeat: "audio_repeat",
    AudioRandom: "audio_random",

    Num0: "num0",
    Num1: "num1",
    Num2: "num2",
    Num3: "num3",
    Num4: "num4",
    Num5: "num5",
    Num6: "num6",
    Num7: "num7",
    Num8: "num8",
    Num9: "num9",
    NumLock: "num_lock",
    NumDecimal: "num.",
    NumPlus: "num+",
    NumMinus: "num-",
    NumMul: "num*",
    NumDiv: "num/",
    NumClear: "num_clear",
    NumEnter: "num_enter",
    NumEqual: "num_equal",
    LightsMonUp: "lights_mon_up",
    LightsMonDown: "lights_mon_down",
    LightsKbdToggle: "lights_kbd_toggle",
    LightsKbdUp: "lights_kbd_up",
    LightsKbdDown: "lights_kbd_down",
}

export function GetShowComponents(key: string) {
    if (/^[a-zA-Z]$/.test(key)) {
        return key.toUpperCase()
    }
    switch (key) {
        case KeyDefine.Spec0:
            return '~ `'
        case KeyDefine.Spec1:
            return '_ -'
        case KeyDefine.Spec2:
            return '+ ='
        case KeyDefine.Spec3:
            return '{ ['
        case KeyDefine.Spec4:
            return '} ]'
        case KeyDefine.Spec5:
            return '| \\'
        case KeyDefine.Spec6:
            return ': ;'
        case KeyDefine.Spec7:
            return '" \''
        case KeyDefine.Spec8:
            return '< ,'
        case KeyDefine.Spec9:
            return '> .'
        case KeyDefine.Spec10:
            return '? /'

        case KeyDefine.Key1:
            return '! 1'
        case KeyDefine.Key2:
            return '@ 2'
        case KeyDefine.Key3:
            return '# 3'
        case KeyDefine.Key4:
            return '$ 4'
        case KeyDefine.Key5:
            return '% 5'
        case KeyDefine.Key6:
            return '^ 6'
        case KeyDefine.Key7:
            return '& 7'
        case KeyDefine.Key8:
            return '* 8'
        case KeyDefine.Key9:
            return '( 9'
        case KeyDefine.Key0:
            return ') 0'

        case KeyDefine.F1:
            return 'F1'
        case KeyDefine.F2:
            return 'F2'
        case KeyDefine.F3:
            return 'F3'
        case KeyDefine.F4:
            return 'F4'
        case KeyDefine.F5:
            return 'F5'
        case KeyDefine.F6:
            return 'F6'
        case KeyDefine.F7:
            return 'F7'
        case KeyDefine.F8:
            return 'F8'
        case KeyDefine.F9:
            return 'F9'
        case KeyDefine.F10:
            return 'F10'
        case KeyDefine.F11:
            return 'F11'
        case KeyDefine.F12:
            return 'F12'
        case KeyDefine.F13:
            return 'F13'
        case KeyDefine.F14:
            return 'F14'
        case KeyDefine.F15:
            return 'F15'
        case KeyDefine.F16:
            return 'F16'
        case KeyDefine.F17:
            return 'F17'
        case KeyDefine.F18:
            return 'F18'
        case KeyDefine.F19:
            return 'F19'
        case KeyDefine.F20:
            return 'F20'
        case KeyDefine.F21:
            return 'F21'
        case KeyDefine.F22:
            return 'F22'
        case KeyDefine.F23:
            return 'F23'
        case KeyDefine.F24:
            return 'F24'

        case KeyDefine.Up:
            return '↑'
        case KeyDefine.Down:
            return '↓'
        case KeyDefine.Left:
            return '←'
        case KeyDefine.Right:
            return '→'

        case KeyDefine.Pageup:
            return 'PG↑'
        case KeyDefine.Pagedown:
            return 'PG↓'

        case KeyDefine.Lcmd:
            return '🅻-⊞'
        case KeyDefine.Rcmd:
            return '🆁-⊞'

        case KeyDefine.Menu:
            return '▤'

        case KeyDefine.Tab:
            return 'TAB↹'
        case KeyDefine.Capslock:
            return 'CAPS'
        case KeyDefine.Lshift:
            return '🅻-SHIFT'
        case KeyDefine.Lctrl:
            return '🅻-CTRL'
        case KeyDefine.Lalt:
            return '🅻-ALT'
        case KeyDefine.Backspace:
            return 'BACK'
        case KeyDefine.Enter:
            return 'ENTER'
        case KeyDefine.Rshift:
            return '🆁-SHIFT'
        case KeyDefine.Ralt:
            return '🆁-ALT'
        case KeyDefine.Rctrl:
            return '🆁-CTRL'
        case KeyDefine.Insert:
            return 'INS'
        case KeyDefine.Delete:
            return 'DEL'
        case KeyDefine.Home:
            return 'HM'
        case KeyDefine.End:
            return 'END'
        case KeyDefine.Num0:
            return "(0)"
        case  KeyDefine.Num1:
            return "(1)"
        case  KeyDefine.Num2:
            return "(2)"
        case  KeyDefine.Num3:
            return "(3)"
        case  KeyDefine.Num4:
            return "(4)"
        case  KeyDefine.Num5:
            return "(5)"
        case  KeyDefine.Num6:
            return "(6)"
        case  KeyDefine.Num7:
            return "(7)"
        case  KeyDefine.Num8:
            return "(8)"
        case  KeyDefine.Num9:
            return "(9)"
        case  KeyDefine.NumLock:
            return "NMLK"
        case  KeyDefine.NumDecimal:
            return "(.)"
        case  KeyDefine.NumPlus:
            return "(+)"
        case  KeyDefine.NumMinus:
            return "(-)"
        case  KeyDefine.NumMul:
            return "(*)"
        case  KeyDefine.NumDiv:
            return "(/)"

        case  KeyDefine.AudioMute:
            return "🎵✖"
        case  KeyDefine.AudioVolDown:
            return "🎵-"
        case  KeyDefine.AudioVolUp:
            return "🎵+"
        case  KeyDefine.AudioPlay:
            return "▶︎"
        case  KeyDefine.AudioStop:
            return "⏹︎"
        case  KeyDefine.AudioPause:
            return "⏸︎"
        case  KeyDefine.AudioPrev:
            return "⏮︎"
        case  KeyDefine.AudioNext:
            return "⏭︎"
    }
    return key
}



export const MouseDefine = {
    MouseLeft: "left",
    MouseRight: "right",
    MouseWheel: "wheel",
}

export const MouseScrollDirection = {
    Up: 1,
    Down: 2,
    Left: 3,
    Right: 4,
}