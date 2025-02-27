<template>
  <div>
    <div style="display: flex;flex-direction: column;align-items: center;justify-content: center;width: 100%">

        <div :id="item.id" v-for="(item, index) in showList" :key="item.id"
             class="detail-list-item-view"
        >
          <div style="display: flex;height: 100%">
            <div style="height: 100%;display: flex;">

              <el-button @click="ClickItemAddBtn(item,index)" size="small"
                         style="background: #b2caf5; width:40px;height: 100%;">
                <el-icon>
                  <plus/>
                </el-icon>
              </el-button>
            </div>
            <div :style="GetKeyListBackground(item.key_type)" style="flex-grow: 1" class="detail-item-desc-view sort_my_key_list">
              <div  style="display: flex;justify-content: center;height: 100%;width: 60%">
                <el-text v-show="item.key_type===KeyType.Text" style="width: 80%" truncated>
                  文本：{{ item.input }}
                </el-text>
                <el-text v-show="item.key_type===KeyType.Delay" truncated>
                  延时: {{ item.delay }} ms
                </el-text>
                <div v-show="item.key_type==KeyType.Default || item.key_type==KeyType.Mouse"
                     style="display: flex;align-items: center;justify-content: center;">
                  <el-text v-if="item.key_type==KeyType.Default">
                   按键({{GetShowKeyTouchType(item.key_type)}}}): {{ GetShowComponents(item.key) }}
                  </el-text>
                  <el-text v-if="item.key_type==KeyType.Mouse">
                   鼠标按键({{GetShowKeyTouchType(item.key_type)}}})： {{ item.key }}
                  </el-text>
                </div>
                <div v-show="item.key_type==KeyType.ShortcutKey" style="display: flex;height: 100%">
                <div  style="display: flex;align-items: center;justify-content: center;">
                 组合键：
                </div>
                    <div v-for="kkeeyy in item.key_list" class="detail-list-item-ShortcutKey-view">
                    <el-text>
                     {{ kkeeyy.key }}
                    </el-text>
                  </div>
                </div>

                <div v-show="item.key_type==KeyType.MouseMove" style="display: flex;height: 100%;">
                  <el-text>
                   鼠标定位:  X: {{ item.point_x }}   Y: {{ item.point_y }}
                  </el-text>
                </div>

                <div v-show="item.key_type==KeyType.MouseMoveStartPos" style="display: flex;height: 100%">
                  <el-text>
                    鼠标定位始起点
                  </el-text>
                </div>

                <div v-show="item.key_type==KeyType.MouseMoveSmooth" style="display: flex;height: 100%;">
                  <el-text>
                    鼠标移动:  X: {{ item.point_x }}   Y: {{ item.point_y }}
                  </el-text>
                </div>

                <div v-show="item.key_type==KeyType.MouseMoveSmoothStartPos" style="display: flex;height: 100%">
                  <el-text>
                    鼠标移动始起点
                  </el-text>
                </div>

                <div v-show="item.key_type == KeyType.MouseScroll" style="display: flex;height: 100%">
                  <el-text>
                    {{ item.scroll }}
                  </el-text>
                  <div style="margin-left: 5px">
                    <ArrowUp v-if="item.scroll_dir == MouseScrollDirection.Up"/>
                    <ArrowDown v-if="item.scroll_dir == MouseScrollDirection.Down"/>
                    <ArrowLeft v-if="item.scroll_dir == MouseScrollDirection.Left"/>
                    <ArrowRight v-if="item.scroll_dir == MouseScrollDirection.Right"/>
                  </div>
                </div>
              </div>
              <div  style="display: flex;height: 100%;width: 30%">
                <el-text>
                    {{item.remark}}
                </el-text>
              </div>
            </div>

            <div style="height: 100%;">
              <el-button @click="ClickItemEditBtn(item)"
                         class="detail-list-item-ctrl-btn-view" style="background: #ce885a;" size="small">
                <el-icon>
                  <Edit/>
                </el-icon>
              </el-button>
              <el-button @click="ClickItemDeleteBtn(item)"
                         class="detail-list-item-ctrl-btn-view" style="background: #e56565;" size="small">
                <el-icon>
                  <Close/>
                </el-icon>
              </el-button>
            </div>

          </div>

        </div>
      <div style="margin-top: 4px;width: 80%">
        <el-button style="width: 100%;background: #409eff"
                   @click="SelectEditVisible = true;AddItemPosition=-1">
          <el-icon>
            <CirclePlus/>
          </el-icon>
        </el-button>
      </div>
    </div>
  </div>


  <!-- 新增 选择类型弹框 -->
  <el-dialog
      v-model="SelectEditVisible"
      align-center
      width="auto">
    <div style="width: 100%;display: flex;align-items: center;margin-bottom: 5px">
      <el-text style="width: 100%;text-align: center">类型选择</el-text>
    </div>
    <div style="">
      <el-button @click="SelectEditMod=KeyType.Default"
                 :style="SelectEditMod==KeyType.Default? {background:'aquamarine'}:{}">单键
      </el-button>
      <el-button @click="SelectEditMod=KeyType.ShortcutKey"
                 :style="SelectEditMod==KeyType.ShortcutKey? {background:'aquamarine'}:{}">组合键
      </el-button>
      <el-button @click="SelectEditMod=KeyType.Text"
                 :style="SelectEditMod==KeyType.Text? {background:'aquamarine'}:{}">文本
      </el-button>
    </div>
    <div style="margin-top: 10px">
      <el-button @click="SelectEditMod=KeyType.Mouse"
                 :style="SelectEditMod==KeyType.Mouse? {background:'aquamarine'}:{}">鼠标按键
      </el-button>
      <el-button @click="SelectEditMod=KeyType.MouseScroll"
                 :style="SelectEditMod==KeyType.MouseScroll? {background:'aquamarine'}:{}">鼠标滚轮
      </el-button>
    </div>
    <div style="margin-top: 10px">

      <el-button @click="SelectEditMod=KeyType.MouseMove"
                 :style="SelectEditMod==KeyType.MouseMove? {background:'aquamarine'}:{}">鼠标定位
      </el-button>
      <el-button @click="SelectEditMod=KeyType.MouseMoveStartPos"
                 :style="SelectEditMod==KeyType.MouseMoveStartPos? {background:'aquamarine'}:{}">鼠标定位起点
      </el-button>
    </div>
    <div style="margin-top: 10px">
      <el-button @click="SelectEditMod=KeyType.MouseMoveSmooth"
                 :style="SelectEditMod==KeyType.MouseMoveSmooth? {background:'aquamarine'}:{}">鼠标移动
      </el-button>
      <el-button @click="SelectEditMod=KeyType.MouseMoveSmoothStartPos"
                 :style="SelectEditMod==KeyType.MouseMoveSmoothStartPos? {background:'aquamarine'}:{}">鼠标移动回起点
      </el-button>
    </div>
    <div style="margin-top: 10px">
      <el-button @click="SelectEditMod=KeyType.Delay"
                 :style="SelectEditMod==KeyType.Delay? {background:'aquamarine'}:{}">延时
      </el-button>
    </div>
    <div v-if="SelectEditMod!==KeyType.Delay" style="margin-top: 10px">
      <el-switch v-model="isAddDelay" active-text="追加延时" inactive-text="取消追加延时"/>
    </div>
    <template #footer>
      <div class="works-dialog-footer">
        <el-button @click="SelectEditVisible = false">取消</el-button>
        <el-button type="primary"
                   @click="ClickAddKeyBtn">
          确认
        </el-button>
      </div>
    </template>
  </el-dialog>


  <!-- 新增 按键框 -->
  <el-dialog
      v-model="EditItemVisible"
      align-center
      width="auto">
    <div>
      <div
          v-if="EditItemInfo.key_type==KeyType.Default || EditItemInfo.key_type==KeyType.ShortcutKey || EditItemInfo.key_type==KeyType.Mouse">
        <!-- 单键/组合键 -->
        <div v-if="EditItemInfo.key_type==KeyType.Default || EditItemInfo.key_type==KeyType.Mouse"
             style="display: flex;align-items: center;width: 100%;justify-content: center">
          <el-button @click="EditItemInfo.key_press=PressType.Click"
                     :style="EditItemInfo.key_press===PressType.Click? {background:'aquamarine'}:{}">单击
          </el-button>
          <el-button @click="EditItemInfo.key_press=PressType.DoubleClick"
                     :style="EditItemInfo.key_press===PressType.DoubleClick? {background:'aquamarine'}:{}">双击
          </el-button>
          <el-button @click="EditItemInfo.key_press=PressType.PressDown"
                     :style="EditItemInfo.key_press===PressType.PressDown? {background:'aquamarine'}:{}">按下
          </el-button>
          <el-button @click="EditItemInfo.key_press=PressType.PressUp"
                     :style="EditItemInfo.key_press===PressType.PressUp? {background:'aquamarine'}:{}">抬起
          </el-button>
        </div>
        <!-- 鼠标按键 -->
        <div v-if="EditItemInfo.key_type==KeyType.Mouse"
             style="display: flex;align-items: center;width: 100%;justify-content: center;margin-top: 10px">
          <el-button @click="EditItemInfo.key = MouseDefine.MouseLeft "
                     :style="EditItemInfo.key===MouseDefine.MouseLeft? {background:'aquamarine'}:{}">左键
          </el-button>
          <el-button @click="EditItemInfo.key=MouseDefine.MouseWheel"
                     :style="EditItemInfo.key===MouseDefine.MouseWheel? {background:'aquamarine'}:{}">中键
          </el-button>
          <el-button @click="EditItemInfo.key=MouseDefine.MouseRight"
                     :style="EditItemInfo.key===MouseDefine.MouseRight? {background:'aquamarine'}:{}">右键
          </el-button>
        </div>
        <DetailEdit v-if="EditItemInfo.key_type==KeyType.Default || EditItemInfo.key_type==KeyType.ShortcutKey"
                    :set-list="EditItemInfo.key_list" :onlyOne="EditItemInfo.key_type==KeyType.Default"
                    :update-list="list =>{OnKeyListUpdate(list)}"></DetailEdit>

      </div>
      <!-- 鼠标移动 -->
      <div v-if="EditItemInfo.key_type==KeyType.MouseMove||EditItemInfo.key_type==KeyType.MouseMoveSmooth"
           style="display: flex;align-items: center;width: 100%;justify-content: center;margin-top: 10px;flex-direction: column">
        <el-form label-width="auto">
          <el-form-item label="X">
            <el-input-number v-model="EditItemInfo.point_x"/>
          </el-form-item>
          <el-form-item label="Y">
            <el-input-number v-model="EditItemInfo.point_y"/>
          </el-form-item>
        </el-form>
        <el-button @click="ClickSetToNowMousePosition(0)">获取当前鼠标坐标</el-button>
        <el-button @click="ClickSetToNowMousePosition(2000)" style="margin-top: 5px">获取当前鼠标坐标（延时2秒）</el-button>
      </div>
      <!-- 鼠标移动 -->
      <div v-if="EditItemInfo.key_type==KeyType.MouseMoveStartPos||EditItemInfo.key_type==KeyType.MouseMoveSmoothStartPos"
           style="display: flex;align-items: center;width: 100%;justify-content: center;margin-top: 10px;flex-direction: column">
        <el-text>鼠标将回到开始起点</el-text>
      </div>
      <!-- 鼠标滚轮 -->
      <div v-if="EditItemInfo.key_type==KeyType.MouseScroll"
           style="display: flex;align-items: center;width: 100%;justify-content: center;margin-top: 10px;flex-direction: column">

        <el-form label-width="auto">
          <el-form-item label="滚动值">
            <el-input-number v-model="EditItemInfo.scroll"/>
          </el-form-item>
        </el-form>
        <div>
          <el-button @click="EditItemInfo.scroll_dir = MouseScrollDirection.Up "
                     :style="EditItemInfo.scroll_dir===MouseScrollDirection.Up? {background:'aquamarine'}:{}">向上
          </el-button>
          <el-button @click="EditItemInfo.scroll_dir = MouseScrollDirection.Down "
                     :style="EditItemInfo.scroll_dir===MouseScrollDirection.Down? {background:'aquamarine'}:{}">向下
          </el-button>
          <el-button @click="EditItemInfo.scroll_dir = MouseScrollDirection.Left "
                     :style="EditItemInfo.scroll_dir===MouseScrollDirection.Left? {background:'aquamarine'}:{}">向左
          </el-button>
          <el-button @click="EditItemInfo.scroll_dir = MouseScrollDirection.Right "
                     :style="EditItemInfo.scroll_dir===MouseScrollDirection.Right? {background:'aquamarine'}:{}">向右
          </el-button>
        </div>
      </div>
      <div v-if="EditItemInfo.key_type==KeyType.Text || EditItemInfo.key_type==KeyType.Delay">
        <el-form>
          <el-form-item v-if="EditItemInfo.key_type==KeyType.Text" label="输入文本">
            <el-input v-model="EditItemInfo.input" type="textarea" placeholder="输入文本"></el-input>
          </el-form-item>
          <el-form-item v-if="EditItemInfo.key_type==KeyType.Delay" label="延时(ms)">
            <el-input-number v-model="EditItemInfo.delay" placeholder="延时(ms)"></el-input-number>
          </el-form-item>
          <el-form-item v-if="EditItemInfo.key_type==KeyType.Delay">
            <el-button @click="EditItemInfo.delay=50">50</el-button>
            <el-button @click="EditItemInfo.delay=100">100</el-button>
            <el-button @click="EditItemInfo.delay=200">200</el-button>
            <el-button @click="EditItemInfo.delay=500">500</el-button>
            <el-button @click="EditItemInfo.delay=1000">1000</el-button>
            <el-button @click="EditItemInfo.delay=3000">3000</el-button>
            <el-button @click="EditItemInfo.delay=5000">5000</el-button>
          </el-form-item>
        </el-form>

      </div>

      <div style="margin-top: 10px">
        <el-form-item label="备注" style="display: flex;align-items: center;">
          <el-input v-model="EditItemInfo.remark"></el-input>
        </el-form-item>

      </div>
    </div>

    <template #footer>
      <div class="works-dialog-footer">
        <el-button @click="EditItemVisible = false">取消</el-button>
        <el-button type="primary"
                   @click="ClickAddDefaultKeyBtn">
          确认
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>

import {
  ControlDetailKey,
  ControlKeyList,
  ControlType,
  CopyControlDetailKeyList,
  CopyControlKeyList,
  CopyControlKeyListList,
  GetNowMousePosition, GetNowMousePositionResp, GetShowKeyTouchType,
  KeyType,
  NewControlDetailKey,
  NewControlDetailKeyList,
  NewControlKeyList,
  PressType
} from "@/components/api/detail.js";
import {ref, watch} from "vue";
import {
  CirclePlus,
  Close,
  Delete,
  Document, Download,
  Edit,
  Key,
  Plus,
  Sort,
  SortDown,
  SortUp,
  Timer, Upload
} from "@element-plus/icons-vue";
import {
  Click, MoveInOne, Mouse, Keyboard, KeyboardOne, ClickTap, ClickTapTwo,
  ArrowUp, ArrowDown, ArrowLeft, ArrowRight
} from "@icon-park/vue-next";
import {VueDraggable} from "vue-draggable-plus";
import DetailEdit from "@/components/detail/detailEdit.vue";
import {GetNewId} from "@/components/common/id";
import {GetShowComponents, MouseDefine, MouseScrollDirection} from "@/components/api/keyDefine";


interface Prop {
  setList: ControlKeyList[],
  updateList: Function,
  baseKeySize?: number,
}

const showList = ref([] as ControlKeyList[])

const props = withDefaults(defineProps<Prop>(), {
  baseKeySize: 15,
  updateList: (showList) => {
  },
})

const SelectEditVisible = ref(false)
const SelectEditMod = ref(KeyType.Default)
const EditItemInfo = ref({} as ControlKeyList)

const EditItemVisible = ref(false)

const AddItemPosition = ref(-1)
const isAddDelay = ref(true)

showList.value = CopyControlKeyListList(props.setList)

watch(() => props.setList, (value: ControlKeyList[]) => {
  showList.value = CopyControlKeyListList(value)
})


// 方法
const ClickAddDefaultKeyBtn = () => {
  if (EditItemInfo.value.id === 0) {
    EditItemInfo.value.id = GetNewId()
    if (AddItemPosition.value === -1) {
      // 添加到最后
      showList.value.push(EditItemInfo.value)
      if (isAddDelay.value && EditItemInfo.value.key_type !== KeyType.Delay) {
        showList.value.push(NewNormalDetail())
      }
    } else {
      // 说明是中途插入
      if (isAddDelay.value && EditItemInfo.value.key_type !== KeyType.Delay) {
        showList.value.splice(AddItemPosition.value, 0, EditItemInfo.value, NewNormalDetail())
      } else {
        showList.value.splice(AddItemPosition.value, 0, EditItemInfo.value)
      }
    }
  } else {
    showList.value.forEach(
        (res, index) => {
          if (res.id == EditItemInfo.value.id) {
            showList.value[index] = CopyControlKeyList(EditItemInfo.value)
            return
          }
        }
    )
  }

  EditItemVisible.value = false
  OnUpdate()
  console.log(showList.value)

}

const ClickSetToNowMousePosition = (tm) => {
  setTimeout(() => {
    GetNowMousePosition().then((res) => {
      if (res.state != 0) {
      } else {
        const data = res.data as GetNowMousePositionResp
        EditItemInfo.value.point_x = data.x
        EditItemInfo.value.point_y = data.y
      }
    })
  }, tm); // 3000毫秒 = 3秒
}
const ClickAddKeyBtn = () => {
  SelectEditVisible.value = false
  EditItemInfo.value = NewControlKeyList()
  EditItemInfo.value.key_type = SelectEditMod.value

  switch (EditItemInfo.value.key_type) {
    case KeyType.Mouse:
      EditItemInfo.value.key = MouseDefine.MouseLeft
      EditItemInfo.value.key_press = PressType.Click
      break
    case KeyType.MouseScroll:
      EditItemInfo.value.scroll_dir = MouseScrollDirection.Down
      EditItemInfo.value.scroll = 5
      break
    case KeyType.Default:
      EditItemInfo.value.key_press = PressType.Click
      break
    case KeyType.Delay:
      EditItemInfo.value.delay = 50
      break
  }

  EditItemVisible.value = true
}
const GetKeyListBackground = (keytype: number) => {
  let sty = {
    background: ''
  }
  switch (keytype) {
    case KeyType.Default:
      sty.background = '#c6e2ff'
      break
    case KeyType.ShortcutKey:
      sty.background = '#a0cfff'
      break
    case KeyType.Text:
      sty.background = '#f3d19e'
      break
    case KeyType.Mouse:
      sty.background = '#95d475'
      break
    case KeyType.MouseMove:
      sty.background = '#e1f3d8'
      break
    case KeyType.MouseMoveSmooth:
      sty.background = '#d4efc7'
      break
    case KeyType.MouseScroll:
      sty.background = '#d1edc4'
      break
    case KeyType.Delay:
      sty.background = '#fcd3d3'
      break
    case KeyType.MouseMoveStartPos:
      sty.background = '#c9c7c7'
      break
    case KeyType.MouseMoveSmoothStartPos:
      sty.background = '#cabfbf'
      break
  }
  return sty
}

const OnKeyListUpdate = (list: ControlDetailKey[]) => {
  if (EditItemInfo.value.key_type == KeyType.Default) {
    EditItemInfo.value.key = ""
    if (list.length > 0) {
      EditItemInfo.value.key = list[0].key
    }
  } else {
    EditItemInfo.value.key_list = list
  }
  OnUpdate()
}
const OnKeyListOrderUpdate = () => {
  OnUpdate()
}
const OnUpdate = () => {
  props.updateList(showList.value)
}

const NewNormalDetail = (): ControlKeyList => {
  let delay = NewControlKeyList()
  delay.key_type = KeyType.Delay
  delay.id = GetNewId()
  delay.delay = 50
  return delay
}

const ClickItemAddBtn = (item: ControlKeyList, index: number) => {
  AddItemPosition.value = index
  SelectEditVisible.value = true

}

const EditItemRemark = (item: ControlKeyList) => {
  // 变化了
  showList.value.forEach(
      (res, index) => {
        if (res.id == item.id) {
          showList.value[index].remark = item.remark
          return
        }
      }
  )
}

const ClickItemEditBtn = (item: ControlKeyList) => {
  EditItemInfo.value = CopyControlKeyList(item)
  if (EditItemInfo.value.key_type === KeyType.Default) {
    EditItemInfo.value.key_list = []
    EditItemInfo.value.key_list.push({
      id: GetNewId(),
      key: EditItemInfo.value.key,
    })
  }
  EditItemVisible.value = true
}

const ClickItemDeleteBtn = (item: ControlKeyList) => {
  showList.value.forEach((value, index) => {
    if (value.id === item.id) {
      showList.value.splice(index, 1)
    }
  })

  OnUpdate()
}


</script>
