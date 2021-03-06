package core

/// <summary>
/// IWorker 业务处理
/// <summary>
type IWorker interface {
	OnInit(args ...interface{})
	OnMessage(cmd uint32, msg interface{}, peer Session)
	//SetProc(c IProc)
	GetProc() IProc
	SetDispatcher(c IProc)
	GetDispatcher() IProc
	ResetDispatcher()
}

type IWorkerCreator interface {
	Create(c IProc) IWorker
}
