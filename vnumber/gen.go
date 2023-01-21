package vnumber

//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_equal.go -out=equal.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_range.go -out=range.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_range_test.go -out=range_test.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_required_with_all.go -out=required_with_all.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_required_with_any.go -out=required_with_any.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_required_without_all.go -out=required_without_all.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_required_without_any.go -out=required_without_any.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
