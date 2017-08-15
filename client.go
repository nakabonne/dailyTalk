package main

import (
	"github.com/gorilla/websocket"
)

// 一人のユーザー
type client struct {
	// このクライアントのためのwebsocket
	socket *websocket.Conn
	// メッセージが送られるチャネル
	send chan []byte
	// このクライアントが参加してるチャットルーム
	room *room
}
