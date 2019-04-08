package main

import (
	"log"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	http.ListenAndServe(":3737", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			log.Printf("upgrade error: %v", err)
			return
		}
		go func() {
			defer conn.Close()

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					log.Printf("read error: %v", err)
					return
				}

				err = wsutil.WriteServerMessage(conn, op, msg)
				if err != nil {
					log.Printf("write error: %v", err)
					return
				}
			}
		}()
	}))
}
