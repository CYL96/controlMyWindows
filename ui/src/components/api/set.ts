import {httpPostReq} from "../mod/http";

export interface SystemSetting {
    run_port: number,
    sound_open: boolean,
}


export const NewSystemSetting = (): SystemSetting => {
    return {
        run_port: 0,
        sound_open: false,
    }
}

export const ApiSetSystemConfig = "/api/SetSystemConfig"

export const SetSystemConfig = (item: SystemSetting) => {
    return httpPostReq(ApiSetSystemConfig, item)
}


export const ApiGetSystemConfig = "/api/GetSystemConfig"

export const GetSystemConfig = () => {
    return httpPostReq(ApiGetSystemConfig, {})
}
