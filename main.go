package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/singleflight"
)

var (
	once sync.Once
	conf *config
)

type config struct {
	host string
}

func GetConfig() *config {
	once.Do(func() {
		conf = &config{
			host: "127.0.0.1",
		}
	})
	return conf
}

func main() {
	// gracefullyShutdown()
	// cancelableSleep()
	// reuseBullet()
	// stringer()
	callSingleflightFunc()
}

func gracefullyShutdown() {
	fmt.Println("gracefullyShutdown")
	ctx, cancel := signal.NotifyContext(
		context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	server := http.Server{
		Addr: ":8080",
	}
	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return server.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		log.Println("Shutdown signal received, gracefully shutting down...")
		return server.Shutdown(context.Background())
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

type Lockable[T any] struct {
	lock sync.Mutex
	data T
}

func (l *Lockable[T]) Get() T {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.data
}

func (l *Lockable[T]) Set(data T) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.data = data
}

type User struct {
	Name string
	Age  int
}

type LockableUser Lockable[User]

func lockableUser() {
	var lockableUser Lockable[User]
	// lockableUser.Set(...)
	_ = lockableUser

	// user := Lockable(LockableUser{
	// 	data: User{
	// 		Name: "Go",
	// 		Age:  20,
	// 	},
	// })
	// user.Set(User{
	// 	Name: "Go",
	// 	Age:  20,
	// })
}

func track() func() {
	start := time.Now()
	return func() {
		fmt.Println("elasped: ", time.Since(start))
	}
}

func track2(start time.Time) func() {
	return func() {
		fmt.Println("elasped: ", time.Since(start))
	}
}

func sleepx(ctx context.Context, duration time.Duration) error {
	// defer track()()
	defer track2(time.Now())()
	timer := time.NewTimer(duration)
	defer timer.Stop()
	select {
	case <-timer.C:
		fmt.Println("timeout")
		return nil
	case <-ctx.Done():
		fmt.Println("cancelled")
		return ctx.Err()
	}
}

func sleepy(ctx context.Context, duration time.Duration) error {
	defer track()()
	select {
	case <-ctx.Done():
		fmt.Println("cancelled")
		return ctx.Err()
	default:
		time.Sleep(duration)
		fmt.Println("timeout")
		return nil
	}
}

func sleep(ctx context.Context, duration time.Duration) error {
	select {
	case <-time.After(duration): // cause short-term memory leak
		fmt.Println("timeout")
		return nil
	case <-ctx.Done():
		fmt.Println("cancelled")
		return ctx.Err()
	}
}

func cancelableSleep() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	// err := sleepy(ctx, 3*time.Second)
	err := sleepx(ctx, 3*time.Second)
	if err != nil {
		println("Error: ", err.Error())
	}
}

type bullet struct{}

func (b *bullet) fire() {
	fmt.Println("fire bullet...")
}

var bulletPool = sync.Pool{
	New: func() interface{} {
		return new(bullet)
	},
}

type userPool struct {
	internal sync.Pool
}

func NewUserPool() *userPool {
	return &userPool{
		internal: sync.Pool{
			New: func() interface{} {
				return new(User)
			},
		},
	}
}

type genericPool[T any] struct {
	internal sync.Pool
}

func NewGenericPool[T any]() *genericPool[T] {
	return &genericPool[T]{
		internal: sync.Pool{
			New: func() interface{} {
				return new(T)
			},
		},
	}
}

func (p *genericPool[T]) Get() *T {
	return p.internal.Get().(*T)
}

func (p *genericPool[T]) Put(t *T) {
	p.internal.Put(t)
}

func reuseBullet() {
	b := bulletPool.Get().(*bullet)
	defer bulletPool.Put(b)
	b.fire()
	p := NewGenericPool[bullet]()
	b2 := p.Get()
	defer p.Put(b2)
	b2.fire()
	// user pool
	uPool := NewUserPool()
	u, ok := uPool.internal.Get().(*User)
	if !ok {
		fmt.Printf("%+v can't cast to User", u)
		return
	}
	fmt.Printf("user: %+v\n", u)
	u.Name = "shen"
	u.Age = 20
	uPool.internal.Put(&User{}) // clean up the user object before put back to pool
	u2 := uPool.internal.Get().(*User)
	fmt.Printf("user: %+v\n", u2)
}

//go:generate stringer -type=HeroType -trimprefix=Hero -linecomment
type HeroType int

const (
	HeroHunter HeroType = iota + 1
	HeroDarwf           // small
	HeroKnight
)

// func (t HeroType) String() string {
// 	return fmt.Sprintf("HeroType[%d]", t)
// }

func stringer() {
	fmt.Println(time.Second)
	// fmt.Println(HeroHunter.String())
	fmt.Println(HeroHunter)
	fmt.Println(HeroDarwf)
	fmt.Println(HeroKnight)
	fmt.Printf("Hero: %s\n", HeroHunter)
	fmt.Printf("Hero: %s\n", HeroDarwf)
}

// const warmUpSeconds int = 10 // don't use it
const warmUpSeconds = 10

// a cleaner representation
const refreshDuration = time.Hour * 24 * 7

func timeDuration() {
	time.Sleep(warmUpSeconds * time.Second)
}

func FetchExpensiveData() (int64, error) {
	fmt.Println("FetchExpensiveData called", time.Now())
	time.Sleep(1 * time.Second)
	return time.Now().Unix() / 1, nil
}

var group singleflight.Group

func doFuncWithSingleflight(key string) {
	// var group singleflight.Group
	v, err, shared := group.Do(key, func() (interface{}, error) {
		return FetchExpensiveData()
	})
	if err != nil {
		fmt.Println("do singleflight function call failed: ", err)
		return
	}
	if shared {
		fmt.Println("do singleflight function call is shared")
	}
	fmt.Println("do singleflight function call success: ", v)
}

func callSingleflightFunc() {
	key := "user:123"
	go doFuncWithSingleflight(key)
	go doFuncWithSingleflight(key)
	go doFuncWithSingleflight(key)

	time.Sleep(2 * time.Second)
	go doFuncWithSingleflight(key)
	go doFuncWithSingleflight(key)
	go doFuncWithSingleflight(key)

	time.Sleep(2 * time.Second)
}
