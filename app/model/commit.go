package model

type CommitReq struct {
	MinerNumber  uint64 `v:"required#矿工号不能为空"`
	SectorNumber uint64 `v:"required#扇区号不能为空"`
	Phase1Output string `v:"required#Phase1Output不能为空"`
}

type CommitRsp struct {
	MinerNumber  uint64
	SectorNumber uint64
	Proof        string
}
