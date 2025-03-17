package model

import (
	"time"
)

type Message struct {
	ID        uint      `gorm:"primaryKey"`
	UserId      uint    `gorm:"size:255"`
	Content   string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
/**
반정규화 형태로 구성 하여, jwt 토큰 을 파싱하여 해당 room 에' 대한 접근 권한 및 user_id 를 받아온다.
룸 넘버에 대한 값을 넘겨 받아야한다.
**/