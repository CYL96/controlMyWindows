<template>
  <header style="height: 10%">
    <div style="display: flex;align-items: center;justify-content: center">
      <el-button @click="GotoHome" style="width: 10vw">
        <el-icon>
          <Back/>
        </el-icon>
      </el-button>


      <el-button @click="GetDetailList" style="width: 30vw">
        <el-icon>
          <Refresh/>
        </el-icon>
      </el-button>

      <div style="width: 45vw;margin-left: 12px;display: flex;align-items: center;justify-content: center">
        <el-button v-show="!isEditMod" @click="isEditMod= true" style="width: 100%">
          <el-icon>
            <Edit/>
          </el-icon>
        </el-button>

        <el-button v-show="isEditMod" @click="selectDialogVisible=true;"
                   style="width: 40%;background:deepskyblue">
          <el-icon>
            <CirclePlus/>
          </el-icon>
        </el-button>

        <el-button v-show="isEditMod" @click="isEditMod=false" style="width: 40%;background: greenyellow">
          <el-icon>
            <Check/>
          </el-icon>
        </el-button>

        <el-button v-show="isEditMod" @click="ClickEditSizeBtn" style="width: 10%;background: #ffce2f">
          <el-icon>
            <FullScreen/>
          </el-icon>
        </el-button>
      </div>


    </div>
  </header>
  <main style="margin-top: 1vh;display: flex;align-items: center;justify-content: center">
    <!-- 左边的列表 -->
    <div v-if="isControlInfoGet" ref="detailListContainer" style="width: auto">
      <el-scrollbar style="display: flex;" >
        <VueDraggable
            v-model="detailList"
            :animation="150"
            @update="OnDetailListOrderChange()"
            :disabled="!isEditMod"
            handle=".sort_detail_key_list"
            id="LabelList"
            class="detail-list-view"
        >
          <div :id="item.detail_id" v-for="item in detailList" :key="item.detail_id"
               class="detail-item ">
            <div v-if="isEditMod" class="sort_detail_key_list"
                style="width: 95%;height: 20px;display: flex;border-radius: 2px;align-content: center;justify-content: center;background: #13ce66">
                <el-icon size="16" style="height: 100%">
                  <Sort/>
                </el-icon>
            </div>
            <div class="" :style="KeySizeStyle()">
              <div style=" width:100%;height:100%;display: flex;flex-direction: column">
                <el-button style="flex-grow: 1;position: relative" :style="{background:item.detail_color}"
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
                  <el-text line-clamp="2" style="text-align: center;white-space: normal;font-weight: bold;"
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
      </el-scrollbar>
    </div>
  </main>


  <!-- 删除 弹框 -->
  <el-dialog
      v-model="delDialogVisible"
      align-center
      width="auto">
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
      <el-form-item label="宽度">
        <el-input-number v-model="sizeEdit.key_width"></el-input-number>
      </el-form-item>
      <el-form-item label="高度">
        <el-input-number v-model="sizeEdit.key_height"></el-input-number>
      </el-form-item>
    </el-form>
    <template #footer>
      <div class="works-dialog-footer">
        <el-button @click="isEditSize = false">取消</el-button>
        <el-button type="primary"
                   @click="ClickUpdateSizeSubBtn()">
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
             width="90vw">
    <div
        style="width: 100%;display: flex;flex-direction: column;align-items: center;justify-content: center;margin-top: 30px;">
      <el-form label-position="right" label-width="auto" >
        <el-form-item label="名称">
          <el-input v-model="editItem.detail_name"></el-input>
        </el-form-item>
        <el-form-item label="背景色">
          <div style="display: flex;flex-direction: row">
            <el-color-picker @active-change="event => {editItem.detail_color = event}"
                             popper-class="hex"
                             v-model="editItem.detail_color" show-alpha :predefine="predefineColors"/>
            <div style="display: flex;flex-grow: 1;margin-left: 5px;flex-wrap: wrap">
              <el-button style="margin-left: 1px" :style="{background:item}" size="small"
                  v-for="item in predefineColorListAuto" @click="editItem.detail_color=item"></el-button>

            </div>
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
  RunState, ExecStopControlDetail, NewStopControlDetailReq
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

const detailListContainer = ref(null)


const KeySizeStyle = () => {
  return {
    height: controlInfo.value.key_height + 'px',
    width: controlInfo.value.key_width + 'px',
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

const sizeEdit = ref({key_height: 0, key_width: 0})

const controlInfo = ref(NewControlClass())
const detailList = ref(NewControlDetailList())

const isEditMod = ref(false)

const isEditSize = ref(false)

const delDialogVisible = ref(false)
const delId = ref(0)


const nowSystemConfig = ref(NewSystemSetting())

const selectDialogVisible = ref(false)
const addControlType = ref(ControlType.Normal)
const editDialogVisible = ref(false)
const editItem = ref(NewControlDetail())

const isControlInfoGet = ref(false)

const tk = ref(0)
const lastColor = ref('')

const predefineColorList = [
  '#00c8c8',
  '#FFFF00',
  '#0000FF',
  '#008000',
  '#FFA500',
  '#4B0082',
  '#800080',
  '#F0F8FF',
  '#FAEBD7',
  '#00FFFF',
  '#7FFFD4',
  '#F0FFFF',
  '#F5F5DC',
  '#FFE4C4',
  '#000000',
  '#FFEBCD',
  '#8A2BE2',
  '#A52A2A',
  '#DEB887',
  '#5F9EA0',
  '#7FFF00',
  '#D2691E',
  '#FF7F50',
  '#6495ED',
  '#FFF8DC',
  '#DC143C',
  '#00FFFF',
  '#00008B',
  '#008B8B',
  '#B8860B',
  '#A9A9A9',
  '#006400',
  '#BDB76B',
  '#8B008B',
  '#556B2F',
  '#FF8C00',
  '#9932CC',
  '#8B0000',
  '#E9967A',
  '#8FBC8F',
  '#483D8B',
  '#2F4F4F',
  '#00CED1',
  '#9400D3',
  '#FF1493',
  '#00BFFF',
  '#696969',
  '#1E90FF',
  '#D19275',
  '#B22222',
  '#FFFAF0',
  '#228B22',
  '#FF00FF',
  '#DCDCDC',
  '#F8F8FF',
  '#FFD700',
  '#DAA520',
  '#808080',
  '#ADFF2F',
  '#F0FFF0',
  '#FF69B4',
  '#CD5C5C',
  '#FFFFF0',
  '#F0E68C',
  '#E6E6FA',
  '#FFF0F5',
  '#7CFC00',
  '#FFFACD',
  '#ADD8E6',
  '#F08080',
  '#E0FFFF',
  '#FAFAD2',
  '#D3D3D3',
  '#90EE90',
  '#FFB6C1',
  '#FFA07A',
  '#20B2AA',
  '#87CEFA',
  '#8470FF',
  '#778899',
  '#B0C4DE',
  '#FFFFE0',
  '#00FF00',
  '#32CD32',
  '#FAF0E6',
  '#FF00FF',
  '#800000',
  '#66CDAA',
  '#0000CD',
  '#BA55D3',
  '#9370D8',
  '#3CB371',
  '#7B68EE',
  '#00FA9A',
  '#48D1CC',
  '#C71585',
  '#191970',
  '#F5FFFA',
  '#FFE4E1',
  '#FFE4B5',
  '#FFDEAD',
  '#000080',
  '#FDF5E6',
  '#808000',
  '#6B8E23',
  '#FF4500',
  '#DA70D6',
  '#EEE8AA',
  '#98FB98',
  '#AFEEEE',
  '#D87093',
  '#FFEFD5',
  '#FFDAB9',
  '#CD853F',
  '#FFC0CB',
  '#DDA0DD',
  '#B0E0E6',
  '#BC8F8F',
  '#4169E1',
  '#8B4513',
  '#FA8072',
  '#F4A460',
  '#2E8B57',
  '#FFF5EE',
  '#A0522D',
  '#C0C0C0',
  '#87CEEB',
  '#6A5ACD',
  '#708090',
  '#FFFAFA',
  '#00FF7F',
  '#4682B4',
  '#D2B48C',
  '#008080',
  '#D8BFD8',
  '#FF6347',
  '#40E0D0',
  '#EE82EE',
  '#D02090',
  '#F5DEB3',
  '#FFFFFF',
  '#F5F5F5',
  '#9ACD32',
]

const predefineColors = ref([

])


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
})
onUpdated(
    () => {
      // SetDetailListContainerSize(controlInfo.value.key_width)
    }
)
onUnmounted(
    () => {
      if (tk.value != 0) {
        clearInterval(tk.value)
        tk.value = 0
      }
    }
)

const GetColorByInt= (num:number)  =>{
  let b = Math.floor(num%255)
  num = num/255
  let g = Math.floor(num%255)
  num = num/255
  let r = Math.floor(num%255)
  console.log(r,g,b)
  return 'rgb('+r+','+g+','+ b+')'
}
const GetAutoColor=(num:number)=>{
  let  color = []
  const totalColor = 255 * 255 * 255
  const colorSpace = totalColor/num
  for (let i = 0; i < totalColor; i=i+colorSpace) {
    color.push(GetColorByInt(i))
  }
return color
}


let predefineColorListAuto =  GetAutoColor(60)

// const SetDetailListContainerSize = (detailWidth: number) => {
//   if (detailListContainer.value == null || detailListContainer.value.style == null) {
//     return [[]]
//   }
//   let canUse = Math.floor(window.innerWidth - window.innerWidth * 0.01)
//   let mayCount = Math.floor(canUse / detailWidth)
//   let realCanUse = canUse - mayCount * 2
//   let realCount = Math.floor(realCanUse / detailWidth)
//   let realHeight = Math.floor(window.innerHeight*0.9)
//
//   let realWidth = realCount * (parseInt(String(detailWidth)) + 2)
//   detailListContainer.value.style.width = realWidth + 'px'
//   detailListContainer.value.style.height = realHeight + 'px'
// }

const GetKeyWidth = (): number => {
  return parseInt(controlInfo.value.key_width)
}
const ClickEditSizeBtn = () => {
  sizeEdit.value.key_height = parseInt(controlInfo.value.key_height)
  sizeEdit.value.key_width = parseInt(controlInfo.value.key_width)
  isEditSize.value = true
}
const audio = new Audio('/touch_001.mp3?ts='+GetNewId()); // 替换成你音频文件的路径

const ExecDetail = (item: ControlDetail) => {
  if ((item.control_type != ControlType.Normal) && (isEditMod.value || item.run_state == RunState.Running)) {
    return
  }
  if (nowSystemConfig.value.sound_open) {
    audio.currentTime = 0
    audio.play();
  }
  if (item.control_type != ControlType.Normal) {
    item.run_state = RunState.Running
  } else {
    item.run_state = RunState.Free
  }
  let req = NewExecControlDetailReq()
  req.detail_id = item.detail_id
  req.control_id = controlInfo.value.control_id
  ExecControlDetail(req).then(resp => {
    if (resp.state != 0) {
      item.run_state = RunState.Free
    }
  })
}
const ClickStopDetailBtn = (item: ControlDetail) => {
  let req = NewStopControlDetailReq()
  req.detail_id = item.detail_id
  req.control_id = controlInfo.value.control_id
  ExecStopControlDetail(req).then(resp => {
    if (resp.state != 0) {
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

const ClickUpdateSizeSubBtn = () => {
  controlInfo.value.key_width = sizeEdit.value.key_width.toString()
  controlInfo.value.key_height = sizeEdit.value.key_height.toString()
  EditControlClass(controlInfo.value).then(
      res => {
        if (res.state !== 0) {
        } else {
          GetControlInfo(controlInfo.value.control_id)
        }
        isEditSize.value = false
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
          if (tk.value == 0) {
            GetDetailList()
            tk.value = setInterval(() => {
              GetDetailList()
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
    }
  })
}


const ClickAddKeyBtn = () => {
  selectDialogVisible.value = false
  editItem.value = NewControlDetail()
  editItem.value.control_type = addControlType.value
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
    } else {
      GetDetailList()
    }
    editDialogVisible.value = false
  })
}

</script>
