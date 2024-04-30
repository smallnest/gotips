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
	reuseBullet()
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
