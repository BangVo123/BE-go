package presenter

type FollowReq struct {
	FollowingId string `json:"followingId" validate:"require"`
}
