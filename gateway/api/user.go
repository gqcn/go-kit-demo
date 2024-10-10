package api

// GeoPoint 地理位置结构
type GeoPoint struct {
	Type        string    `json:"type"`        // 类型
	Coordinates []float64 `json:"coordinates"` // 坐标点
}

type UserData struct {
	Name        string    `json:"name"`          // 姓名
	Age         int       `json:"age"`           // 年龄
	Gender      int       `json:"gender"`        // 性别
	Location    *GeoPoint `json:"location"`      // GPS位置
	MatchGender int       `json:"match_gender"`  // 性取向
	MatchMinAge int       `json:"match_min_age"` // 年龄偏好范围(min)
	MatchMaxAge int       `json:"match_max_age"` // 年龄偏好范围(max)
}

type CreateRequest struct {
	User *UserData `json:"user"` // 用户数据结构
}
type CreateResponse struct{}

type SearchRequest struct {
	MinAge      int `json:"min_age"`       // 年龄范围(min)
	MaxAge      int `json:"max_age"`       // 年龄范围(max)
	Gender      int `json:"gender"`        // 性别
	MatchGender int `json:"match_gender"`  // 性取向
	MatchMinAge int `json:"match_min_age"` // 年龄偏好(min)
	MatchMaxAge int `json:"match_max_age"` // 年龄偏好(max)
}

type SearchResponse struct {
	Users []*UserData `json:"users"` // 查询匹配到的用户列表
}
