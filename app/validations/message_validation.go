package validations

type Message struct {
	AppId        string `json:"app_id" `         //binding:"required"
	AvatarUrl    string `json:"avatar_url" `     //binding:"required"
	Email        string `json:"email" `          //binding:"required"
	IsNewSession bool   `json:"is_new_session" ` //binding:"required"
	IsResolved   bool   `json:"is_resolved" `    //binding:"required"
	Name         string `json:"name" `           //binding:"required"
	RoomId       string `json:"room_id" `        //binding:"required"
	Source       string `json:"source" `         //binding:"required"
}
