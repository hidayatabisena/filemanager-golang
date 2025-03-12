package fileutils

// File categories
const (
	CAT_DOCUMENTS = "documents"
	CAT_IMAGES    = "images"
	CAT_AUDIO     = "audio"
	CAT_VIDEO     = "video"
	CAT_ARCHIVES  = "archives"
	CAT_CODE      = "code"
	CAT_DATA      = "data"
	CAT_OTHER     = "other"
)

// Extension to category mappings
var extensionCategories = map[string]string{
	// Documents
	".pdf":  CAT_DOCUMENTS,
	".doc":  CAT_DOCUMENTS,
	".docx": CAT_DOCUMENTS,
	".xls":  CAT_DOCUMENTS,
	".xlsx": CAT_DOCUMENTS,
	".ppt":  CAT_DOCUMENTS,
	".pptx": CAT_DOCUMENTS,
	".txt":  CAT_DOCUMENTS,
	".rtf":  CAT_DOCUMENTS,
	".odt":  CAT_DOCUMENTS,

	// Images
	".jpg":  CAT_IMAGES,
	".jpeg": CAT_IMAGES,
	".png":  CAT_IMAGES,
	".gif":  CAT_IMAGES,
	".bmp":  CAT_IMAGES,
	".svg":  CAT_IMAGES,
	".webp": CAT_IMAGES,
	".tiff": CAT_IMAGES,
	".psd":  CAT_IMAGES,

	// Audio
	".mp3":  CAT_AUDIO,
	".wav":  CAT_AUDIO,
	".flac": CAT_AUDIO,
	".aac":  CAT_AUDIO,
	".ogg":  CAT_AUDIO,
	".m4a":  CAT_AUDIO,

	// Video
	".mp4":  CAT_VIDEO,
	".avi":  CAT_VIDEO,
	".mkv":  CAT_VIDEO,
	".mov":  CAT_VIDEO,
	".wmv":  CAT_VIDEO,
	".flv":  CAT_VIDEO,
	".webm": CAT_VIDEO,

	// Archives
	".zip": CAT_ARCHIVES,
	".rar": CAT_ARCHIVES,
	".7z":  CAT_ARCHIVES,
	".tar": CAT_ARCHIVES,
	".gz":  CAT_ARCHIVES,
	".bz2": CAT_ARCHIVES,

	// Code
	".go":    CAT_CODE,
	".py":    CAT_CODE,
	".js":    CAT_CODE,
	".html":  CAT_CODE,
	".css":   CAT_CODE,
	".c":     CAT_CODE,
	".cpp":   CAT_CODE,
	".java":  CAT_CODE,
	".php":   CAT_CODE,
	".rb":    CAT_CODE,
	".rs":    CAT_CODE,
	".swift": CAT_CODE,

	// Data
	".json": CAT_DATA,
	".xml":  CAT_DATA,
	".csv":  CAT_DATA,
	".sql":  CAT_DATA,
	".db":   CAT_DATA,
	".yaml": CAT_DATA,
	".toml": CAT_DATA,
}

func GetFileCategory(extension string) string {
	if category, ok := extensionCategories[extension]; ok {
		return category
	}
	return CAT_OTHER
}
