package channelsdto

type EditChannelRequest struct {
  Channelname     string `json:"channelName" form:"channelName"`
  Email    string `json:"email" form:"email"`
  Password string `json:"password" form:"password"`
  Image string `json:"image" form:"image"`

}