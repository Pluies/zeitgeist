// Code generated by counterfeiter. DO NOT EDIT.
package gitlabfakes

import (
	"sync"

	gitlaba "github.com/xanzy/go-gitlab"
	"sigs.k8s.io/zeitgeist/pkg/gitlab"
)

type FakeClient struct {
	ListBranchesStub        func(string, string, *gitlaba.ListBranchesOptions) ([]*gitlaba.Branch, *gitlaba.Response, error)
	listBranchesMutex       sync.RWMutex
	listBranchesArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *gitlaba.ListBranchesOptions
	}
	listBranchesReturns struct {
		result1 []*gitlaba.Branch
		result2 *gitlaba.Response
		result3 error
	}
	listBranchesReturnsOnCall map[int]struct {
		result1 []*gitlaba.Branch
		result2 *gitlaba.Response
		result3 error
	}
	ListReleasesStub        func(string, string, *gitlaba.ListReleasesOptions) ([]*gitlaba.Release, *gitlaba.Response, error)
	listReleasesMutex       sync.RWMutex
	listReleasesArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *gitlaba.ListReleasesOptions
	}
	listReleasesReturns struct {
		result1 []*gitlaba.Release
		result2 *gitlaba.Response
		result3 error
	}
	listReleasesReturnsOnCall map[int]struct {
		result1 []*gitlaba.Release
		result2 *gitlaba.Response
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) ListBranches(arg1 string, arg2 string, arg3 *gitlaba.ListBranchesOptions) ([]*gitlaba.Branch, *gitlaba.Response, error) {
	fake.listBranchesMutex.Lock()
	ret, specificReturn := fake.listBranchesReturnsOnCall[len(fake.listBranchesArgsForCall)]
	fake.listBranchesArgsForCall = append(fake.listBranchesArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *gitlaba.ListBranchesOptions
	}{arg1, arg2, arg3})
	stub := fake.ListBranchesStub
	fakeReturns := fake.listBranchesReturns
	fake.recordInvocation("ListBranches", []interface{}{arg1, arg2, arg3})
	fake.listBranchesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeClient) ListBranchesCallCount() int {
	fake.listBranchesMutex.RLock()
	defer fake.listBranchesMutex.RUnlock()
	return len(fake.listBranchesArgsForCall)
}

func (fake *FakeClient) ListBranchesCalls(stub func(string, string, *gitlaba.ListBranchesOptions) ([]*gitlaba.Branch, *gitlaba.Response, error)) {
	fake.listBranchesMutex.Lock()
	defer fake.listBranchesMutex.Unlock()
	fake.ListBranchesStub = stub
}

func (fake *FakeClient) ListBranchesArgsForCall(i int) (string, string, *gitlaba.ListBranchesOptions) {
	fake.listBranchesMutex.RLock()
	defer fake.listBranchesMutex.RUnlock()
	argsForCall := fake.listBranchesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeClient) ListBranchesReturns(result1 []*gitlaba.Branch, result2 *gitlaba.Response, result3 error) {
	fake.listBranchesMutex.Lock()
	defer fake.listBranchesMutex.Unlock()
	fake.ListBranchesStub = nil
	fake.listBranchesReturns = struct {
		result1 []*gitlaba.Branch
		result2 *gitlaba.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) ListBranchesReturnsOnCall(i int, result1 []*gitlaba.Branch, result2 *gitlaba.Response, result3 error) {
	fake.listBranchesMutex.Lock()
	defer fake.listBranchesMutex.Unlock()
	fake.ListBranchesStub = nil
	if fake.listBranchesReturnsOnCall == nil {
		fake.listBranchesReturnsOnCall = make(map[int]struct {
			result1 []*gitlaba.Branch
			result2 *gitlaba.Response
			result3 error
		})
	}
	fake.listBranchesReturnsOnCall[i] = struct {
		result1 []*gitlaba.Branch
		result2 *gitlaba.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) ListReleases(arg1 string, arg2 string, arg3 *gitlaba.ListReleasesOptions) ([]*gitlaba.Release, *gitlaba.Response, error) {
	fake.listReleasesMutex.Lock()
	ret, specificReturn := fake.listReleasesReturnsOnCall[len(fake.listReleasesArgsForCall)]
	fake.listReleasesArgsForCall = append(fake.listReleasesArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *gitlaba.ListReleasesOptions
	}{arg1, arg2, arg3})
	stub := fake.ListReleasesStub
	fakeReturns := fake.listReleasesReturns
	fake.recordInvocation("ListReleases", []interface{}{arg1, arg2, arg3})
	fake.listReleasesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeClient) ListReleasesCallCount() int {
	fake.listReleasesMutex.RLock()
	defer fake.listReleasesMutex.RUnlock()
	return len(fake.listReleasesArgsForCall)
}

func (fake *FakeClient) ListReleasesCalls(stub func(string, string, *gitlaba.ListReleasesOptions) ([]*gitlaba.Release, *gitlaba.Response, error)) {
	fake.listReleasesMutex.Lock()
	defer fake.listReleasesMutex.Unlock()
	fake.ListReleasesStub = stub
}

func (fake *FakeClient) ListReleasesArgsForCall(i int) (string, string, *gitlaba.ListReleasesOptions) {
	fake.listReleasesMutex.RLock()
	defer fake.listReleasesMutex.RUnlock()
	argsForCall := fake.listReleasesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeClient) ListReleasesReturns(result1 []*gitlaba.Release, result2 *gitlaba.Response, result3 error) {
	fake.listReleasesMutex.Lock()
	defer fake.listReleasesMutex.Unlock()
	fake.ListReleasesStub = nil
	fake.listReleasesReturns = struct {
		result1 []*gitlaba.Release
		result2 *gitlaba.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) ListReleasesReturnsOnCall(i int, result1 []*gitlaba.Release, result2 *gitlaba.Response, result3 error) {
	fake.listReleasesMutex.Lock()
	defer fake.listReleasesMutex.Unlock()
	fake.ListReleasesStub = nil
	if fake.listReleasesReturnsOnCall == nil {
		fake.listReleasesReturnsOnCall = make(map[int]struct {
			result1 []*gitlaba.Release
			result2 *gitlaba.Response
			result3 error
		})
	}
	fake.listReleasesReturnsOnCall[i] = struct {
		result1 []*gitlaba.Release
		result2 *gitlaba.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listBranchesMutex.RLock()
	defer fake.listBranchesMutex.RUnlock()
	fake.listReleasesMutex.RLock()
	defer fake.listReleasesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ gitlab.Client = new(FakeClient)
