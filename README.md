Repositório com estudos da linguagem GO


realextenso:

Objetivo: Informado um valor monetário em Real o pacote devolverá este valor por extenso. 

Por exemplo: 1234.56 --> Um Mil e Duzentos e Trinta e Quatro Reais e Cincoenta e Seis Centavos

Versão: 0.1 - para testes de implementação, futura versão resultaro num package go. 

numtotext:
Retornar o valor monetário em Dolar por extenso. 

numtoreal: 

Objetivo: Informar um valor em Real e obter como resultado a grafia do valor em Real por extenso. Este pacote é uma memlhoria do pacote realextenso. 

Versão: 0.1 - para testes de implementação, futura versão resultaro num package go. 


Uso: 
$ go run numtoreal/main.go 15685.65
Quinze Mil, Seiscentos e Oitenta e Cinco Reais e Sessenta e Cinco Centavos

$ go run numtoreal/main.go -u 15685.65
QUINZE MIL, SEISCENTOS E OITENTA E CINCO REAIS E SESSENTA E CINCO CENTAVOS



