package main

import "fmt"

// 面积错误
type areaError struct {
	err    string  //错误信息
	length float64 //长度
	width  float64 //宽度
}

// 错误
func (e *areaError) Error() string {
	return e.err
}

// 长度是否小于0
func (e *areaError) lengthNegative() bool {
	return e.length < 0
}

// 宽度是否小于0
func (e *areaError) widthNegative() bool {
	return e.width < 0
}

// 矩形面积
func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero" //长度小于0
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero" // 宽度小于0
		} else {
			err += ",width is less than zero" // 宽度小于0
		}
	}
	//错误信息不等于0
	if err != "" {
		return 0, &areaError{err, length, width}
	}
	return length * width, nil
}

func main() {
	length, width := -5.0, -9.0
	area, err := rectArea(length, width)
	if err != nil {

		// 错误类型断言
		if err, ok := err.(*areaError); ok {
			// 判断长度是否是负数
			if err.lengthNegative() {
				fmt.Printf("error:长度%0.2f小于0\n", err.length)
			}
			// 判断宽度是否是负数
			if err.widthNegative() {
				fmt.Printf("error:宽度%0.2f小于0\n", err.width)
			}
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("面积是", area)
}
