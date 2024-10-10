package api

// GeoPoint 地理位置结构
type GeoPoint struct {
	Type        string    `json:"type"`        // 类型
	Coordinates []float64 `json:"coordinates"` // 坐标点
}

type Gender int

const (
	GenderMale   Gender = 1
	GenderFemale Gender = 2
	GenderSecret Gender = 3
)

type UserData struct {
	Name        string    `v:"required"          json:"name"`          // 姓名
	Age         int       `v:"required|min:1"    json:"age"`           // 年龄
	Gender      Gender    `v:"required|in:1,2,3" json:"gender"`        // 性别
	Location    *GeoPoint `v:"required"          json:"location"`      // GPS位置
	MatchGender int       `v:"required|min:1"    json:"match_gender"`  // 性取向
	MatchMinAge int       `v:"required|min:1"    json:"match_min_age"` // 年龄偏好范围(min)
	MatchMaxAge int       `v:"required|min:1"    json:"match_max_age"` // 年龄偏好范围(max)
}

type CreateRequest struct {
	User *UserData `v:"required" json:"user"` // 用户数据结构
}
type CreateResponse struct{}

type SearchRequest struct {
	MinAge      int    `v:"min:1"    json:"min_age"`       // 年龄范围(min)
	MaxAge      int    `v:"min:1"    json:"max_age"`       // 年龄范围(max)
	Gender      Gender `v:"in:1,2,3" json:"gender"`        // 性别
	MatchGender int    `v:"min:1"    json:"match_gender"`  // 性取向
	MatchMinAge int    `v:"min:1"    json:"match_min_age"` // 年龄偏好(min)
	MatchMaxAge int    `v:"min:1"    json:"match_max_age"` // 年龄偏好(max)
}

type SearchResponse struct {
	Users []*UserData `json:"users"` // 查询匹配到的用户列表
}
