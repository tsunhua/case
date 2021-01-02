## case

已知有 Go 文件 [main.go](compile_001_flex/main.go)，請使用 lex 工具對其進行詞法分析。

## solution

（1）安裝 lex 工具

在 Mac 中執行以下命令安裝 flex (fast lex)：

```bash
brew install flex
```

（2）編寫 [lexer.l](compile_001_flex/lexer.l) 文件，設定詞法規則

（3）生成詞法分析器

```bash
# 生成 c 代碼文件
lex lexer.l
# 生成二進制文件
cc lex.yy.c -o lexer -ll
```

（4）分析 [main.go](compile_001_flex/main.go) 中詞法

```bash
cat main.go | ./lexer
```

執行結果爲：

```
PACKAGE  IDENT 

IMPORT  LPAREN 
        QUOTE IDENT QUOTE 
RPAREN 

IDENT  IDENT LPAREN RPAREN  LBRACE 
        IDENT DOT IDENT LPAREN QUOTE IDENT QUOTE RPAREN 
RBRACE 
```
