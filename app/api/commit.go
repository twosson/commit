package api

import (
	"commit/app/model"
	"commit/extern/response"
	"encoding/base64"
	ffi "github.com/filecoin-project/filecoin-ffi"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"time"
)

var Commit = commitApi{}

type commitApi struct{}

// Index is a demonstration route handler for output "Hello World!".
func (*commitApi) Index(r *ghttp.Request) {
	var data *model.CommitReq
	if err := r.Parse(&data); err != nil {
		g.Log("commit").Errorf("commit err: %s", err.Error())
		response.JsonExit(r, 1, err.Error())
		return
	}
	g.Log("commit").Infof("commit start miner: %d, sector: %d, time: %s", data.MinerNumber, data.SectorNumber, time.Now().String())
	phase1Output, err := base64.StdEncoding.DecodeString(data.Phase1Output)
	if err != nil {
		g.Log("commit").Errorf("commit err: %s", err.Error())
		response.JsonExit(r, 1, err.Error())
		return
	}
	proof, err := ffi.SealCommitPhase2(phase1Output, abi.SectorNumber(data.SectorNumber), abi.ActorID(data.MinerNumber))
	if err != nil {
		g.Log("commit").Errorf("commit err: %s", err.Error())
		response.JsonExit(r, 1, err.Error())
		return
	}

	rsp := model.CommitRsp{
		MinerNumber:  data.MinerNumber,
		SectorNumber: data.SectorNumber,
		Proof:        base64.StdEncoding.EncodeToString(proof),
	}

	g.Log("commit").Infof("commit end: %s", time.Now().String())
	response.Json(r, 0, "success", rsp)
}
