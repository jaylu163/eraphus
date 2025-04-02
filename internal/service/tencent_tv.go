package service

import (
	"bytes"
	"context"
	"github.com/jaylu163/eraphus/internal/hades/logging"
	"github.com/jaylu163/eraphus/internal/manager"
	"github.com/jaylu163/eraphus/internal/models"

	"github.com/PuerkitoBio/goquery"
)

// GetHotRec 热门推荐
func GetHotRec(ctx context.Context, urlCover string) ([]models.HotRec, error) {
	ret, err := manager.GetRestCli().R().Get(urlCover)
	if err != nil {
		logging.WithFor(ctx, "func:", "GetHotRec").Errorf("resty get err:%v", err)
		return nil, err
	}
	if ret.StatusCode() != 200 {
		logging.Errorf("status code error: %d  %s", ret.StatusCode, ret.Status)
		return nil, err
	}
	docReader, err := goquery.NewDocumentFromReader(bytes.NewReader(ret.Body()))
	if err != nil {
		logging.Errorf("NewDocumentFromReader err:%v", err)
	}
	hotRecList := []models.HotRec{}
	// 获取热门视频列表
	docReader.Find("div.video-banner-item").Each(func(i int, selection *goquery.Selection) {
		subSelection := selection.Find("a")
		videoCoverHref, _ := subSelection.Attr("href")
		coverIcon, _ := subSelection.Children().Find("div.right-top-tag-area img").Attr("src")
		coverImg, _ := subSelection.Children().Find("img").Attr("src")
		episode := subSelection.Children().Find("img").Children().Find("span").Text()

		nameBrief := selection.Find("div.twice-title-area").Find("a").Text()
		videoBrief := selection.Find("div.twice-title-area").Find("span").Text()

		//fmt.Println("热门电视剧cover链接:", videoCoverHref, coverIcon, coverImg, episode, nameBrief, videoBrief)
		hotRecList = append(hotRecList, models.HotRec{
			CoverUrl:     videoCoverHref,
			CoverIcon:    "https://" + coverIcon,
			CoverImg:     "https://" + coverImg,
			UpEpisode:    episode,
			VideoNameDes: nameBrief,
			VideoBrief:   videoBrief,
		})
	})
	return hotRecList, nil
}

// GetVideoList 获取腾讯视频电视
func GetVideoList(ctx context.Context, tcDetailUrl string) ([]models.TVPlayInfo, error) {
	ret, err := manager.GetRestCli().R().Get(tcDetailUrl)
	log := logging.WithFor(ctx, "func:", "GetVideoList")
	if err != nil {
		log.Errorf("resty get err:%v", err)
		return nil, err
	}
	if ret.StatusCode() != 200 {
		log.Errorf("status code error: %d  %s", ret.StatusCode, ret.Status)
		return nil, err
	}
	docReader, err := goquery.NewDocumentFromReader(bytes.NewReader(ret.Body()))
	if err != nil {
		log.Errorf("NewDocumentFromReader err:%v", err)
	}
	tvList := []models.TVPlayInfo{}
	docReader.Find("div.episode-list-rect__item").Each(func(i int, selection *goquery.Selection) {
		cid, _ := selection.Children().Attr("data-cid")
		vid, _ := selection.Children().Attr("data-vid")

		imgIcon, _ := selection.Children().Find("img").Attr("src")
		curentJIshu := selection.Children().Find("span").Text()
		//fmt.Println("cid", cid, "vid", vid, "剧集:", curentJIshu)
		//logs.WithFor(ctx, "func", "GetVideoList").Infof("video info cid:%s vid:%s 剧集:%s", cid, vid, curentJIshu)
		tvList = append(tvList, models.TVPlayInfo{
			Cid:      cid,
			Vid:      vid,
			VIcon:    imgIcon,
			HumanVid: curentJIshu,
		})
	})
	log.Infof("video info cid:%s vid:%s", "aaa", "bbb")

	manager.GetHotList(ctx, "123")
	return tvList, nil
}

func GetVideoStream(ctx context.Context, cid, vid string) {
	// 视频url
	//videoUrl := fmt.Sprintf("https://v.qq.com/x/cover/%s/%s.html", cid, vid)
	// 拿到vid拼接后返回 stream
	// https://vv.video.qq.com/getinfo?defn=fhd&platform=10801&otype=ojson&sdtfrom=v4138&appVer=7&vid=e0045qljjwf&newnettype=1&fhdswitch=1&show1080p=1&dtype=3&sphls=2

}
