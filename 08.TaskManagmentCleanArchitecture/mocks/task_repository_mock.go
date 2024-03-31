package mocks

import (
	"TaskManger/domain"
	"context"

	"github.com/stretchr/testify/mock"
)

type TaskRepositoryMock struct {
    mock.Mock
}


func NewTaskRepositoryMock() domain.TaskRepository {
    return &TaskRepositoryMock{}
}



func (m *TaskRepositoryMock) Create(c context.Context, task *domain.Task) (*domain.Task, error) {
    ret := m.Called(c, task)
    var r0 *domain.Task
    if rf, ok := ret.Get(0).(func(context.Context, *domain.Task) *domain.Task); ok {
        r0 = rf(c, task)
    } else {
        if ret.Get(0) != nil {
            r0 = ret.Get(0).(*domain.Task)
        }
    }

    var r1 error
    if rf, ok := ret.Get(1).(func(context.Context, *domain.Task) error); ok {
        r1 = rf(c, task)
    } else {
        r1 = ret.Error(1)
    }
    return r0, r1
}


func (m *TaskRepositoryMock) Delete(c context.Context, task *domain.Task) (*domain.Task, error) {
    ret := m.Called(c, task)
    var r0 *domain.Task
    if rf, ok := ret.Get(0).(func(context.Context, *domain.Task) *domain.Task); ok {
        r0 = rf(c, task)
    } else {
        if ret.Get(0) != nil {
            r0 = ret.Get(0).(*domain.Task)
        }
    }

    var r1 error
    if rf, ok := ret.Get(1).(func(context.Context, *domain.Task) error); ok {
        r1 = rf(c, task)
    } else {
        r1 = ret.Error(1)
    }
    return r0, r1
}


func (m *TaskRepositoryMock) GetById(c context.Context, taskId string) (*domain.Task, error) {
    ret := m.Called(c, taskId)
    var r0 *domain.Task
    if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Task); ok {
        r0 = rf(c, taskId)
    } else {
        if ret.Get(0) != nil {
            r0 = ret.Get(0).(*domain.Task)
        }
    }

    var r1 error
    if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
        r1 = rf(c, taskId)
    } else {
        r1 = ret.Error(1)
    }
    return r0, r1
}


func (m *TaskRepositoryMock) GetByUserId(c context.Context, userId string) (*[]*domain.Task, error) {
    ret := m.Called(c, userId)
    var r0 *[]*domain.Task
    if rf, ok := ret.Get(0).(func(context.Context, string) *[]*domain.Task); ok {
        r0 = rf(c, userId)
    } else {
        if ret.Get(0) != nil {
            r0 = ret.Get(0).(*[]*domain.Task)
        }
    }

    var r1 error
    if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
        r1 = rf(c, userId)
    } else {
        r1 = ret.Error(1)
    }
    return r0, r1
}

func (m *TaskRepositoryMock) Update(c context.Context, task *domain.Task) (*domain.Task, error) {
    ret := m.Called(c, task)
    var r0 *domain.Task
    if rf, ok := ret.Get(0).(func(context.Context, *domain.Task) *domain.Task); ok {
        r0 = rf(c, task)
    } else {
        if ret.Get(0) != nil {
            r0 = ret.Get(0).(*domain.Task)
        }
    }

    var r1 error
    if rf, ok := ret.Get(1).(func(context.Context, *domain.Task) error); ok {
        r1 = rf(c, task)
    } else {
        r1 = ret.Error(1)
    }
    return r0, r1
}