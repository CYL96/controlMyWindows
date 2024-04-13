import {ControlClassId} from "./control";
import {httpPostReq} from "../mod/http";

export const ApiAddControlDetail = "/api/AddControlDetail"
export const ApiUpdateControlDetail = "/api/UpdateControlDetail"
export const ApiDeleteControlDetail = "/api/DeleteControlDetail"
export const ApiUpdateControlDetailOrder = "/api/UpdateControlDetailOrder"
export const ApiGetControlDetailList = "/api/GetControlDetailList"
export const ApiExecControlDetail = "/api/ExecControlDetail"
export const ApiStopControlDetail = "/api/StopControlDetail"
export const ApiGetNowMousePosition = "/api/GetNowMousePosition"

export const ControlType = {
    Normal: 1,// 1:快捷键
    Script: 2,// 2：脚本
    Explorer: 3,// 3：打开文件夹目录
    Website: 4,// 4:打开网页
    RunExe: 5,// 4:打开exe
    RunCmd: 6,// 4:打开cmd
}

export const KeyType = {
    Default: 1,// 单键
    Text: 2,// 文本
    ShortcutKey: 3,// 快捷键
    Mouse: 4, // 鼠标按键
    MouseMove: 5, // 鼠标移动
    MouseScroll: 6, // 鼠标移动
    Delay: 99,// 延迟
}


export const RunState = {
    Free: 1, // 空闲
    Running: 2 // 运行中
}

export const PressType = {
    Click: 1,// 单击
    DoubleClick: 2,// 双击
    PressDown: 3,// 按下
    PressUp: 4,// 抬起
}

export interface ControlDetail extends ControlDetailId {
    detail_name: string
    detail_color: string
    control_type: number // 1:快捷键，2：脚本 3：打开文件夹目录  4:打开网页
    path: string
    run_state: number
    detail_key: ControlKeyList[]
}

export interface ControlDetailId {
    detail_id: number
}

export interface ControlKeyList extends ControlDetailKey {
    key_type: number
    key_press: number
    input: string
    key_list: ControlDetailKey[]
    delay: number
    point_x: number
    point_y: number
    scroll: number
    scroll_dir: number
}

export interface ControlDetailKey {
    id: number
    key: string
}

export function NewControlDetail(): ControlDetail {
    return {
        path: "",
        run_state: 0,
        control_type: 0,
        detail_color: "",
        detail_id: 0,
        detail_key: [] as ControlKeyList[],
        detail_name: ""

    }
}

export function CopyControlDetail(item: ControlDetail): ControlDetail {
    return {
        path: "",
        run_state: 0,
        control_type: item.control_type,
        detail_color: item.detail_color,
        detail_id: item.detail_id,
        detail_key: CopyControlKeyListList(item.detail_key),
        detail_name: item.detail_name
    }

}

export function NewControlKeyListList(): ControlKeyList[] {
    return []
}

export function CopyControlKeyListList(item: ControlKeyList[]): ControlKeyList[] {
    let tmp = NewControlKeyListList()
    item.forEach(res => {
        tmp.push(CopyControlKeyList(res))
    })
    return tmp
}

export function CopyControlKeyList(item: ControlKeyList): ControlKeyList {
    let tmp = NewControlKeyList()
    tmp.delay = item.delay
    tmp.input = item.input
    tmp.key_list = CopyControlDetailKeyList(item.key_list)
    tmp.key_press = item.key_press
    tmp.key_type = item.key_type
    tmp.id = item.id
    tmp.key = item.key
    tmp.point_x = item.point_x
    tmp.point_y = item.point_y
    tmp.scroll = item.scroll
    tmp.scroll_dir = item.scroll_dir
    return tmp
}


export function NewControlDetailId(): ControlDetailId {
    return {
        detail_id: 0
    }
}

export function NewControlDetailList(): ControlDetail[] {
    return []
}

export function CopyControlDetailKeyList(item: ControlDetailKey[]): ControlDetailKey[] {
    let tmp = NewControlDetailKeyList()
    item.forEach(res => {
        tmp.push({
            id: res.id,
            key: res.key
        })
    })
    return tmp

}

export function NewControlDetailKeyList(): ControlDetailKey[] {
    return []
}

export function NewControlDetailKey(): ControlDetailKey {
    return {
        id: 0,
        key: ''
    }
}

export function NewControlKeyList(): ControlKeyList {
    return {
        id: 0, key: "",
        key_type: 0,
        key_press: 0,
        input: '',
        key_list: [] as ControlDetailKey[],
        delay: 0,
        point_x: 0, point_y: 0, scroll: 0, scroll_dir: 0,
    }
}

export function NewControlKeyListNormal(): ControlKeyList {
    return {
        id: 0, key: "",
        key_type: KeyType.Default,
        key_press: PressType.Click,
        input: '',
        key_list: [] as ControlDetailKey[],
        delay: 0,
        point_x: 0, point_y: 0, scroll: 0, scroll_dir: 0,
    }

}

export function GetKeyStr(item: ControlDetailKey[]): string {
    let tmp = ""
    item.forEach(res => {
        if (tmp != "") {
            tmp += " "
        }
        tmp += res.key
    })
    return tmp
}

//接口
// 添加明细
export async function EditControlDetail(req: EditControlDetailReq) {
    if (req.detail.detail_id === 0) {
        return httpPostReq(ApiAddControlDetail, req)
    } else {
        return httpPostReq(ApiUpdateControlDetail, req)
    }
}

export interface EditControlDetailReq extends ControlClassId {
    detail: ControlDetail
}

export function NewEditControlDetailReq(): EditControlDetailReq {
    return {
        control_id: 0,
        detail: NewControlDetail()
    }

}


// 删除明细
export async function DeleteControlDetail(req: DeleteControlDetailReq) {
    return httpPostReq(ApiDeleteControlDetail, req)
}

export interface DeleteControlDetailReq extends ControlDetailId, ControlClassId {
}

export function NewDeleteControlDetailReq(): DeleteControlDetailReq {
    return {
        control_id: 0,
        detail_id: 0
    }
}

// 更新顺序
export async function UpdateControlDetailOrder(req: UpdateControlDetailOrderReq) {
    return httpPostReq(ApiUpdateControlDetailOrder, req)
}

export interface UpdateControlDetailOrderReq extends ControlClassId {
    detail: ControlDetailId[]
}

export function NewUpdateControlDetailOrderReq(): UpdateControlDetailOrderReq {
    return {
        control_id: 0,
        detail: []
    }

}

// 获取详情
export async function GetControlDetailList(req: GetControlDetailListReq) {
    return httpPostReq(ApiGetControlDetailList, req)
}

export interface GetControlDetailListReq extends ControlClassId {
}

export function NewGetControlDetailListReq(): GetControlDetailListReq {
    return {
        control_id: 0
    }
}

export interface GetControlDetailListResp {
    detail: ControlDetail[]
}

// 执行detail
export async function ExecControlDetail(req: ExecControlDetailReq) {
    return httpPostReq(ApiExecControlDetail, req)
}

export interface ExecControlDetailReq extends ControlClassId, ControlDetailId {
}

export function NewExecControlDetailReq(): ExecControlDetailReq {
    return {
        control_id: 0,
        detail_id: 0
    }
}


// 停止detail
export async function ExecStopControlDetail(req: StopControlDetailReq) {
    return httpPostReq(ApiStopControlDetail, req)
}


export interface StopControlDetailReq extends ControlClassId, ControlDetailId {
}

export function NewStopControlDetailReq(): ExecControlDetailReq {
    return {
        control_id: 0,
        detail_id: 0
    }
}

export async function GetNowMousePosition() {
    return httpPostReq(ApiGetNowMousePosition, {})
}

export interface GetNowMousePositionResp {
    x: number
    y: number
}
