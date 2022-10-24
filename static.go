package main

import "embed"

// 静态文件 static/* template/*
//
//go:embed static/*
//go:embed favicon.ico
var StaticEmbed embed.FS

//go:embed views/*
var TemplatesEmbed embed.FS
