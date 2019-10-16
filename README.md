# graphql-pubsub

### command
####  mod init
```
go mod init github.com/sora083/graphql-pubsub
```

####  generate
```
cd graphql/scripts
go build
./scripts
```

```
cd graphql
go run scripts/gqlgen.go init
```

### reference
####
[GraphQL SubscriptionsとRedis PubSubを使ってリアルタイムチャットサーバーを作る](https://qiita.com/p1ass/items/462209fe73ece1238d85)
[graphql-redis-realtime-chat](https://github.com/p1ass/graphql-redis-realtime-chat)

#### redis
[MacにRedisをインストールする](https://qiita.com/sawa-@github/items/1f303626bdc219ea8fa1)

```
brew services start redis
```

### Check
```
http://localhost:8000
```

### queries
#### reqd
```
query {
  users
}
```

```
query all {
  users
}
```

#### publish
```
mutation createUser {
  createUser(user: "test")
}
```

```
mutation postMessage {
  postMessage(user: "test", message: "test message") {
    user,
    message
  }
}
```

#### subscribe
```
subscription {
  messagePosted(user: "test") {
    user, message
  }
}
```

```
subscription {
  userJoined(user: "test")
}
```

### TODO
* バイナリをgit管理外にする
* generateで「SubscriptionResolver」が期待と違うものになる・・
* 環境変数化
* messagePostedのsubscriptionがうまく動かない・・