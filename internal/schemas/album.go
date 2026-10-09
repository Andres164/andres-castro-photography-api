package schemas

type AlbumSummaryResponse struct  {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	CoverImageURI string `json:"coverImageUri"`
}

type AlbumDetailedResponse struct  {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	CoverImageURI string `json:"coverImageUri"`
	Photos []GetPhotoResponse `json:"photos"`
}

type GetAlbumOutput struct {
	Body AlbumDetailedResponse
}

type GetAlbumsOutput struct {
	Body []AlbumSummaryResponse
}