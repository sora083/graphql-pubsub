package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/sora083/graphql-pubsub/graphql"
)

type config struct {
	RedisURL string `envconfig:"REDIS_URL"`
	Port     int    `envconfig:"PORT"`
}

type GraphQLServer struct {
	redisClient *redis.Client
}

func main() {

	log.Print("START!!")
	//var config config
	// err := envconfig.Process("", &config)
	// if err != nil {
	// 	log.Println(err)
	// }
	// config("localhost", 6379)
	// log.Print("conf: %v", config)

	//	client, err := RedisClient(config.RedisURL)
	client, err := RedisClient("localhost:6379")
	if err != nil {
		panic(err)
	}
	log.Print("redis client: ", client)
	defer client.Close()

	s := NewGraphQLServer(client)
	log.Print("graphql server: ", &s)
	//log.Fatal(s.Serve("/query", config.Port))
	log.Fatal(s.Serve("/query", 8000))

	log.Print("END!!")
}

// RedisClient returns a client for redis.
func RedisClient(url string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()

	return client, err
}

// NewGraphQLServer returns GraphQL server.
func NewGraphQLServer(redisClient *redis.Client) *GraphQLServer {

	return &GraphQLServer{
		redisClient: redisClient,
	}
}

// Serve starts GraphQL server.
func (s *GraphQLServer) Serve(route string, port int) error {

	log.Println("runnning server...")

	mux := http.NewServeMux()
	mux.Handle(
		route,
		handler.GraphQL(graphql.NewExecutableSchema(graphql.NewGraphQLConfig(s.redisClient)),
			handler.WebsocketUpgrader(websocket.Upgrader{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			}),
		),
	)

	mux.Handle("/", handler.Playground("GraphQL playground", route))

	handler := cors.AllowAll().Handler(mux)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
}



//
