package structure
type Messageboard struct {
AuthorId int32/*numeric*/ `json:"author_id" xml:"author_id" ` // Message author ID

CreatedAt int64/*timestamp*/ `json:"created_at" xml:"created_at" ` // Message creation date

EditorId int32/*numeric*/ `json:"editor_id" xml:"editor_id" ` // ID of a player who has changed the message

IsRead bool/*boolean*/ `json:"is_read" xml:"is_read" ` // Indicates if the message has been read

Message string/*string*/ `json:"message" xml:"message" ` // Message text

UpdatedAt int64/*timestamp*/ `json:"updated_at" xml:"updated_at" ` // Date when message was updated

}