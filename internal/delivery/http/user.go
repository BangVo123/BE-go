package http

import (
	"net/http"
	"project/internal/handlers"
	"project/internal/models"
	"project/internal/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	UserCase         usecase.UserCase
	FollowUseCase    usecase.FollowUseCase
	FollowerUseCase  usecase.FollowerInfoUseCase
	FollowingUseCase usecase.FollowingInfoUseCase
	LikeUseCase      usecase.LikeUseCase
	LoveUseCase      usecase.LoveUseCase
}

func NewUserHandler(userCase usecase.UserCase, followUseCase usecase.FollowUseCase, followerUseCase usecase.FollowerInfoUseCase, followingUseCase usecase.FollowingInfoUseCase, likeUseCase usecase.LikeUseCase, loveUseCase usecase.LoveUseCase) handlers.UserHandler {
	return &UserHandler{UserCase: userCase, FollowUseCase: followUseCase, FollowerUseCase: followerUseCase, FollowingUseCase: followingUseCase, LikeUseCase: likeUseCase, LoveUseCase: loveUseCase}
}

func (uh *UserHandler) GetMe(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusForbidden, "User not authenticate")
		return
	}

	foundUser := user.(models.User)

	//combine other service into 1 endpoint
	allFollow, err := uh.FollowUseCase.GetAll(c.Request.Context(), foundUser.Id.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Something went wrong when get follow")
		return
	}

	var followers *[]models.FollowerInfo
	var following *[]models.FollowingInfo

	followers, err = uh.FollowerUseCase.GetFollower(c.Request.Context(), foundUser.Id.Hex())
	if err != nil && err != mongo.ErrNilCursor {
		c.JSON(http.StatusInternalServerError, map[string]any{"message": "Something went wrong when get followers", "err": err})
		return
	}

	following, err = uh.FollowingUseCase.GetFollowing(c.Request.Context(), foundUser.Id.Hex())
	if err != nil && err != mongo.ErrNoDocuments {
		c.JSON(http.StatusInternalServerError, "Something went wrong when get followings")
		return
	}

	var friends []models.FollowerInfo

	for _, val1 := range *followers {
		for _, val2 := range *following {
			if val1.Id == val2.Id {
				friends = append(friends, val1)
			}
		}
	}

	likes, err := uh.LikeUseCase.GetLike(c.Request.Context(), foundUser.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Something went wrong when get likes")
	}

	loves, err := uh.LoveUseCase.GetLove(c.Request.Context(), foundUser.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Something went wrong when get loves")
	}

	resData := gin.H{
		"data": gin.H{
			"user":   foundUser,
			"follow": allFollow,
			"favorite": gin.H{
				"likes": likes,
				"loves": loves,
			},
			"accRelations": gin.H{
				"followers": func() any {
					if followers != nil && len(*followers) > 0 {
						// return (*followers)[0].UserId
						return *followers
					}
					return nil
				}(),
				"following": func() any {
					if following != nil && len(*following) > 0 {
						return *following
					}
					return nil
				}(),
				"friends": friends,
			}},
	}

	c.JSON(http.StatusOK, resData)
}

func (uh *UserHandler) GetUserInfo(c *gin.Context) {
	userIdString := c.Param("userId")

	FoundUser, err := uh.UserCase.GetUserById(c.Request.Context(), userIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Something went wrong when get userinfo")
	}

	c.JSON(http.StatusOK, map[string]any{"data": FoundUser})
}

func (uh *UserHandler) UpdateMe(c *gin.Context) {
	var payload map[string]any

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if len(payload) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No data update"})
		return
	}

	user, _ := c.Get("user")
	userObj, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	err := uh.UserCase.UpdateUser(c.Request.Context(), userObj.Id.Hex(), payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "Success"})
}
