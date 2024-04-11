<template>
  <div>
    <div style="display: flex;flex-direction: column;align-items: center;justify-content: center;width: 100%">
      <VueDraggable
          v-model="showList"
          :animation="150"
          @update="OnKeyListOrderUpdate"
          handle=".sort_my_key_list"
          id="LabelList"
          style="display: flex"
          class="detail-list-script-view"
      >
        <div :id="item.id" v-for="(item, index) in showList" :key="item.id"
             class="detail-list-item-view" :style="GetKeyListBackground(item.key_type)"
        >
          <el-icon class="sort_my_key_list" style="height:100%;width: 30px;background: #b3e19d;">
            <Sort/>
          </el-icon>

          <div class="detail-item-icon-view" style="margin-left: 5px">
            <el-icon>
              <Key v-show="item.key_type===KeyType.Default"/>
              <Key v-show="item.key_type===KeyType.ShortcutKey"/>
              <Key v-show="item.key_type===KeyType.ShortcutKey"/>
              <Key v-show="item.key_type===KeyType.ShortcutKey"/>
              <Document v-show="item.key_type===KeyType.Text"/>
              <Timer v-show="item.key_type===KeyType.Delay"/>
            </el-icon>
          </div>

          <div class="detail-item-desc-view">

            <el-text v-show="item.key_type===KeyType.Text" style="width: 80%" truncated>
              {{ item.input }}
            </el-text>
            <el-text v-show="item.key_type===KeyType.Delay" truncated>
              {{ item.delay }}
            </el-text>
            <div v-show="item.key_type==KeyType.Default"
                 style="display: flex;align-items: center;justify-content: center;">
              <el-text>
                {{ item.key }}
              </el-text>
              <el-icon :size="24" style="margin-left: 10px; background: #a0b3c7">
                <Download v-show="item.key_press===PressType.Click|| item.key_press===PressType.DoubleClick "/>
                <Upload v-show="item.key_press===PressType.Click|| item.key_press===PressType.DoubleClick "/>

                <Download v-show="item.key_press===PressType.DoubleClick"/>
                <Upload v-show="item.key_press===PressType.DoubleClick"/>

                <Download v-show="item.key_press===PressType.PressDown"/>
                <Upload v-show="item.key_press===PressType.PressUp"/>
              </el-icon>
            </div>
            <div v-show="item.key_type==KeyType.ShortcutKey" style="display: flex;height: 100%">
              <div v-for="kkeeyy in item.key_list" class="detail-list-item-ShortcutKey-view">
                <el-text>
                  {{ kkeeyy.key }}
                </el-text>
              </div>
            </div>

          </div>


          <div>
            <div class="detail-list-item-ctrl-view">
              <el-button @click="ClickItemAddBtn(item,index)"
                         class="detail-list-item-ctrl-btn-view" style="background: #b2caf5" size="default">
                <el-icon>
                  <Plus/>
                </el-icon>
              </el-button>
              <el-button @click="ClickItemEditBtn(item)"
                         class="detail-list-item-ctrl-btn-view" style="background: #ce885a" size="default">
                <el-icon>
                  <Edit/>
                </el-icon>
              </el-button>
              <el-button @click="ClickItemDeleteBtn(item)"
                         class="detail-list-item-ctrl-btn-view" style="background: #e56565" size="default">
                <el-icon>
                  <Close/>
                </el-icon>
              </el-button>
            </div>
          </div>
        </div>
      </VueDraggable>
      <div style="margin-top: 4px;width: 80%">
        <el-button style="width: 100%;background: #409eff"
                   @click="  SelectEditMod = KeyType.Default;SelectEditVisible = true;AddItemPosition=-1">
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
    <div>
      <el-button @click="SelectEditMod=KeyType.Default"
                 :style="SelectEditMod==KeyType.Default? {background:'aquamarine'}:{}">单键
      </el-button>
      <el-button @click="SelectEditMod=KeyType.ShortcutKey"
                 :style="SelectEditMod==KeyType.ShortcutKey? {background:'aquamarine'}:{}">组合键
      </el-button>
      <el-button @click="SelectEditMod=KeyType.Text"
                 :style="SelectEditMod==KeyType.Text? {background:'aquamarine'}:{}">文本
      </el-button>
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
      <div v-if="EditItemInfo.key_type==KeyType.Default || EditItemInfo.key_type==KeyType.ShortcutKey">
        <!-- 单键/组合键 -->
        <div v-if="EditItemInfo.key_type==KeyType.Default"
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
        <DetailEdit
            :set-list="EditItemInfo.key_list" :onlyOne="EditItemInfo.key_type==KeyType.Default"
            :update-list="list =>{OnKeyListUpdate(list)}"></DetailEdit>
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
  ControlKeyList, ControlType,
  CopyControlDetailKeyList, CopyControlKeyList,
  CopyControlKeyListList, KeyType, NewControlDetailKey, NewControlDetailKeyList, NewControlKeyList, PressType
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
import {VueDraggable} from "vue-draggable-plus";
import DetailEdit from "@/components/detail/detailEdit.vue";


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
const SelectEditMod = ref(0)
const EditItemInfo = ref({} as ControlKeyList)

const EditItemVisible = ref(false)

const AddItemPosition = ref(-1)
const isAddDelay = ref(true)

showList.value = CopyControlKeyListList(props.setList)

watch(() => props.setList, (value: ControlDetailKey[]) => {
  showList.value = CopyControlKeyListList(value)
})


// 方法
const ClickAddDefaultKeyBtn = () => {
  console.log(EditItemInfo.value)
  if (EditItemInfo.value.id === 0) {
    EditItemInfo.value.id = new Date().getTime()
    if (AddItemPosition.value === -1) {
      showList.value.push(EditItemInfo.value)
      if (isAddDelay.value && EditItemInfo.value.key_type !== KeyType.Delay) {
        showList.value.push(NewNormalDetail())
      }
    } else {
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
  isAddDelay.value
  OnUpdate()
}
const ClickAddKeyBtn = () => {
  SelectEditVisible.value = false
  EditItemInfo.value = NewControlKeyList()
  EditItemInfo.value.key_type = SelectEditMod.value

  switch (EditItemInfo.value.key_type) {
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
      sty.background = '#82d900'
      break
    case KeyType.Text:
      sty.background = '#c08080'
      break
    case KeyType.ShortcutKey:
      sty.background = '#72e0a9'
      break
    case KeyType.Delay:
      sty.background = '#e0e5ee'
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
  delay.id = new Date().getTime()
  delay.delay = 50
  return delay
}

const ClickItemAddBtn = (item: ControlKeyList, index: number) => {
  AddItemPosition.value = index
  SelectEditVisible.value = true

}

const ClickItemEditBtn = (item: ControlKeyList) => {
  EditItemInfo.value = CopyControlKeyList(item)
  if (EditItemInfo.value.key_type === KeyType.Default) {
    EditItemInfo.value.key_list = []
    EditItemInfo.value.key_list.push({
      id: new Date().getTime(),
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
