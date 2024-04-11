import axios from "axios";
import {runConfig} from "./config";

export interface ServerResponse {
    state: number,
    msg: string,
    data: any
}

export async function httpPostReq(api:string, para:any){
    const data = JSON.stringify(para);
    return httpPostReqByHost(runConfig.server,api,para)
}

export async function httpPostReqByHost(host:string,api:string, para:any){
    const data = JSON.stringify(para);

    const config = {
        method: 'post',
        url: host+api,
        headers: {
            'Content-Type': 'application/json'
        },
        data: data
    };
    let resp:ServerResponse = {
        state:-1,
        msg:"失败",
        data:null,
    }
    try {
        await axios(config)
            .then(function (response) {
                resp.state = response.data.state;
                resp.msg = response.data.msg;
                resp.data = response.data.data;
            })
            .catch(function (error) {
                console.log(error)
            });
    }catch (err){
        console.log(err)
    }
    return resp
}