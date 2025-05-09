/**************************************************
*文件名：hd.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/3/26
**************************************************/

package hd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"server/src/mod"
	"server/src/module"
	"server/src/msg"
	"server/src/runCtx"
)

// AddControlClass
// @Summary	添加一个控制模块
// @Accept        json
// @Produce       json
// @Description	添加一个控制模块
// @Tags			Control
// @Param body body AddControlClassReq true "请求"
// @success 200 {object} GinResponse{data=AddControlClassResp} "desc"
// @Router			/api/AddControlClass [post]
func AddControlClass(c *gin.Context) {
	var err error
	var result AddControlClassResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req AddControlClassReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = mod.AddControlClass(ctx, req.AddControlClassPara)
	return
}

type (
	AddControlClassReq struct {
		mod.AddControlClassPara
	}
	AddControlClassResp struct {
	}
)

// UpdateControlClass
// @Summary	更新控制模块
// @Accept        json
// @Produce       json
// @Description	更新控制模块
// @Tags			Control
// @Param body body UpdateControlClassReq true "请求"
// @success 200 {object} GinResponse{data=UpdateControlClassResp} "desc"
// @Router			/api/UpdateControlClass [post]
func UpdateControlClass(c *gin.Context) {
	var err error
	var result UpdateControlClassResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req UpdateControlClassReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = mod.UpdateControlClass(ctx, req.UpdateControlClassPara)
	return
}

type (
	UpdateControlClassReq struct {
		mod.UpdateControlClassPara
	}
	UpdateControlClassResp struct {
	}
)

// DeleteControlClass
// @Summary	删除控制模块
// @Accept        json
// @Produce       json
// @Description	删除控制模块
// @Tags			Control
// @Param body body DeleteControlClassReq true "请求"
// @success 200 {object} GinResponse{data=DeleteControlClassResp} "desc"
// @Router			/api/DeleteControlClass [post]
func DeleteControlClass(c *gin.Context) {
	var err error
	var result DeleteControlClassResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req DeleteControlClassReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = mod.DeleteControlClass(ctx, req.DeleteControlClassPara)
	return
}

type (
	DeleteControlClassReq struct {
		mod.DeleteControlClassPara
	}
	DeleteControlClassResp struct {
	}
)

// UpdateControlClassOrder
// @Summary	更新控制顺序
// @Accept        json
// @Produce       json
// @Description	更新控制顺序
// @Tags			Control
// @Param body body UpdateControlClassOrderReq true "请求"
// @success 200 {object} GinResponse{data=UpdateControlClassOrderResp} "desc"
// @Router			/api/UpdateControlClassOrder [post]
func UpdateControlClassOrder(c *gin.Context) {
	var err error
	var result UpdateControlClassOrderResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req UpdateControlClassOrderReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = mod.UpdateControlClassOrder(ctx, req.UpdateControlClassOrderPara)
	return
}

type (
	UpdateControlClassOrderReq struct {
		mod.UpdateControlClassOrderPara
	}
	UpdateControlClassOrderResp struct {
	}
)

// GetControlClassList
// @Summary	获取控制模块列表
// @Accept        json
// @Produce       json
// @Description	获取控制模块列表
// @Tags			Control
// @Param body body GetControlClassListReq true "请求"
// @success 200 {object} GinResponse{data=GetControlClassListResp} "desc"
// @Router			/api/GetControlClassList [post]
func GetControlClassList(c *gin.Context) {
	var err error
	var result GetControlClassListResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req GetControlClassListReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	result.GetControlClassListResult, err = mod.GetControlClassList(ctx, req.GetControlClassListPara)
	return
}

type (
	GetControlClassListReq struct {
		mod.GetControlClassListPara
	}
	GetControlClassListResp struct {
		mod.GetControlClassListResult
	}
)

// GetControlClassInfo
// @Summary	获取控制模块详情
// @Accept        json
// @Produce       json
// @Description	获取控制模块详情
// @Tags			Control
// @Param body body GetControlClassInfoReq true "请求"
// @success 200 {object} GinResponse{data=GetControlClassInfoResp} "desc"
// @Router			/api/GetControlClassInfo [post]
func GetControlClassInfo(c *gin.Context) {
	var err error
	var result GetControlClassInfoResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req GetControlClassInfoReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	result.GetControlClassInfoResult, err = mod.GetControlClassInfo(ctx, req.GetControlClassInfoPara)
	return
}

type (
	GetControlClassInfoReq struct {
		mod.GetControlClassInfoPara
	}
	GetControlClassInfoResp struct {
		mod.GetControlClassInfoResult
	}
)

// ActivateControlClassCombinationKey
// @Summary	启用当前控制模块组合键
// @Accept        json
// @Produce       json
// @Description	启用当前控制模块组合键
// @Tags			Control
// @Param body body ActivateControlClassCombinationKeyReq true "请求"
// @success 200 {object} GinResponse{data=ActivateControlClassCombinationKeyResp} "desc"
// @Router			/api/ActivateControlClassCombinationKey [post]
func ActivateControlClassCombinationKey(c *gin.Context) {
	var err error
	var result ActivateControlClassCombinationKeyResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req ActivateControlClassCombinationKeyReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = mod.ActivateControlClassCombinationKey(ctx, req.ActivateControlClassCombinationKeyPara)
	return
}

type (
	ActivateControlClassCombinationKeyReq struct {
		mod.ActivateControlClassCombinationKeyPara
	}
	ActivateControlClassCombinationKeyResp struct {
	}
)

// StopCombinationKey
// @Summary	启用当前控制模块组合键
// @Accept        json
// @Produce       json
// @Description	启用当前控制模块组合键
// @Tags			Control
// @Param body body StopCombinationKeyReq true "请求"
// @success 200 {object} GinResponse{data=StopCombinationKeyResp} "desc"
// @Router			/api/StopCombinationKey [post]
func StopCombinationKey(c *gin.Context) {
	var err error
	var result StopCombinationKeyResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req StopCombinationKeyReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = mod.StopCombinationKey(ctx, req.StopCombinationKeyPara)
	return
}

type (
	StopCombinationKeyReq struct {
		mod.StopCombinationKeyPara
	}
	StopCombinationKeyResp struct {
	}
)

// GetNowCombinationState
// @Summary	获取控制模块详情
// @Accept        json
// @Produce       json
// @Description	获取控制模块详情
// @Tags			Control
// @Param body body GetNowCombinationStateReq true "请求"
// @success 200 {object} GinResponse{data=GetNowCombinationStateResp} "desc"
// @Router			/api/GetNowCombinationState [post]
func GetNowCombinationState(c *gin.Context) {
	var err error
	var result GetNowCombinationStateResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req GetNowCombinationStateReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	result.GetNowCombinationStateResult, err = mod.GetNowCombinationState(ctx, req.GetNowCombinationStatePara)
	return
}

type (
	GetNowCombinationStateReq struct {
		mod.GetNowCombinationStatePara
	}
	GetNowCombinationStateResp struct {
		mod.GetNowCombinationStateResult
	}
)

// AddControlDetail
// @Summary	添加一个控制模块-键
// @Accept        json
// @Produce       json
// @Description	控制模块-键
// @Tags			控制模块-键
// @Param body body AddControlDetailReq true "请求"
// @success 200 {object} GinResponse{data=AddControlDetailResp} "desc"
// @Router			/api/AddControlDetail [post]
func AddControlDetail(c *gin.Context) {
	var err error
	var result AddControlDetailResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req AddControlDetailReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = mod.AddControlDetail(ctx, req.AddControlDetailPara)
	return
}

type (
	AddControlDetailReq struct {
		mod.AddControlDetailPara
	}
	AddControlDetailResp struct {
	}
)

// UpdateControlDetail
// @Summary	更新控制模块
// @Accept        json
// @Produce       json
// @Description	更新控制模块
// @Tags			控制模块-键
// @Param body body UpdateControlDetailReq true "请求"
// @success 200 {object} GinResponse{data=UpdateControlDetailResp} "desc"
// @Router			/api/UpdateControlDetail [post]
func UpdateControlDetail(c *gin.Context) {
	var err error
	var result UpdateControlDetailResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req UpdateControlDetailReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = mod.UpdateControlDetail(ctx, req.UpdateControlDetailPara)

	return
}

type (
	UpdateControlDetailReq struct {
		mod.UpdateControlDetailPara
	}
	UpdateControlDetailResp struct {
	}
)

// DeleteControlDetail
// @Summary	删除控制模块
// @Accept        json
// @Produce       json
// @Description	删除控制模块
// @Tags			控制模块-键
// @Param body body DeleteControlDetailReq true "请求"
// @success 200 {object} GinResponse{data=DeleteControlDetailResp} "desc"
// @Router			/api/DeleteControlDetail [post]
func DeleteControlDetail(c *gin.Context) {
	var err error
	var result DeleteControlDetailResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req DeleteControlDetailReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = mod.DeleteControlDetail(ctx, req.DeleteControlDetailPara)
	return
}

type (
	DeleteControlDetailReq struct {
		mod.DeleteControlDetailPara
	}
	DeleteControlDetailResp struct {
	}
)

// UpdateControlDetailOrder
// @Summary	更新控制顺序
// @Accept        json
// @Produce       json
// @Description	更新控制顺序
// @Tags			控制模块-键
// @Param body body UpdateControlDetailOrderReq true "请求"
// @success 200 {object} GinResponse{data=UpdateControlDetailOrderResp} "desc"
// @Router			/api/UpdateControlDetailOrder [post]
func UpdateControlDetailOrder(c *gin.Context) {
	var err error
	var result UpdateControlDetailOrderResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req UpdateControlDetailOrderReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = mod.UpdateControlDetailOrder(ctx, req.UpdateControlDetailOrderPara)
	return
}

type (
	UpdateControlDetailOrderReq struct {
		mod.UpdateControlDetailOrderPara
	}
	UpdateControlDetailOrderResp struct {
	}
)

// GetControlDetailList
// @Summary	获取键列表
// @Accept        json
// @Produce       json
// @Description	获取键列表
// @Tags			控制模块-键
// @param body body GetControlDetailListReq true "请求"
// @success 200 {object} GinResponse{data=GetControlDetailListResp} "desc"
// @Router			/api/GetControlDetailList [post]
func GetControlDetailList(c *gin.Context) {
	var err error
	var result GetControlDetailListResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req GetControlDetailListReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	result.GetControlDetailListResult, err = mod.GetControlDetailList(ctx, req.GetControlDetailListPara)
}

type (
	GetControlDetailListReq struct {
		mod.GetControlDetailListPara
	}
	GetControlDetailListResp struct {
		mod.GetControlDetailListResult
	}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// GetMsgWS
// @Summary         建立消息推送长连接
// @Accept          json
// @Produce         json
// @Description	    建立消息推送长连接
// @Tags            长链接
// @param   request body    GetMsgWSRequest   true    "请求"
// @success 200 {object}    models.GinResponse{data=GetMsgWSResponse}    "desc"
// @Router  /runMsg    [get]
func GetMsgWS(c *gin.Context) {
	var err error
	ctx := runCtx.FromContext(c)
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Info("设备推送WS："+conn.RemoteAddr().String(), conn.RemoteAddr().Network())
	defer conn.Close()
	msg.MsgCenter.AddDevicePusher(conn)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			ctx.Error(err.Error())
			return
		}
	}
}

// ExecControlDetail
// @Summary	执行key
// @Accept        json
// @Produce       json
// @Description	执行key
// @Tags			控制模块-键
// @param body body ExecControlDetailReq true "请求"
// @success 200 {object} GinResponse{data=ExecControlDetailResp} "desc"
// @Router			/api/ExecControlDetail [post]
func ExecControlDetail(c *gin.Context) {
	var err error
	var result ExecControlDetailResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req ExecControlDetailReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = mod.ExecControlDetail(ctx, req.ExecControlDetailPara)
	return
}

type (
	ExecControlDetailReq struct {
		mod.ExecControlDetailPara
	}
	ExecControlDetailResp struct {
	}
)

// StopControlDetail
// @Summary	停止执行key
// @Accept        json
// @Produce       json
// @Description	执行key
// @Tags			控制模块-键
// @param body body StopControlDetailReq true "请求"
// @success 200 {object} GinResponse{data=StopControlDetailResp} "desc"
// @Router			/api/StopControlDetail [post]
func StopControlDetail(c *gin.Context) {
	var err error
	var result StopControlDetailResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req StopControlDetailReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = mod.StopControlDetail(ctx, req.StopControlDetailPara)
	return
}

type (
	StopControlDetailReq struct {
		mod.StopControlDetailPara
	}
	StopControlDetailResp struct {
	}
)

// GetNowMousePosition
// @Summary	获取当前鼠标位置
// @Accept        json
// @Produce       json
// @Description	执行key
// @Tags			控制模块-键
// @param body body GetNowMousePositionReq true "请求"
// @success 200 {object} GinResponse{data=GetNowMousePositionResp} "desc"
// @Router			/api/GetNowMousePosition [post]
func GetNowMousePosition(c *gin.Context) {
	var err error
	var result GetNowMousePositionResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req GetNowMousePositionReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	result.GetNowMousePositionResult, err = mod.GetNowMousePosition(ctx, req.GetNowMousePositionPara)
	return
}

type (
	GetNowMousePositionReq struct {
		mod.GetNowMousePositionPara
	}
	GetNowMousePositionResp struct {
		mod.GetNowMousePositionResult
	}
)

// GetSystemConfig
// @Summary	获取当前系统设置
// @Accept        json
// @Produce       json
// @Description	获取当前系统设置
// @Tags			系统设置
// @param body body GetSystemConfigReq true "请求"
// @success 200 {object} GinResponse{data=GetSystemConfigResp} "desc"
// @Router			/api/GetSystemConfig [post]
func GetSystemConfig(c *gin.Context) {
	var err error
	var result GetSystemConfigResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req GetSystemConfigReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}

	result.GetNowSystemConfigResp, err = mod.GetNowSystemConfig(ctx)
	return
}

type (
	GetSystemConfigReq struct {
	}
	GetSystemConfigResp struct {
		mod.GetNowSystemConfigResp
	}
)

// SetSystemConfig
// @Summary	设置当前系统设置
// @Accept        json
// @Produce       json
// @Description	设置当前系统设置
// @Tags			系统设置
// @param body body SetSystemConfigReq true "请求"
// @success 200 {object} GinResponse{data=SetSystemConfigResp} "desc"
// @Router			/api/SetSystemConfig [post]
func SetSystemConfig(c *gin.Context) {
	var err error
	var result SetSystemConfigResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req SetSystemConfigReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = mod.EditSystemConfig(ctx, req.EditSystemConfigPara)
	return
}

type (
	SetSystemConfigReq struct {
		mod.EditSystemConfigPara
	}
	SetSystemConfigResp struct {
	}
)

// GetNowIconList
// @Summary	获取当前图标列表
// @Accept        json
// @Produce       json
// @Description	获取当前图标列表
// @Tags			系统设置
// @param body body GetNowIconListReq true "请求"
// @success 200 {object} GinResponse{data=GetNowIconListResp} "desc"
// @Router			/api/GetNowIconList [post]
func GetNowIconList(c *gin.Context) {
	var err error
	var result GetNowIconListResp
	ctx := runCtx.FromContext(c)
	defer func() {
		if err != nil {
			GinResponseWithStateAndMsg(c, StateFailed, err.Error())
		} else {
			GinResponseOk(c, result)
		}
	}()
	var req GetNowIconListReq
	err = c.BindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}

	result.List, err = module.GetNowIcon(ctx, "icon")
	return
}

type (
	GetNowIconListReq struct {
	}
	GetNowIconListResp struct {
		List []module.IconT
	}
)
