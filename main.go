package main

import (
	"github.com/go-redis/redis"
	//"github.com/naoki-kishi/graphql-redis-realtime-chat/infrastructure"
	"log"
	// TODO ここもはがせるようにしたい
)

type config struct {
	RedisURL string `envconfig:"REDIS_URL"`
	Port     int    `envconfig:"PORT"`
}

// type GraphQLServer struct {
// 	redisClient *redis.Client
// }

func main() {

	log.Fatal("error")
	var config config
	// err := envconfig.Process("", &config)
	// if err != nil {
	// 	log.Println(err)
	// }
	config("localhost", 6379)
	log.Print("conf: %v", config)

	client, err := infrastructure.NewRedisClient(config.RedisURL)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// s := infrastructure.NewGraphQLServer(client)
	// log.Fatal(s.Serve("/query", config.Port))
}

// NewRedisClient returns a client for redis.
func NewRedisClient(url string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()

	return client, err
}

// // NewGraphQLServer returns GraphQL server.
// func NewGraphQLServer(redisClient *redis.Client) *GraphQLServer {

// 	return &GraphQLServer{
// 		redisClient: redisClient,
// 	}
// }

// // Serve starts GraphQL server.
// func (s *GraphQLServer) Serve(route string, port int) error {

// 	log.Println("runnning server...")

// 	mux := http.NewServeMux()
// 	mux.Handle(
// 		route,
// 		handler.GraphQL(graphql.NewExecutableSchema(graphql.NewGraphQLConfig(s.redisClient)),
// 			handler.WebsocketUpgrader(websocket.Upgrader{
// 				CheckOrigin: func(r *http.Request) bool {
// 					return true
// 				},
// 			}),
// 		),
// 	)

// 	mux.Handle("/", handler.Playground("GraphQL playground", route))

// 	handler := cors.AllowAll().Handler(mux)
// 	return http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
// }
