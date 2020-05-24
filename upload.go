func UploadVideo(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]string)

	file, video, err := c.Request.FormFile("video")
	if err != nil {
		logging.Warn(err)
		code = e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}

	if video == nil {
		code = e.INVALID_PARAMS
	} else {
		videoName := upload.GetVideoName(image.Filename)
		fullPath := upload.GetVideoFullPath()
		savePath := upload.GetVideoPath()

		src := fullPath + videoName
		if ! upload.CheckVideoExt(videoName) || ! upload.CheckVideoSize(file) {
			code = e.ERROR_UPLOAD_CHECK_Video_FORMAT
		} else {
			err := upload.CheckVideo(fullPath)
			if err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_CHECK_Video_FAIL
			} else if err := c.SaveUploadedFile(video, src); err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_SAVE_Video_FAIL
			} else {
				data["video_url"] = upload.GetVideoFullUrl(imageName)
				data["video_save_url"] = savePath + videoName
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

