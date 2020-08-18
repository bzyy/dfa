# dfa
detect sensitive words

## Installation
```shell script
$ go get -u github.com/bzyy/dfa
```

## Base Usage
- English
```
d := dfa.NewDFA()
d.Add("sb", "fuck")
fmt.Println(d.Filter("You are sb?")) // => You are **?
fmt.Println(d.Filter("fuck you"))    // =>**** you
```
- Chinese
```
d := dfa.NewDFA()
d.Add("傻逼", "煞笔")
fmt.Println(d.Filter("你是傻逼吗？煞笔说谁呢？")) // => 你是**吗？**说谁呢？
```

##  Noise words
Like "f,u*ck", we can't detect it.How solve?
```
d := dfa.NewDFA()
d.Add("fuck")
fmt.Println(d.Filter("f,u*ck you")) // => f,u*ck you
d.AddSkip("*", ",")
fmt.Println(d.Filter("f,u*ck you")) // => ****** you
d.Add("王八蛋")
d.AddSkip("。")
fmt.Println(d.Filter("王。八蛋.")) // => ****.
```