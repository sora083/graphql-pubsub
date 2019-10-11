package main

import (
	"log"

	"github.com/go-redis/redis"
	// TODO ここもはがせるようにしたい
	//"github.com/naoki-kishi/graphql-redis-realtime-chat/infrastructure"
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

	// s := NewGraphQLServer(client)
	// log.Print("graphql server: ", &s)
	// //log.Fatal(s.Serve("/query", config.Port))
	// log.Fatal(s.Serve("/query", 8000))

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
// 		handler.GraphQL(NewExecutableSchema(NewGraphQLConfig(s.redisClient)),
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

// type executableSchema struct {
// 	resolvers  ResolverRoot
// 	directives DirectiveRoot
// 	complexity ComplexityRoot
// }

// type Config struct {
// 	Resolvers  ResolverRoot
// 	Directives DirectiveRoot
// 	Complexity ComplexityRoot
// }

// type ResolverRoot interface {
// 	Mutation() MutationResolver
// 	Query() QueryResolver
// 	Subscription() SubscriptionResolver
// }

// type DirectiveRoot struct {
// }

// type ComplexityRoot struct {
// 	Message struct {
// 		User    func(childComplexity int) int
// 		Message func(childComplexity int) int
// 	}

// 	Mutation struct {
// 		PostMessage func(childComplexity int, user string, message string) int
// 		CreateUser  func(childComplexity int, user string) int
// 	}

// 	Query struct {
// 		Users func(childComplexity int) int
// 	}

// 	Subscription struct {
// 		MessagePosted func(childComplexity int, user string) int
// 		UserJoined    func(childComplexity int, user string) int
// 	}
// }

// type MutationResolver interface {
// 	PostMessage(ctx context.Context, user string, message string) (*Message, error)
// 	CreateUser(ctx context.Context, user string) (string, error)
// }
// type QueryResolver interface {
// 	Users(ctx context.Context) ([]string, error)
// }
// type SubscriptionResolver interface {
// 	MessagePosted(ctx context.Context, user string) (<-chan Message, error)
// 	UserJoined(ctx context.Context, user string) (<-chan string, error)
// }

// type Message struct {
// 	User    string `json:"user"`
// 	Message string `json:"message"`
// }

// // NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
// func NewExecutableSchema(cfg Config) ExecutableSchema {
// 	return &executableSchema{
// 		resolvers:  cfg.Resolvers,
// 		directives: cfg.Directives,
// 		complexity: cfg.Complexity,
// 	}
// }

// // NewGraphQLConfig returns Config and start subscribing redis pubsub.
// func NewGraphQLConfig(redisClient *redis.Client) Config {
// 	resolver := newResolver(redisClient)

// 	resolver.startSubscribingRedis()

// 	return Config{
// 		Resolvers: resolver,
// 	}
// }

// // Resolver implements ResolverRoot interface.
// type Resolver struct {
// 	redisClient     *redis.Client
// 	messageChannels map[string]chan Message
// 	userChannels    map[string]chan string
// 	mutex           sync.Mutex
// }

// func newResolver(redisClient *redis.Client) *Resolver {
// 	return &Resolver{
// 		redisClient:     redisClient,
// 		messageChannels: map[string]chan Message{},
// 		userChannels:    map[string]chan string{},
// 		mutex:           sync.Mutex{},
// 	}
// }

// // Mutation returns a resolver for mutation.
// func (r *Resolver) Mutation() MutationResolver {
// 	return &mutationResolver{r}
// }

// // Query returns a resolver for query.
// func (r *Resolver) Query() QueryResolver {
// 	return &queryResolver{r}
// }

// // Subscription returns a resolver for subsctiption.
// func (r *Resolver) Subscription() SubscriptionResolver {
// 	return &subscriptionResolver{r}
// }

// func (r *Resolver) startSubscribingRedis() {
// 	log.Println("Start Subscribing Redis...")

// 	go func() {
// 		pubsub := r.redisClient.Subscribe("room")
// 		defer pubsub.Close()

// 		for {
// 			msgi, err := pubsub.Receive()
// 			if err != nil {
// 				panic(err)
// 			}

// 			switch msg := msgi.(type) {
// 			case *redis.Message:

// 				// Convert recieved string to Message.
// 				m := Message{}
// 				if err := json.Unmarshal([]byte(msg.Payload), &m); err != nil {
// 					log.Println(err)
// 					continue
// 				}

// 				// Notify new message.
// 				r.mutex.Lock()
// 				for _, ch := range r.messageChannels {
// 					ch <- m
// 				}
// 				r.mutex.Unlock()

// 			default:
// 			}
// 		}
// 	}()
// }
