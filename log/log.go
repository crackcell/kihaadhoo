/***************************************************************
 *
 * Copyright (c) 2017, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the Apache licence
 *
 **************************************************************/

/**
 * 
 *
 * @file log.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Mon May  8 15:39:03 2017
 *
 **/

package log

import (
	"github.com/crackcell/nusadua/config"
	log "github.com/Sirupsen/logrus"
	"os"
)

//===================================================================
// Public APIs
//===================================================================

func Init() (err error) {
	log.SetFormatter(&log.JSONFormatter{})
	lvl, err := log.ParseLevel(config.GlobalConfig.LogLevel)
	if err != nil {
		lvl = log.InfoLevel
	}
	log.SetLevel(lvl)
	log.SetOutput(os.Stdout)
	return nil
}

func Debugf(format string, v ...interface{}) {
	log.Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	log.Infof(format, v...)
}

func Warnf(format string, v ...interface{}) {
	log.Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	log.Errorf(format, v...)
}

func Panicf(format string, v ...interface{}) {
	log.Panicf(format, v...)
}

func Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

//===================================================================
// Private
//===================================================================
