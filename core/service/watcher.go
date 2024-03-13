package service

import (
	"github.com/fsnotify/fsnotify"
	"lacerate/core/log"
	"os"
	"strings"
	"time"
)

var (
	// 文件事件与事件时间字典
	eventTime = make(map[string]int64)
	// 触发编译时间
	scheduleTime time.Time
)

// Watch 文件监控
type Watch struct {
	Paths  []string // 监控文件路径
	Suffix []string // 监控文件后缀
}

// NewWatch 新建文件监控
func NewWatch(paths []string, suffix []string) *Watch {
	return &Watch{paths, suffix}
}

// Watcher 文件监控
func (w *Watch) Watcher() {
	// 初始化监听器
	log.Log.Info("initialize the file listener...")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic("failed to initialize the file listener: " + err.Error())
	}
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				build := true
				if !w.checkFileSuffix(event.Name) {
					continue
				}
				if event.Op&fsnotify.Chmod == fsnotify.Chmod {
					log.Log.Infof(" skin %s ", event)
					continue
				}
				mt := w.getFileModTime(event.Name)
				if t := eventTime[event.Name]; mt == t {
					log.Log.Infof(" skin %s ", event.String())
					build = false
				}
				eventTime[event.Name] = mt
				if build {
					go func() {
						scheduleTime = time.Now().Add(1 * time.Second)
						for {
							time.Sleep(scheduleTime.Sub(time.Now()))
							if time.Now().After(scheduleTime) {
								break
							}
							return
						}
						log.Log.Infof("triggers a compilation event: %s ", event)
						Compile()
					}()
				}
			case err := <-watcher.Errors:
				log.Log.Errorf("monitoring failed %s ", err)
			}
		}
	}()

	for _, path := range w.Paths {
		log.Log.Infof("listen to folders: [%s] ", path)
		err = watcher.Add(path)
		if err != nil {
			log.Log.Errorf("failed to monitor folder: [%s] ", err)
			os.Exit(2)
		}
	}
	log.Log.Debug("the monitoring is successfully initialized...")
}

// 校验文件后缀名
func (w *Watch) checkFileSuffix(name string) bool {
	for _, s := range w.Suffix {
		if strings.HasSuffix(name, "."+s) {
			return true
		}
	}
	return false
}

// 获取文件最后更新时间
func (w *Watch) getFileModTime(path string) int64 {
	path = strings.Replace(path, "\\", "/", -1)
	f, err := os.Open(path)
	if err != nil {
		log.Log.Errorf("the file failed to open [ %s ]", err)
		return time.Now().Unix()
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	fi, err := f.Stat()
	if err != nil {
		log.Log.Errorf("unable to get file information [ %s ]", err)
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}
