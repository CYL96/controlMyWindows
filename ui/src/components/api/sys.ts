import {httpPostReq} from "../mod/http";
import {ApiAddControlClass, ControlClass} from "./control";
import {runConfig} from "../mod/config";

export const ApiExit = "/api/Exit"


export async function ApiExitControl() {
    return httpPostReq(ApiExit, {})
}

export const  GetIconSrc = (src: string,ts :number): string => {
    return runConfig.server + '/' + src + '?ts=' + ts
}