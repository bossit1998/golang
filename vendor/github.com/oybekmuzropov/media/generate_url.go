package media

import "fmt"

const URL  = "https://s3.amazonaws.com"
const BUCKET  = "fiesta-storage-resized"
const PROMO_VIDEO_PATH  = "videos/original"
const RESIZED_VIDEO_PATH  = "videos/720p"
const THUMBNAIL_480x858_PATH  = "thumbnails/480x858"
const THUMBNAIL_720x1280_PATH  = "thumbnails/720x1280"
const FLYER_ORIGINAL_PATH  = "flyers/original"
const FLYER_500x500_PATH  = "flyers/500x500"
const AVATAR_300x300_PATH  = "avatars/300x300"
const AVATAR_500x500_PATH  = "avatars/500x500"

func GetPromoVideoUrl(filename string) string {
	return fmt.Sprintf("%s/%s/%s/%s.mp4", URL, BUCKET, PROMO_VIDEO_PATH, filename)
}

func GetResizedVideoUrl(filename string) string {
	return fmt.Sprintf("%s/%s/%s/%s.mp4", URL, BUCKET, RESIZED_VIDEO_PATH, filename)
}

func GetThumbnail480x858Url(filename string) string {
	return fmt.Sprintf("%s/%s/%s/%s", URL, BUCKET, THUMBNAIL_480x858_PATH, filename)
}

func GetThumbnail720x1280Url(filename string) string {
	return fmt.Sprintf("%s/%s/%s/%s", URL, BUCKET, THUMBNAIL_720x1280_PATH, filename)
}

func GetFlyerOriginalUrl(filename string) string {
	return fmt.Sprintf("%s/%s/%s/%s", URL, BUCKET, FLYER_ORIGINAL_PATH, filename)
}

func GetFlyer500x500Url(filename string) string {
	return fmt.Sprintf("%s/%s/%s/%s", URL, BUCKET, FLYER_500x500_PATH, filename)
}

func GetAvatars300x300Url(filename string) string {
	return fmt.Sprintf("%s/%s/%s/%s", URL, BUCKET, AVATAR_300x300_PATH, filename)
}

func GetAvatars500x500Url(filename string) string {
	return fmt.Sprintf("%s/%s/%s/%s", URL, BUCKET, AVATAR_500x500_PATH, filename)
}