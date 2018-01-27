package entity

type JobInfo struct {
	Id       string
	Location string
	Date     string
	Title    string
	Url      string
}

type JobInfoList []JobInfo

func (list *JobInfoList) Index(item JobInfo) (int, bool) {
	for i, v := range *list {
		if v.Id == item.Id {
			return i, true
		}
	}
	return -1, false
}


func (list *JobInfoList) Add(item JobInfo){
	*list =append(*list,item)
}

func InitJobInfoList() *JobInfoList {
	var list JobInfoList
	list = make([]JobInfo,0)
	return &list
}
