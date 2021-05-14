package err

import (
	"fmt"

	"github.com/pkg/errors"
)

//PanicOnError - panic with error wrapped with StackTrace
func PanicOnError(err error) {
	if err != nil {
		err = errors.WithStack(err)
		panic(err)
	}
}

//Catch - Catch fatal error and returns as error
//r - inerface{} - recover() or any value
//from - method which executing, for better debuging
func Catch(r interface{}, from string) (err error) {
	if r != nil {
		switch r.(type) {
		case error:
			err = r.(error)
		default:
			err = fmt.Errorf("%+v", r)
		}
	}
	if err != nil {
		err = errors.Wrap(err, from+":")
		err = errors.WithStack(err)
	}
	return
}

//func CatchWithLogging(log log.Service, r interface{})(err error){
//	if r != nil {
//		switch r.(type) {
//		case error:
//			err = r.(error)
//		default:
//			err = fmt.Errorf("%+v", r)
//		}
//	}
//	if(err != nil){
//		log.PanicOnError(models.LogMessage{
//			Message:  fmt.Sprintf("%v", err),
//			Detailed: fmt.Sprintf("%+v", err),
//		})
//		err = errors.WithStack(err)
//	}
//	return
//}
//
//func CatchHttp(c echo.Context, r interface{},from string)(err error){
//	if r != nil {
//		switch r.(type) {
//		case error:
//			err = r.(error)
//		default:
//			err = fmt.Errorf("%+v", r)
//		}
//	}
//	if err != nil{
//		err = errors.WithStack(err)
//		//c.PanicOnError(err)
//		c.JSON(http.StatusInternalServerError, models.BaseResponse{
//			Message: "Internal server error",
//			PanicOnError: fmt.Sprintf("%v", err),
//			Detailed: fmt.Sprintf("%+v", err),
//		})
//	}
//	return
//}
