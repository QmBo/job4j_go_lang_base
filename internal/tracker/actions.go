package tracker

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

type UseCase interface {
	Done(in Input, out Output, tracker *Tracker)
	Desc() string
}

type AddUseCase struct{}

func (u AddUseCase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter name:")
	name := in.Get()
	id := uuid.New().String()
	_, err := tracker.AddItem(Item{Name: name, ID: id})
	if err != nil {
		out.Out("failed to add item")
	}
}
func (u AddUseCase) Desc() string {
	return "Add Element"
}

type GetUseCase struct{}

func (u GetUseCase) Done(_ Input, out Output, tracker *Tracker) {
	items := tracker.GetItems()
	if len(items) == 0 {
		out.Out("no elements")
		return
	}
	out.Out("all elements:")
	printItems(items, out)
}

func (u GetUseCase) Desc() string {
	return "Get all Elements"
}

type FindUseCase struct{}

func (u FindUseCase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter name fo find:")
	name := in.Get()
	find, err := tracker.Find(name)
	if err != nil {
		out.Out(err.Error())
		return
	}
	out.Out("found elements:")
	printItems(find, out)
}

func (u FindUseCase) Desc() string {
	return "Find Elements"
}

type DeleteUseCase struct{}

func (u DeleteUseCase) Done(in Input, out Output, tracker *Tracker) {
	for {
		out.Out("enter element number for delete:")
		position := in.Get()
		number, err := strconv.Atoi(position)
		if err != nil {
			out.Out("you entered an invalid number")
			continue
		}
		item, err := tracker.RemoveItem(number)
		if err != nil {
			out.Out(err.Error())
			return
		}
		out.Out(fmt.Sprintf("deleted item: %s", item.toString()))
		return
	}
}

func (u DeleteUseCase) Desc() string {
	return "Delete Element"
}

type UpdateUseCase struct{}

func (u UpdateUseCase) Done(in Input, out Output, tracker *Tracker) {
	for {
		out.Out("enter number position for update:")
		number := in.Get()
		position, err := strconv.Atoi(number)
		if err != nil {
			out.Out("you entered an invalid number")
			continue
		}
		item, err := tracker.GetByPosition(position)
		if err != nil {
			out.Out(err.Error())
			return
		}
		out.Out(fmt.Sprintf("enter new name for update: %s", item.toString()))
		newName := in.Get()
		item.Name = newName
		tracker.Update(position-1, item)
		out.Out(fmt.Sprintf("updated item: %s", item.toString()))
		return
	}
}

func (u UpdateUseCase) Desc() string {
	return "Update Element"
}

func printItems(items []Item, out Output) {
	for _, item := range items {
		out.Out(item.toString())
	}
}
