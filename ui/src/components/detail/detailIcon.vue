<template>
  <div style="display: flex;flex-direction: column">
    <div ref="detailIconHeader" class="detail-icon-header">
      <div>
        <el-button @click="GetIconList">
          <el-icon><Refresh /></el-icon>
        </el-button>
      </div>
      <div v-show="NowSelect.icon_path != ''"
           style="padding: 2px;background:white;border: #72767b solid 1px; display: flex;align-items: center;justify-content: center">
        <el-image style="width: 40px;height: 40px;" :src="GetIconSrc(NowSelect.icon_path,Ts)" fit="fill"></el-image>
        <el-button style="height: 40px;width: 20px;background: #E30F0F" size="small"
                   @click="ClickSelectPic(NewIconT())">
          <el-icon>
            <Delete/>
          </el-icon>
        </el-button>
      </div>
    </div>
    <div style=" display: flex;justify-content: center;flex-grow: 1
        ">
      <div style="display: flex;flex-direction: row;flex-wrap: wrap;" :style="GetIconListStyle()">
        <el-scrollbar view-class="detail-icon-body">
          <div v-for="item in IconList" class="detail-icon-body" @click="ClickSelectPic(item)"
               :style="SelectIconStyle(item.icon_path)">
            <el-image :src="GetIconSrc(item.icon_path,Ts)" fit="fill"
                      style="width: 80px;height: 80px;"></el-image>
          </div>
        </el-scrollbar>
      </div>
    </div>

  </div>

</template>
<script setup lang="ts">
import {onMounted, ref, watch} from "vue";
import {GetNowIconList, IconResp, IconT, NewIconT, NewIconT2, NewIconTList} from "@/components/api/set.js";
import {runConfig} from "@/components/mod/config";
import {GetNewId} from "@/components/common/id";
import {ControlDetailKey} from "@/components/api/detail";
import {GetIconSrc} from "@/components/api/sys";
import {Bicycle, Delete, Refresh} from "@element-plus/icons-vue";

interface Prop {
  pic: string
  OnPicUpdate: (pic: string) => void
}

const props = withDefaults(defineProps<Prop>(), {
  pic: '',
  OnPicUpdate: (pic: string) => {
  }
})

const IconList = ref(NewIconTList())
const NowSelect = ref(NewIconT2(props.pic))
const detailIconHeader = ref()
const Ts = GetNewId()


const ClickSelectPic = (item: IconT) => {
  NowSelect.value.icon_path = item.icon_path
  NowSelect.value.icon_name = item.icon_name
  props.OnPicUpdate(NowSelect.value.icon_path)
}

const SelectIconStyle = (path: string) => {
  let sty = {
    border: '#605f5f solid 2px',
    background: '',
  }
  if (path == NowSelect.value.icon_path) {
    sty.border = '2px solid red'
    sty.background = '#409EFF'
  }
  return sty
}

const GetIconListStyle = () => {
  let sty = {
    height: '80%'
  }
  // if (detailIconHeader.value != null) {
  //   sty.height = (window.innerHeight - detailIconHeader.value.clientHeight) + 'px'
  // }
  return sty
}

onMounted(() => {
  GetIconList()
})
const GetIconList = () => {
  GetNowIconList().then(
      res => {
        if (res.state == 0) {
          let data = res.data as IconResp
          IconList.value = data.List
        }
      }
  )
}
</script>