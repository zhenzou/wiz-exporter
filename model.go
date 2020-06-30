package wiz

const (
	DocumentTableName    = "wiz_document"
	TagTableName         = "wiz_tag"
	DocumentTagTableName = "wiz_document_tag"
)

type Document struct {
	Guid            string `json:"guid" db:"DOCUMENT_GUID"`
	Title           string `json:"title" db:"DOCUMENT_TITLE"`
	Location        string `json:"location" db:"DOCUMENT_LOCATION"`
	Name            string `json:"name" db:"DOCUMENT_NAME"`
	SEO             string `json:"seo" db:"DOCUMENT_SEO"`
	URL             string `json:"url" db:"DOCUMENT_URL"`
	Author          string `json:"author" db:"DOCUMENT_AUTHOR"`
	Keywords        string `json:"keywords" db:"DOCUMENT_KEYWORDS"`
	Owner           string `json:"owner" db:"DOCUMENT_OWNER"`
	StyleGuid       string `json:"style_guid" db:"STYLE_GUID"`
	IconIndex       int    `json:"icon_index"db:"DOCUMENT_ICON_INDEX"`
	Sync            int    `json:"sync" db:"DOCUMENT_SYNC"`
	Protect         int    `json:"protect" db:"DOCUMENT_PROTECT"`
	ReadCount       int    `json:"read_count" db:"DOCUMENT_READ_COUNT"`
	Indexed         int    `json:"indexed" db:"DOCUMENT_INDEXED"`
	DTInfoModified  string `json:"dt_info_modified" db:"DT_INFO_MODIFIED"`
	InfoMd5         string `json:"info_md5" db:"DOCUMENT_INFO_MD5"`
	DTDataModified  string `json:"dt_data_modified" db:"DT_DATA_MODIFIED"`
	DataMd5         string `json:"data_md5"db:"DOCUMENT_DATA_MD5"`
	DTParamModified string `json:"dt_param_modified" db:"DT_PARAM_MODIFIED"`
	WizVersion      int64  `json:"wiz_version" db:"WIZ_VERSION"`
	InfoChanged     int    `json:"info_changed" db:"INFO_CHANGED"`
	DataChanged     int    `json:"data_changed" db:"DATA_CHANGED"`
	DTCreated       string `json:"dt_created" db:"DT_CREATED"`
	DTModified      string `json:"dt_modified" db:"DT_MODIFIED"`
	DTAccessed      string `json:"dt_accessed" db:"DT_ACCESSED"`
}

func (d *Document) TableName() string {
	return DocumentTableName
}

type Tag struct {
	Guid        string `json:"guid" db:"TAG_GUID"`
	GroupGuid   string `json:"group_guid" db:"TAG_GROUP_GUID"`
	Name        string `json:"name" db:"TAG_NAME"`
	Pos         int    `json:"pos" db:"tag_pos"`
	Description string `json:"description" db:"DESCRIPTION"`
	DTModified  string `json:"dt_modified" db:"DT_MODIFIED"`
	WizVersion  int64  `json:"wiz_version" db:"WIZ_VERSION"`
}

func (t *Tag) TableName() string {
	return TagTableName
}

type DocumentTag struct {
	DocumentGuid string `json:"document_guid" db:"document_guid"`
	TagGuid      string `json:"tag_guid" db:"tag_guid"`
}

func (t *DocumentTag) TableName() string {
	return DocumentTagTableName
}
