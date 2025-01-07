package chain_test

import (
	chain "github.com/magicgopher/go-design-pattern/14_chain"
	"testing"
)

// TestLeaveApprovalChain
func TestLeaveApprovalChain(t *testing.T) {
	// 创建责任链
	teamLeader := chain.NewTeamLeader()
	projectManager := chain.NewProjectManager()
	departmentManager := chain.NewDepartmentManager()
	ceo := chain.NewCEO()

	// 设置责任链顺序
	teamLeader.SetNext(projectManager)
	projectManager.SetNext(departmentManager)
	departmentManager.SetNext(ceo)

	tests := []struct {
		name     string
		request  chain.LeaveRequest
		expected bool
	}{
		{
			name: "2天请假-组长审批",
			request: chain.LeaveRequest{
				Name:   "张三",
				Days:   2,
				Reason: "参加朋友婚礼",
			},
			expected: true,
		},
		{
			name: "5天请假-项目经理审批",
			request: chain.LeaveRequest{
				Name:   "李四",
				Days:   5,
				Reason: "回家探亲",
			},
			expected: true,
		},
		{
			name: "10天请假-部门经理审批",
			request: chain.LeaveRequest{
				Name:   "王五",
				Days:   10,
				Reason: "出国旅游",
			},
			expected: true,
		},
		{
			name: "20天请假-CEO审批",
			request: chain.LeaveRequest{
				Name:   "赵六",
				Days:   20,
				Reason: "长期休养",
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 处理请求
			teamLeader.Handle(&tt.request)

			if tt.request.Approved != tt.expected {
				t.Errorf("期望审批结果为 %v，但得到 %v", tt.expected, tt.request.Approved)
			}
		})
	}
}

// TestLeaveApprovalChainCalls
func TestLeaveApprovalChainCalls(t *testing.T) {
	// 使用链式调用创建责任链
	teamLeader := chain.NewTeamLeader()
	teamLeader.SetNext(chain.NewProjectManager()).
		SetNext(chain.NewDepartmentManager()).
		SetNext(chain.NewCEO())

	tests := []struct {
		name     string
		request  chain.LeaveRequest
		expected bool
	}{
		{
			name: "2天请假-组长审批",
			request: chain.LeaveRequest{
				Name:   "张三",
				Days:   2,
				Reason: "参加朋友婚礼",
			},
			expected: true,
		},
		{
			name: "5天请假-项目经理审批",
			request: chain.LeaveRequest{
				Name:   "李四",
				Days:   5,
				Reason: "回家探亲",
			},
			expected: true,
		},
		{
			name: "10天请假-部门经理审批",
			request: chain.LeaveRequest{
				Name:   "王五",
				Days:   10,
				Reason: "出国旅游",
			},
			expected: true,
		},
		{
			name: "20天请假-CEO审批",
			request: chain.LeaveRequest{
				Name:   "赵六",
				Days:   20,
				Reason: "长期休养",
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 处理请求
			teamLeader.Handle(&tt.request)

			if tt.request.Approved != tt.expected {
				t.Errorf("期望审批结果为 %v，但得到 %v", tt.expected, tt.request.Approved)
			}
		})
	}
}

// 新增的测试函数：测试特殊情况
func TestLeaveApprovalChainSpecialCases(t *testing.T) {
	// 使用链式调用创建责任链
	teamLeader := chain.NewTeamLeader()
	teamLeader.SetNext(chain.NewProjectManager()).
		SetNext(chain.NewDepartmentManager()).
		SetNext(chain.NewCEO())

	tests := []struct {
		name     string
		request  chain.LeaveRequest
		expected bool
	}{
		{
			name: "0天请假-不合理请求",
			request: chain.LeaveRequest{
				Name:   "张三",
				Days:   0,
				Reason: "测试零天请假",
			},
			expected: true, // 组长会处理
		},
		{
			name: "3.5天请假-边界情况",
			request: chain.LeaveRequest{
				Name:   "李四",
				Days:   4,
				Reason: "测试边界情况",
			},
			expected: true, // 项目经理会处理
		},
		{
			name: "30天请假-超长假期",
			request: chain.LeaveRequest{
				Name:   "王五",
				Days:   30,
				Reason: "长期休假",
			},
			expected: true, // CEO会处理
		},
		{
			name: "7天请假-带病假条",
			request: chain.LeaveRequest{
				Name:   "赵六",
				Days:   7,
				Reason: "病假，已附医院证明",
			},
			expected: true, // 项目经理会处理
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重置审批状态
			tt.request.Approved = false

			// 处理请求
			teamLeader.Handle(&tt.request)

			if tt.request.Approved != tt.expected {
				t.Errorf("%s: 期望审批结果为 %v，但得到 %v",
					tt.name, tt.expected, tt.request.Approved)
			}
		})
	}
}
