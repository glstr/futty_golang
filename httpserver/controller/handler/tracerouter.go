package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/glstr/futty_golang/context"
	"github.com/glstr/futty_golang/httpserver/controller/middleware"
	"github.com/glstr/futty_golang/httpserver/controller/views"
	"github.com/glstr/futty_golang/logger"
	"github.com/glstr/futty_golang/repo"
	"github.com/glstr/futty_golang/repo/models"
	"github.com/glstr/futty_golang/service/nettool"
	"github.com/glstr/futty_golang/service/task"
)

type TraceRouterRequest struct {
	IP string `json:"ip"`
}

type TraceRouterResponse struct {
	TaskID int32 `json:"task_id"`
	middleware.CommonResponse
}

func TraceRouterHandler(c *gin.Context) {
	ctx := context.GetContext()
	var res TraceRouterResponse
	res.CommonResponse.RequestId = ctx.Logid
	var err error
	defer func() {
		res.CommonResponse.ErrorCode, res.CommonResponse.ErrorMsg =
			views.GetErrInfoFromErr(err)
		c.JSON(200, res)
		logger.Notice(ctx.LogBuffer.String())
		context.PutContext(ctx)
	}()

	var req TraceRouterRequest
	err = middleware.GetJsonParam(c, &req)
	if err != nil {
		ctx.LogBuffer.WriteLog("get_param[failed] error_msg[%s]", err.Error())
		return
	}
	ctx.LogBuffer.WriteLog("IP[%s]", req.IP)

	traceRouterFunc := func() error {
		newLogger := context.NewLogBuffer()
		newLogger.WriteLog("tracerouter start ")
		defer func() {
			logger.Notice(newLogger.String())
		}()
		result, err := nettool.TraceRouter(newLogger, req.IP)
		if err != nil {
			newLogger.WriteLog("trace router failed, err:%s", err.Error())
			return err
		}

		traceDataRepo := repo.GetTraceDataRepo()
		resultStr, err := json.Marshal(result)
		if err != nil {
			newLogger.WriteLog("result marshal failed:%s", err.Error())
			return err
		}

		newLogger.WriteLog("trace router result:%s", resultStr)
		data := &models.TraceData{
			IP:   req.IP,
			Data: resultStr,
		}
		err = traceDataRepo.SaveData(req.IP, data)
		if err != nil {
			newLogger.WriteLog("save data failed:%s", err.Error())
			return err
		}

		return nil
	}

	var extraInfo []byte
	extraInfo, err = json.Marshal(req)
	if err != nil {
		ctx.LogBuffer.WriteLog("marshal_req[failed] error_msg[%s]", err.Error())
		return
	}

	taskSer := task.GetTaskService()
	var taskID int32
	taskID, err = taskSer.SetTask(traceRouterFunc, extraInfo)
	if err != nil {
		ctx.LogBuffer.WriteLog("set_task[failed] error_msg[%s]", err.Error())
		return
	}

	res.TaskID = taskID
}

type GetRouterInfoRequest struct {
	TaskID int32 `json:"task_id"`
}

type GetRouterInfoResponse struct {
	Result string         `json:"result"`
	State  task.TaskState `json:"state"`
	middleware.CommonResponse
}

func GetRouterInfoHandler(c *gin.Context) {
	ctx := context.GetContext()
	var res GetRouterInfoResponse
	res.CommonResponse.RequestId = ctx.Logid
	var err error
	defer func() {
		res.CommonResponse.ErrorCode, res.CommonResponse.ErrorMsg =
			views.GetErrInfoFromErr(err)
		c.JSON(200, res)
		logger.Notice(ctx.LogBuffer.String())
		context.PutContext(ctx)
	}()

	var req GetRouterInfoRequest
	err = middleware.GetJsonParam(c, &req)
	if err != nil {
		ctx.LogBuffer.WriteLog("get_param[failed] error_msg[%s]", err.Error())
		return
	}
	ctx.LogBuffer.WriteLog("TaskID[%d]", req.TaskID)

	taskSer := task.GetTaskService()
	var result *task.TaskResult
	result, err = taskSer.GetTask(req.TaskID)
	if err != nil {
		ctx.LogBuffer.WriteLog("get_task[failed] error_msg[%s]", err.Error())
		return
	}

	res.State = result.GetState()
	ctx.LogBuffer.WriteLog("task_state[%d]", res.State)
	if res.State == task.TaskStatInit || res.State == task.TaskStatFailed {
		return
	}

	var traceRouteReq TraceRouterRequest
	err = json.Unmarshal(result.GetExtraInfo(), &traceRouteReq)
	if err != nil {
		ctx.LogBuffer.WriteLog("unmarshal_extra_info[failed] error_msg[%s]", err.Error())
		return
	}

	traceDataRepo := repo.GetTraceDataRepo()
	data, err := traceDataRepo.GetData(traceRouteReq.IP)
	if err != nil {
		ctx.LogBuffer.WriteLog("get_data[failed] error_msg[%s]", err.Error())
		return

	}

	res.Result = string(data.Data)
}
