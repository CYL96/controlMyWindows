import {createRouter, createWebHashHistory, RouteRecordRaw} from 'vue-router'

import {Page_Detail, Page_Home} from "./define";
// @ts-ignore
import Control from "../control/control.vue";
import Detail from "../detail/detail.vue";

const routes:Readonly<RouteRecordRaw[]> = [
    {
        path: "/",
        redirect:Page_Home,
    },
    {
        path: Page_Home,
        name: 'Home',
        component: Control
    },
    {
        path: Page_Detail,
        name: 'Detail',
        component: Detail,
    },
]

const router = createRouter({
    history:createWebHashHistory(),
    routes: routes,
})
export default router
