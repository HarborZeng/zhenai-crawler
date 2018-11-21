package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strconv"
	"strings"
	"zhenai-crawler/crawler/common/constant"
	"zhenai-crawler/crawler/common/reporter"
	"zhenai-crawler/crawler/common/util"
	"zhenai-crawler/crawler/engine"
	"zhenai-crawler/crawler/model"
)

func ParseProfile(contents []byte) engine.ParseResult {

	profile := model.Profile{}

	document, err := goquery.NewDocumentFromReader(bytes.NewReader(contents))
	if err != nil {
		reporter.ReportError("goquery解析网页出错", err)
	}

	idSelection := document.Find(".info > .id")
	if idSelection.Length() == 0 {
		reporter.ReportMessage("查到一个不开放的用户")
		log.Print("查到一个不开放的用户")
		return engine.ParseResult{}
	}

	if _, ok := constant.DeduplicationBoolMap.Load(idSelection.Text()); !ok {

		extractAvatar(document, &profile)
		extractId(document, &profile)
		extractIntroduction(&profile, document)
		extractBasis(&profile, document)
		extractDetails(&profile, document)
		extractHobbies(&profile, document)

		constant.DeduplicationBoolMap.Store(idSelection.Text(), true)

		return engine.ParseResult{
			Items: []model.Profile{profile},
		}
	} else {
		log.Printf("检测到一个重复的%s\n", idSelection.Text())
		reporter.ReportMessage("检测到一个重复的" + idSelection.Text())
		return engine.ParseResult{}
	}
}

func extractHobbies(profile *model.Profile, document *goquery.Document) {
	document.Find(".m-userInfoDetail > .m-interest > div > .answer").
		Each(func(i int, selection *goquery.Selection) {
			switch i {
			case 0:
				profile.Hobby.Dish = selection.Text()
			case 1:
				profile.Hobby.Celebrity = selection.Text()
			case 2:
				profile.Hobby.Song = selection.Text()
			case 3:
				profile.Hobby.Book = selection.Text()
			case 4:
				profile.Hobby.OutstandingHobby = selection.Text()
			}
		})
}

func extractDetails(profile *model.Profile, document *goquery.Document) {
	document.Find(".pink-btns > div").
		Each(func(i int, selection *goquery.Selection) {
			text := selection.Text()

			switch {
			case strings.Contains(text, "族"):
				profile.Detail.Nationality = text
			case strings.Contains(text, "籍贯"):
				profile.Detail.Birthplace = text[7:]
			case strings.Contains(text, "体型"):
				profile.Detail.FigureType = text[7:]
			case strings.Contains(text, "烟"):
				profile.Detail.Smoking = text
			case strings.Contains(text, "酒"):
				profile.Detail.Drinking = text
			case strings.Contains(text, "购房") || strings.Contains(text, "宿舍") ||
				strings.Contains(text, "租房") || strings.Contains(text, "同住"):
				profile.Detail.Housing = text
			case util.Contain(text, constant.ChildCondition):
				profile.Detail.Child = text
			case util.Contain(text, constant.WillGiveBirthCondition):
				profile.Detail.WillGiveBirth = text[19:]
			case strings.Contains(text, "时机成熟") ||
				strings.Contains(text, "年内") || text == "认同闪婚":
				profile.Detail.WhenToMarry = text[13:]
			case strings.Contains(text, "买车"):
				profile.Detail.Caring = text
			}
		})
}

func extractBasis(profile *model.Profile, document *goquery.Document) {
	var err error

	document.Find(".m-userInfoDetail > .m-title").
		Each(func(i int, selection *goquery.Selection) {
			if strings.Contains(selection.Text(), "他") {
				profile.Basis.Gender = "男"
			} else if strings.Contains(selection.Text(), "她") {
				profile.Basis.Gender = "女"
			}
		})

	//昵称
	profile.Basis.Nickname = document.
		Find(".right > .info > .name > .nickName").First().Text()
	//实名认证
	profile.Basis.IsRealName = document.
		Find(".right > .info > .name > .realname").Length() != 0
	//珍爱网会员
	profile.Basis.IsVIP = document.
		Find(".right > .info > .name > .zhenxin").Length() != 0

	document.Find(".purple-btns > div").
		Each(func(i int, selection *goquery.Selection) {
			text := selection.Text()

			switch {
			case util.Contain(text, constant.EducationCondition):
				profile.Basis.Education = text
			case util.Contain(text, constant.OccupationCondition):
				profile.Basis.Occupation = text
			case strings.HasPrefix(text, "月收入"):
				profile.Basis.Income = text[10:]
			case strings.HasSuffix(text, "岁"):
				profile.Basis.Age, err = strconv.Atoi(
					strings.Replace(text, "岁", "", -1))
				if err != nil {
					reporter.ReportError("获取年龄出错", err)
				}
			case strings.HasSuffix(text, "cm"):
				profile.Basis.Height, err = strconv.Atoi(
					strings.Replace(text, "cm", "", -1))
				if err != nil {
					reporter.ReportError("获取身高出错", err)
				}
			case strings.HasSuffix(text, "kg"):
				profile.Basis.Weight, err = strconv.Atoi(
					strings.Replace(text, "kg", "", -1))
				if err != nil {
					reporter.ReportError("获取体重出错", err)
				}
			case strings.Contains(text, "(") &&
				strings.HasSuffix(text, ")"):
				profile.Basis.Sigh = text[:9]
			case strings.HasPrefix(text, "工作地"):
				profile.Basis.WorkPlace = text[10:]
			case util.Contain(text, constant.MaritalStatusCondition):
				profile.Basis.MaritalStatus = text
			}
		})
}

func extractIntroduction(profile *model.Profile, document *goquery.Document) {
	//内心独白、简介
	profile.Introduction = document.Find(".m-des > span").Text()
}

func extractId(document *goquery.Document, profile *model.Profile) {
	//user id
	var err error
	profile.Id, err = strconv.ParseInt(
		document.Find(".info > .id").Text()[5:], 10, 64)
	if err != nil {
		reporter.ReportError("获取用户ID出错", err)
	}
}

func extractAvatar(document *goquery.Document, profile *model.Profile) {
	//头像URL
	mainAvatar, ex := document.
		Find(".top > .logo").Attr("style")
	if !ex {
		reporter.ReportMessage("未成功获取用户头像URL: " + profile.Basis.Nickname)
	}
	if strings.Contains(mainAvatar, "?") {
		mainAvatar = mainAvatar[:strings.Index(mainAvatar, "?")]
	}
	profile.Basis.Avatar = strings.Replace(
		mainAvatar, "background-image:url(", "", -1)
}
