package jsonwrap

import (
	"fmt"
	"reflect"
	"strings"
)

type FormatComment func(wraper *JsonWraper, kind reflect.Kind, name string, comment string) string

type JsonWraper struct {
	NewLine           string
	NewField          string
	KeyValueDelimeter string
	CommentTagName    string
	Format            FormatComment
}

func formatComment(wraper *JsonWraper, kind reflect.Kind, name string, comment string) string {
	str := ""
	if strings.Contains(wraper.NewLine, "\n") {
		str = fmt.Sprintf("%s%s, %s%s", wraper.KeyValueDelimeter, kind.String(), comment, wraper.KeyValueDelimeter)
	} else {
		str = fmt.Sprintf("%s%s, %s%s", wraper.KeyValueDelimeter, kind.String(), comment, wraper.KeyValueDelimeter)
	}
	return str
}

func New() *JsonWraper {
	out := JsonWraper{
		NewLine:           "\n",
		NewField:          "\t",
		KeyValueDelimeter: " ",
		CommentTagName:    "jscmm",
		Format:            formatComment,
	}
	return &out
}

// TODO перехватить MarshalJSON
func (wraper *JsonWraper) marshal(sb *strings.Builder, level int, fieldsType reflect.Type) (exitError error) {
	defer func() {
		if r := recover(); r != nil {
			exitError = fmt.Errorf("unsupport error")
		}
	}()
	// fmt.Println(fieldsType)
	switch fieldsType.Kind() {
	case reflect.Struct:
		sb.WriteString("{")
		sb.WriteString(wraper.NewLine)
		count := fieldsType.NumField()
		for i := 0; i < count; i++ {
			fieldType := fieldsType.Field(i)
			name := fieldType.Name
			if val, find := fieldType.Tag.Lookup("json"); find {
				name = val
			}
			comment := ""
			if val, find := fieldType.Tag.Lookup(wraper.CommentTagName); find {
				comment = val
			}
			comment = wraper.Format(wraper, fieldsType.Field(i).Type.Kind(), name, comment)
			fieldsDelimeter := ""
			if i < count-1 {
				fieldsDelimeter = ","
			}
			if strings.Contains(wraper.NewLine, "\n") {
				for i := 0; i < level+1; i++ {
					sb.WriteString(wraper.NewField)
				}
			} else {
				sb.WriteString(wraper.NewField)
			}
			sb.WriteString(fmt.Sprintf("\"%s\":%s", name, wraper.KeyValueDelimeter))
			if exitError = wraper.marshal(sb, level+1, fieldType.Type); exitError != nil {
				return
			}
			sb.WriteString(fmt.Sprintf("%s%s", fieldsDelimeter, wraper.KeyValueDelimeter))
			if strings.Contains(wraper.NewLine, "\n") {
				sb.WriteString(fmt.Sprintf("//%s", comment))
			} else {
				sb.WriteString(fmt.Sprintf("/*%s*/", comment))
			}
			sb.WriteString(wraper.NewLine)
		}
		if strings.Contains(wraper.NewLine, "\n") {
			for i := 0; i < level; i++ {
				sb.WriteString(wraper.NewField)
			}
		} else {
			sb.WriteString(wraper.NewField)
		}
		sb.WriteString("}")
	case reflect.Slice, reflect.Array:
		sb.WriteString("[")
		sb.WriteString(wraper.NewLine)
		if strings.Contains(wraper.NewLine, "\n") {
			for i := 0; i < level+1; i++ {
				sb.WriteString(wraper.NewField)
			}
		} else {
			sb.WriteString(wraper.NewField)
		}
		if exitError = wraper.marshal(sb, level+1, fieldsType.Elem()); exitError != nil {
			return
		}
		sb.WriteString(wraper.NewLine)
		if strings.Contains(wraper.NewLine, "\n") {
			for i := 0; i < level; i++ {
				sb.WriteString(wraper.NewField)
			}
		} else {
			sb.WriteString(wraper.NewField)
		}
		sb.WriteString("]")
	case reflect.Map:
		sb.WriteString("{")
		sb.WriteString(wraper.NewLine)
		if strings.Contains(wraper.NewLine, "\n") {
			for i := 0; i < level+1; i++ {
				sb.WriteString(wraper.NewField)
			}
		} else {
			sb.WriteString(wraper.NewField)
		}
		switch fieldsType.Key().Kind() {
		case reflect.String:
			sb.WriteString(fmt.Sprintf("\"%s\":%s", "strKeyMap", wraper.KeyValueDelimeter))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			sb.WriteString(fmt.Sprintf("\"%s\":%s", "intKeyMap", wraper.KeyValueDelimeter))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			sb.WriteString(fmt.Sprintf("\"%s\":%s", "uintKeyMap", wraper.KeyValueDelimeter))
		default:
			exitError = fmt.Errorf("unsupport map key type")
			return
		}
		if exitError = wraper.marshal(sb, level+1, fieldsType.Elem()); exitError != nil {
			return
		}
		sb.WriteString(wraper.NewLine)
		if strings.Contains(wraper.NewLine, "\n") {
			for i := 0; i < level; i++ {
				sb.WriteString(wraper.NewField)
			}
		} else {
			sb.WriteString(wraper.NewField)
		}
		sb.WriteString("}")
	case reflect.Pointer:
		wraper.marshal(sb, level, fieldsType.Elem())
	case reflect.String:
		sb.WriteString("\"\"")
	case reflect.Bool:
		sb.WriteString("false")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		sb.WriteString("0")
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		sb.WriteString("0")
	case reflect.Float32, reflect.Float64:
		sb.WriteString("0.0")
	case reflect.Interface:
		sb.WriteString("interface{}")
	default:
		exitError = fmt.Errorf("unsupport type")
		return
	}
	return
}

func (wraper *JsonWraper) MarshalByType(t reflect.Type) ([]byte, error) {
	sb := strings.Builder{}
	err := wraper.marshal(&sb, 0, t)
	return []byte(sb.String()), err
}

func (wraper *JsonWraper) Marshal(v interface{}) ([]byte, error) {
	return wraper.MarshalByType(reflect.TypeOf(v))
}
