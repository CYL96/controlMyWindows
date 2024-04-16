import {httpPostReq, httpPostReqCommonNotify} from "../mod/http";
import {ApiGetNowMousePosition} from "./detail";

export interface SystemSetting {
    run_port: number,
    sound_open: boolean,
    is_scale: boolean,
}


export const NewSystemSetting = (): SystemSetting => {
    return {
        run_port: 0,
        sound_open: false,
        is_scale: false,
    }
}

export const ApiSetSystemConfig = "/api/SetSystemConfig"

export async function SetSystemConfig(item: SystemSetting) {
    return httpPostReqCommonNotify(ApiSetSystemConfig, item)
}


export const ApiGetSystemConfig = "/api/GetSystemConfig"

export async function GetSystemConfig() {
    return httpPostReqCommonNotify(ApiGetSystemConfig, {})
}


export const ApiGetNowIconList = "/api/GetNowIconList"


export async function GetNowIconList() {
    return httpPostReqCommonNotify(ApiGetNowIconList, {})
}


export interface IconT {
    icon_name: string
    icon_path: string
}

export function NewIconT(): IconT {
    return {
        icon_name: "",
        icon_path: "",
    }
}

export function NewIconT2(path:string): IconT {
    return {
        icon_name: "",
        icon_path: path,
    }
}
export interface IconResp {
    List: IconT[]
}

export const NewIconTList = (): IconT[] => {
    return []
}

