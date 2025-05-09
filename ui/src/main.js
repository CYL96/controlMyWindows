import 'core-js/actual/array/flat-map'
import {createApp} from "vue";
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import '@icon-park/vue-next/styles/index.css';
import router from "@/components/router/router.ts";
import App from "./App.vue";
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
import {runConfig} from "@/components/mod/config.ts";
import {install} from '@icon-park/vue-next/es/all';
import { ElMessage } from 'element-plus'
const app = createApp(App)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

install(app, 'icon');

app.use(ElementPlus, {
    locale: zhCn,
})
app.use(router)

app.mount('#app')

export  let baseHost = "10.5.10.87:55001"
if (import.meta.env.PROD){
    runConfig.server =  window.location.origin
    runConfig.host =  window.location.host
}else {
    runConfig.server = "http://"+baseHost
    runConfig.host = baseHost
}