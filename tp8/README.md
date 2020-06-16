###Creación de una API en Go
#####Evolucionar la calculadora del ejercicio del TP4 para que vía API se puedan efectuar las operaciones enviando valores como parámetros y devolviendo el resultado.
####Endpoints

######GET http://localhost:8080/calc/sum?a=X&b=Y

######GET http://localhost:8080/calc/res?a=X&b=Y

######GET http://localhost:8080/calc/muli?a=X&b=Y

######GET http://localhost:8080/calc/div?a=X&b=Y

####Validaciones
#####http://localhost:8080/calc/all
######HTTP 400 Missing params a,b

#####http://localhost:8080/calc/all?a=string&b=string 
######HTTP 400 Invalid params a,b

#####http://localhost:8080/calc/div?a=X&b=0
######HTTP 400 Invalid param b, cannot be 0 in div

####Retorno
######HTTP 200 {“result”: $number}

#####Nota: Usar NGROKpara crear una url pública para nuestra calculadora