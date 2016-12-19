package tasks

import "github.com/medcl/gopa/core/types"
import (
	"github.com/asdine/storm"
	log "github.com/cihub/seelog"
	"time"
	"github.com/medcl/gopa/core/global"
	"path"
)


var db *storm.DB
var inited bool
func Start() error  {
	var err error
	file:= path.Join(global.Env().RuntimeConfig.PathConfig.Data,"task_db")
	db, err = storm.Open(file)
	inited=true
	return err

}

func Stop()  {
	db.Close()
}

func CreateSeed(task types.TaskSeed)  {
	if(!inited){Start()}
	log.Trace("start create seed")
	time:=time.Now()
	task.CreateTime=&time
	err := db.Save(&task)
	if(err!=nil){
		panic(err)
	}
	global.Env().Channels.PushUrlToCheck(task)
}

func DeleteSeed(id int)  {
	if(!inited){Start()}
	log.Trace("start delete seed: ",id )
	task:=types.TaskSeed{ID:id}
	err := db.DeleteStruct(&task)
	if(err!=nil){
		panic(err)
	}
}

func GetSeed(id int) (types.TaskSeed,error)  {
	if(!inited){Start()}
	log.Trace("start get seed: ",id)
	task:=types.TaskSeed{}
	err := db.One("ID", id, &task)
	return task,err
}

func GetSeedList()[]types.TaskSeed {
	if(!inited){Start()}
	log.Trace("start get all seeds")
	var tasks []types.TaskSeed
	err := db.AllByIndex("CreateTime",&tasks)
	if(err!=nil){
		panic(err)
	}
	return tasks
}



func CreateTask(task types.CrawlerTask)  {
	if(!inited){Start()}
	log.Trace("start create crawler task")
	time:=time.Now()
	task.CreateTime=&time
	err := db.Save(&task)
	if(err!=nil){
		panic(err)
	}
}

func DeleteTask(id int)  {
	if(!inited){Start()}
	log.Trace("start delete crawler task: ",id )
	task:=types.CrawlerTask{ID:id}
	err := db.DeleteStruct(&task)
	if(err!=nil){
		panic(err)
	}
}

func GetTask(id int) (types.CrawlerTask,error)  {
	if(!inited){Start()}
	log.Trace("start get seed: ",id)
	task:=types.CrawlerTask{}
	err := db.One("ID", id, &task)
	return task,err
}

func GetTaskList(from,size int)(int,[]types.CrawlerTask,error) {
	if(!inited){Start()}
	log.Trace("start get all crawler tasks")
	var tasks []types.CrawlerTask
	total,err:=db.Count(&types.CrawlerTask{})
	if(err!=nil){
		log.Error(err)
	}
	err= db.AllByIndex("CreateTime",&tasks,storm.Skip(from),storm.Limit(size))
	return total,tasks,err
}