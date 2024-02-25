package util

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// type any = interface{}

// copy copy struct field
func CopyStruct(dst, src any) error {
	// err := checkCopy(dst, src)
	// if err != nil {
	// 	return err
	// }
	sv := reflect.ValueOf(src)
	dv := reflect.ValueOf(dst)
	// reflect.Indirect(sv)
	sv = reflect.Indirect(sv)
	dv = reflect.Indirect(dv)
	// sv = sv.Elem() // reflect.Indirect(sv)
	// dv = dv.Elem() //reflect.Indirect(dv)

	fields := modelFields(sv)
	// fmt.Println(fields)
	for _, f := range fields {
		// f.Index
		sfv := sv.FieldByName(f.Name)
		// get dst field by name
		dfv := dv.FieldByName(f.Name)

		if !dfv.IsValid() {
			continue
		}
		err := copyVal(dfv, sfv)
		if err != nil {
			fmt.Println(err, "***", f, "***")
		}
		// fmt.Println("hh***", sfv, "*", dfv, "***")
	}
	return nil
}

func copyVal(dfv reflect.Value, sfv reflect.Value) error {
	dt := dfv.Type()
	st := sfv.Type()
	// st.AssignableTo()
	// fmt.Println("dt.Kind()", dt.Kind()) fmt.Println("st.Kind()", st.Kind())
	if st.Kind() == dt.Kind() {
		dfv.Set(sfv)
		return nil
	} else {
		// 不一样的
		if dt.Kind() == reflect.String {
			switch st.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				intStr := strconv.FormatInt(sfv.Int(), 10)
				dfv.SetString(intStr)
				return nil
			case reflect.Struct:
				// // fmt.Println("特殊123类型", st.String(), st.Name()) 特殊123类型 time.Time Time
				if st.String() == "time.Time" {
					// sfv.Pointer()
					time1 := sfv.Interface().(time.Time)
					// 结构体不能这么写 fff := *(*time.Time)(sfv.UnsafePointer())
					// *(*float64)(v.ptr)
					time1Str := time1.Format("2006-01-02 15:04:05")
					dfv.SetString(time1Str)
				}
				return nil
			case reflect.Bool:
				bool1 := sfv.Bool()
				if bool1 {
					dfv.SetString("true")
				} else {
					dfv.SetString("false")
				}
				return nil
			default:
				return nil
			}
		}
	}
	return nil
}

func modelFields(v reflect.Value) []reflect.StructField {
	t := v.Type()
	fs := make([]reflect.StructField, 0, t.NumField())
	//
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		// 比较类型 f.Type.Kind() == reflect.Chan
		// fmt.Println("PkgPath:"+f.PkgPath+";", f.Name, f.Type, "***")
		// 小写字段的PkgPath会带上包名,比如main
		if f.PkgPath == "" {
			fs = append(fs, f)
		}
	}
	return fs
}

func checkCopy(dst, src any) error {
	if src == nil {
		return errors.New("source is nil")
	}
	if dst == nil {
		return errors.New("destination is nil")
	}

	sv := reflect.ValueOf(src)
	dv := reflect.ValueOf(dst)

	if sv.Kind() != reflect.Pointer {
		return errors.New("src is not ptr")
	}

	if dv.Kind() != reflect.Pointer {
		return errors.New("dest is not ptr")
	}

	if sv.Elem().Kind() != reflect.Struct ||
		dv.Elem().Kind() != reflect.Struct {
		return errors.New("elem is not struct")
	}

	return nil
}
