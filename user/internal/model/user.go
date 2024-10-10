package model

type User struct {
	Name        string    `json:"name"          bson:"name"`          // 姓名
	Age         int       `json:"age"           bson:"age"`           // 年龄
	Gender      int       `json:"gender"        bson:"gender"`        // 性别
	Location    *GeoPoint `json:"location"      bson:"location"`      // GPS位置
	MatchGender int       `json:"match_gender"  bson:"match_gender"`  // 性取向
	MatchMinAge int       `json:"match_min_age" bson:"match_min_age"` // 年龄偏好范围(min)
	MatchMaxAge int       `json:"match_max_age" bson:"match_max_age"` // 年龄偏好范围(max)
}

// GeoPoint 地理位置结构
type GeoPoint struct {
	Type        string    `json:"type"        bson:"type"`        // 类型
	Coordinates []float64 `json:"coordinates" bson:"coordinates"` // 坐标点
}
