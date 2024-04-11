import {httpPostReq} from "../mod/http";
import {ApiAddControlClass, ControlClass} from "./control";

export const ApiExit = "/api/Exit"


export async function ApiExitControl() {
    return httpPostReq(ApiExit, {})
}