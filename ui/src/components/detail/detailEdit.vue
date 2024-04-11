<template>
  <div style="height: 40px">
    <div style="display: flex;width: 90vw;align-items: center;justify-content: center">
      <el-scrollbar :always="true"
                    style="width: 80vw;display: flex;align-items: center;justify-content: center">
        <VueDraggable
            v-model="showList"
            :animation="150"
            id="LabelList"
            :on-update="OnUpdate"
            style="display: flex;"
        >
          <div v-for="item in showList" :key="item.id"
               class="detail-key-item">
            <div style="text-align: center">
              <div class="detail-list-item-ShortcutKey-view">
                <el-text truncated style="width: 55px">
                  {{ item.key }}
                </el-text>
              </div>
            </div>
          </div>
        </VueDraggable>
      </el-scrollbar>
    </div>
  </div>
<div style="display: flex;flex-direction: row;align-items: center;flex-wrap:wrap">
  <div style="display: flex;flex-direction: column;align-items: flex-start;margin-top: 5px">
    <div>
      <el-button size="small" :class="CheckKeyHas('esc')" @click="ClickKey('esc')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">esc
      </el-button>
      <el-button size="small" :class="CheckKeyHas('f1')" @click="ClickKey('f1')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">F1
      </el-button>
      <el-button size="small" :class="CheckKeyHas('f2')" @click="ClickKey('f2')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">F2
      </el-button>
      <el-button size="small" :class="CheckKeyHas('f3')" @click="ClickKey('f3')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">F3
      </el-button>
      <el-button size="small" :class="CheckKeyHas('f4')" @click="ClickKey('f4')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">F4
      </el-button>
      <el-button size="small" :class="CheckKeyHas('f5')" @click="ClickKey('f5')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">F5
      </el-button>
      <el-button size="small" :class="CheckKeyHas('f6')" @click="ClickKey('f6')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">F6
      </el-button>
      <el-button size="small" :class="CheckKeyHas('f7')" @click="ClickKey('f7')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">F7
      </el-button>
    </div>
    <div style="margin-top: 3px">
      <el-button size="small" :class="CheckKeyHas('`')" @click="ClickKey('`')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">~ `
      </el-button>
      <el-button size="small" :class="CheckKeyHas('1')" @click="ClickKey('1')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">! 1
      </el-button>
      <el-button size="small" :class="CheckKeyHas('2')" @click="ClickKey('2')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">@ 2
      </el-button>
      <el-button size="small" :class="CheckKeyHas('3')" @click="ClickKey('3')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }"># 3
      </el-button>
      <el-button size="small" :class="CheckKeyHas('4')" @click="ClickKey('4')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">$ 4
      </el-button>
      <el-button size="small" :class="CheckKeyHas('5')" @click="ClickKey('5')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">% 5
      </el-button>
      <el-button size="small" :class="CheckKeyHas('6')" @click="ClickKey('6')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">^ 6
      </el-button>
      <el-button size="small" :class="CheckKeyHas('7')" @click="ClickKey('7')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">& 7
      </el-button>
    </div>
    <div style="margin-top: 1px">
      <el-button size="small" :class="CheckKeyHas('tab')" @click="ClickKey('tab')"
                 style="margin-left: 1px;width: 36px"
                 :style="{ width: 3*baseKeySize + 'px' }">tab
      </el-button>
      <el-button size="small" :class="CheckKeyHas('q')" @click="ClickKey('q')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">q
      </el-button>
      <el-button size="small" :class="CheckKeyHas('w')" @click="ClickKey('w')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">w
      </el-button>
      <el-button size="small" :class="CheckKeyHas('e')" @click="ClickKey('e')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">e
      </el-button>
      <el-button size="small" :class="CheckKeyHas('r')" @click="ClickKey('r')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">r
      </el-button>
      <el-button size="small" :class="CheckKeyHas('t')" @click="ClickKey('t')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">t
      </el-button>
      <el-button size="small" :class="CheckKeyHas('y')" @click="ClickKey('y')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 3*baseKeySize+1 + 'px' }">y
      </el-button>
    </div>
    <div style="margin-top: 1px">
      <el-button size="small" :class="CheckKeyHas('capslock')" @click="ClickKey('capslock')"
                 style="margin-left: 1px;width: 48px"
                 :style="{ width: 4*baseKeySize + 'px' }">caps
      </el-button>
      <el-button size="small" :class="CheckKeyHas('a')" @click="ClickKey('a')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">a
      </el-button>
      <el-button size="small" :class="CheckKeyHas('s')" @click="ClickKey('s')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">s
      </el-button>
      <el-button size="small" :class="CheckKeyHas('d')" @click="ClickKey('d')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">d
      </el-button>
      <el-button size="small" :class="CheckKeyHas('f')" @click="ClickKey('f')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">f
      </el-button>
      <el-button size="small" :class="CheckKeyHas('g')" @click="ClickKey('g')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">g
      </el-button>
      <el-button size="small" :class="CheckKeyHas('h')" @click="ClickKey('h')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">h
      </el-button>
    </div>
    <div style="margin-top: 1px">
      <el-button size="small" :class="CheckKeyHas('lshift')" @click="ClickKey('lshift')"
                 style="margin-left: 1px;width: 60px"
                 :style="{ width: 5*baseKeySize + 'px' }">lsft
      </el-button>
      <el-button size="small" :class="CheckKeyHas('z')" @click="ClickKey('z')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">z
      </el-button>
      <el-button size="small" :class="CheckKeyHas('x')" @click="ClickKey('x')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">x
      </el-button>
      <el-button size="small" :class="CheckKeyHas('c')" @click="ClickKey('c')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">c
      </el-button>
      <el-button size="small" :class="CheckKeyHas('v')" @click="ClickKey('v')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">v
      </el-button>
      <el-button size="small" :class="CheckKeyHas('b')" @click="ClickKey('b')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 3*baseKeySize+1 + 'px' }">b
      </el-button>

    </div>
    <div style="margin-top: 1px">
      <el-button size="small" :class="CheckKeyHas('lctrl')" @click="ClickKey('lctrl')"
                 style="margin-left: 1px;width: 36px"
                 :style="{ width: 3*baseKeySize + 'px' }">lctrl
      </el-button>
      <el-button size="small" :class="CheckKeyHas('cmd')" @click="ClickKey('cmd')"
                 style="margin-left: 1px;width: 36px"
                 :style="{ width: 3*baseKeySize + 'px' }">win
      </el-button>
      <el-button size="small" :class="CheckKeyHas('lalt')" @click="ClickKey('lalt')"
                 style="margin-left: 1px;width: 36px"
                 :style="{ width: 3*baseKeySize + 'px' }">lalt
      </el-button>
      <el-button size="small" :class="CheckKeyHas('space')" @click="ClickKey('space')"
                 style="margin-left: 1px;width: 90px"
                 :style="{ width: 7*baseKeySize+3 + 'px' }">space
      </el-button>
    </div>
  </div>
  <div style="display: flex;flex-direction: column;align-items: flex-end;margin-top: 5px">
    <div>
      <el-button size="small" :class="CheckKeyHas('f8')" @click="ClickKey('f8')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 3*baseKeySize+1 + 'px' }">F8
      </el-button>
      <el-button size="small" :class="CheckKeyHas('f9')" @click="ClickKey('f9')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">F9
      </el-button>
      <el-button size="small" :class="CheckKeyHas('f10')" @click="ClickKey('f10')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">F10
      </el-button>
      <el-button size="small" :class="CheckKeyHas('f11')" @click="ClickKey('f11')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">F11
      </el-button>
      <el-button size="small" :class="CheckKeyHas('f12')" @click="ClickKey('f12')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">F12
      </el-button>

      <el-button size="small" :class="CheckKeyHas('insert')" @click="ClickKey('insert')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">ins
      </el-button>
      <el-button size="small" :class="CheckKeyHas('delete')" @click="ClickKey('delete')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">del
      </el-button>
      <el-button size="small" :class="CheckKeyHas('f2')" @click="ClickKey('home')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">hm
      </el-button>

    </div>
    <div style="margin-top: 3px">

      <el-button size="small" :class="CheckKeyHas('8')" @click="ClickKey('8')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 3*baseKeySize+1 + 'px' }">* 8
      </el-button>
      <el-button size="small" :class="CheckKeyHas('9')" @click="ClickKey('9')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">( 9
      </el-button>
      <el-button size="small" :class="CheckKeyHas('0')" @click="ClickKey('0')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">) 0
      </el-button>
      <el-button size="small" :class="CheckKeyHas('-')" @click="ClickKey('-')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">_ -
      </el-button>
      <el-button size="small" :class="CheckKeyHas('=')" @click="ClickKey('=')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">+ =
      </el-button>
      <el-button size="small" :class="CheckKeyHas('backspace')" @click="ClickKey('backspace')"
                 style="margin-left: 1px;width: 49px"
                 :style="{ width: 4*baseKeySize+1 + 'px' }">back
      </el-button>

      <el-button size="small" :class="CheckKeyHas('end')" @click="ClickKey('end')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">end
      </el-button>

    </div>
    <div style="margin-top: 1px">

      <el-button size="small" :class="CheckKeyHas('u')" @click="ClickKey('u')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">u
      </el-button>
      <el-button size="small" :class="CheckKeyHas('i')" @click="ClickKey('i')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">i
      </el-button>
      <el-button size="small" :class="CheckKeyHas('o')" @click="ClickKey('o')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">o
      </el-button>
      <el-button size="small" :class="CheckKeyHas('p')" @click="ClickKey('p')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">p
      </el-button>
      <el-button size="small" :class="CheckKeyHas('[')" @click="ClickKey('[')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">{ [
      </el-button>
      <el-button size="small" :class="CheckKeyHas(']')" @click="ClickKey(']')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">} ]
      </el-button>
      <el-button size="small" :class="CheckKeyHas('\\')" @click="ClickKey('\\')"
                 style="margin-left: 1px;width: 37px"
                 :style="{ width: 3*baseKeySize+1 + 'px' }">| \
      </el-button>

      <el-button size="small" :class="CheckKeyHas('pageup')" @click="ClickKey('pageup')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">pgu
      </el-button>

    </div>
    <div style="margin-top: 1px">
      <el-button size="small" :class="CheckKeyHas('j')" @click="ClickKey('j')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 3*baseKeySize+1 + 'px' }">j
      </el-button>
      <el-button size="small" :class="CheckKeyHas('k')" @click="ClickKey('k')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">k
      </el-button>
      <el-button size="small" :class="CheckKeyHas('l')" @click="ClickKey('l')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">l
      </el-button>
      <el-button size="small" :class="CheckKeyHas(';')" @click="ClickKey(';')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">: ;
      </el-button>
      <el-button size="small" :class="CheckKeyHas('\'')" @click="ClickKey('\'')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">" '
      </el-button>
      <el-button size="small" :class="CheckKeyHas('enter')" @click="ClickKey('enter')"
                 style="margin-left: 1px;width: 50px"
                 :style="{ width: 4*baseKeySize+2 + 'px' }">enter
      </el-button>

      <el-button size="small" :class="CheckKeyHas('pagedown')" @click="ClickKey('pagedown')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">pgd
      </el-button>

    </div>
    <div style="margin-top: 1px">

      <el-button size="small" :class="CheckKeyHas('n')" @click="ClickKey('n')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">n
      </el-button>
      <el-button size="small" :class="CheckKeyHas('m')" @click="ClickKey('m')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">m
      </el-button>
      <el-button size="small" :class="CheckKeyHas(',')" @click="ClickKey(',')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">< ,
      </el-button>
      <el-button size="small" :class="CheckKeyHas('.')" @click="ClickKey('.')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">> .
      </el-button>
      <el-button size="small" :class="CheckKeyHas('/')" @click="ClickKey('/')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">? /
      </el-button>
      <el-button size="small" :class="CheckKeyHas('rshift')" @click="ClickKey('rshift')"
                 style="margin-left: 1px;width: 38px"
                 :style="{ width: 3*baseKeySize+2 + 'px' }">rsft
      </el-button>

      <el-button size="small" :class="CheckKeyHas('up')" @click="ClickKey('up')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">↑
      </el-button>
      <el-button size="small"
                 :disabled="true"
                 style="margin-left: 1px;width: 24px;background: #8a8a8a"
                 :style="{ width: 2*baseKeySize + 'px' }">
      </el-button>

    </div>
    <div style="margin-top: 1px">
      <el-button size="small" :class="CheckKeyHas('ralt')" @click="ClickKey('ralt')"
                 style="margin-left: 1px;width: 36px"
                 :style="{ width: 5*baseKeySize+4 + 'px' }">ralt
      </el-button>
      <el-button size="small" :class="CheckKeyHas('menu')" @click="ClickKey('menu')"
                 style="margin-left: 1px;width: 36px"
                 :style="{ width: 3*baseKeySize + 'px' }">menu
      </el-button>
      <el-button size="small" :class="CheckKeyHas('rcltl')" @click="ClickKey('rcltl')"
                 style="margin-left: 1px;width: 48px"
                 :style="{ width: 3*baseKeySize + 'px' }">rcltl
      </el-button>

      <el-button size="small" :class="CheckKeyHas('left')" @click="ClickKey('left')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">←
      </el-button>
      <el-button size="small" :class="CheckKeyHas('down')" @click="ClickKey('down')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">↓
      </el-button>
      <el-button size="small" :class="CheckKeyHas('right')" @click="ClickKey('right')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">→
      </el-button>

    </div>
  </div>
  <div style="display: flex;flex-direction: column;align-items: flex-start;margin-left: 10px;margin-top: 5px">
    <div>
      <el-button size="small" :class="CheckKeyHas('num-')" @click="ClickKey('num-')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">-
      </el-button>
      <el-button size="small" :class="CheckKeyHas('num+')" @click="ClickKey('num+')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">+
      </el-button>
    </div>
    <div style="margin-top: 3px">
      <el-button size="small" :class="CheckKeyHas('num_lock')" @click="ClickKey('num_lock')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">nlk
      </el-button>
      <el-button size="small" :class="CheckKeyHas('num/')" @click="ClickKey('num/')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">/
      </el-button>
      <el-button size="small" :class="CheckKeyHas('num*')" @click="ClickKey('num*')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">*
      </el-button>
    </div>
    <div style="margin-top: 1px">
      <el-button size="small" :class="CheckKeyHas('num7')" @click="ClickKey('num7')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">7
      </el-button>
      <el-button size="small" :class="CheckKeyHas('num8')" @click="ClickKey('num8')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">8
      </el-button>
      <el-button size="small" :class="CheckKeyHas('num9')" @click="ClickKey('num9')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">9
      </el-button>
    </div>
    <div style="margin-top: 1px">
      <el-button size="small" :class="CheckKeyHas('num4')" @click="ClickKey('num4')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">4
      </el-button>
      <el-button size="small" :class="CheckKeyHas('num5')" @click="ClickKey('num5')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">5
      </el-button>
      <el-button size="small" :class="CheckKeyHas('num6')" @click="ClickKey('num6')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">6
      </el-button>
    </div>
    <div style="margin-top: 1px">
      <el-button size="small" :class="CheckKeyHas('num1')" @click="ClickKey('num1')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">1
      </el-button>
      <el-button size="small" :class="CheckKeyHas('num2')" @click="ClickKey('num2')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">2
      </el-button>
      <el-button size="small" :class="CheckKeyHas('num3')" @click="ClickKey('num3')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">3
      </el-button>

    </div>
    <div style="margin-top: 1px">
      <el-button size="small" :class="CheckKeyHas('num0')" @click="ClickKey('num0')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 4*baseKeySize+1 + 'px' }">0
      </el-button>
      <el-button size="small" :class="CheckKeyHas('num.')" @click="ClickKey('num.')"
                 style="margin-left: 1px;width: 24px"
                 :style="{ width: 2*baseKeySize + 'px' }">.
      </el-button>

    </div>
  </div>
</div>

</template>


<script lang="ts" setup>
import "./detail.css"
import {VueDraggable} from "vue-draggable-plus";
import {
  ControlDetailKey,
  CopyControlDetailKeyList,
  NewControlDetailKey,
} from "@/components/api/detail";
import {ref, watch} from "vue";

interface Prop {
  setList: ControlDetailKey[],
  onlyOne?: boolean,
  updateList: Function,
  baseKeySize?: number,
}

const showList = ref([])
let OneKey =  ref(false)

const props = withDefaults(defineProps<Prop>(), {
  baseKeySize: 15,
  onlyOne: false,
  updateList: (showList) => {
  },
})
OneKey.value = props.onlyOne
showList.value = CopyControlDetailKeyList(props.setList)

watch(() => props.setList, (value: ControlDetailKey[]) => {
  showList.value = CopyControlDetailKeyList(value)
})

watch(() => props.onlyOne, (value: boolean) => {
  if (value){
    OneKey.value = true
  }else {
    OneKey.value = false
  }
})
const CheckKeyHas = (key: string) => {
  let css = ""
  showList.value.forEach(
      item => {
        if (item.key == key) {
          css = "detail-key-press"
          return
        }
      }
  )
  return css
}

const ClickKey = (key: string) => {
  let find = false
  showList.value.forEach((value, index) => {
    if (value.key === key) {
      showList.value.splice(index, 1)
      find = true
    }
  })
  if (OneKey.value) {
    showList.value = [] as ControlDetailKey[]
  }

  if (!find) {
    let it = NewControlDetailKey()
    it.id = new Date().getTime()
    it.key = key
    showList.value.push(it)
  }
  OnUpdate()
  // emit('updateList', showList);
}

const OnUpdate = () => {
  props.updateList(showList.value)
}
</script>