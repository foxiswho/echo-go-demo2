package main

import "embed"

// ้ๆๆไปถ static/* template/*
//
//go:embed static/*
//go:embed favicon.ico
var StaticEmbed embed.FS

//go:embed views/*
var TemplatesEmbed embed.FS
