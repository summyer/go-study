package mmm

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var url = "pig:pLlGJSpdg1RCp^80@tcp(10.20.152.211:3306)/mmm?charset=utf8mb4&parseTime=True&loc=Local"

func init() {
	url = "pig:pLlGJSpdg1RCp^80@tcp(10.20.152.213:3306)/mmm?charset=utf8mb4&parseTime=True&loc=Local"
}
func GetConnection(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //日志打印
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		fmt.Printf("连接数据库%s出错", dsn)
		panic("连接数据库出错")
	}
	return db
}

type Archives struct {
	Code         string `gorm:"primary_key"`
	Detail       string
	DetailEn     string
	RepairPlan   string
	RepairPlanEn string
}

//来源于211的翻译数据
type ArchivesCopy struct {
	Code         string `gorm:"primary_key"`
	Detail       string
	DetailEn     string
	RepairPlan   string
	RepairPlanEn string
}

func Notify() {
	url = "pig:pLlGJSpdg1RCp^80@tcp(10.20.152.211:3306)/mmm?charset=utf8mb4&parseTime=True&loc=Local"
	var db *gorm.DB = GetConnection(url)
	var count int64
	//不要忘记表名  db.Model(&user) or db.Table("users")
	queryCount := "((detail_en is not null and reserved_flag=0)or(reserved_flag=1)) and repair_plan_en != '0'"
	res := db.Table("archives").Where(queryCount).Count(&count)
	fmt.Println(res.Error)
	fmt.Println(count)
	var flag bool = true
	for {
		time.Sleep(5 * time.Minute)
		var count2 int64
		//不要忘记表名  db.Model(&user) or db.Table("users")
		res := db.Table("archives").Where(queryCount).Count(&count2)
		fmt.Println(res.Error)
		fmt.Println(count2)

		if count2 == count && flag {
			log.Println("翻译通知")
			SendDingMsg(fmt.Sprintf("翻译通知:【翻译停止，保留档案不翻译】,当前:%d", count2))
			SendDingMsg(fmt.Sprintf("翻译通知:【翻译停止，保留档案不翻译】！,当前:%d", count2))
			SendDingMsg(fmt.Sprintf("翻译通知:【翻译停止，保留档案不翻译】！！,当前:%d", count2))
			flag = false
		}

		if count2 > count && !flag {
			log.Println("翻译通知")
			SendDingMsg(fmt.Sprintf("翻译通知:【再次新增】,当前:%d", count2))
			SendDingMsg(fmt.Sprintf("翻译通知:【再次新增】！,当前:%d", count2))
			SendDingMsg(fmt.Sprintf("翻译通知:【再次新增】！！,当前:%d", count2))
			flag = true
		}
		count = count2
	}

}

func Start() {
	var db *gorm.DB = GetConnection(url)
	//var archives []Archives
	//db.Find(&archives, []string{"DM-202203-000355", "DM-202203-000357"})
	//fmt.Println(len(archives))
	var count int64
	//不要忘记表名  db.Model(&user) or db.Table("users")
	res := db.Table("archives").Where("detail_en is not null").Count(&count)
	fmt.Println(res.Error)
	fmt.Println(count)
	var pageSize int = 20
	fmt.Println()
	var totalPage int
	if int(count)%pageSize == 0 {
		totalPage = int(count) / pageSize
	} else {
		totalPage = int(count)/pageSize + 1
	}
	for page := 1; page <= totalPage; page++ {
		FindPage(db, page, pageSize)
	}

}
func GetArchivesCopyMap(db *gorm.DB, codes []string) map[string]ArchivesCopy {
	if len(codes) == 0 {
		return make(map[string]ArchivesCopy)
	}
	var archivesCopy []ArchivesCopy
	db.Find(&archivesCopy, codes)
	var archivesCopyMap map[string]ArchivesCopy = map[string]ArchivesCopy{}
	for _, v := range archivesCopy {
		archivesCopyMap[v.Code] = v
	}
	return archivesCopyMap
}
func FindPage(db *gorm.DB, page, pageSize int) {
	log.Println("当前页:", page)
	var archives []Archives
	db.Limit(pageSize).Offset((page - 1) * pageSize).Where("detail_en is null").Find(&archives)
	codes := make([]string, len(archives))
	for i := 0; i < len(codes); i++ {
		codes[i] = archives[i].Code
	}
	var archivesCopyMap map[string]ArchivesCopy
	//获取archives_copy表map数据
	archivesCopyMap = GetArchivesCopyMap(db, codes)
	//for k, v := range archivesCopyMap {
	//	fmt.Println(k, v)
	//}
	//遍历archives表

	for k, archive := range archives {
		fmt.Println(archive.Code)
		//m中有key为zero对应的值为0，所以直接通过m["zero"]取值，获取到的ok为true，且value=0，但是在使用m["three"]取值时，由于m中没有three这个key，所以获取到的ok为false，value同样也是0
		archiveCopy, ok := archivesCopyMap[archive.Code]
		if ok {
			if archive.Detail == archiveCopy.DetailEn {
				archives[k].Detail = archiveCopy.Detail
				archives[k].DetailEn = archiveCopy.DetailEn
			} else {
				archives[k].DetailEn = archiveCopy.DetailEn
			}

			if archive.RepairPlan == archiveCopy.RepairPlan {
				archives[k].RepairPlanEn = archiveCopy.RepairPlanEn
			}
		}
	}
	SaveBatch(db, archives)
	fmt.Println(len(archives))
}

func SaveBatch(db *gorm.DB, archives []Archives) {
	//好像没有像mybatisplus那样的批量更新
	for _, archive := range archives {
		db.Save(archive)
	}
}
