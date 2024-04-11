import {httpPostReq} from "../mod/http";

export const ApiAddControlClass = "/api/AddControlClass"
export const ApiUpdateControlClass = "/api/UpdateControlClass"
export const ApiDeleteControlClass = "/api/DeleteControlClass"
export const ApiUpdateControlClassOrder = "/api/UpdateControlClassOrder"
export const ApiGetControlClassList = "/api/GetControlClassList"
export const ApiGetControlClassInfo = "/api/GetControlClassInfo"


export interface ControlClass extends ControlClassId {
    control_name: string
    key_width: string
    key_height: string
}

export interface ControlClassId {
    control_id: number
}

export function NewControlClassId(): ControlClassId {
    return {
        control_id: 0
    }
}

export function NewControlClass(): ControlClass {
    return {
        key_height: "", key_width: "",
        control_name: '',
        control_id: 0
    }
}



export function NewControlClassList(): ControlClass[] {
    return []
}

// 新增coco
export async function AddControlClass(item: ControlClass) {
    return httpPostReq(ApiAddControlClass, item)
}



// 删除coco
export async function DeleteControlClass(id: number) {
    return httpPostReq(ApiDeleteControlClass, {control_id: id})
}


// 于服务器读取token
export async function GetControlClassList(req: GetControlClassListReq) {
    return httpPostReq(ApiGetControlClassList, req)
}

export interface GetControlClassListReq extends ControlClassId {
}

export function NewGetControlClassListReq(): GetControlClassListReq {
    return {
        control_id: 0
    }
}

export interface GetControlClassListResp {
    list: ControlClass[]
}

// 获取详情
export async function GetControlClassInfo(req: ControlClassId) {
    return httpPostReq(ApiGetControlClassInfo, req)
}

// 于服务器读取token
export async function EditControlClass(item: ControlClass) {
    if (item.control_id === 0) {
        return httpPostReq(ApiAddControlClass, item)
    } else {
        return httpPostReq(ApiUpdateControlClass, item)
    }
}

export async function UpdateControlClassOrder(item: UpdateControlClassOrderReq) {
    return httpPostReq(ApiUpdateControlClassOrder, item)
}

export interface UpdateControlClassOrderReq {
    order_list: ControlClassId[]
}

export function NewUpdateControlClassOrderReq(): UpdateControlClassOrderReq {
    return {
        order_list: [] as ControlClassId[]
    }

}