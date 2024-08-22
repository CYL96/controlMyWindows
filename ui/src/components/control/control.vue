<template>
  <div id="control_div" style="width: 100%;height: 100%;display: flex;flex-direction: column;">
    <header id="control_header_div" style="padding: 1% 0 1% 0;height: 5%;min-height: 30px;width: 100%">
      <div style="display: flex;align-items: center;justify-content: center;height: 100%">
        <el-button @click="GetControlClassListFromServer" style="width: 30%;background: #d1edc4">
          <el-icon>
            <Refresh/>
          </el-icon>
        </el-button>
        <el-button @click="ClickAddClassBtn" style="width: 30%;background: #79bbff">
          <el-icon>
            <CirclePlus/>
          </el-icon>
        </el-button>
        <el-button @click="ClickEditSystemBtn" style="width: 10%;background: #c8c9cc">
          <Icon.SettingConfig/>
        </el-button>

        <el-button @click="exitDialogVisible=true" style="width: 10%;background: #e72222">
          <el-icon>
            <Icon.Power/>
          </el-icon>
        </el-button>


      </div>
    </header>
    <main id="control_main_div" style="display: flex;height: 95%;width: 100%">
      <!-- 左边的列表 -->
      <div style="display: flex;flex: 1;align-items: flex-start;justify-content: center;padding:0 0 20px 0; overflow-y: scroll;">
          <VueDraggable
              handle=".copilot-sort-icon"
              v-model="list"
              :animation="150"
              @update="OnClassListOrderChange()"
              id="LabelList"
              class="control-for-list-div"
          >
            <div :id="item.control_id" v-for="item in list" :key="item.control_id"
                 class="control-for-list">
              <div  class="control-item control-cm-border">
                <div class="copilot-sort-icon">
                  <el-icon>
                    <Sort/>
                  </el-icon>
                </div>

                <div @click="ClickControlDetail(item)"
                     style="display: flex;flex-grow: 1;margin-left: 1%;text-align: center;justify-content: center;height: 100%">
                  <el-text style="width: 100%;">{{ item.control_name }}</el-text>
                </div>
                <div style="display: flex;flex-direction: column; justify-content: space-around;align-items: flex-end;padding-right: 1%;height: 100%">
                  <!--              排序按钮-->
                  <!--                编辑按钮-->
                  <el-button class="control-btn-001" style="background: #dea942" @click="ClickEditClassBtn(item)">
                    <el-icon>
                      <Edit/>
                    </el-icon>
                  </el-button>
                  <!--                删除按钮-->
                  <el-button class="control-btn-001" style="background: #d75555"
                             @click="delDialogVisible=true;delId=item.control_id">
                    <el-icon>
                      <Delete/>
                    </el-icon>
                  </el-button>
                </div>

              </div>

            </div>
          </VueDraggable>
      </div>
    </main>
  </div>

  <!-- 新增copilot 弹框 -->
  <el-dialog align-center
             v-model="editDialogVisible"

             width="80%">
    <el-form label-width="auto">
      <el-form-item label="名称">
        <el-input v-model="editItem.control_name" style="width:90%"></el-input>
      </el-form-item>
    </el-form>

    <template #footer>
      <div class="works-dialog-footer">
        <el-button @click="editDialogVisible = false">取消</el-button>
        <el-button type="primary"
                   @click="ClickEditClassSubBtn()">
          确认
        </el-button>
      </div>
    </template>
  </el-dialog>

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

  <!-- 删除 弹框 -->
  <el-dialog
      v-model="exitDialogVisible"
      align-center
      width="40%">
    <el-text>退出程序？</el-text>
    <template #footer>
      <div class="works-dialog-footer">
        <el-button @click="exitDialogVisible = false">取消</el-button>
        <el-button type="primary"
                   @click="ClickExitSubBtn()">
          确认
        </el-button>
      </div>
    </template>
  </el-dialog>

  <!-- 系统设置 弹框 -->
  <el-dialog align-center
             v-model="editSystemDialogVisible"
             width="80%">
    <el-form label-width="auto">
      <el-form-item label="端口">
        <el-input-number v-model="editSystemItem.run_port"></el-input-number>
      </el-form-item>
      <el-form-item label="点击音效">
        <el-switch
            v-model="editSystemItem.sound_open"
            style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
        />
      </el-form-item>
      <el-form-item label="屏幕是否设置缩放">
        <el-switch
            v-model="editSystemItem.is_scale"
            style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949"
        />
      </el-form-item>
    </el-form>

    <template #footer>
      <div class="works-dialog-footer">
        <el-button @click="editSystemDialogVisible = false">取消</el-button>
        <el-button type="primary"
                   @click="ClickEditSystemSubBtn">
          确认
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import "./control.css"
import * as Icon from "@icon-park/vue-next"
import {defineComponent, onMounted, ref} from "vue";
import {
  ControlClass, ControlClassId, DeleteControlClass, EditControlClass,
  GetControlClassList,
  GetControlClassListResp, NewControlClass, NewControlClassId,
  NewControlClassList, NewGetControlClassListReq, NewUpdateControlClassOrderReq, UpdateControlClassOrder,
} from "@/components/api/control.js";
import {VueDraggable} from "vue-draggable-plus";
import {GotoDetail, NewDetailQuery} from "@/components/router/define";
import {ApiExitControl} from "@/components/api/sys";
import {GetSystemConfig, NewSystemSetting, SetSystemConfig, SystemSetting} from "@/components/api/set";
import {CirclePlus, Delete, Edit, Refresh, Sort} from "@element-plus/icons-vue";
import {onUpdated} from "@vue/runtime-core";
import {MessageErr} from "@/components/mod/msg";

const list = ref(NewControlClassList())

const delDialogVisible = ref(false)
const delId = ref(0)

const control_div = ref()

const control_header_div = ref()

const editDialogVisible = ref(false)
const editItem = ref(NewControlClass())

const exitDialogVisible = ref(false)

const editSystemDialogVisible = ref(false)
const editSystemItem = ref(NewSystemSetting())

onMounted(() => {
  GetControlClassListFromServer()
  // SetMainHeight()
})

const ClickControlDetail = (item: ControlClass) => {
  let query = NewDetailQuery()
  query.control_id = String(item.control_id)
  GotoDetail(query)
}
const GetControlClassListFromServer = () => {
  GetControlClassList(NewGetControlClassListReq()).then(
      res => {
        if (res.state !== 0) {
        } else {
          let data = res.data as GetControlClassListResp
          list.value = data.list
        }
      }
  )
}

const SetMainHeight = () => {
  var ctl = document.getElementById("control_div");
  var head = document.getElementById("control_header_div");
  var main = document.getElementById("control_main_div");
  if (main){
    // main.style.height='800px'
    main.style.height=ctl.offsetHeight - head.offsetHeight+"px"
  }
  MessageErr("ctl:"+ctl.offsetHeight)
  MessageErr("head:"+head.offsetHeight)
  MessageErr("main:"+main.offsetHeight)
  MessageErr("main:"+window.innerHeight+","+window.innerWidth)
}

const ClickEditSystemBtn = () => {
  editSystemDialogVisible.value = true
  editSystemItem.value = NewSystemSetting()
  GetSystemConfig().then(
      res => {
        if (res.state !== 0) {
        } else {
          editSystemItem.value = res.data as SystemSetting
        }
      }
  )
}
const ClickEditSystemSubBtn = () => {
  SetSystemConfig(editSystemItem.value).then(
      res => {
        if (res.state !== 0) {
        }
        editSystemDialogVisible.value = false
      }
  )
}

// 点击退出确认按钮
const ClickExitSubBtn = () => {
  ApiExitControl()
  exitDialogVisible.value = false
}
const ClickAddClassBtn = () => {
  editDialogVisible.value = true
  editItem.value = NewControlClass()
}
const ClickEditClassBtn = (item: ControlClass) => {
  editDialogVisible.value = true
  editItem.value = NewControlClass()
  editItem.value.control_id = item.control_id
  editItem.value.control_name = item.control_name
}
const ClickEditClassSubBtn = () => {
  EditControlClass(editItem.value).then(
      res => {
        if (res.state !== 0) {
        } else {
          GetControlClassListFromServer()
        }
        editDialogVisible.value = false
      },)
}
const ClickDeleteBtn = () => {
  DeleteControlClass(delId.value).then(
      res => {
        if (res.state !== 0) {
        } else {
          GetControlClassListFromServer()
        }
        delDialogVisible.value = false
      },
  )
}
const OnClassListOrderChange = () => {
  let para = NewUpdateControlClassOrderReq()
  list.value.forEach((item, index) => {
    let order = NewControlClassId()
    order.control_id = item.control_id

    para.order_list.push(order)
  })

  UpdateControlClassOrder(para).then(
      res => {
        if (res.state !== 0) {
        }
      }
  )
}


</script>