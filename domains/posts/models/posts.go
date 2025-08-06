package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Title   string
	Content string
}

type PostResponse struct {
	ID          uuid.UUID `pg:"type:uuid,default:uuid_generate_v4(),pk"`
	Title       string    `pg:"type:varchar(50)"`
	Content     string    `pg:"type:varchar(1200)"`
	PublishDate time.Time `pg:"type:date,default:CURRENT_DATE"`
}
