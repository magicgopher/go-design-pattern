package chain

// 责任链模式

// LeaveRequest 请假请求
type LeaveRequest struct {
	Name     string // 请假人姓名
	Days     int    // 请假天数
	Reason   string // 请假原因
	Approved bool   // 是否批准
}

// Handler 处理者接口
type Handler interface {
	SetNext(handler Handler) Handler // 设置下一个处理者
	Handle(request *LeaveRequest)    // 请假请求
}

// BaseHandler 处理者基础结构
type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) Handler {
	h.next = handler
	return handler // 返回下一个处理者以支持链式调用
}

// TeamLeader 组长 - 可以批准 3 天以内的请假
type TeamLeader struct {
	BaseHandler
}

func NewTeamLeader() *TeamLeader {
	return &TeamLeader{}
}

func (t *TeamLeader) Handle(request *LeaveRequest) {
	if request.Days <= 3 {
		request.Approved = true
		return
	}
	if t.next != nil {
		t.next.Handle(request)
	}
}

// ProjectManager 项目经理 - 可以批准 7 天以内的请假
type ProjectManager struct {
	BaseHandler
}

func NewProjectManager() *ProjectManager {
	return &ProjectManager{}
}

func (p *ProjectManager) Handle(request *LeaveRequest) {
	if request.Days <= 7 {
		request.Approved = true
		return
	}
	if p.next != nil {
		p.next.Handle(request)
	}
}

// DepartmentManager 部门经理 - 可以批准 14 天以内的请假
type DepartmentManager struct {
	BaseHandler
}

func NewDepartmentManager() *DepartmentManager {
	return &DepartmentManager{}
}

func (d *DepartmentManager) Handle(request *LeaveRequest) {
	if request.Days <= 14 {
		request.Approved = true
		return
	}
	if d.next != nil {
		d.next.Handle(request)
	}
}

// CEO 公司CEO - 可以批准 14 天以上的请假
type CEO struct {
	BaseHandler
}

func NewCEO() *CEO {
	return &CEO{}
}

func (c *CEO) Handle(request *LeaveRequest) {
	request.Approved = true
}
