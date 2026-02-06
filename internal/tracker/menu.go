package tracker

import (
	"fmt"
	"sort"
)

type UI struct {
	In      Input
	Out     Output
	Tracker *Tracker
}

func (u UI) Run() {
	actions := initUseCase()
	keys := sortKeys(actions)

	for {
		u.printHeader(keys, actions)
		selected := u.In.Get()

		if selected == "exit" {
			break
		}

		action, ok := actions[selected]
		if !ok {
			u.Out.Out("not found action")
			continue
		}

		action.Done(u.In, u.Out, u.Tracker)
	}
}

func sortKeys(actions map[string]UseCase) []string {
	keys := make([]string, 0, len(actions))
	for k := range actions {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func initUseCase() map[string]UseCase {
	actions := map[string]UseCase{
		"add":  AddUseCase{},
		"get":  GetUseCase{},
		"find": FindUseCase{},
		"del":  DeleteUseCase{},
		"updt": UpdateUseCase{},
	}
	return actions
}

func (u UI) printHeader(keys []string, actions map[string]UseCase) {
	u.Out.Out("\n\rSelect action")
	for _, k := range keys {
		u.Out.Out(fmt.Sprintf("\"%s\" for %s", k, actions[k].Desc()))
	}
	u.Out.Out("\"exit\" for exit")
}
