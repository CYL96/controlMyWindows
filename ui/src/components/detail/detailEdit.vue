<template>
  <div >
    <div style="display: flex;height: 90%;max-width: 90vw;align-items: center;justify-content: center;">
        <VueDraggable
            v-model="showList"
            :animation="150"
            id="LabelList"
            :on-update="OnUpdate"
            style="display: flex;flex-wrap: wrap;"
        >
          <div v-for="item in showList" :key="item.id"
               class="detail-key-item">
            <div style="text-align: center">
              <div >
                <el-button  @click="ClickDelKey(item)"
                    style="height: 20px;width: 100%;background: rgba(231,34,34,0.47)">
                  <Icon.Delete />
                </el-button>
              </div>
              <div class="detail-list-item-ShortcutKey-view">
                <el-text>
                  {{ GetShowComponents(item.key) }}
                </el-text>
              </div>
            </div>
          </div>
        </VueDraggable>
    </div>
  </div>
  <div style="background: #e9e9eb;padding: 3px;margin-top: 3px">
    <div style="display: flex;flex-direction: row;align-items: center;justify-content: center;">
      <div style="margin-top: 1px">

        <el-button size="small" :class="CheckKeyHas(KeyDefine.AudioVolDown)" @click="ClickKey(KeyDefine.AudioVolDown)"
                   style="margin-left: 1px;width: 24px"
                   :style="{ width: 2*baseKeySize + 'px' }">
          {{ GetShowComponents(KeyDefine.AudioVolDown) }}
        </el-button>
        <el-button size="small" :class="CheckKeyHas(KeyDefine.AudioMute)" @click="ClickKey(KeyDefine.AudioMute)"
                   style="margin-left: 1px;width: 24px;" :disabled="false"
                   :style="{ width: 2*baseKeySize + 'px' }">
          {{ GetShowComponents(KeyDefine.AudioMute) }}
        </el-button>
        <el-button size="small" :class="CheckKeyHas(KeyDefine.AudioVolUp)" @click="ClickKey(KeyDefine.AudioVolUp)"
                   style="margin-left: 1px;width: 24px"
                   :style="{ width: 2*baseKeySize + 'px' }">
          {{ GetShowComponents(KeyDefine.AudioVolUp) }}
        </el-button>
        <el-button size="small" :class="CheckKeyHas(KeyDefine.AudioPrev)" @click="ClickKey(KeyDefine.AudioPrev)"
                   style="margin-left: 1px;width: 24px"
                   :style="{ width: 2*baseKeySize + 'px' }">
          {{ GetShowComponents(KeyDefine.AudioPrev) }}
        </el-button>
        <el-button size="small" :class="CheckKeyHas(KeyDefine.AudioStop)" @click="ClickKey(KeyDefine.AudioStop)"
                   style="margin-left: 1px;width: 24px"
                   :style="{ width: 2*baseKeySize + 'px' }">
          {{ GetShowComponents(KeyDefine.AudioStop) }}
        </el-button>
        <el-button size="small" :class="CheckKeyHas(KeyDefine.AudioPlay)" @click="ClickKey(KeyDefine.AudioPlay)"
                   style="margin-left: 1px;width: 24px"
                   :style="{ width: 2*baseKeySize + 'px' }">
          {{ GetShowComponents(KeyDefine.AudioPlay) }}
        </el-button>
        <el-button size="small" :class="CheckKeyHas(KeyDefine.AudioPause)" @click="ClickKey(KeyDefine.AudioPause)"
                   style="margin-left: 1px;width: 24px"
                   :style="{ width: 2*baseKeySize + 'px' }">
          {{ GetShowComponents(KeyDefine.AudioPause) }}
        </el-button>

        <el-button size="small" :class="CheckKeyHas(KeyDefine.AudioNext)" @click="ClickKey(KeyDefine.AudioNext)"
                   style="margin-left: 1px;width: 24px"
                   :style="{ width: 2*baseKeySize + 'px' }">
          {{ GetShowComponents(KeyDefine.AudioNext) }}
        </el-button>
      </div>
    </div>
    <div style="display: flex;flex-direction: row;align-items: center;justify-content: center;flex-wrap:wrap">
      <div style="display: flex;flex-direction: column;align-items: flex-start;margin-top: 5px">
        <div>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Esc)" @click="ClickKey(KeyDefine.Esc)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Esc) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.F1)" @click="ClickKey(KeyDefine.F1)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.F1) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.F2)" @click="ClickKey(KeyDefine.F2)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.F2) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.F3)" @click="ClickKey(KeyDefine.F3)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.F3) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.F4)" @click="ClickKey(KeyDefine.F4)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.F4) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.F5)" @click="ClickKey(KeyDefine.F5)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.F5) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.F6)" @click="ClickKey(KeyDefine.F6)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.F6) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.F7)" @click="ClickKey(KeyDefine.F7)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.F7) }}
          </el-button>
        </div>
        <div style="margin-top: 3px">
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Spec0)" @click="ClickKey(KeyDefine.Spec0)"
                     style="margin-left: 1px;width: 24px;" :disabled="false"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Spec0) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Key1)" @click="ClickKey(KeyDefine.Key1)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Key1) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Key2)" @click="ClickKey(KeyDefine.Key2)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Key2) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Key3)" @click="ClickKey(KeyDefine.Key3)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Key3) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Key4)" @click="ClickKey(KeyDefine.Key4)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Key4) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Key5)" @click="ClickKey(KeyDefine.Key5)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Key5) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Key6)" @click="ClickKey(KeyDefine.Key6)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Key6) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Key7)" @click="ClickKey(KeyDefine.Key7)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Key7) }}
          </el-button>
        </div>
        <div style="margin-top: 1px">
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Tab)" @click="ClickKey(KeyDefine.Tab)"
                     style="margin-left: 1px;width: 36px"
                     :style="{ width: 3*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Tab) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyQ)" @click="ClickKey(KeyDefine.KeyQ)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.KeyQ) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyW)" @click="ClickKey(KeyDefine.KeyW)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.KeyW) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyE)" @click="ClickKey(KeyDefine.KeyE)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.KeyE) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyR)" @click="ClickKey(KeyDefine.KeyR)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.KeyR) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyT)" @click="ClickKey(KeyDefine.KeyT)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.KeyT) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyY)" @click="ClickKey(KeyDefine.KeyY)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 3*baseKeySize+1 + 'px' }">{{ GetShowComponents(KeyDefine.KeyY) }}
          </el-button>
        </div>
        <div style="margin-top: 1px">
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Capslock)" @click="ClickKey(KeyDefine.Capslock)"
                     style="margin-left: 1px;width: 48px"
                     :style="{ width: 4*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Capslock) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyA)" @click="ClickKey(KeyDefine.KeyA)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.KeyA) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyS)" @click="ClickKey(KeyDefine.KeyS)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.KeyS) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyD)" @click="ClickKey(KeyDefine.KeyD)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.KeyD) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyF)" @click="ClickKey(KeyDefine.KeyF)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.KeyF) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyG)" @click="ClickKey(KeyDefine.KeyG)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.KeyG) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyH)" @click="ClickKey(KeyDefine.KeyH)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.KeyH) }}
          </el-button>
        </div>
        <div style="margin-top: 1px">
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Lshift)" @click="ClickKey(KeyDefine.Lshift)"
                     style="margin-left: 1px;width: 60px"
                     :style="{ width: 5*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Lshift) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyZ)" @click="ClickKey(KeyDefine.KeyZ)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.KeyZ) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyX)" @click="ClickKey(KeyDefine.KeyX)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.KeyX) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyC)" @click="ClickKey(KeyDefine.KeyC)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.KeyC) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyV)" @click="ClickKey(KeyDefine.KeyV)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.KeyV) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyB)" @click="ClickKey(KeyDefine.KeyB)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 3*baseKeySize+1 + 'px' }">{{ GetShowComponents(KeyDefine.KeyB) }}
          </el-button>

        </div>
        <div style="margin-top: 1px">
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Lctrl)" @click="ClickKey(KeyDefine.Lctrl)"
                     style="margin-left: 1px;width: 36px"
                     :style="{ width: 3*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Lctrl) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Lcmd)" @click="ClickKey(KeyDefine.Lcmd)"
                     style="margin-left: 1px;width: 36px"
                     :style="{ width: 3*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Lcmd) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Lalt)" @click="ClickKey(KeyDefine.Lalt)"
                     style="margin-left: 1px;width: 36px"
                     :style="{ width: 3*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Lalt) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Space)" @click="ClickKey(KeyDefine.Space)"
                     style="margin-left: 1px;width: 90px"
                     :style="{ width: 7*baseKeySize+3 + 'px' }">{{ GetShowComponents(KeyDefine.Space) }}
          </el-button>
        </div>
      </div>
      <div style="display: flex;flex-direction: column;align-items: flex-start;margin-top: 5px">
        <div>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.F8)" @click="ClickKey(KeyDefine.F8)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 3*baseKeySize+1 + 'px' }">{{ GetShowComponents(KeyDefine.F8) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.F9)" @click="ClickKey(KeyDefine.F9)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.F9) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.F10)" @click="ClickKey(KeyDefine.F10)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.F10) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.F11)" @click="ClickKey(KeyDefine.F11)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.F11) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.F12)" @click="ClickKey(KeyDefine.F12)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.F12) }}
          </el-button>

          <el-button size="small" :class="CheckKeyHas(KeyDefine.Insert)" @click="ClickKey(KeyDefine.Insert)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Insert) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Delete)" @click="ClickKey(KeyDefine.Delete)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Delete) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Home)" @click="ClickKey(KeyDefine.Home)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Home) }}
          </el-button>

        </div>
        <div style="margin-top: 3px">

          <el-button size="small" :class="CheckKeyHas(KeyDefine.Key8)" @click="ClickKey(KeyDefine.Key8)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 3*baseKeySize+1 + 'px' }">{{ GetShowComponents(KeyDefine.Key8) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Key9)" @click="ClickKey(KeyDefine.Key9)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Key9) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Key0)" @click="ClickKey(KeyDefine.Key0)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Key0) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Spec1)" @click="ClickKey(KeyDefine.Spec1)"
                     style="margin-left: 1px;width: 24px;" :disabled="false"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Spec1) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Spec2)" @click="ClickKey(KeyDefine.Spec2)"
                     style="margin-left: 1px;width: 24px;" :disabled="false"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Spec2) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Backspace)" @click="ClickKey(KeyDefine.Backspace)"
                     style="margin-left: 1px;width: 49px"
                     :style="{ width: 4*baseKeySize+1 + 'px' }">{{ GetShowComponents(KeyDefine.Backspace) }}
          </el-button>

          <el-button size="small" :class="CheckKeyHas(KeyDefine.End)" @click="ClickKey(KeyDefine.End)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.End) }}
          </el-button>

        </div>
        <div style="margin-top: 1px">

          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyU)" @click="ClickKey(KeyDefine.KeyU)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.KeyU) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyI)" @click="ClickKey(KeyDefine.KeyI)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.KeyI) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyO)" @click="ClickKey(KeyDefine.KeyO)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.KeyO) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyP)" @click="ClickKey(KeyDefine.KeyP)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.KeyP) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Spec3)" @click="ClickKey(KeyDefine.Spec3)"
                     style="margin-left: 1px;width: 24px;" :disabled="false"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Spec3) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Spec4)" @click="ClickKey(KeyDefine.Spec4)"
                     style="margin-left: 1px;width: 24px;" :disabled="false"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Spec4) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Spec5)" @click="ClickKey(KeyDefine.Spec5)"
                     style="margin-left: 1px;width: 37px;" :disabled="false"
                     :style="{ width: 3*baseKeySize+1 + 'px' }">{{ GetShowComponents(KeyDefine.Spec5) }}
          </el-button>

          <el-button size="small" :class="CheckKeyHas(KeyDefine.Pageup)" @click="ClickKey(KeyDefine.Pageup)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Pageup) }}
          </el-button>

        </div>
        <div style="margin-top: 1px">
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyJ)" @click="ClickKey(KeyDefine.KeyJ)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 3*baseKeySize+1 + 'px' }">{{ GetShowComponents(KeyDefine.KeyJ) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyK)" @click="ClickKey(KeyDefine.KeyK)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.KeyK) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyL)" @click="ClickKey(KeyDefine.KeyL)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.KeyL) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Spec6)" @click="ClickKey(KeyDefine.Spec6)"
                     style="margin-left: 1px;width: 24px;" :disabled="false"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Spec6) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Spec7)" @click="ClickKey(KeyDefine.Spec7)"
                     style="margin-left: 1px;width: 24px;" :disabled="false"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Spec7) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Enter)" @click="ClickKey(KeyDefine.Enter)"
                     style="margin-left: 1px;width: 50px"
                     :style="{ width: 4*baseKeySize+1 + 'px' }">{{ GetShowComponents(KeyDefine.Enter) }}
          </el-button>

          <el-button size="small" :class="CheckKeyHas(KeyDefine.Pagedown)" @click="ClickKey(KeyDefine.Pagedown)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Pagedown) }}
          </el-button>

        </div>
        <div style="margin-top: 1px">

          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyN)" @click="ClickKey(KeyDefine.KeyN)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.KeyN) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.KeyM)" @click="ClickKey(KeyDefine.KeyM)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.KeyM) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Spec8)" @click="ClickKey(KeyDefine.Spec8)"
                     style="margin-left: 1px;width: 24px;" :disabled="false"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Spec8) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Spec9)" @click="ClickKey(KeyDefine.Spec9)"
                     style="margin-left: 1px;width: 24px;" :disabled="false"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Spec9) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Spec10)" @click="ClickKey(KeyDefine.Spec10)"
                     style="margin-left: 1px;width: 24px;" :disabled="false"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Spec10) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Rshift)" @click="ClickKey(KeyDefine.Rshift)"
                     style="margin-left: 1px;width: 38px"
                     :style="{ width: 3*baseKeySize+1 + 'px' }">{{ GetShowComponents(KeyDefine.Rshift) }}
          </el-button>

          <el-button size="small" :class="CheckKeyHas(KeyDefine.Up)" @click="ClickKey(KeyDefine.Up)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Up) }}
          </el-button>
          <!--      <el-button size="small"-->
          <!--                 :disabled="true"-->
          <!--                 style="margin-left: 1px;width: 24px;background: #8a8a8a"-->
          <!--                 :style="{ width: 2*baseKeySize + 'px' }">-->
          <!--      </el-button>-->

        </div>
        <div style="margin-top: 1px">
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Ralt)" @click="ClickKey(KeyDefine.Ralt)"
                     style="margin-left: 1px;width: 36px"
                     :style="{ width: 3*baseKeySize+1 + 'px' }">{{ GetShowComponents(KeyDefine.Ralt) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Rcmd)" @click="ClickKey(KeyDefine.Rcmd)"
                     style="margin-left: 1px;width: 36px"
                     :style="{ width: 2.5*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Rcmd) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Menu)" @click="ClickKey(KeyDefine.Menu)"
                     style="margin-left: 1px;width: 36px"
                     :style="{ width: 2.5*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Menu) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Rctrl)" @click="ClickKey(KeyDefine.Rctrl)"
                     style="margin-left: 1px;width: 48px"
                     :style="{ width: 3*baseKeySize+1 + 'px' }">{{ GetShowComponents(KeyDefine.Rctrl) }}
          </el-button>

          <el-button size="small" :class="CheckKeyHas(KeyDefine.Left)" @click="ClickKey(KeyDefine.Left)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Left) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Down)" @click="ClickKey(KeyDefine.Down)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Down) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Right)" @click="ClickKey(KeyDefine.Right)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Right) }}
          </el-button>

        </div>
      </div>
      <div style="display: flex;flex-direction: column;align-items: flex-start;margin-left: 10px;margin-top: 5px">
        <div>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.NumMinus)" @click="ClickKey(KeyDefine.NumMinus)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.NumMinus) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.NumPlus)" @click="ClickKey(KeyDefine.NumPlus)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.NumPlus) }}
          </el-button>
        </div>
        <div style="margin-top: 3px">
          <el-button size="small" :class="CheckKeyHas(KeyDefine.NumLock)" @click="ClickKey(KeyDefine.NumLock)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.NumLock) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.NumDiv)" @click="ClickKey(KeyDefine.NumDiv)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.NumDiv) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.NumMul)" @click="ClickKey(KeyDefine.NumMul)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.NumMul) }}
          </el-button>
        </div>
        <div style="margin-top: 1px">
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Num7)" @click="ClickKey(KeyDefine.Num7)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Num7) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Num8)" @click="ClickKey(KeyDefine.Num8)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Num8) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Num9)" @click="ClickKey(KeyDefine.Num9)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Num9) }}
          </el-button>
        </div>
        <div style="margin-top: 1px">
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Num4)" @click="ClickKey(KeyDefine.Num4)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Num4) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Num5)" @click="ClickKey(KeyDefine.Num5)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Num5) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Num6)" @click="ClickKey(KeyDefine.Num6)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Num6) }}
          </el-button>
        </div>
        <div style="margin-top: 1px">
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Num1)" @click="ClickKey(KeyDefine.Num1)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Num1) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Num2)" @click="ClickKey(KeyDefine.Num2)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Num2) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Num3)" @click="ClickKey(KeyDefine.Num3)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.Num3) }}
          </el-button>

        </div>
        <div style="margin-top: 1px">
          <el-button size="small" :class="CheckKeyHas(KeyDefine.Num0)" @click="ClickKey(KeyDefine.Num0)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 4*baseKeySize+1 + 'px' }">{{ GetShowComponents(KeyDefine.Num0) }}
          </el-button>
          <el-button size="small" :class="CheckKeyHas(KeyDefine.NumDecimal)" @click="ClickKey(KeyDefine.NumDecimal)"
                     style="margin-left: 1px;width: 24px"
                     :style="{ width: 2*baseKeySize + 'px' }">{{ GetShowComponents(KeyDefine.NumDecimal) }}
          </el-button>

        </div>
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
import {GetNewId} from "@/components/common/id";
import {GetShowComponents, KeyDefine} from "@/components/api/keyDefine";
import * as Icon from "@icon-park/vue-next"
import {ElScrollbar} from "element-plus";
import {Delete} from "@element-plus/icons-vue";

interface Prop {
  setList: ControlDetailKey[],
  onlyOne?: boolean,
  updateList: Function,
  baseKeySize?: number,
}


const showList = ref([])
const OneKey = ref(false)

const props = withDefaults(defineProps<Prop>(), {
  baseKeySize: 18,
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
  if (value) {
    OneKey.value = true
  } else {
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
    it.id = GetNewId()
    it.key = key
    showList.value.push(it)
  }
  OnUpdate()
  // emit('updateList', showList);
}

const ClickDelKey = (item :ControlDetailKey)=>{
  showList.value.forEach((value, index) => {
    if (value.id === item.id) {
      showList.value.splice(index, 1)
    }
  })
}
const OnUpdate = () => {
  props.updateList(showList.value)
}
</script>