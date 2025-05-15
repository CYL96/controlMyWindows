<template>
  <div style="display: flex;flex-direction: column;height: 100%;width: 100%">
    <header style="height: 5%;padding: 1% 0 1% 0;min-height: 30px;width: 100%">
      <el-row :gutter="5" style="padding: 5px;width: 98%">
        <el-col :span="6">
          <el-button @click="GotoHome" style="width: 100%">
            <el-icon>
              <Back/>
            </el-icon>
          </el-button>
        </el-col>
        <el-col :span="4">
          <el-button @click="GetDetailList" style="width: 100%">
            <el-icon>
              <Refresh/>
            </el-icon>
          </el-button>
        </el-col>

        <el-col v-show="!isEditMod" :span="4">
          <el-button :disabled="isEditMod" @click="setBindNormalKeyVisibleVisible = true;setBindNormalKeyVisibleState=1"
                     style="width: 100%;background:#00ffd4">
            <el-icon>
              <Link/>
            </el-icon>
          </el-button>
        </el-col>

        <el-col v-show="!isEditMod" :span="4">
          <el-button @click="setBindNormalKeyVisibleVisible = true;setBindNormalKeyVisibleState=2"
                     style="width: 100%;background:#e72222">
            <el-icon>
              <Link/>
            </el-icon>
          </el-button>
        </el-col>
        <!--        编辑按钮-->
        <el-col v-show="!isEditMod" :span="6">
          <el-button @click="isEditMod= true" style="width: 100%">
            <el-icon>
              <Edit/>
            </el-icon>
          </el-button>
        </el-col>
        <el-col v-show="isEditMod" :span="4">
          <el-button  @click="selectDialogVisible=true;"
                     style="width: 100%;background:deepskyblue">
            <el-icon>
              <CirclePlus/>
            </el-icon>
          </el-button>
        </el-col>
        <el-col v-show="isEditMod" :span="4">
          <el-button @click="ClickEditSizeBtn" style="width: 100%;background: #ffce2f">
            <el-icon>
              <FullScreen/>
            </el-icon>
          </el-button>
        </el-col>
        <el-col v-show="isEditMod" :span="6">
          <el-button @click="isEditMod=false" style="width: 100%;background: greenyellow">
            <el-icon>
              <Check/>
            </el-icon>
          </el-button>
        </el-col>
      </el-row>
    </header>
    <main style="display: flex;overflow-y: scroll;height: 95%;width: 100%">
      <!-- 左边的列表 -->
      <div style="width: 100%;height: 100%;display: flex;flex-direction: column;align-items: center;">
        <div v-if="isControlInfoGet" style="display: flex;  width: 96%">
          <VueDraggable
              v-model="detailList"
              :animation="150"
              @update="OnDetailListOrderChange()"
              :disabled="!isEditMod"
              handle=".sort_detail_key_list"
              class="detail-list-view"
              id="LabelList"
          >
            <div :id="item.detail_id" v-for="item in detailList" :key="item.detail_id"
                 class="detail-item ">
              <div v-if="isEditMod" class="sort_detail_key_list"
                   style="width: 95%;height: 30px;display: flex;border-radius: 2px;align-content: center;justify-content: center;background: #13ce66">
                <el-icon size="16" style="height: 100%;width: 100%">
                  <Sort/>
                </el-icon>
              </div>
              <div class="" :style="KeySizeStyle(0,0)">
                <div style=" width:100%;height:100%;display: flex;flex-direction: column">
                  <el-button style="flex-grow: 1;position: relative"
                             :style=" item.detail_show_type==ShowType.Color?{background:item.detail_color}:{}"
                             :disabled="item.run_state==RunState.Running"
                             @click="ExecDetail(item)">
                    <div class="detail-item-control-icon" :style="GetControlTypeStyle(item.control_type)">
                      <Icon.KeyboardOne size="18" v-show="item.control_type==ControlType.Normal"/>
                      <Icon.RobotOne size="18" v-show="item.control_type==ControlType.Script"/>
                      <Icon.FolderOpen size="18" v-show="item.control_type==ControlType.Explorer"/>
                      <Icon.BrowserChrome size="18" v-show="item.control_type==ControlType.Website"/>
                      <Icon.Application size="18" v-show="item.control_type==ControlType.RunExe"/>
                      <Icon.Terminal size="18" v-show="item.control_type==ControlType.RunCmd"/>
                    </div>
                    <el-image
                        v-if="item.detail_show_type==ShowType.Pic"
                        class="detail-item-control-back-pic"
                        :src="GetIconSrc(item.detail_pic,nowTs)"
                        fit="fill" :style="KeySizeStyle(-4,-4)"
                    ></el-image>

                    <el-text line-clamp="2" v-show="item.detail_show_name"
                             style="text-align: center;white-space: normal;font-weight: bold;z-index: 500"
                             :style="{width:(GetKeyWidth()-10)+'px'}">{{
                        item.detail_name
                      }}
                    </el-text>
                  </el-button>

                  <div v-if="item.run_state==RunState.Running && item.control_type== ControlType.Script"
                       style="height: 30%">
                    <el-button style="width: 100%;height: 100%" :style="runStateLight(item)"
                               @click="ClickStopDetailBtn(item)">
                      <el-icon>
                        <SwitchButton/>
                      </el-icon>
                    </el-button>
                  </div>
                </div>
              </div>
              <div v-show="isEditMod" class="detail-edit-btn-001"
                   style="background: #e3dddd;border: #b4b4b4 solid 1px;" :style="{width:GetKeyWidth()-4+'px'}">
                <!--                编辑按钮-->
                <el-button size="small" style="width: 50%;background: #dea942"
                           @click="ClickEditKeyBtn(item)">
                  <el-icon>
                    <Edit/>
                  </el-icon>
                </el-button>
                <!--                删除按钮-->
                <el-button size="small" style="width: 50%;margin-left: 0px;background: #d75555"
                           @click="delDialogVisible=true;delId=item.detail_id">
                  <el-icon>
                    <Delete/>
                  </el-icon>
                </el-button>
              </div>

            </div>
          </VueDraggable>
        </div>

      </div>
    </main>
  </div>


  <!-- 删除 弹框 -->
  <el-dialog
      v-model="delDialogVisible"
      align-center
      style="width: 80%;max-width: 400px"
  >
    <el-text>确认删除？</el-text>
    <template #footer>
      <div class="works-dialog-footer">
        <el-button @click="delDialogVisible = false">取消</el-button>
        <el-button type="primary"
                   @click="ClickDeleteBtn()">
          确认
        </el-button>
      </div>
    </template>
  </el-dialog>

  <!-- 修改大小 弹框 -->
  <el-dialog
      v-model="isEditSize"
      align-center
      width="auto">
    <div style="width: 100%;display: flex;align-items: center;margin-bottom: 5px">
      <el-text style="width: 100%;text-align: center">修改Key大小</el-text>
    </div>

    <el-form>
      <el-form-item label="同步修改">
        <el-switch v-model="sizeEdit.sync_edit"/>
      </el-form-item>
      <el-form-item label="步进">
        <el-button @click="sizeEdit.add_step=1" :style="sizeEdit.add_step==1?'background: #79bbff':''">1</el-button>
        <el-button @click="sizeEdit.add_step=5" :style="sizeEdit.add_step==5?'background: #79bbff':''">5</el-button>
        <el-button @click="sizeEdit.add_step=10" :style="sizeEdit.add_step==10?'background: #79bbff':''">10</el-button>
      </el-form-item>
      <el-form-item label="宽度">
        <el-input-number :min="1" :step="sizeEdit.add_step" v-model="sizeEdit.key_width"
                         @change="()=>{if(sizeEdit.sync_edit)sizeEdit.key_height=sizeEdit.key_width}"
        ></el-input-number>
      </el-form-item>
      <el-form-item label="高度">
        <el-input-number :min="1" :step="sizeEdit.add_step" v-model="sizeEdit.key_height"
                         @change="()=>{if(sizeEdit.sync_edit)sizeEdit.key_width=sizeEdit.key_height}"
        ></el-input-number>
      </el-form-item>

      <el-text>鼠标偏移</el-text>
      <el-form-item label="X">
        <el-input-number :min="0" v-model="sizeEdit.offset_x"
        ></el-input-number>
      </el-form-item>
      <el-form-item label="Y">
        <el-input-number :min="0" v-model="sizeEdit.offset_y"
        ></el-input-number>
      </el-form-item>
    </el-form>
    <template #footer>
      <div class="works-dialog-footer">
        <el-button @click="isEditSize = false">取消</el-button>
        <el-button style="background: #79bbff"
                   @click="ClickUpdateSizeSubBtn(false)">
          应用
        </el-button>
        <el-button type="primary"
                   @click="ClickUpdateSizeSubBtn(true)">
          确认
        </el-button>
      </div>
    </template>
  </el-dialog>

  <!-- 新增 选择类型弹框 -->
  <el-dialog
      v-model="selectDialogVisible"
      align-center
      width="auto">
    <div style="width: 100%;display: flex;align-items: center;margin-bottom: 5px">
      <el-text style="width: 100%;text-align: center">类型选择</el-text>
    </div>
    <div>
      <el-button @click="addControlType=ControlType.Normal"
                 :style="addControlType==ControlType.Normal? {background:'aquamarine'}:{}">组合键
      </el-button>
      <el-button @click="addControlType=ControlType.Script"
                 :style="addControlType==ControlType.Script? {background:'aquamarine'}:{}">脚本
      </el-button>
      <el-button @click="addControlType=ControlType.Explorer"
                 :style="addControlType==ControlType.Explorer? {background:'aquamarine'}:{}">文件夹
      </el-button>
      <el-button @click="addControlType=ControlType.Website"
                 :style="addControlType==ControlType.Website? {background:'aquamarine'}:{}">浏览器
      </el-button>
    </div>
    <div style="margin-top: 10px">
      <el-button @click="addControlType=ControlType.RunExe"
                 :style="addControlType==ControlType.RunExe? {background:'aquamarine'}:{}">程序
      </el-button>
      <el-button @click="addControlType=ControlType.RunCmd"
                 :style="addControlType==ControlType.RunCmd? {background:'aquamarine'}:{}">命令
      </el-button>

    </div>
    <template #footer>
      <div class="works-dialog-footer">
        <el-button @click="selectDialogVisible = false">取消</el-button>
        <el-button type="primary"
                   @click="ClickAddKeyBtn">
          确认
        </el-button>
      </div>
    </template>
  </el-dialog>

  <!-- 新增/编辑 弹框 -->
  <el-dialog align-center
             v-model="editDialogVisible"
             :fullscreen="true"
             width="90%">
    <div
        style="width: 100%;display: flex;flex-direction: column;align-items: center;justify-content: center;margin-top: 30px;">
      <el-form label-position="right" label-width="auto" style="width: 100%;">
        <el-form-item label="名称">
          <el-input v-model="editItem.detail_name"></el-input>
        </el-form-item>
        <el-form-item label="是否显示名称">
          <el-switch
              v-model="editItem.detail_show_name"
              style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
          />
        </el-form-item>
        <el-form-item label="绑定快捷键">
          <el-button @click="bindNormalKey = editItem.combination_key;bindNormalKeyVisible=true">
            <div v-for="kkeeyy in editItem.combination_key" class="detail-list-item-bind-view">
              <el-text>
                {{ kkeeyy.key }}
              </el-text>
            </div>
          </el-button>


          <el-button @click="editItem.combination_key=[]" style="background: #c45656">X</el-button>
        </el-form-item>
        <el-form-item label="显示样式" style="display: flex;align-items: center;">
          <el-radio-group v-model="editItem.detail_show_type" class="ml-4">
            <el-radio :value="1" size="large">背景色</el-radio>
            <el-radio :value="2" size="large">图片</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="背景色" v-if="editItem.detail_show_type==ShowType.Color">
          <div style="display: flex;flex-direction: row">
            <el-color-picker @active-change="event => {editItem.detail_color = event}"
                             popper-class="hex"
                             v-model="editItem.detail_color" show-alpha :predefine="predefineColors"/>

          </div>
        </el-form-item>
        <div style="display: flex;flex-grow: 1;margin-left: 5px;flex-wrap: wrap"
             v-if="editItem.detail_show_type==ShowType.Color">
          <el-button style="margin: 1px" :style="{background:item}" size="small"
                     v-for="item in predefineColorListAuto" @click="editItem.detail_color=item"></el-button>
        </div>
        <el-form-item label="图片" v-if="editItem.detail_show_type==ShowType.Pic">
          <div style="border: #72767b solid 1px;width: 40px;height: 40px;display: flex;justify-content: center"
               @click="selectIconVisible=true">
            <el-image v-if="editItem.detail_pic!=''" style="width: 40px;height: 40px;"
                      :src="GetIconSrc(editItem.detail_pic,nowTs)"
                      fit="fill"></el-image>
          </div>

        </el-form-item>
        <el-form-item label="文件夹路径" v-if="editItem.control_type==ControlType.Explorer">
          <el-input type="textarea" rows="4" v-model="editItem.path"></el-input>
        </el-form-item>
        <el-form-item label="网址链接" v-if="editItem.control_type==ControlType.Website">
          <el-input type="textarea" rows="4" v-model="editItem.path"></el-input>
        </el-form-item>
        <el-form-item label="程序运行路径" v-if="editItem.control_type==ControlType.RunExe">
          <el-input type="textarea" rows="4" v-model="editItem.path"></el-input>
        </el-form-item>
        <el-form-item label="命令" v-if="editItem.control_type==ControlType.RunCmd">
          <el-input type="textarea" rows="4" v-model="editItem.path"></el-input>
        </el-form-item>
      </el-form>

    </div>
    <DetailEdit v-if="editItem.control_type==ControlType.Normal"
                :set-list="transferToKeyList(editItem.detail_key)"
                :update-list="list =>{normalKeyUpdate(list)}"></DetailEdit>


    <DetailScript v-if="editItem.control_type==ControlType.Script" :set-list="editItem.detail_key"
                  :update-list="list =>{scriptKeyUpdate(list)}"></DetailScript>


    <template #footer>
      <div class=" works-dialog-footer">
        <el-button @click="editDialogVisible = false">取消</el-button>
        <el-button type="primary"
                   @click="ClickEditSubBtn()">
          确认
        </el-button>
      </div>
    </template>
  </el-dialog>

  <!-- 图片 选择弹框 -->
  <el-dialog
      v-model="selectIconVisible"
      align-center
      :fullscreen="true">
    <div style="width: 100%;display: flex;align-items: center;margin-bottom: 5px">
      <el-text style="width: 100%;text-align: center">图片选择</el-text>
    </div>
    <DetailIcon :pic="selectIcon" :OnPicUpdate="pic => {selectIcon=pic}"></DetailIcon>
    <template #footer>
      <div class="works-dialog-footer">
        <el-button @click="selectIconVisible = false">取消</el-button>
        <el-button type="primary"
                   @click="ClickSubPicSelectBtn">
          确认
        </el-button>
      </div>
    </template>
  </el-dialog>

  <!-- 绑定快捷键 -->
  <el-dialog
      v-model="bindNormalKeyVisible"
      align-center
      width="auto">
    <div>
      <DetailEdit :set-list="bindNormalKey" :onlyOne="false"
                  :update-list="list =>{bindNormalKey = list;console.log(bindNormalKey)}"></DetailEdit>
    </div>

    <template #footer>
      <div class="works-dialog-footer">
        <el-button @click="bindNormalKeyVisible = false">取消</el-button>
        <el-button type="primary"
                   @click="confirmBindShortcutKey()">
          确认
        </el-button>
      </div>
    </template>
  </el-dialog>

  <!-- 启动快捷键弹框 -->
  <el-dialog
      v-model="setBindNormalKeyVisibleVisible"
      align-center
      width="auto">
    <div>
      <el-text v-show="setBindNormalKeyVisibleState==1">是否绑定当前控制快捷键？</el-text>
      <el-text v-show="setBindNormalKeyVisibleState==2">是否解除绑定控制快捷键？</el-text>
    </div>

    <template #footer>
      <div class="works-dialog-footer">
        <el-button @click="setBindNormalKeyVisibleVisible = false;setBindNormalKeyVisibleState=0">取消</el-button>
        <el-button type="primary"
                   @click="confirmSetBindNormalKey()">
          确认
        </el-button>
      </div>
    </template>
  </el-dialog>

</template>

<script lang="ts" setup>
import {
  Back,
  Check, ChromeFilled,
  CirclePlus, Connection,
  Delete,
  Edit, FolderOpened,
  FullScreen, Key,
  Refresh, Sort,
  SwitchButton
} from "@element-plus/icons-vue";
import * as Icon from "@icon-park/vue-next"
import {GotoHome} from "@/components/router/define";
import {
  ControlClass, EditControlClass,
  GetControlClassInfo,
  NewControlClass, NewControlClassId, UpdateControlClassOrder,
} from "@/components/api/control.js";
import {VueDraggable} from "vue-draggable-plus";
import {
  GetControlDetailList,
  GetControlDetailListResp,
  NewControlDetail,
  NewControlDetailId,
  NewControlDetailList,
  NewGetControlDetailListReq,
  NewUpdateControlDetailOrderReq,
  UpdateControlDetailOrder,
  ControlDetail,
  CopyControlDetail,
  EditControlDetail,
  NewEditControlDetailReq,
  NewControlDetailKey,
  DeleteControlDetail,
  NewDeleteControlDetailReq,
  ExecControlDetail,
  NewExecControlDetailReq,
  ControlType,
  ControlDetailKey,
  ControlKeyList,
  NewControlKeyListNormal,
  RunState,
  ExecStopControlDetail,
  NewStopControlDetailReq,
  ShowType,
  KeyType,
  PressType,
  StopCombinationKey,
  ActivateControlClassCombinationKey, ActivateControlClassCombinationKeyReq, NewActivateControlClassCombinationKeyReq
} from "@/components/api/detail";
import DetailEdit from "@/components/detail/detailEdit.vue";
import {onMounted, ref} from "vue";
import {useRoute} from "vue-router";
import DetailScript from "@/components/detail/detailScript.vue";
import {onUnmounted, onUpdated} from "@vue/runtime-core";
import {ElColorPicker, ElMessage} from "element-plus";
import {GetSystemConfig, NewSystemSetting, SystemSetting} from "@/components/api/set";
import {MessageErr} from "@/components/mod/msg";
import {GetNewId} from "@/components/common/id";
import DetailIcon from "@/components/detail/detailIcon.vue";
import {GetIconSrc} from "@/components/api/sys";
import {MouseDefine, MouseScrollDirection} from "@/components/api/keyDefine";
import {runConfig} from "@/components/mod/config";

const detailListContainer = ref(null)


const KeySizeStyle = (hOff: number, wOff: number) => {
  return {
    height: (parseInt(controlInfo.value.key_height) + hOff) + 'px',
    width: (parseInt(controlInfo.value.key_width) + wOff) + 'px',
  }
}

const runStateLight = (item: ControlDetail) => {
  let sty = {
    background: "",

  }
  switch (item.run_state) {
    case RunState.Free:
      sty.background = '#82d900'
      break
    case RunState.Running:
      sty.background = '#ffda02'
      break
  }
  return sty
}

const sizeEdit = ref({sync_edit: true, add_step: 1, key_height: 0, key_width: 0, offset_x: 0, offset_y: 0})

const controlInfo = ref(NewControlClass())
const detailList = ref(NewControlDetailList())

const isEditMod = ref(false)

const isEditSize = ref(false)

let socket = ref(null)

const delDialogVisible = ref(false)
const delId = ref(0)


const nowSystemConfig = ref(NewSystemSetting())

const selectDialogVisible = ref(false)
const addControlType = ref(ControlType.Normal)
const editDialogVisible = ref(false)
const editItem = ref(NewControlDetail())

const isControlInfoGet = ref(false)

const bindNormalKeyVisible = ref(false)
const bindNormalKey = ref([])

const tk = ref(0)
const lastColor = ref('')

const selectIconVisible = ref(false)
const selectIcon = ref('')


const predefineColors = ref([])

const setBindNormalKeyVisibleVisible = ref(false)
const setBindNormalKeyVisibleState = ref(0)


const GetControlTypeStyle = (tp: number) => {
  let sty = {
    background: ""
  }
  switch (tp) {
    case ControlType.Normal:
      sty.background = "#79bbff"
      break
    case ControlType.Script:
      sty.background = "#95d475"
      break
    case ControlType.Explorer:
      sty.background = "#eebe77"
      break
    case ControlType.Website:
      sty.background = "#f89898"
      break
    case ControlType.RunExe:
      sty.background = "#2BDAB1"
      break
    case ControlType.RunCmd:
      sty.background = "#c8c9cc"
      break
  }
  return sty
}
const onColorPickerBlur = () => {

  lastColor.value = editItem.value.detail_color
  console.log(lastColor.value)
}
onMounted(() => {
  const route = useRoute();
  let query = route.query
  if (query == undefined || query.control_id == undefined || query.control_id == '') {
    MessageErr("数据错误")
    GotoHome()
    return
  }
  GetSystemConfig().then(
      res => {
        if (res.state == 0) {
          nowSystemConfig.value = res.data as SystemSetting
        }
      })
  GetControlInfo(parseInt(<string>query.control_id))
  connectWebSocket()
})
onUpdated(
    () => {
      // SetDetailListContainerSize(controlInfo.value.key_width)
    }
)
onUnmounted(
    () => {
      if (socket.value) {
        socket.value.close()
      }
      if (tk.value != 0) {
        clearInterval(tk.value)
        tk.value = 0
      }
    }
)

const ClickSubPicSelectBtn = () => {
  editItem.value.detail_show_type = ShowType.Pic
  editItem.value.detail_pic = selectIcon.value
  selectIconVisible.value = false
}

const connectWebSocket = () => {
  // 连接到 WebSocket 服务器
  let sc = new WebSocket('ws://' + runConfig.host + "/api/GetMsgWS");
  socket.value = sc

  // 监听消息
  socket.value.addEventListener('message', (event) => {
    // console.log(event.data)
    let msg = JSON.parse(event.data);
    if (msg.api == "scriptStateChange") {
      if (msg.Data.control_id != controlInfo.value.control_id) {
        return
      }
      detailList.value.forEach((item, index) => {
        if (item.detail_id == msg.Data.detail_id) {
          detailList.value[index].run_state = msg.Data.state
        }
      })

    }
  });

  // 监听连接关闭
  socket.value.addEventListener('close', () => {
    console.log('WebSocket connection closed');
  });

  // 监听错误
  socket.value.addEventListener('error', (error) => {
    console.error('WebSocket error:', error);
  });
}

const sendMessage = () => {
  if (socket.value && socket.value.readyState === WebSocket.OPEN) {
    socket.value.send(Date.now())
  }
}
const GetColorByInt = (num: number) => {
  let b = Math.floor(num % 256)
  num = num / 256
  let g = Math.floor(num % 256)
  num = num / 256
  let r = Math.floor(num % 256)
  return 'rgb(' + r + ',' + g + ',' + b + ')'
}

const GetColorByRGB = (r: number, g: number, b: number) => {

  return 'rgb(' + r + ',' + g + ',' + b + ')'
}
const GetAutoColor = (num: number) => {
  let color = []
  for (let i = 0; i < 10; i++) {
    var c = 256 / 10 * i
    color.push(GetColorByRGB(c, c, c))
  }

  const colorSpace = Math.floor(256 / (num / 6) + 1)
  let now = "r"
  let add = true
  let r = 0
  let g = 0
  let b = 0
  for (let i = 0; i < 6; i++) {
    for (let j = 0; j < num / 6; j++) {
      switch (now) {
        case 'r':
          // brg
          r = 255
          if (add) {
            g = j * colorSpace
            b = 0
          } else {
            g = 0
            b = 255 - j * colorSpace
          }
          break
        case 'g':
          // rgb
          g = 255
          if (add) {
            b = j * colorSpace
            r = 0
          } else {
            b = 0
            r = 255 - j * colorSpace
          }
          break
        case 'b':
          // gbr
          b = 255
          if (add) {
            r = j * colorSpace
            g = 0
          } else {
            r = 0
            g = 255 - j * colorSpace
          }
          break
      }
      color.push(GetColorByRGB(r, g, b))
    }
    add = !add
    if (!add) {
      switch (now) {
        case 'r':
          now = 'g'
          break
        case 'g':
          now = 'b'
          break
        case 'b':
          now = 'r'
          break
      }
    }
  }
  return color
}

const nowTs = GetNewId()

let predefineColorListAuto = GetAutoColor(80)


const GetKeyWidth = (): number => {
  return parseInt(controlInfo.value.key_width)
}
const ClickEditSizeBtn = () => {
  sizeEdit.value.key_height = parseInt(controlInfo.value.key_height)
  sizeEdit.value.key_width = parseInt(controlInfo.value.key_width)
  sizeEdit.value.offset_x = parseInt(controlInfo.value.mouse_off_set_x)
  sizeEdit.value.offset_y = parseInt(controlInfo.value.mouse_off_set_y)
  isEditSize.value = true
}
const audio = new Audio('/touch_001.mp3?ts=' + nowTs); // 替换成你音频文件的路径

const ExecDetail = (item: ControlDetail) => {
  if (isEditMod.value) {
    return;
  }
  if (item.control_type != ControlType.Normal && item.run_state == RunState.Running) {
    return
  }
  if (nowSystemConfig.value.sound_open) {
    audio.currentTime = 0
    audio.play();
  }
  if (item.control_type == ControlType.Script) {
    item.run_state = RunState.Running
  } else {
    item.run_state = RunState.Free
  }
  let req = NewExecControlDetailReq()
  req.detail_id = item.detail_id
  req.control_id = controlInfo.value.control_id
  ExecControlDetail(req).then(resp => {
    if (resp.state != 0) {
      MessageErr(resp.msg)
    }
  })
}
const ClickStopDetailBtn = (item: ControlDetail) => {
  let req = NewStopControlDetailReq()
  req.detail_id = item.detail_id
  req.control_id = controlInfo.value.control_id
  ExecStopControlDetail(req).then(resp => {
    if (resp.state != 0) {
      MessageErr(resp.msg)
    }
  })
}
const transferToKeyList = (list: ControlKeyList[]): ControlDetailKey[] => {
  let keyList: ControlDetailKey[] = []
  list.forEach(
      (value, index) => {
        let key = NewControlDetailKey()
        key.id = value.id
        key.key = value.key
        keyList.push(key)
      }
  )
  return keyList
}

const scriptKeyUpdate = (list: ControlKeyList[]) => {
  editItem.value.detail_key = list
}

const normalKeyUpdate = (list: ControlDetailKey[]) => {
  let newList: ControlKeyList[] = []
  list.forEach(
      (value, index) => {
        let key = NewControlKeyListNormal()
        key.key = value.key
        key.id = value.id
        newList.push(key)
      }
  )
  editItem.value.detail_key = newList
}

const ClickUpdateSizeSubBtn = (close: boolean) => {
  controlInfo.value.key_width = sizeEdit.value.key_width.toString()
  controlInfo.value.key_height = sizeEdit.value.key_height.toString()
  controlInfo.value.mouse_off_set_x = sizeEdit.value.offset_x
  controlInfo.value.mouse_off_set_y = sizeEdit.value.offset_y
  EditControlClass(controlInfo.value).then(
      res => {
        if (res.state !== 0) {
          MessageErr(res.msg)
        } else {
          GetControlInfo(controlInfo.value.control_id)
        }
        if (close) {
          isEditSize.value = false
        }
      },
  )
}

const GetDetailList = () => {
  let req = NewGetControlDetailListReq()
  req.control_id = controlInfo.value.control_id
  GetControlDetailList(req).then(resp => {
    if (resp.state != 0) {
      GotoHome()
    } else {
      let data = resp.data as GetControlDetailListResp
      detailList.value = data.detail
    }
  })
}
const ClickDeleteBtn = () => {
  let req = NewDeleteControlDetailReq()
  req.control_id = controlInfo.value.control_id
  req.detail_id = delId.value
  DeleteControlDetail(req).then(
      res => {
        if (res.state != 0) {
        } else {
          GetDetailList()
        }
        delDialogVisible.value = false
      }
  )
}
const GetControlInfo = (id: number) => {
  isControlInfoGet.value = false
  let req = NewControlClassId()
  req.control_id = id
  GetControlClassInfo(req).then(resp => {
        if (resp.state != 0) {
          GotoHome()
        } else {
          isControlInfoGet.value = true
          controlInfo.value = resp.data as ControlClass
          // SetDetailListContainerSize(controlInfo.value.key_width)
          GetDetailList()
          if (tk.value == 0) {
            tk.value = setInterval(() => {
              sendMessage()
            }, 1000)
          }

        }
      }
  )
}

const OnDetailListOrderChange = () => {
  let req = NewUpdateControlDetailOrderReq()
  req.control_id = controlInfo.value.control_id
  detailList.value.forEach((value, index) => {
    let order = NewControlDetailId()
    order.detail_id = value.detail_id
    req.detail.push(order)
  })
  UpdateControlDetailOrder(req).then(resp => {
    if (resp.state != 0) {
      MessageErr(resp.msg)
    }
  })
}

const confirmSetBindNormalKey = () => {

  if (setBindNormalKeyVisibleState.value == 1) {
    let req = NewActivateControlClassCombinationKeyReq()
    req.control_id = controlInfo.value.control_id

    ActivateControlClassCombinationKey(req).then(resp => {
      if (resp.state != 0) {
        MessageErr(resp.msg)
      }
    })
  } else if (setBindNormalKeyVisibleState.value == 2) {
    StopCombinationKey().then(resp => {
      if (resp.state != 0) {
        MessageErr(resp.msg)
      }
    })
  }

  setBindNormalKeyVisibleVisible.value = false
  setBindNormalKeyVisibleState.value = 0

}

const ClickAddKeyBtn = () => {
  selectDialogVisible.value = false
  editItem.value = NewControlDetail()
  editItem.value.detail_show_name = true
  editItem.value.detail_show_type = ShowType.Color
  editItem.value.control_type = addControlType.value

  switch (editItem.value.control_type) {
    case ControlType.Script:
      editItem.value.mouse_back_to_origin = true
  }
  editDialogVisible.value = true
}
const ClickEditKeyBtn = (item: ControlDetail) => {
  editItem.value = CopyControlDetail(item)
  editDialogVisible.value = true
}

const ClickEditSubBtn = () => {
  let req = NewEditControlDetailReq()
  req.control_id = controlInfo.value.control_id
  req.detail = editItem.value

  EditControlDetail(req).then(resp => {
    if (resp.state != 0) {
      MessageErr(resp.msg)
    } else {
      GetDetailList()
      editDialogVisible.value = false
    }
  })
}

const confirmBindShortcutKey = () => {
  bindNormalKeyVisible.value = false
  editItem.value.combination_key = bindNormalKey.value

}

</script>
