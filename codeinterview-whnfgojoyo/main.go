//калькулятор работает с арабскими цифрами при условии что число не больше 10
//0 1 2 3 4 5 6 7 8 9 10
//а так же с римскими цифрами при условии что вводимое число не больше 10
//i ii iii iv v vi vii viii ix x
//можно использовать операции + - * /

package main
import (
  "fmt" 
  "bufio" 
  "os" 
  //"strings" 
  "strconv"
)
func main() { 
    fmt.Println("Введите два числа и операцию над ними...")
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    i := 0; // счетчик символов строки
    s := "" // сюда будем записывать очередной перебираемый символ строки
    oper := "" // oper!="" означает что мы встретили один из операторав +-*/ и записали его символ в "oper"    
    a1 := ""; // первое арабское число
    a2 := ""; // второе арабское число    
    r_index := -1; // счетчик индекса текущего записываемого рисмского числа
    r1 := ""; // значение первого римского числа
    r2 := ""; // значение второго римского числа
    mess := "" // текст сообщения об ошибке. Если он не нулевой, то значит ошибка есть
    
    var ToNumber1, ToNumber2, result int
    // переберем все введенные с консоли символы
    for i<len(text)-1 {
    s = string(text[i]); // смотрим далее чем является символ s     
    if (s == "+" || s == "-" || s == "*" || s == "/") { 
      if (i==0 && s!="+") {mess="оператор стоит в самом начале выражения и этот оператор не +"; break} else // проверяем не стоит ли оператор в самом начале выражения. Если это "+" то все в порядке, в остальных случаях выводим ошибку  
      if (oper!="") {mess="появление второго оператора"; break} else 
      if (oper=="" && i!=0) {oper=s; r_index = -1;}} else // сбрасываем r_index после полного прочтение первого числа. oper!=0 - значит первое число уже записано
    if (s == "0" || s == "1" || s == "2" || s == "3" || s == "4" || s == "5"  || s == "6" || s == "7" || s == "8" || s == "9") { 
      if (r1!="") {mess="присутствует арабская цифра после римской"; break} else 
      if (oper=="" && a1!="" && s!="0")||(oper!="" && a2!="" && s!="0") {a1=""; a2=""; mess="одно из чисел выражения превышает 10"; break} else 
      if (a1=="" && a2=="" && s=="0") {mess="одно из чисел выражения нулевое"; break} else 
      if (oper=="") {a1=a1+s} else { a2=a2+s}; // выписываем арабские числа в a1 b a2
      if (oper=="" && a1=="1" && s=="0") { a1="10"} else if (oper!="" && a2=="1" && s=="0") { a2="10"}} else // только число 10 двухзначно
    if (s == "i" || s == "v" || s == "x") {
      r_index++; {if (oper=="") {r1=r1+s;} else {r2=r2+s;}}
      if (a1!="") {mess="присутствует римская цифра после арабской"; break} else 
      if (len(r1)>1) {if (string(r1[len(r1)-2]) == "x") {mess="римское число превышает римскую x"; break}};  
      if (len(r2)>1) {if (string(r2[len(r2)-2]) == "x") {mess="римское число превышает римскую x"; break}}
      if (len(r1)>3) {if (oper=="" && string(r1[len(r1)-1])=="i" && string(r1[len(r1)-2])=="i" &&  string(r1[len(r1)-3])=="i" &&  string(r1[len(r1)-4])=="i") {r1=""; r2=""; mess="некорректная запись iiii римского числа"; break}}
      if (len(r2)>3) {if (oper!="" && string(r2[len(r2)-1])=="i" && string(r2[len(r2)-2])=="i" &&  string(r2[len(r2)-3])=="i" &&  string(r2[len(r2)-4])=="i") {r1=""; r2=""; mess="некорректная запись iiii римского числа"; break}}
      //fmt.Println("r2=", r2)
      if (len(r1)>1) {if (oper=="" && string(r1[len(r1)-1]) == "v" && string(r1[len(r1)-2])=="v") {r1=""; r2=""; mess="некорректная запись vv римского числа"; break}}
      if (len(r2)>1) {if (oper!="" && string(r2[len(r2)-1]) == "v" && string(r2[len(r2)-2])=="v") {r1=""; r2=""; mess="некорректная запись vv римского числа"; break}}
      if (len(r1)>1) {if (oper=="" && string(r1[len(r1)-1]) == "v" && string(r1[len(r1)-2])=="x") {r1=""; r2=""; mess="некорректная запись vx римского числа"; break}}
      if (len(r2)>1) {if (oper!="" && string(r2[len(r2)-1]) == "v" && string(r2[len(r2)-2])=="x") {r1=""; r2=""; mess="некорректная запись vx римского числа"; break}}} else
    if (s != " " && s != "  ") {mess="найден символ, который не допустим"}
    i++}
    
      if (a1!="" && a2!= "") {      
      ToNumber1, _ = strconv.Atoi(a1); ToNumber2, _ = strconv.Atoi(a2);      
      if oper=="+" {result = ToNumber1 + ToNumber2} else
      if oper=="-" {result = ToNumber1 - ToNumber2} else
      if oper=="*" {result = ToNumber1 * ToNumber2} else
      if oper=="/" {result = ToNumber1 / ToNumber2} // remainder := ToNumber1 % ToNumber2;      
      } else
    // проверяем(в случае римских чисел) не является ли результат выражения отрицательным либо нулем
    if (r1!="" && r2!= "") {
    // переводим в обычный формат для вычислений - находим ToNumber1 и ToNumber2
    if (r1=="i") {ToNumber1=1} else if r1=="ii" {ToNumber1=2} else if r1=="iii" {ToNumber1=3} else if r1=="iv" {ToNumber1=4} else if r1=="v" {ToNumber1=5} else if r1=="vi" {ToNumber1=6} else if r1=="vii" {ToNumber1=7} else if r1=="viii" {ToNumber1=8} else if r1=="ix" {ToNumber1=9} else if r1=="x" {ToNumber1=10}
    if (r2=="i") {ToNumber2=1} else if r2=="ii" {ToNumber2=2} else if r2=="iii" {ToNumber2=3} else if r2=="iv" {ToNumber2=4} else if r2=="v" {ToNumber2=5} else if r2=="vi" {ToNumber2=6} else if r2=="vii" {ToNumber2=7} else if r2=="viii" {ToNumber2=8} else if r2=="ix" {ToNumber2=9} else if r2=="x" {ToNumber2=10} 
    if oper=="+" {result = ToNumber1 + ToNumber2} else
    if oper=="-" {result = ToNumber1 - ToNumber2; if result<0 {mess="результат вычитание римских чисел не должен быть отрицательным"} else if result==0 {mess="результат вычитания римских чисел равен нулю"}} else
    if oper=="*" {result = ToNumber1 * ToNumber2} else
    if oper=="/" {result = ToNumber1 / ToNumber2}
    }
    if (a1!="" && a2!= "") {fmt.Println(" = ", result)} else // выводим результат операции в случае арабских чисел
    if (r1!="" && r2!= "") {
      // выводим на печать результат в римском формате 
      fmt.Print("= ")
      i=result/10
      j:=0; for j<i {fmt.Print("x"); j++}
      rest := result%10
      if (rest==1) {fmt.Print("i")} else if (rest==2) {fmt.Print("ii")} else if (rest==3) {fmt.Print("iii")} else if (rest==4) {fmt.Print("iv")} else if (rest==5) {fmt.Print("v")} else if (rest==6) {fmt.Print("vi")} else if (rest==7) {fmt.Print("vii")} else if (rest==8) {fmt.Print("viii")} else if (rest==9) {fmt.Print("ix")} else if (rest==10) {fmt.Print("x")}      
      }

  if mess!="" {fmt.Println("ОШИБКА!!! - " + mess)} // печатаем ошибку

}