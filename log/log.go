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
 * @file logrus.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Mon May  8 15:39:03 2017
 *
 **/

package log

import (
	"github.com/Sirupsen/logrus"
	"os"
)

//===================================================================
// Public APIs
//===================================================================

func Init(loglevel string) (err error) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	lvl, err := logrus.ParseLevel(loglevel)
	if err != nil {
		lvl = logrus.InfoLevel
	}
	logrus.SetLevel(lvl)
	logrus.SetOutput(os.Stdout)
	return nil
}

func Debugf(format string, v ...interface{}) {
	logrus.Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	logrus.Infof(format, v...)
}

func Warnf(format string, v ...interface{}) {
	logrus.Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	logrus.Errorf(format, v...)
}

func Panicf(format string, v ...interface{}) {
	logrus.Panicf(format, v...)
}

func Fatalf(format string, v ...interface{}) {
	logrus.Fatalf(format, v...)
}

//===================================================================
// Private
//===================================================================
