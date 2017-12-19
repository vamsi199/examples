package lorem

import (
	"golang.org/x/net/context"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	httptransport "github.com/go-kit/kit/transport/http"

	"net/http"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
)

func main() {

}
func MakeHttpHandler(ctx context.Context,endpoints Endpoints,logger log.Logger) http.Handler{
	router := mux.NewRouter()
	options:= []httptransport.ServerOption{httptransport.ServerErrorLogger(logger)/*,httptransport.ServerErrorEncoder(encodeError)*/}
	router.Methods("POST").Path("/lorem/{type}/{min}/{max}").Handler(httptransport.NewServer(endpoints.MyEndpoint,decodeLoremReuest, ))

	return nil


}
/*func encodeError(_ context.Context,err error,w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil err")
	}
	w.Header().Set("Content-Type","application/json; charset = utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{} {
		"error":err.Error(),
	})
}*/