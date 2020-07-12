package main

func ToUpper(s []byte) []byte
func ToLower(s []byte) []byte
func ToTitle(s []byte) []byte

func Title(s []byte) []byte
nil = []byte{}

func Equal(a, b,[]byte) bool
func EqualFold(a,b,[]byte) bool

func Trim(s []byte, cutset string) []byte
func TrimLeft(s []byte, custer string) []byte
func TrimRight(s []byte,custer string) []byte

func TrimFunc(s []byte, f func(r rune) bool) []byte
func TrimSpace(s []byte) []byte

func Split(s,sp []byte) [][]byte
func Split(s,sep []byte, n int) [][]byte

func Count(s,seq []byte) int
func Replace(s,old,new []byte,n int) []byte

func (b *Buffer) len() int
func (b *Buffer) Cap() int
func (b *Buffer) Next(n int) []byte

func (b *Buffer) Reset