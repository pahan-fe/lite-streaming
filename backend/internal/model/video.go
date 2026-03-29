package model

import "time"

type Video struct {
	ID 					string 		`json:"id" db:"id"`
	OriginalFilename 	string 		`json:"originalFilename" db:"original_filename"`
	ContentType 		string 		`json:"contentType" db:"content_type"`
	Size 				int64 		`json:"size" db:"size"`
	Status 				string 		`json:"status" db:"status"`
	S3RawKey 			string 		`json:"s3RawKey" db:"s3_raw_key"`
	S3HLSKey 			string 		`json:"s3HLSKey" db:"s3_hls_key"`
	CreatedAt 			time.Time 	`json:"createdAt" db:"created_at"`
	UpdatedAt 			time.Time 	`json:"updatedAt" db:"updated_at"`
}