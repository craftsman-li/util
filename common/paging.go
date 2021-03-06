package common

import (
	"math"

	"github.com/craftsman-li/util"
	"github.com/emirpasic/gods/lists"
	"github.com/emirpasic/gods/lists/arraylist"
)

type Paging struct {
	Total int64

	Data     interface{}
	PageNo   int64
	PageSize int64
}

func NewPaging(total, pageNo, pageSize int64, data interface{}) *Paging {
	return &Paging{Total: total, Data: data, PageNo: pageNo, PageSize: pageSize}
}

type PagNum struct {
	Key   string
	Value string
}

func (p *Paging) Num() lists.List {
	pp := arraylist.New()
	// 总页数
	totalNo := int64(math.Ceil(float64(p.Total) / float64(p.PageSize)))
	if p.PageNo != 1 && totalNo >= p.PageNo {
		pp.Add(&PagNum{Key: "第一页", Value: "1"})
	}
	if p.PageNo > 2 {
		pp.Add(&PagNum{Key: "...", Value: "javascript;"})
	}
	for i := p.PageNo - 3; i < p.PageNo && i > 0; i++ {
		ii := util.CastInt64ToString(i)
		pp.Add(&PagNum{Key: ii, Value: ii})
	}
	ppp := util.CastInt64ToString(p.PageNo)
	pp.Add(&PagNum{Key: ppp, Value: ppp})

	for i := p.PageNo + 1; i <= totalNo && i-p.PageNo <= 2; i++ {
		ii := util.CastInt64ToString(i)
		pp.Add(&PagNum{Key: ii, Value: ii})
	}
	if totalNo-p.PageNo > 2 {
		pp.Add(&PagNum{Key: "...", Value: "javascript;"})
	}
	if p.PageNo < totalNo && totalNo >= p.PageNo {
		pp.Add(&PagNum{Key: "最后一页", Value: util.CastInt64ToString(totalNo)})
	}
	return pp
}
