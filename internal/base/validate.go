package base

type ValidateRequest struct {
	UserID      string
	Title       string
	Description string
}

func Validate(req *ValidateRequest) []string {
	res := make([]string, 0)
	if req == nil {
		return append(res, "nil is passed")
	}
	if req.UserID == "" {
		res = append(res, "UserID is empty")
	}
	if req.Title == "" {
		res = append(res, "Title is empty")
	}
	if req.Description == "" {
		res = append(res, "Description is empty")
	}
	return res
}
