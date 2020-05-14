# gopathname

root 폴더 내 전체 파일 긁어오는 라이브러리   
   
install   
`go get github.com/sosisusy/gopathname`   
   
### functions   
<hr>   
##### func GetFileNames(root string) []string   
root폴더 기준으로 전체 파일 이름 배열로 리턴   
<hr>
##### func GetTemplate(root string) *template.Template   
root폴더 기준으로 전체 파일을 template에 담아서 리턴
