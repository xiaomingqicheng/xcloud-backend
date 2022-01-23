package task

func GetResult(taskUuid string) string {
	bk,_:=taskserver.GetBackend().GetState(taskUuid)
	return bk.State
}