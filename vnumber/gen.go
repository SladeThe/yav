package vnumber

//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_equal.go -out=equal.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"

//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_range.go -out=range.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_range_test.go -out=range_test.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"

//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_required_if.go -out=required_if.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_required_if_test.go -out=required_if_test.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_required_unless.go -out=required_unless.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_required_unless_test.go -out=required_unless_test.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_required_with_all.go -out=required_with_all.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_required_with_all_test.go -out=required_with_all_test.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_required_with_any.go -out=required_with_any.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_required_with_any_test.go -out=required_with_any_test.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_required_without_all.go -out=required_without_all.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_required_without_all_test.go -out=required_without_all_test.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_required_without_any.go -out=required_without_any.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_required_without_any_test.go -out=required_without_any_test.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"

//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_excluded_if.go -out=excluded_if.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_excluded_unless.go -out=excluded_unless.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_excluded_with_all.go -out=excluded_with_all.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_excluded_with_any.go -out=excluded_with_any.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_excluded_without_all.go -out=excluded_without_all.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
//go:generate go run github.com/cheekybits/genny@3e22f1a -in=_excluded_without_any.go -out=excluded_without_any.go -pkg=vnumber gen "Element=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"
