package goroutemanager

import (
	"net/http"
	"regexp"
	"strings"
)

type routeManager struct {
	routes map[string]map[string]Route
}

type Route struct {
	Method  string
	Uri     string
	Execute func(http.ResponseWriter, *http.Request, map[string]interface{})
}

func (rm *routeManager) add(method string, uri string, execute func(http.ResponseWriter, *http.Request, map[string]interface{})) *routeManager {

	rt := Route{
		Uri:     uri,
		Method:  "GET",
		Execute: execute,
	}

	rm.routes[method][rt.Uri] = rt

	return rm
}

func RouteManagerInit() *routeManager {
	routes := map[string]map[string]Route{
		"DELETE": {},
		"GET":    {},
		"POST":   {},
	}

	return &routeManager{
		routes: routes,
	}
}

func (rm *routeManager) Delete(uri string, exec func(http.ResponseWriter, *http.Request, map[string]interface{})) *routeManager {
	return rm.add("DELETE", uri, exec)
}

func (rm *routeManager) Get(uri string, exec func(http.ResponseWriter, *http.Request, map[string]interface{})) *routeManager {
	return rm.add("GET", uri, exec)
}

func (rm *routeManager) Post(uri string, exec func(http.ResponseWriter, *http.Request, map[string]interface{})) *routeManager {
	return rm.add("POST", uri, exec)
}

func (rm *routeManager) HandleFunc(w http.ResponseWriter, r *http.Request) {
	routes := rm.routes[r.Method]
	path := r.URL.Path

	for _, route := range routes {

		var uriSplit []string = strings.Split(route.Uri, "/")
		var pathSplit []string = strings.Split(path, "/")
		var equals bool = true
		var fields map[string]interface{} = map[string]interface{}{}
		// a uri da rota tem que possuir o mesmo tamanho de
		// seções do path ex:
		// /user/select/field:id[0-9]+ == /user/select/1
		if len(uriSplit) == len(pathSplit) {
			// faz um loop comparando seção por sessão
			// casa seção é o conteúdo entre as barras
			for index, sectionUri := range uriSplit {
				sectionPath := pathSplit[index]
				// Verifica se seção passada é um parametro
				if strings.Contains(sectionUri, "field:") {

					var split []string = strings.Split(sectionUri, ":")
					var field string = split[1]
					// Caso tenha o comprimento de 3  significa que este campo
					// possui um regex de comparação caso não tenha
					// significa que o valor do campo pode ser qualquer coisa
					if len(split) == 3 {
						var regex string = split[2]
						matched, err := regexp.MatchString("^"+regex+"$", sectionPath)

						if !matched || err != nil {
							equals = false
							break
						}

						fields[field] = sectionPath
					}

					fields[field] = sectionPath

				} else if !strings.EqualFold(sectionUri, sectionPath) {
					equals = false
					break
				}
			}

			if equals {
				route.Execute(w, r, fields)
				break
			}
		}
	}

}
