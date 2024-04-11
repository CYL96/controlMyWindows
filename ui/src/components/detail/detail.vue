<template>
  <header>
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

        <el-button v-show="isEditMod" @click="selectDialogVisible=true;addControlType=ControlType.Script"
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
  <main style="margin-top: 1vh">
    <!-- 左边的列表 -->
    <div v-if="isControlInfoGet" style="display: flex;width: 98vw;align-items: center;justify-content: center">
      <el-scrollbar style="display: flex;">
        <VueDraggable
            v-model="detailList"
            :animation="150"
            @update="OnDetailListOrderChange()"
            :disabled="!isEditMod"
            id="LabelList"
            class="detail-list-view"
        >
          <div :id="item.detail_id" v-for="item in detailList" :key="item.detail_id"
               class="detail-item ">
            <div class="" :style="KeySizeStyle(item)">
              <div style=" width:100%;height:100%;display: flex;flex-direction: column">
                <el-button style="flex-grow: 1;" :style="{background:item.detail_color}"
                           @click="ExecDetail(item)">
                  <el-text line-clamp="2" style="text-align: center;white-space: normal;"
                           :style="{width:(GetKeyWidth()-10)+'px'}">({{ item.control_type }}){{
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
            <!--            <div style="display: flex;flex-direction: row;margin-bottom: 1px">-->
            <!--              <div v-for="keys in item.detail_key" style="margin-left: 1px">-->
            <!--                <div style="width: 24px;height: 24px;display: flex;align-items: center" class="detail-cm-border">-->
            <!--                  <el-tag style="font-size: small;width: 100%;text-align: center;" >{{ keys.key }}</el-tag>-->
            <!--                </div>-->
            <!--              </div>-->
            <!--            </div>-->
            <div v-show="isEditMod" class="detail-edit-btn-001"
                 style="background: #e3dddd;border: #b4b4b4 solid 1px;" :style="{width:GetKeyWidth()+'px'}">
              <!--                编辑按钮-->
              <el-button size="small" style="width: 40%"
                         @click="ClickEditKeyBtn(item)">
                <el-icon>
                  <Edit/>
                </el-icon>
              </el-button>
              <!--                删除按钮-->
              <el-button size="small" style="width: 40%;margin-left: 0px"
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

  <!-- 新增 弹框 -->
  <el-dialog align-center
             v-model="editDialogVisible"
             :fullscreen="true" style="min-height: 150vh"
             width="90vw">
    <div style="width: 100%;display: flex;flex-direction: column;align-items: center;justify-content: center;margin-top: 30px;">
      <el-form label-position="right" label-width="auto" style="width: 90%">
        <el-form-item label="名称">
          <el-input v-model="editItem.detail_name" ></el-input>
        </el-form-item>
        <el-form-item label="背景色">
          <el-color-picker @active-change="event => {editItem.detail_color = event}"
                           popper-class="hex detail-color-picker"
                           v-model="editItem.detail_color" show-alpha :predefine="predefineColors"/>
        </el-form-item>
        <el-form-item label="文件夹路径" v-if="editItem.control_type==ControlType.Explorer">
          <el-input type="textarea" rows="4" v-model="editItem.path" ></el-input>
        </el-form-item>
        <el-form-item label="网址链接" v-if="editItem.control_type==ControlType.Website">
          <el-input type="textarea" rows="4" v-model="editItem.path" ></el-input>
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


</template>

<script lang="ts" setup>
import {
  Back,
  Check,
  CirclePlus,
  Close,
  CloseBold,
  Delete,
  Edit,
  FullScreen, Orange,
  Refresh,
  Sort, SwitchButton
} from "@element-plus/icons-vue";
import {DetailQuery, GotoHome} from "@/components/router/define";
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
  GetKeyStr,
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
  NewControlKeyList,
  NewControlKeyListNormal,
  RunState, ExecStopControlDetail, NewStopControlDetailReq
} from "@/components/api/detail";
import Hahah from "@/components/detail/detailEdit.vue";
import DetailEdit from "@/components/detail/detailEdit.vue";
import {onMounted, ref} from "vue";
import router from "@/components/router/router";
import {routes} from "vue-router/vue-router-auto-routes";
import {useRoute} from "vue-router";
import DetailScript from "@/components/detail/detailScript.vue";
import {onUnmounted} from "@vue/runtime-core";
import {ElColorPicker} from "element-plus";

const KeySizeStyle = (item: ControlDetail) => {
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


const selectDialogVisible = ref(false)
const addControlType = ref(0)
const editDialogVisible = ref(false)
const editItem = ref(NewControlDetail())

const isControlInfoGet = ref(false)

const tk = ref(0)
const lastColor = ref('')

const predefineColors = ref([
  '#ff4500',
  '#ff8c00',
  '#ffd700',
  '#90ee90',
  '#00ced1',
  '#1e90ff',
  '#c71585',
  'rgba(255, 69, 0, 0.68)',
  'rgb(255, 120, 0)',
  'hsv(51, 100, 98)',
  'hsva(120, 40, 94, 0.5)',
  'hsl(181, 100%, 37%)',
  'hsla(209, 100%, 56%, 0.73)',
  '#c7158577',
])

const onColorPickerBlur = () => {

  lastColor.value = editItem.value.detail_color
  console.log(lastColor.value)
}
onMounted(() => {
  const route = useRoute();
  let query = route.query
  if (query == undefined || query.control_id == undefined || query.control_id == '') {
    alert("数据错误")
    GotoHome()
    return
  }
  GetControlInfo(parseInt(<string>query.control_id))
})
onUnmounted(
    () => {
      if (tk.value != 0) {
        clearInterval(tk.value)
        tk.value = 0
      }
    }
)

const GetKeyWidth = (): number => {
  return parseInt(controlInfo.value.key_width)
}
const ClickEditSizeBtn = () => {
  sizeEdit.value.key_height = parseInt(controlInfo.value.key_height)
  sizeEdit.value.key_width = parseInt(controlInfo.value.key_width)
  isEditSize.value = true
}
const ExecDetail = (item: ControlDetail) => {
  if (isEditMod.value) {
    return
  }
  let req = NewExecControlDetailReq()
  req.detail_id = item.detail_id
  req.control_id = controlInfo.value.control_id
  ExecControlDetail(req).then(resp => {
    if (resp.state != 0) {
      alert(resp.msg)
    }
  })
}
const ClickStopDetailBtn = (item: ControlDetail) => {
  let req = NewStopControlDetailReq()
  req.detail_id = item.detail_id
  req.control_id = controlInfo.value.control_id
  ExecStopControlDetail(req).then(resp => {
    if (resp.state != 0) {
      alert(resp.msg)
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
          alert(res.msg)
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
      alert(resp.msg)
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
          alert(res.msg)
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
          alert(resp.msg)
          GotoHome()
        } else {
          isControlInfoGet.value = true
          controlInfo.value = resp.data as ControlClass
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
      alert(resp.msg)
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
      alert(resp.msg)
    } else {
      GetDetailList()
    }
    editDialogVisible.value = false
  })
}

</script>
