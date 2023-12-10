package models

// HotRec 热门推荐的
type HotRec struct {
	CoverUrl     string `json:"cover_url"`      //视频url
	CoverIcon    string `json:"cover_icon"`     //热播视频标识
	CoverImg     string `json:"cover_img"`      // 视频cover
	UpEpisode    string `json:"up_episode"`     // 视频更新至哪页面
	VideoNameDes string `json:"video_name_des"` //视频名字描述
	VideoBrief   string `json:"video_des"`      //视频简介
}

// TVPlayInfo 电视剧集信息
type TVPlayInfo struct {
	Cid      string `json:"cid"`       //当前电视剧id
	Vid      string `json:"vid"`       //视频集数id
	VIcon    string `json:"v_icon"`    // 视频icon信息 预，vip等标识
	HumanVid string `json:"human_vid"` //人可以读懂的电视集数
}
