package trackerstore

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"job4j.ru/go-lang-base/internal/tracker"
)

type Store interface {
	Create(ctx context.Context, item Item) error
	List(ctx context.Context) ([]Item, error)
	Get(ctx context.Context, id string) (Item, error)
	GetByPosition(ctx context.Context, position int) (Item, error)
	GetLastPosition(ctx context.Context) (int, error)
	Find(ctx context.Context, name string) ([]Item, error)
	Update(ctx context.Context, position int, item Item) error
	Delete(ctx context.Context, position int) error
	Reorder(ctx context.Context, position int) error
}

type UseCase interface {
	Done(ctx context.Context, in tracker.Input, out tracker.Output, store Store) error
	Desc() string
}

type AddUseCase struct{}

func (u AddUseCase) Done(
	ctx context.Context,
	in tracker.Input,
	out tracker.Output,
	store Store,
) error {
	out.Out("enter name:")
	name := in.Get()
	id := uuid.New().String()

	lastPosition, err := store.GetLastPosition(ctx)
	if err != nil {
		return ErrGetLastPosition(err)
	}
	if err := store.Create(
		ctx,
		Item{ID: id, Name: name, Position: lastPosition + 1},
	); err != nil {
		return ErrCreate(err)
	}
	return nil
}

func (u AddUseCase) Desc() string {
	return "Add Element"
}

type GetUseCase struct{}

func (u GetUseCase) Done(
	ctx context.Context,
	_ tracker.Input,
	out tracker.Output,
	store Store,
) error {
	items, err := store.List(ctx)
	if err != nil {
		return ErrGet(err)
	}
	printItems(items, out)
	return nil
}

func (u GetUseCase) Desc() string {
	return "Get all Elements"
}

type UpdateUseCase struct{}

func (u UpdateUseCase) Done(
	ctx context.Context,
	in tracker.Input,
	out tracker.Output,
	store Store,
) error {
	for {
		out.Out("enter number position for update:")
		number := in.Get()
		position, err := strconv.Atoi(number)
		if err != nil {
			out.Out("you entered an invalid number")
			continue
		}
		item, err := store.GetByPosition(ctx, position)
		if err != nil {
			out.Out(err.Error())
			return nil
		}
		out.Out(fmt.Sprintf("enter new name for position %d:", position))
		newName := in.Get()
		item.Name = newName
		if err = store.Update(ctx, position, item); err != nil {
			return ErrUpdate(err)
		}
		out.Out(fmt.Sprintf("element on position %d updated", position))
		return nil
	}
}

func (u UpdateUseCase) Desc() string {
	return "Update Element"
}

type DeleteUseCase struct{}

func (u DeleteUseCase) Done(
	ctx context.Context,
	in tracker.Input,
	out tracker.Output,
	store Store,
) error {
	for {
		out.Out("enter element position for delete:")
		number := in.Get()
		position, err := strconv.Atoi(number)
		if err != nil {
			out.Out("you entered an invalid number")
			continue
		}
		if err = store.Delete(ctx, position); err != nil {
			return ErrDelete(err)
		}
		out.Out(fmt.Sprintf("element on position %d deleted", position))
		if err = store.Reorder(ctx, position); err != nil {
			return ErrReorderAfterDelete(err)
		}
		return nil
	}
}

func (u DeleteUseCase) Desc() string {
	return "Delete Elements"
}

type FindUseCase struct{}

func (u FindUseCase) Done(
	ctx context.Context,
	in tracker.Input,
	out tracker.Output,
	store Store,
) error {
	out.Out("enter name fo find:")
	name := in.Get()
	find, err := store.Find(ctx, name)
	if err != nil {
		return ErrFind(err)
	}
	out.Out("found elements:")
	printItems(find, out)
	return nil
}

func (u FindUseCase) Desc() string {
	return "Find Elements"
}

func printItems(items []Item, out tracker.Output) {
	for _, item := range items {
		out.Out(fmt.Sprintf("%d) %s\t%s", item.Position, item.ID, item.Name))
	}
}
