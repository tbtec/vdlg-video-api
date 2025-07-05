package presenter

import (
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/dto"
)

type VideoPresenter struct {
}

func NewVideoPresenter() *VideoPresenter {
	return &VideoPresenter{}
}

func (presenter *VideoPresenter) BuildVideoContentResponse(videos []entity.Video) dto.VideoContent {
	response := []dto.Video{}

	for _, video := range videos {
		response = append(
			response,
			presenter.BuildVideoResponse(video),
		)
	}

	return dto.VideoContent{Content: response}
}

func (presenter *VideoPresenter) BuildVideoResponse(video entity.Video) dto.Video {
	response := dto.Video{
		ID:             video.ID,
		CustomerId:     video.CustomerId,
		Status:         string(video.Status),
		FileNameInput:  video.FileNameInput,
		FileNameOutput: video.FileNameOutput,
		CreatedAt:      video.CreatedAt,
		UpdatedAt:      video.UpdatedAt,
	}
	if video.UploadUrl != nil {
		response.UploadUrl = video.UploadUrl
	}
	if video.DownloadUrl != nil {
		response.DownloadUrl = video.DownloadUrl
	}

	return response
}
