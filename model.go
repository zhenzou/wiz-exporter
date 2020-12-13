package wiz

const (
	DocumentTableName    = "wiz_document"
	TagTableName         = "wiz_tag"
	DocumentTagTableName = "wiz_document_tag"
)

type documentEntity struct {
	Guid            string `json:"guid" gorm:"column:DOCUMENT_GUID"`
	Title           string `json:"title" gorm:"column:DOCUMENT_TITLE"`
	Location        string `json:"location" gorm:"column:DOCUMENT_LOCATION"`
	Name            string `json:"name" gorm:"column:DOCUMENT_NAME"`
	SEO             string `json:"seo" gorm:"column:DOCUMENT_SEO"`
	URL             string `json:"url" gorm:"column:DOCUMENT_URL"`
	Author          string `json:"author" gorm:"column:DOCUMENT_AUTHOR"`
	Keywords        string `json:"keywords" gorm:"column:DOCUMENT_KEYWORDS"`
	Type            string `json:"type" gorm:"column:DOCUMENT_TYPE"`
	Owner           string `json:"owner" gorm:"column:DOCUMENT_OWNER"`
	FileType        string `json:"file_type" gorm:"column:DOCUMENT_FILE_TYPE"`
	StyleGuid       string `json:"style_guid" gorm:"column:STYLE_GUID"`
	IconIndex       int    `json:"icon_index"gorm:"column:DOCUMENT_ICON_INDEX"`
	Sync            int    `json:"sync" gorm:"column:DOCUMENT_SYNC"`
	Protect         int    `json:"protect" gorm:"column:DOCUMENT_PROTECT"`
	ReadCount       int    `json:"read_count" gorm:"column:DOCUMENT_READ_COUNT"`
	AttachmentCount int    `json:"attachment_count" gorm:"column:DOCUMENT_ATTACHMENT_COUNT"`
	Indexed         int    `json:"indexed" gorm:"column:DOCUMENT_INDEXED"`
	DTInfoModified  string `json:"dt_info_modified" gorm:"column:DT_INFO_MODIFIED"`
	InfoMd5         string `json:"info_md5" gorm:"column:DOCUMENT_INFO_MD5"`
	DTDataModified  string `json:"dt_data_modified" gorm:"column:DT_DATA_MODIFIED"`
	DataMd5         string `json:"data_md5"gorm:"column:DOCUMENT_DATA_MD5"`
	DTParamModified string `json:"dt_param_modified" gorm:"column:DT_PARAM_MODIFIED"`
	ParamModified   string `json:"param_modified" gorm:"column:DOCUMENT_PARAM_MODIFIED"`
	WizVersion      int64  `json:"wiz_version" gorm:"column:WIZ_VERSION"`
	InfoChanged     int    `json:"info_changed" gorm:"column:INFO_CHANGED"`
	DataChanged     int    `json:"data_changed" gorm:"column:DATA_CHANGED"`
	DTCreated       string `json:"dt_created" gorm:"column:DT_CREATED"`
	DTModified      string `json:"dt_modified" gorm:"column:DT_MODIFIED"`
	DTAccessed      string `json:"dt_accessed" gorm:"column:DT_ACCESSED"`
}

func (d *documentEntity) TableName() string {
	return DocumentTableName
}

type tagEntity struct {
	Guid        string `json:"guid" gorm:"column:TAG_GUID"`
	GroupGuid   string `json:"group_guid" gorm:"column:TAG_GROUP_GUID"`
	Name        string `json:"name" gorm:"column:TAG_NAME"`
	Pos         int    `json:"pos" gorm:"column:tag_pos"`
	Description string `json:"description" gorm:"column:DESCRIPTION"`
	DTModified  string `json:"dt_modified" gorm:"column:DT_MODIFIED"`
	WizVersion  int64  `json:"wiz_version" gorm:"column:WIZ_VERSION"`
}

func (t *tagEntity) TableName() string {
	return TagTableName
}

type documentTagEntity struct {
	DocumentGuid string `json:"document_guid" gorm:"column:document_guid"`
	TagGuid      string `json:"tag_guid" gorm:"column:tag_guid"`
}

func (t *documentTagEntity) TableName() string {
	return DocumentTagTableName
}
