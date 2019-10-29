package wiz

const (
	DocumentTableName    = "wiz_document"
	TagTableName         = "wiz_tag"
	DocumentTagTableName = "wiz_document_tag"
)

type Document struct {
	Guid            string `json:"guid" gorm:"column:document_guid"`
	Title           string `json:"title" gorm:"column:document_title"`
	Location        string `json:"location" gorm:"column:document_location"`
	Name            string `json:"name" gorm:"column:document_name"`
	SEO             string `json:"seo" gorm:"column:document_seo"`
	URL             string `json:"url" gorm:"column:document_url"`
	Author          string `json:"author" gorm:"column:document_author"`
	Keywords        string `json:"keywords" gorm:"column:document_keywords"`
	Type            string `json:"type" gorm:"column:document_type"`
	Owner           string `json:"owner" gorm:"column:document_owner"`
	FileType        string `json:"file_type" gorm:"column:document_file_type"`
	StyleGuid       string `json:"style_guid" gorm:"column:style_guid"`
	IconIndex       int    `json:"icon_index"gorm:"column:document_icon_index"`
	Sync            int    `json:"sync" gorm:"column:document_sync"`
	Protect         int    `json:"protect" gorm:"column:document_protect"`
	ReadCount       int    `json:"read_count" gorm:"column:document_read_count"`
	AttachmentCount int    `json:"attachment_count" gorm:"column:document_attachment_count"`
	Indexed         int    `json:"indexed" gorm:"column:document_indexed"`
	DTInfoModified  string `json:"dt_info_modified" gorm:"column:dt_info_modified"`
	InfoMd5         string `json:"info_md5" gorm:"column:document_info_md5"`
	DTDataModified  string `json:"dt_data_modified" gorm:"column:dt_data_modified"`
	DataMd5         string `json:"data_md5"gorm:"column:document_data_md5"`
	DTParamModified string `json:"dt_param_modified" gorm:"column:dt_param_modified"`
	ParamModified   string `json:"param_modified" gorm:"column:document_param_modified"`
	WizVersion      int64  `json:"wiz_version" gorm:"column:wiz_version"`
	InfoChanged     int    `json:"info_changed" gorm:"column:info_changed"`
	DataChanged     int    `json:"data_changed" gorm:"column:data_changed"`
	DTCreated       string `json:"dt_created" gorm:"column:dt_created"`
	DTModified      string `json:"dt_mofified" gorm:"column:dt_modified"`
	DTAccessed      string `json:"dt_accessed" gorm:"column:dt_accessed"`
}

func (d *Document) TableName() string {
	return DocumentTableName
}

type Tag struct {
	Guid        string `json:"guid" gorm:"column:tag_guid"`
	GroupGuid   string `json:"group_guid" gorm:"column:tag_group_guid"`
	Name        string `json:"name" gorm:"column:tag_name"`
	Pos         int    `json:"pos" gorm:"column:tag_pos"`
	Description string `json:"description" gorm:"column:description"`
	DTModified  string `json:"dt_modified" gorm:"column:dt_modified"`
	WizVersion  int64  `json:"wiz_version" gorm:"column:wiz_version"`
}

func (t *Tag) TableName() string {
	return TagTableName
}

type DocumentTag struct {
	DocumentGuid string `json:"document_guid" gorm:"column:document_guid"`
	TagGuid      string `json:"tag_guid" gorm:"column:tag_guid"`
}

func (t *DocumentTag) TableName() string {
	return DocumentTagTableName
}
