import router from "./router";

export const Page_Home = "/Home"
export function GotoHome(){
    router.push(Page_Home).then(r => {})
}


export const Page_Detail = "/Detail"
export function GotoDetail(info :DetailQuery){
    router.push({
        path:Page_Detail,
        query:{
            ...info
        }
    }).then(r => {})
}
export interface DetailQuery {
    control_id: string
}
export function NewDetailQuery(): DetailQuery {
    return {
        control_id:''
    }
}