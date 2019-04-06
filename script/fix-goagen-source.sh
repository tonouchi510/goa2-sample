#!/usr/bin/env bash

PKG_NAME=${1}
REPO=${2}
APP_NAME=${3}

FUNC_DEV_CONSOLE_HANDLER='
func newDevConsoleHandler(pathPrefix string, directory string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fs := http.FileServer(http.Dir(directory))
		http.StripPrefix(pathPrefix, fs).ServeHTTP(w, r)
	}
}'

# package名の修正
find ${PKG_NAME}/*.go | xargs gsed -i 's|package goa2sample|package '${PKG_NAME}'|g'

# main.go/importの修正
gsed -i 's|goa2sample "'${REPO}'"|goa2sample "'${REPO}/${PKG_NAME}'"|g' cmd/${APP_NAME}/main.go
gsed -i '5i \\t"database/sql"' cmd/${APP_NAME}/main.go
gsed -i '15i \\t_ "github.com/go-sql-driver/mysql"' cmd/${APP_NAME}/main.go

# main.go/sqlClientの追加
gsed -i 's|logger \*log.Logger|logger \*log.Logger\n\t\tdb *sql.DB\n\t\terr error|g' cmd/${APP_NAME}/main.go
gsed -i 's|logger = log.New(os.Stderr, "\[goa2sample\] ", log.Ltime)|logger = log.New(os.Stderr, "\[goa2sample\] ", log.Ltime)\n\t\tdb, err = sql.Open("mysql", "test:test@/sampledb")\n\t\tif err != nil {\n\t\t\tlog.Fatal(err.Error())\n\t\t}\n\t\tdefer db.Close()|g' cmd/${APP_NAME}/main.go

# dbClientが必要なServiceを追加した時は以下と同様に置換処理を加える
gsed -i 's|usersSvc = goa2sample.NewUsers(logger)|usersSvc = goa2sample.NewUsers(logger, db)|g' cmd/${APP_NAME}/main.go
gsed -i 's|adminSvc = goa2sample.NewAdmin(logger)|adminSvc = goa2sample.NewAdmin(logger, db)|g' cmd/${APP_NAME}/main.go

# http.Serverのハンドリング変えるための修正（独自handlerを加えるため）
gsed -i 's|{Addr: u.Host, Handler: handler}|{Addr: u.Host}\n\thttp.Handle("/", handler)|g' cmd/${APP_NAME}/http.go
gsed -i 's|http.Handle("/", handler)|http.Handle("/", handler)\n\thttp.HandleFunc("/_dev/console/", newDevConsoleHandler("/_dev/console/","./server/swagger-ui/"))\n|g' cmd/${APP_NAME}/http.go

echo -e "${FUNC_DEV_CONSOLE_HANDLER}" >> cmd/${APP_NAME}/http.go
