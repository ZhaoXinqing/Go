type FoldersModel struct {
	BaseModel

	Name      string
	Type      int64 //0:保存数据集，1:保存看板,2:保存标签
	ParentID  int64
	Path      string // 文件夹路径
	Level     int64  // 路径级别
	IsDelete  bool
	CreatorID int64 // 创建人ID
	CreAtUpdAtField
}

// MaterialsModel ...
type MaterialsModel struct {
	BaseModel

	Name      string
	Type      int64  // 素材类型：1、图形，2、图片，3、视频
	Width     int64  // 宽度
	Height    int64  // 高度
	URL       string // 素材链接
	CreatorID int64  // 创建人ID
	IsPublic  bool   // 是否公开
	IsDelete  bool
	CreAtUpdAtField
}

// TagsModel ...
type TagsModel struct {
	BaseModel
	Name        string
	Description string
	Rule        string
	Style       string
	GroupID     int64
	IsDelete    bool
	CreAtUpdAtField
	CreatorID int64
}

// ElementsModel ...
type ElementsModel struct {
	BaseModel

	Name      string
	Type      int64 // 元素类型：1、地图，2、图表，3、文字，4、图形, 5、图片
	BoardID   int64
	X         float64 // 坐标 X
	Y         float64 // 坐标 Y
	Z         float64 // 坐标 Z 层级
	Width     int64   // 宽度
	Height    int64   // 高度
	Options   string  // 配置项 json 类型
	CreatorID int64   // 创建人ID
	IsLock    bool    // 是否锁定
	IsDelete  bool
	CreAtUpdAtField
}

// DatasetsModel ...
type DatasetsModel struct {
	BaseModel

	Title        string
	FeatureType  string // Point, LineString, Polygon, MultiPoint, MultiLineString, and MultiPolygon
	FeatureCount int64
	CRS          string // 坐标参考系
	IsDelete     bool
	HasParsed    bool  // 是否空间化
	ShareType    int64 // 归属 0：个人 1：团队 2：公共
	CreatorID    int64 // 创建人id
	TitleFieldID int64 // 标题id
	FolderID     int64 // 所属文件夹ID
	CreAtUpdAtField
}

// BoardsModel ...
type BoardsModel struct {
	BaseModel

	Name       string
	FolderID   int64  // 所属文件夹ID
	Width      int64  // 宽度
	Height     int64  // 高度
	Thumbnail  string // 缩略图
	Options    string // 配置项 json 类型
	Production string // 已发布的看板信息
	CreatorID  int64  // 创建人ID
	IsDelete   bool
	CreAtUpdAtField
}

// DatasetFieldsModel ...
type DatasetFieldsModel struct {
	BaseModel

	Name       string
	Type       int64 //0:文本，1:数值，2:日期
	DatasetID  int64
	IsDelete   bool
	IsVisible  bool
	OrderValue float64 //排序字段
	CreatorID  int64
	CreAtUpdAtField
}

// FeaturesModel ...
type FeaturesModel struct {
	BaseModel

	DatasetID int64
	Geom      *geometry.Geom
	Infos     string // jsonb
	CreAtUpdAtField
	Tags     string // jsonb
	IsDelete bool
}